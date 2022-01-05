package runtime

import (
	"bytes"
	"reflect"
	"text/template"
)

func RenderTemplate(templateStruct interface{}, data *map[string]interface{}) {
	v := reflect.ValueOf(templateStruct).Elem()

	renderTemplateReflectValue(&v, data)
}

func renderTemplateReflectValue(v *reflect.Value, data *map[string]interface{}) {
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldType := typeOfS.Field(i).Type
		field := v.Field(i)

		if fieldType.Kind() == reflect.String {
			field.SetString(renderTemplateString(field.String(), data))
		} else if fieldType.Kind() == reflect.Struct {
			renderTemplateReflectValue(&field, data)
		} else if fieldType.Kind() == reflect.Slice {
			for i := 0; i < field.Len(); i++ {
				sliceField := field.Index(i)

				renderTemplateReflectValue(&sliceField, data)
			}
		} else if fieldType.Kind() == reflect.Map {
			// Go through all the fields of the map and render the templates if they're string values
			iter := field.MapRange()
			for iter.Next() {
				key := iter.Key()
				value := iter.Value()

				if value.Kind() == reflect.String {
					field.SetMapIndex(key, reflect.ValueOf(renderTemplateString(value.String(), data)))
				}
			}
		}
	}
}

func renderTemplateString(templateString string, data *map[string]interface{}) string {
	var renderedTemplate bytes.Buffer

	t := template.Must(template.New("").Parse(templateString))
	t.Execute(&renderedTemplate, data)
	return renderedTemplate.String()
}
