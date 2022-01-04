package runtime

import (
	"bytes"
	"reflect"
	"text/template"
)

func RenderTemplate(templateStruct interface{}, payload *map[string]interface{}) {
	v := reflect.ValueOf(templateStruct).Elem()

	renderTemplateReflectValue(&v, payload)
}

func renderTemplateReflectValue(v *reflect.Value, payload *map[string]interface{}) {
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldType := typeOfS.Field(i).Type
		field := v.Field(i)

		if fieldType.Kind() == reflect.String {
			var renderedTemplate bytes.Buffer

			t := template.Must(template.New("").Parse(field.String()))
			t.Execute(&renderedTemplate, payload)

			field.SetString(renderedTemplate.String())
		} else if fieldType.Kind() == reflect.Struct {
			renderTemplateReflectValue(&field, payload)
		} else if fieldType.Kind() == reflect.Slice {
			for i := 0; i < field.Len(); i++ {
				sliceField := field.Index(i)

				renderTemplateReflectValue(&sliceField, payload)
			}
		}
	}
}
