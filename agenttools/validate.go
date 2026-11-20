package agenttools

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
)

// checkSchema checks a schema and everything under it. path names it in messages, such as
// "body.reason". isProperty is whether the schema is a property (or a param), so Omit needs it
// optional.
func checkSchema(schema Schema, path string, isProperty bool, issues *[]Issue) {
	s := schema.state()
	add := func(code IssueCode, message string) {
		*issues = append(*issues, errorIssue(code, path, message))
	}
	// The API reads "" as unset, so an empty description or variable name isn't a value source.
	described := s.description != nil && *s.description != ""
	sources := 0
	for _, set := range []bool{
		described,
		s.dynamicVariable != nil && *s.dynamicVariable != "",
		s.hasConstant,
		s.systemProvided,
		s.omitted,
	} {
		if set {
			sources++
		}
	}
	if sources > 1 {
		add(ConflictingValueSource, fmt.Sprintf("%q sets more than one of Describe, FromVariable, Constant, SystemProvided and Omit", path))
	}
	if s.omitted && isProperty && !s.optional {
		add(OmittedRequired, fmt.Sprintf("%q uses Omit, so it must be Optional", path))
	}
	if s.hasChoices && s.kind != SchemaKindString {
		add(InvalidChoices, fmt.Sprintf("%q uses Choices, which only applies to strings", path))
	}
	if (s.allowedFrom != nil || s.allowedValues != nil) && !described {
		add(AllowedValuesNeedsDescription, fmt.Sprintf("%q uses AllowedFrom, which needs Describe so the LLM provides the value", path))
	}
	if s.hasConstant {
		if message := constantProblem(s); message != "" {
			add(ConstantTypeMismatch, fmt.Sprintf("%q %s", path, message))
		}
	}
	// The API rejects a literal with no value source wherever it appears, response schemas too.
	if sources == 0 && s.kind.IsLiteral() {
		add(MissingValueSource, fmt.Sprintf("%q needs one of Describe, FromVariable, Constant, SystemProvided or Omit", path))
	}

	switch s.kind {
	case SchemaKindObject:
		for _, groups := range [][]any{s.anyOf, s.allOf} {
			for _, group := range groups {
				names, _ := groupNames(group)
				for _, name := range names {
					if _, i := s.property(name); i < 0 {
						add(UnknownRequiredProperty, fmt.Sprintf("%q requires %q in RequireAnyOf/RequireAllOf, but has no such property", path, name))
					}
				}
			}
		}
		for _, p := range s.properties {
			checkSchema(p.schema, path+"."+p.name, true, issues)
		}
	case SchemaKindArray:
		switch {
		case s.items != nil:
			checkSchema(s.items, path+"[]", false, issues)
		case contains(s.malformed, "items"):
			// Loaded items of the wrong shape are reported by Build as invalid_model.
		default:
			add(MissingArrayItems, fmt.Sprintf("Array %q has no items schema", path))
		}
	}
}

// constantProblem says what is wrong with a schema's constant, or "" if nothing is. A constant
// whose JSON type doesn't match the schema is rejected, and so is one the generated request model
// would change rather than send exactly.
func constantProblem(s *schemaState) string {
	constant := s.constant
	switch {
	case s.constantErr != "":
		return "has a constant that can't be converted to JSON: " + s.constantErr
	case !matchesKind(s.kind, constant, isNullable(s.literalType)):
		return "has a constant that isn't a " + string(s.kind)
	case constant == "":
		return `can't have "" as a constant: the API reads it as unset`
	}
	if bad := badNumber(constant); bad != "" {
		return fmt.Sprintf("has the constant number %s, which isn't a finite JSON number", bad)
	}
	if n, ok := constant.(json.Number); ok && s.kind == SchemaKindInteger && !fitsInt64(string(n)) {
		// The request model holds an integer constant as an int64; beyond that it would become a
		// float64 and Build would send a different value from ToJSON.
		return fmt.Sprintf("has the constant %s, which is outside the 64-bit integer range and can't be sent exactly", n)
	}
	return ""
}

func matchesKind(kind SchemaKind, value any, nullable bool) bool {
	if value == nil && nullable {
		return true
	}
	switch kind {
	case SchemaKindString:
		_, ok := value.(string)
		return ok
	case SchemaKindInteger:
		// Judged from every digit of the text, so 3.0 and 1e2 are integers and 1.5 is not.
		n, ok := value.(json.Number)
		return ok && wirejson.IsInteger(string(n))
	case SchemaKindNumber:
		_, ok := value.(json.Number)
		return ok
	case SchemaKindBoolean:
		_, ok := value.(bool)
		return ok
	case SchemaKindObject:
		_, ok := value.(map[string]any)
		return ok
	case SchemaKindArray:
		_, ok := value.([]any)
		return ok
	}
	return false
}

// badNumber returns the first number in value that isn't valid JSON number text or doesn't fit a
// float64 (such as 1e400, which JavaScript reads as Infinity), or "".
func badNumber(value any) string {
	switch v := value.(type) {
	case json.Number:
		text := string(v)
		var check json.Number
		if json.Unmarshal([]byte(text), &check) != nil || check != v {
			return text
		}
		if f, err := strconv.ParseFloat(text, 64); err != nil && math.IsInf(f, 0) {
			return text
		}
	case []any:
		for _, item := range v {
			if bad := badNumber(item); bad != "" {
				return bad
			}
		}
	case map[string]any:
		for _, item := range v {
			if bad := badNumber(item); bad != "" {
				return bad
			}
		}
	}
	return ""
}

// fitsInt64 reports whether integer-valued number text is within the int64 range.
func fitsInt64(text string) bool {
	f, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return false
	}
	switch abs := math.Abs(f); {
	case abs < 9e18:
		return true
	case abs > 9.3e18:
		return false
	}
	// Close to the limit, decide exactly. The value is near 2^63, so its exponent is small.
	r, ok := new(big.Rat).SetString(text)
	return ok && r.IsInt() && r.Num().IsInt64()
}

// checkModelShape reports each loaded field under a schema that has the wrong shape for the
// generated request model (kept as loaded, so ToJSON writes it back), as invalid_model.
func checkModelShape(schema Schema, path string, issues *[]Issue) {
	s := schema.state()
	for _, key := range s.malformed {
		*issues = append(*issues, errorIssue(InvalidModel, path, fmt.Sprintf("%q has a loaded %s of the wrong shape, which the generated tool config model can't take. Fix it in the source JSON", path, key)))
	}
	for _, p := range s.properties {
		checkModelShape(p.schema, path+"."+p.name, issues)
	}
	if s.items != nil {
		checkModelShape(s.items, path+"[]", issues)
	}
}
