package model

import (
	"errors"
	"fmt"
	"reflect"
)

type Base struct {
	Table string
}

func (t *Base) SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)
	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}
	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}
	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}
	structFieldValue.Set(val)
	return nil
}

func (t *Base) MapToModel(obj interface{}, m map[string]interface{}) interface{} {
	for k, v := range m {
		err := t.SetField(obj, k, v)
		if err != nil {
			return nil
		}
	}

	return obj
}
