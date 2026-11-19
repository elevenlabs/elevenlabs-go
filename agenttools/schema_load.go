package agenttools

import "fmt"

func errNotObject(what string) error {
	return fmt.Errorf("agenttools: the %s JSON must be an object", what)
}

// kindOf is the kind of a loaded schema. A literal's type may be a list such as
// ["integer", "null"]; its kind is the non-null member.
func kindOf(o *object) SchemaKind {
	t, hasType := o.get("type")
	hasType = hasType && t != nil
	_, hasProperties := o.get("properties")
	_, hasItems := o.get("items")
	if t == "object" || (!hasType && hasProperties) {
		return SchemaKindObject
	}
	if t == "array" || (!hasType && hasItems) {
		return SchemaKindArray
	}
	var name string
	switch v := t.(type) {
	case string:
		name = v
	case []any:
		for _, item := range v {
			if s, ok := item.(string); ok && s != "null" {
				name = s
				break
			}
		}
	}
	switch name {
	case "integer":
		return SchemaKindInteger
	case "number":
		return SchemaKindNumber
	case "boolean":
		return SchemaKindBoolean
	}
	return SchemaKindString
}

// isNullable reports whether a literal's type is a list that includes "null".
func isNullable(t any) bool {
	list, ok := t.([]any)
	if !ok {
		return false
	}
	for _, item := range list {
		if item == "null" {
			return true
		}
	}
	return false
}

// containerSchema loads a schema the API only allows to be an object (a body, a response, the
// query params or a client tool's params). One with no `type` is an object, even when empty, and
// is written with "type": "object". A declared `type` is kept, so validation can report a schema
// that isn't an object.
func containerSchema(o *object) Schema {
	if t, ok := o.get("type"); !ok || t == nil {
		return loadSchema(o, false, SchemaKindObject)
	}
	return loadSchema(o, false, kindOf(o))
}

// isModeled reports whether the builder models key for a schema of kind. Anything else is kept as
// loaded.
func isModeled(kind SchemaKind, key string) bool {
	switch key {
	case "type", "description", "dynamic_variable", "constant_value", "is_omitted":
		return true
	}
	switch kind {
	case SchemaKindObject:
		return key == "properties" || key == "required" || key == "required_constraints"
	case SchemaKindArray:
		return key == "items"
	}
	return key == "is_system_provided" || key == "enum" || key == "allowed_values"
}

// isUnset reports whether value is what the API returns for an unset field. A literal's unset
// constant is "", so "" can't be a literal constant.
func isUnset(key string, value any, kind SchemaKind) bool {
	switch key {
	case "description", "dynamic_variable":
		return value == ""
	case "constant_value":
		return kind.IsLiteral() && value == ""
	case "is_system_provided", "is_omitted":
		return value == false
	case "required":
		list, ok := value.([]any)
		return ok && len(list) == 0
	case "properties":
		obj, ok := value.(*object)
		return ok && len(obj.keys) == 0
	}
	return false
}

// loadSchema loads a wire-format schema whose nulls are already stripped (stripNulls). optional is
// whether its parent leaves it out of `required`.
//
// The API fills unset fields with defaults ("", false, []). They aren't value sources, but are
// kept as extra fields so a pulled config writes back unchanged. A modeled field of the wrong
// shape is kept in extra as loaded and recorded as malformed, so ToJSON writes it back and Build
// reports it rather than the builder dropping it.
func loadSchema(o *object, optional bool, kind SchemaKind) Schema {
	s := newState(kind)
	s.optional = optional
	s.extra = map[string]any{}
	keep := func(key string, value any, malformed bool) {
		s.extra[key] = plain(value)
		if malformed {
			s.malformed = append(s.malformed, key)
		}
	}
	set := map[string]any{}
	for _, key := range o.keys {
		value := o.vals[key]
		if value == nil {
			// Only a constant_value is left null by stripNulls. It is a value on a literal whose
			// type list includes "null"; anywhere else it means unset.
			if key == "constant_value" && kind.IsLiteral() {
				if t, _ := o.get("type"); isNullable(t) {
					s.hasConstant = true
					s.constant = nil
				}
			}
			continue
		}
		if !isModeled(kind, key) || isUnset(key, value, kind) {
			keep(key, value, false)
			continue
		}
		set[key] = value
	}

	if t, ok := set["type"]; ok {
		s.typeLoaded = true
		if kind.IsLiteral() {
			s.literalType = plain(t)
		}
	}
	for _, key := range []string{"description", "dynamic_variable"} {
		value, ok := set[key]
		if !ok {
			continue
		}
		text, isString := value.(string)
		if !isString {
			keep(key, value, true)
			continue
		}
		if key == "description" {
			s.description = &text
		} else {
			s.dynamicVariable = &text
		}
	}
	if value, ok := set["constant_value"]; ok {
		s.hasConstant = true
		s.constant = plain(value)
	}
	if value, ok := set["is_omitted"]; ok {
		if value == true {
			s.omitted = true
		} else {
			keep("is_omitted", value, true)
		}
	}

	switch kind {
	case SchemaKindObject:
		loadObject(s, set, keep)
	case SchemaKindArray:
		if value, ok := set["items"]; ok {
			if items, isObject := value.(*object); isObject {
				s.items = loadSchema(items, true, kindOf(items))
			} else {
				keep("items", value, true)
			}
		}
	default:
		if value, ok := set["is_system_provided"]; ok {
			if value == true {
				s.systemProvided = true
			} else {
				keep("is_system_provided", value, true)
			}
		}
		if value, ok := set["enum"]; ok {
			if _, isList := stringList(value); isList {
				s.hasChoices = true
				s.choices = plain(value)
			} else {
				keep("enum", value, true)
			}
		}
		if value, ok := set["allowed_values"]; ok {
			allowed, isObject := value.(*object)
			if !isObject {
				keep("allowed_values", value, true)
				break
			}
			s.allowedValues = plain(allowed).(map[string]any)
			if variable, has := allowed.get("dynamic_variable"); has {
				if name, isString := variable.(string); isString {
					s.allowedFrom = &name
				} else {
					s.malformed = append(s.malformed, "allowed_values")
				}
			}
		}
	}

	switch kind {
	case SchemaKindObject:
		return ObjectSchema{s: s}
	case SchemaKindArray:
		return ArraySchema{s: s}
	}
	return LiteralSchema{s: s}
}

func loadObject(s *schemaState, set map[string]any, keep func(string, any, bool)) {
	propertiesValue, hasProperties := set["properties"]
	requiredValue, hasRequired := set["required"]
	properties, propertiesOK := propertiesValue.(*object)
	if propertiesOK {
		for _, name := range properties.keys {
			if _, isObject := properties.vals[name].(*object); !isObject {
				propertiesOK = false
				break
			}
		}
	}
	required, requiredOK := stringList(requiredValue)
	if (hasProperties && !propertiesOK) || (hasRequired && !requiredOK) {
		// Each needs the other: a property's optional() comes from `required`.
		if hasProperties {
			keep("properties", propertiesValue, !propertiesOK)
		}
		if hasRequired {
			keep("required", requiredValue, !requiredOK)
		}
	} else {
		if hasRequired {
			s.loadedRequired = required
		}
		if hasProperties {
			for _, name := range properties.keys {
				child := properties.vals[name].(*object)
				s.properties = append(s.properties, property{
					name:   name,
					schema: loadSchema(child, !contains(required, name), kindOf(child)),
				})
			}
		}
	}

	value, ok := set["required_constraints"]
	if !ok {
		return
	}
	constraints, isObject := value.(*object)
	if !isObject {
		keep("required_constraints", value, true)
		return
	}
	// The loaded object keeps its other fields, and each group is kept whole.
	s.constraintsExtra = map[string]any{}
	malformed := false
	for _, key := range constraints.keys {
		item := constraints.vals[key]
		groups, isList := item.([]any)
		if (key != "any_of" && key != "all_of") || !isList {
			s.constraintsExtra[key] = plain(item)
			malformed = malformed || key == "any_of" || key == "all_of"
			continue
		}
		for _, group := range groups {
			if _, ok := groupNames(plain(group)); !ok {
				malformed = true
			}
		}
		if key == "any_of" {
			s.anyOf, s.hasAnyOf = plain(item).([]any), true
		} else {
			s.allOf, s.hasAllOf = plain(item).([]any), true
		}
	}
	if malformed {
		s.malformed = append(s.malformed, "required_constraints")
	}
}

// groupNames reads a constraint group, {"required": [names]}, reporting false for another shape.
func groupNames(group any) ([]string, bool) {
	m, ok := group.(map[string]any)
	if !ok {
		return nil, false
	}
	return stringList(m["required"])
}
