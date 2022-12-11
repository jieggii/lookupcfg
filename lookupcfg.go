package lookupcfg

import (
	"errors"
	"reflect"
	"strings"
)

const ignoranceTag = "lookupcfg:\"ignore\""

type FieldMeta struct {
	Participate bool // indicates if this field participates in all stuff that this lib does

	ValueSources map[string]string // map of sources of value. E.g {"env": "HOST", "json": "host"}
	DefaultValue string            // default value is stored as string because we parse it from string
}

func parseFieldTag(fieldTag reflect.StructTag) (error, *FieldMeta) {
	fieldTagString := string(fieldTag)

	fieldMeta := &FieldMeta{Participate: true}
	fieldMeta.ValueSources = make(map[string]string)

	if len(fieldTagString) == 0 || strings.Contains(fieldTagString, ignoranceTag) {
		// skips fields without (or with empty) tags and fields with `lookupcfg:"ignore"` tag

		// todo: think about length check. Maybe it is not necessary and panic must be
		// triggered even on empty tags

		fieldMeta.Participate = false
		return nil, fieldMeta
	}

	tags := strings.Split(fieldTagString, " ")
	for _, tag := range tags {
		parts := strings.Split(tag, ":")
		if len(parts) != 2 {
			return errors.New("invalid tag format"), nil
		}
		key := parts[0]
		value := strings.Trim(parts[1], "\"")
		if key == "$default" {
			// todo: check if user tries to set default value multiple times
			// todo: check if type of default value matches field type
			fieldMeta.DefaultValue = value
		} else {
			// todo: check if key already exists
			fieldMeta.ValueSources[key] = value
		}
	}
	return nil, fieldMeta
}

type FieldOfIncorrectType struct {
	StructFieldName string
	NameInSource    string

	Value             string
	ExpectedValueType string

	ConversionError error
}

type ConfigPopulationResult struct {
	//Ok                    bool                   // is true if len() of the next two variables == 0
	MissingFields         []string               // list of fields that are missing
	FieldsOfIncorrectType []FieldOfIncorrectType // array of fields of incorrect type
}

func PopulateConfig(
	source string,
	lookupFunction func(string) (string, bool),
	object interface{},
) *ConfigPopulationResult {
	result := &ConfigPopulationResult{}

	configType := reflect.Indirect(reflect.ValueOf(object)).Type()

	for i := 0; i < configType.NumField(); i++ { // iterating over struct fields
		field := configType.Field(i)
		err, fieldMeta := parseFieldTag(field.Tag)
		if err != nil {
			panicf("Error parsing %v.%v's tag: %v", configType.Name(), field.Name, err)
		}
		if !fieldMeta.Participate {
			//skip fields which do not participate
			continue
		}
		fieldValue := reflect.ValueOf(object).Elem().Field(i)

		valueName, ok := fieldMeta.ValueSources[source]
		if !ok { // if `source` provided as function param is not present in the struct's field tag
			panicf(
				"%v.%v does not have tag \"%v\" (for the source \"%v\"). Use `%v` tag if you want to ignore this field.",
				configType.Name(),
				field.Name,
				source,
				source,
				ignoranceTag,
			)
		}
		value, ok := lookupFunction(valueName)
		if !ok { // if value was not received from the provided source
			if fieldMeta.DefaultValue == "" { // if default value of the field was not indicated
				result.MissingFields = append(result.MissingFields, valueName)
				continue
			}
			value = fieldMeta.DefaultValue
		}
		fieldType := field.Type
		err = universalSet(fieldType.Kind(), fieldValue, value)
		if err != nil {
			result.FieldsOfIncorrectType = append(
				result.FieldsOfIncorrectType, FieldOfIncorrectType{
					StructFieldName:   field.Name,
					NameInSource:      valueName,
					Value:             value,
					ExpectedValueType: fieldType.String(),
					ConversionError:   err,
				},
			)
			//fmt.Printf("Error setting %v (%v) to %v. Source: %v\n", field.Name, field.Type, value, source)
		} //else {
		//fmt.Printf("Set %v (%v) to %v. Source: %v\n", field.Name, field.Type, value, source)
		//}
	}
	return result
}
