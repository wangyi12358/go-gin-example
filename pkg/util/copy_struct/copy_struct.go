package copy_struct

import "reflect"

func CopyStruct(source, dest interface{}) error {
	sVal := reflect.ValueOf(source).Elem()
	dVal := reflect.ValueOf(dest).Elem()

	for i := 0; i < sVal.NumField(); i++ {
		field := sVal.Type().Field(i)

		dField := dVal.FieldByName(field.Name)
		if !dField.IsValid() {
			continue
		}

		dField.Set(sVal.Field(i))
	}

	return nil
}
