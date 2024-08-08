package codec

import (
	"encoding/json"
	"fmt"
	reflect "reflect"
	"strings"
)

type ABIField struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type SingleActionABI struct {
	ID    uint8                 `json:"id"`
	Name  string                `json:"name"`
	Types map[string][]ABIField `json:"types"`
}

type HavingTypeId interface {
	GetTypeID() uint8
}

func GetVmABIString(actions []HavingTypeId) ([]byte, error) {
	vmABI := make([]SingleActionABI, 0)
	for _, action := range actions {
		actionABI, err := getActionABI(action)
		if err != nil {
			return nil, err
		}
		vmABI = append(vmABI, actionABI)
	}
	return json.MarshalIndent(vmABI, "", "  ")
}

func getActionABI(action HavingTypeId) (SingleActionABI, error) {
	t := reflect.TypeOf(action)

	result := SingleActionABI{
		ID:    action.GetTypeID(),
		Name:  t.Name(),
		Types: make(map[string][]ABIField),
	}

	typesleft := []reflect.Type{t}
	typesAlreadyProcessed := make(map[reflect.Type]bool)

	for i := 0; i < 1000; i++ { //curcuit breakers are always good
		var nextType reflect.Type
		nextTypeFound := false
		for _, anotherType := range typesleft {
			if !typesAlreadyProcessed[anotherType] {
				nextType = anotherType
				nextTypeFound = true
				break
			}
		}
		if !nextTypeFound {
			break
		}

		fields, moreTypes, err := describeStruct(nextType)
		if err != nil {
			return SingleActionABI{}, err
		}

		result.Types[nextType.Name()] = fields
		typesleft = append(typesleft, moreTypes...)

		typesAlreadyProcessed[nextType] = true
	}

	return result, nil
}

func describeStruct(t reflect.Type) ([]ABIField, []reflect.Type, error) { //reflect.Type returns other types to describe
	kind := t.Kind()

	if kind != reflect.Struct {
		return nil, nil, fmt.Errorf("type %s is not a struct", t.String())
	}

	fields := make([]ABIField, 0)
	otherStructsSeen := make([]reflect.Type, 0)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldType := field.Type
		fieldName := field.Name

		if jsonTag := field.Tag.Get("json"); jsonTag != "" {
			parts := strings.Split(jsonTag, ",")
			fieldName = parts[0]
		}

		typeName := fieldType.Name()
		if fieldType.Kind() == reflect.Slice {
			typeName = "[]" + fieldType.Elem().Name()

			if fieldType.Elem().Kind() == reflect.Struct {
				otherStructsSeen = append(otherStructsSeen, fieldType.Elem())
			}
		} else if fieldType.Kind() == reflect.Ptr {
			otherStructsSeen = append(otherStructsSeen, fieldType.Elem())
		}

		fields = append(fields, ABIField{
			Name: fieldName,
			Type: typeName,
		})
	}

	return fields, otherStructsSeen, nil
}
