package faker

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
	"github.com/omniql/tools/fieldgen"
	"github.com/omniql/reflect"
)

func (j *Json) fakeStruct(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer, structType reflect.StructContainer) (err error) {
	var shouldGenerateField bool
	var i int

	structObject := map[string]interface{}{}
	for i = 0; i < structType.LookupFields().FieldCount(); i++ {
		shouldGenerateField, err = fg.ShouldGenerateField(path, fieldType)
		if err != nil {
			err = &Error{
				Path:       path,
				HybridType: fieldType.HybridType(),
				OmniID:     fieldType.ID(),
				ErrorMsg:   err.Error(),
			}
			return err
		}
		if !shouldGenerateField {
			continue
		}

		childField, _ := structType.LookupFields().FieldByPosition(hybrids.FieldNumber(i))
		if childField.ValueType() != nil && childField.ValueType().Kind() == reflect.Enumeration {
			enumType := childField.ValueType().Enumeration()
			err = j.fakeEnum(fg, path+"."+childField.Name(), structObject, childField, enumType)
			if err != nil {
				return
			}
		} else {
			err = j.fakeScalar(fg, path+"."+childField.Name(), structObject, childField)
			if err != nil {
				return
			}
		}
	}

	out[fieldType.Name()] = structObject
	return
}

func (j *Json) fakeVectorStruct(fg fieldgen.Generator, path string, out map[string]interface{}, ft reflect.FieldContainer, structType reflect.StructContainer) (err error) {
	var vLen int
	var field int
	var shouldGenerateField bool

	//generate vector len
	vLen, err = fg.VectorLen(path, ft)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: ft.HybridType(),
			OmniID:     ft.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	r := make([]map[string]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		structObject := map[string]interface{}{}

		for field = 0; field < structType.LookupFields().FieldCount(); field++ {

			childFieldType, _ := structType.LookupFields().FieldByPosition(hybrids.FieldNumber(field))
			shouldGenerateField, err = fg.ShouldGenerateField(path, childFieldType)
			if err != nil {
				err = &Error{
					Path:       path + fmt.Sprintf("[%d].%s", i, childFieldType.Name()),
					HybridType: childFieldType.HybridType(),
					OmniID:     childFieldType.ID(),
					ErrorMsg:   err.Error(),
				}
				return err
			}
			if !shouldGenerateField {
				continue
			}
			err = j.fakeScalar(fg, path+fmt.Sprintf("[%d].%s", i, childFieldType.Name()), structObject, childFieldType)
			if err != nil {
				return
			}
		}
		r = append(r, structObject)
	}

	out[ft.Name()] = r

	return
}
