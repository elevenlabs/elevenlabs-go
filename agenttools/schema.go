package agenttools

import (
	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
)

// SchemaKind is the kind of a schema: one of the four literals, an object or an array. Its value is
// the wire "type" name.
type SchemaKind string

const (
	SchemaKindString  SchemaKind = "string"
	SchemaKindInteger SchemaKind = "integer"
	SchemaKindNumber  SchemaKind = "number"
	SchemaKindBoolean SchemaKind = "boolean"
	SchemaKindObject  SchemaKind = "object"
	SchemaKindArray   SchemaKind = "array"
)

// IsLiteral reports whether the kind is a string, integer, number or boolean.
func (k SchemaKind) IsLiteral() bool {
	return k != SchemaKindObject && k != SchemaKindArray
}

// Schema is a tool param's JSON schema, in the dialect the API uses, and where its value comes
// from. It is implemented by LiteralSchema (String, Integer, Number, Bool), ObjectSchema (Object)
// and ArraySchema (Array); methods that only apply to one kind live on that type, so they don't
// compile on another. Use a type switch or assertion to edit a Schema of unknown kind.
//
// Schemas are values: every method returns a new schema and never changes the one it is called
// on, so one schema can be reused across tools and properties.
//
// Every literal needs exactly one value source, wherever it appears: Describe (the LLM provides
// it), FromVariable, Constant, SystemProvided or Omit. An object or array may have at most one,
// and needs none; on a container Describe documents it for the LLM.
type Schema interface {
	// Kind is the schema's kind.
	Kind() SchemaKind
	// IsOptional reports whether the parent leaves this property out of its `required` list.
	IsOptional() bool
	// ToJSON returns the schema in API wire format. It doesn't validate.
	ToJSON() map[string]any

	state() *schemaState
}

// schemaState is everything a schema holds. Modeled fields are typed; loaded fields the builder
// doesn't model (and loaded API defaults for unset fields) stay in extra and are written back
// unchanged. A state is never changed once a Schema holds it.
type schemaState struct {
	kind SchemaKind
	// literalType is a loaded literal's `type`, as loaded (the API also allows a list such as
	// ["integer", "null"]), or nil.
	literalType any
	// typeLoaded is whether the loaded schema declared a `type`.
	typeLoaded      bool
	description     *string
	dynamicVariable *string
	hasConstant     bool
	// constant is the constant in wire format; it may be nil (JSON null).
	constant any
	// constantErr says why Constant couldn't convert its value to JSON.
	constantErr    string
	systemProvided bool
	omitted        bool
	optional       bool
	hasChoices     bool
	choices        any
	allowedFrom    *string
	// allowedValues is a loaded allowed_values object, so its other fields survive edits.
	allowedValues map[string]any
	items         Schema
	properties    []property
	// loadedRequired is a loaded object's `required`, so its order survives edits.
	loadedRequired []string
	// constraintsExtra holds a loaded required_constraints object's fields other than any_of and
	// all_of (non-nil when one was loaded).
	constraintsExtra map[string]any
	// anyOf and allOf are the constraint groups in wire format ({"required": [...]}, plus any
	// loaded fields), each kept whole. hasAnyOf / hasAllOf say whether the list is present, so a
	// loaded empty list round-trips.
	anyOf, allOf       []any
	hasAnyOf, hasAllOf bool
	extra              map[string]any
	// malformed lists the loaded modeled fields of the wrong shape, kept in extra as loaded. Build
	// reports each one (invalid_model), since the generated model might drop them silently.
	malformed []string
}

type property struct {
	name   string
	schema Schema
}

func newState(kind SchemaKind) *schemaState {
	return &schemaState{kind: kind}
}

func (s *schemaState) clone() *schemaState {
	c := *s
	return &c
}

func (s *schemaState) with(change func(*schemaState)) *schemaState {
	c := s.clone()
	change(c)
	return c
}

func (s *schemaState) setConstant(value any) {
	s.hasConstant = true
	s.constantErr = ""
	wire, err := wirejson.Normalize(value)
	if err != nil {
		s.constant = nil
		s.constantErr = err.Error()
		return
	}
	s.constant = wire
}

func (s *schemaState) property(name string) (Schema, int) {
	for i, p := range s.properties {
		if p.name == name {
			return p.schema, i
		}
	}
	return nil, -1
}

// required is a loaded object's `required` order for the required properties it still has, then
// new ones in property order.
func (s *schemaState) required() []string {
	var required []string
	for _, p := range s.properties {
		if !p.schema.IsOptional() {
			required = append(required, p.name)
		}
	}
	var out []string
	for _, name := range s.loadedRequired {
		if contains(required, name) && !contains(out, name) {
			out = append(out, name)
		}
	}
	for _, name := range required {
		if !contains(out, name) {
			out = append(out, name)
		}
	}
	return out
}

// toWire returns the schema in wire format (unsorted). Modeled fields that are unset leave the
// matching extra field (an API default) in place.
func (s *schemaState) toWire() map[string]any {
	out := cloneMap(s.extra)
	if s.description != nil {
		out["description"] = *s.description
	}
	if s.dynamicVariable != nil {
		out["dynamic_variable"] = *s.dynamicVariable
	}
	if s.omitted {
		out["is_omitted"] = true
	}
	if s.hasConstant {
		out["constant_value"] = wirejson.Clone(s.constant)
	}
	switch s.kind {
	case SchemaKindObject:
		out["type"] = "object"
		if len(s.properties) > 0 {
			properties := make(map[string]any, len(s.properties))
			for _, p := range s.properties {
				properties[p.name] = p.schema.state().toWire()
			}
			out["properties"] = properties
		}
		if required := s.required(); len(required) > 0 {
			list := make([]any, len(required))
			for i, name := range required {
				list[i] = name
			}
			out["required"] = list
		}
		if s.constraintsExtra != nil || s.hasAnyOf || s.hasAllOf {
			constraints := cloneMap(s.constraintsExtra)
			if s.hasAnyOf {
				constraints["any_of"] = wirejson.Clone(s.anyOf)
			}
			if s.hasAllOf {
				constraints["all_of"] = wirejson.Clone(s.allOf)
			}
			out["required_constraints"] = constraints
		}
	case SchemaKindArray:
		out["type"] = "array"
		if s.items != nil {
			out["items"] = s.items.state().toWire()
		}
	default:
		if s.literalType != nil {
			out["type"] = wirejson.Clone(s.literalType)
		} else {
			out["type"] = string(s.kind)
		}
		if s.systemProvided {
			out["is_system_provided"] = true
		}
		if s.hasChoices {
			out["enum"] = wirejson.Clone(s.choices)
		}
		if s.allowedValues != nil || s.allowedFrom != nil {
			// A loaded object keeps its other fields; the variable name is overlaid.
			allowed := cloneMap(s.allowedValues)
			if s.allowedFrom != nil {
				allowed["dynamic_variable"] = *s.allowedFrom
			}
			out["allowed_values"] = allowed
		}
	}
	return out
}

func canonicalSchema(s *schemaState) map[string]any {
	canonical, err := wirejson.Canonical(s.toWire())
	if err != nil {
		// A wire tree always converts; this is unreachable.
		return s.toWire()
	}
	return canonical.(map[string]any)
}

// String is a string schema.
//
// It is a package function, not a method, so no type has a String method that go vet would read
// as a broken fmt.Stringer.
func String() LiteralSchema { return LiteralSchema{s: newState(SchemaKindString)} }

// Integer is an integer schema.
func Integer() LiteralSchema { return LiteralSchema{s: newState(SchemaKindInteger)} }

// Number is a number schema.
func Number() LiteralSchema { return LiteralSchema{s: newState(SchemaKindNumber)} }

// Bool is a boolean schema. It is named Bool, not Boolean, to match every SDK (boolean is reserved
// in Java).
func Bool() LiteralSchema { return LiteralSchema{s: newState(SchemaKindBoolean)} }

// Object is an object schema. Add properties with Property, in order: `required` is written in
// property order.
func Object() ObjectSchema { return ObjectSchema{s: newState(SchemaKindObject)} }

// Array is an array schema whose elements match items.
func Array(items Schema) ArraySchema {
	s := newState(SchemaKindArray)
	s.items = items
	return ArraySchema{s: s}
}

// LiteralSchema is a string, integer, number or boolean schema. The zero value is a string schema.
type LiteralSchema struct{ s *schemaState }

func (l LiteralSchema) state() *schemaState {
	if l.s == nil {
		return newState(SchemaKindString)
	}
	return l.s
}

// Kind is SchemaKindString, SchemaKindInteger, SchemaKindNumber or SchemaKindBoolean.
func (l LiteralSchema) Kind() SchemaKind { return l.state().kind }

// IsOptional reports whether the parent leaves this property out of `required`.
func (l LiteralSchema) IsOptional() bool { return l.state().optional }

// ToJSON returns the schema in API wire format. It doesn't validate.
func (l LiteralSchema) ToJSON() map[string]any { return canonicalSchema(l.state()) }

// Describe makes the LLM provide the value, guided by text.
func (l LiteralSchema) Describe(text string) LiteralSchema {
	return LiteralSchema{l.state().with(func(s *schemaState) { s.description = &text })}
}

// FromVariable reads the value from the dynamic variable name when the tool runs.
func (l LiteralSchema) FromVariable(name string) LiteralSchema {
	return LiteralSchema{l.state().with(func(s *schemaState) { s.dynamicVariable = &name })}
}

// Constant makes the value always value: a JSON value (any value encoding/json marshals) matching
// the schema's type, checked by Build.
func (l LiteralSchema) Constant(value any) LiteralSchema {
	return LiteralSchema{l.state().with(func(s *schemaState) { s.setConstant(value) })}
}

// SystemProvided makes the system fill the value in at runtime.
func (l LiteralSchema) SystemProvided() LiteralSchema {
	return LiteralSchema{l.state().with(func(s *schemaState) { s.systemProvided = true })}
}

// Omit leaves the param out of the request. Only allowed on an optional property.
func (l LiteralSchema) Omit() LiteralSchema {
	return LiteralSchema{l.state().with(func(s *schemaState) { s.omitted = true })}
}

// Optional leaves this property out of its parent's `required`. Properties are required unless
// marked optional.
func (l LiteralSchema) Optional() LiteralSchema {
	return LiteralSchema{l.state().with(func(s *schemaState) { s.optional = true })}
}

// Choices sets the values the LLM may choose from (`enum`). Strings only. It is named Choices
// because enum is reserved in several SDK languages.
func (l LiteralSchema) Choices(values ...string) LiteralSchema {
	list := make([]any, len(values))
	for i, v := range values {
		list[i] = v
	}
	return LiteralSchema{l.state().with(func(s *schemaState) {
		s.hasChoices = true
		s.choices = list
	})}
}

// AllowedFrom makes the runtime reject an LLM value that isn't in the list held by the dynamic
// variable variable (`allowed_values`). The list isn't shown to the LLM. Needs Describe.
func (l LiteralSchema) AllowedFrom(variable string) LiteralSchema {
	return LiteralSchema{l.state().with(func(s *schemaState) { s.allowedFrom = &variable })}
}

// ObjectSchema is an object schema with named properties. The zero value is an empty object.
type ObjectSchema struct{ s *schemaState }

func (o ObjectSchema) state() *schemaState {
	if o.s == nil {
		return newState(SchemaKindObject)
	}
	return o.s
}

// Kind is SchemaKindObject.
func (o ObjectSchema) Kind() SchemaKind { return SchemaKindObject }

// IsOptional reports whether the parent leaves this property out of `required`.
func (o ObjectSchema) IsOptional() bool { return o.state().optional }

// ToJSON returns the schema in API wire format. It doesn't validate.
func (o ObjectSchema) ToJSON() map[string]any { return canonicalSchema(o.state()) }

// Describe documents the object for the LLM. It isn't a value source for its properties.
func (o ObjectSchema) Describe(text string) ObjectSchema {
	return ObjectSchema{o.state().with(func(s *schemaState) { s.description = &text })}
}

// FromVariable reads the whole object from the dynamic variable name when the tool runs.
func (o ObjectSchema) FromVariable(name string) ObjectSchema {
	return ObjectSchema{o.state().with(func(s *schemaState) { s.dynamicVariable = &name })}
}

// Constant makes the value always value, a JSON object.
func (o ObjectSchema) Constant(value any) ObjectSchema {
	return ObjectSchema{o.state().with(func(s *schemaState) { s.setConstant(value) })}
}

// Omit leaves the param out of the request. Only allowed on an optional property.
func (o ObjectSchema) Omit() ObjectSchema {
	return ObjectSchema{o.state().with(func(s *schemaState) { s.omitted = true })}
}

// Optional leaves this property out of its parent's `required`.
func (o ObjectSchema) Optional() ObjectSchema {
	return ObjectSchema{o.state().with(func(s *schemaState) { s.optional = true })}
}

// Property adds the property name, or replaces it in place if it exists. A nil schema is an empty
// string schema, which Build reports as having no value source.
func (o ObjectSchema) Property(name string, schema Schema) ObjectSchema {
	if schema == nil {
		schema = LiteralSchema{}
	}
	return ObjectSchema{o.state().with(func(s *schemaState) {
		properties := append([]property(nil), s.properties...)
		if _, i := s.property(name); i >= 0 {
			properties[i] = property{name: name, schema: schema}
		} else {
			properties = append(properties, property{name: name, schema: schema})
		}
		s.properties = properties
	})}
}

// Without removes the property name. Removing one that doesn't exist does nothing.
func (o ObjectSchema) Without(name string) ObjectSchema {
	return ObjectSchema{o.state().with(func(s *schemaState) {
		var properties []property
		for _, p := range s.properties {
			if p.name != name {
				properties = append(properties, p)
			}
		}
		s.properties = properties
	})}
}

// RequireAnyOf requires at least one of these properties: it appends {required: names} to
// `required_constraints.any_of`.
func (o ObjectSchema) RequireAnyOf(names ...string) ObjectSchema {
	group := constraintGroup(names)
	return ObjectSchema{o.state().with(func(s *schemaState) {
		s.anyOf = append(append([]any(nil), s.anyOf...), group)
		s.hasAnyOf = true
	})}
}

// RequireAllOf requires all of these properties: it appends {required: names} to
// `required_constraints.all_of`.
func (o ObjectSchema) RequireAllOf(names ...string) ObjectSchema {
	group := constraintGroup(names)
	return ObjectSchema{o.state().with(func(s *schemaState) {
		s.allOf = append(append([]any(nil), s.allOf...), group)
		s.hasAllOf = true
	})}
}

func constraintGroup(names []string) map[string]any {
	list := make([]any, len(names))
	for i, name := range names {
		list[i] = name
	}
	return map[string]any{"required": list}
}

// PropertyNames are the property names, in order.
func (o ObjectSchema) PropertyNames() []string {
	s := o.state()
	names := make([]string, len(s.properties))
	for i, p := range s.properties {
		names[i] = p.name
	}
	return names
}

// PropertySchema is the schema of property name, or nil if there is none.
func (o ObjectSchema) PropertySchema(name string) Schema {
	schema, _ := o.state().property(name)
	return schema
}

// ArraySchema is an array schema whose elements match one schema.
type ArraySchema struct{ s *schemaState }

func (a ArraySchema) state() *schemaState {
	if a.s == nil {
		return newState(SchemaKindArray)
	}
	return a.s
}

// Kind is SchemaKindArray.
func (a ArraySchema) Kind() SchemaKind { return SchemaKindArray }

// IsOptional reports whether the parent leaves this property out of `required`.
func (a ArraySchema) IsOptional() bool { return a.state().optional }

// ToJSON returns the schema in API wire format. It doesn't validate.
func (a ArraySchema) ToJSON() map[string]any { return canonicalSchema(a.state()) }

// Describe documents the array for the LLM. It isn't a value source for its items.
func (a ArraySchema) Describe(text string) ArraySchema {
	return ArraySchema{a.state().with(func(s *schemaState) { s.description = &text })}
}

// FromVariable reads the whole array from the dynamic variable name when the tool runs.
func (a ArraySchema) FromVariable(name string) ArraySchema {
	return ArraySchema{a.state().with(func(s *schemaState) { s.dynamicVariable = &name })}
}

// Constant makes the value always value, a JSON array.
func (a ArraySchema) Constant(value any) ArraySchema {
	return ArraySchema{a.state().with(func(s *schemaState) { s.setConstant(value) })}
}

// Omit leaves the param out of the request. Only allowed on an optional property.
func (a ArraySchema) Omit() ArraySchema {
	return ArraySchema{a.state().with(func(s *schemaState) { s.omitted = true })}
}

// Optional leaves this property out of its parent's `required`.
func (a ArraySchema) Optional() ArraySchema {
	return ArraySchema{a.state().with(func(s *schemaState) { s.optional = true })}
}

// Items is the element schema. It is nil only for a loaded array without one.
func (a ArraySchema) Items() Schema { return a.state().items }

// SchemaFromJSON loads a schema in API wire format (snake_case keys): JSON text (a string, []byte
// or json.RawMessage) or a parsed value such as a map. The schema's kind comes from its `type`
// (for a list such as ["integer", "null"], the non-null member); one without a `type` is an object
// if it has properties, an array if it has items, and a string otherwise.
func SchemaFromJSON(data any) (Schema, error) {
	tree, err := toTree(data)
	if err != nil {
		return nil, err
	}
	obj, ok := stripNulls(tree).(*object)
	if !ok {
		return nil, errNotObject("schema")
	}
	return loadSchema(obj, false, kindOf(obj)), nil
}
