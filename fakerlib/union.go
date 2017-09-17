package faker

import (
	"github.com/omniql/reflect"
	"strings"
	"fmt"
	"github.com/omniql/tools/fieldgen"
)

func (j *Json) fakeUnion(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer, union reflect.UnionContainer) (err error) {
	var childFieldType reflect.FieldContainer

	deep := len(strings.Split(path, "."))

	if deep > j.maxDeep {
		out[fieldType.Name()] = nil
		return
	}

	unionObject := map[string]interface{}{}

	childFieldType, err = fg.Union.SelectTable(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	table := childFieldType.Field().Type().Table()
	err = j.fakeTable(path+"."+childFieldType.Field().Name(), unionObject, childFieldType, table)
	if err != nil {
		return
	}

	out[fieldType.Field().Name()] = unionObject
	return
}

func (j *Json) fakeVectorUnion(path string, out map[string]interface{}, fieldType oreflection.OType, union oreflection.Union) (err error) {
	var shouldNil bool
	var vLen int
	var childFieldType oreflection.OType

	deep := len(strings.Split(path, "."))

	if deep > j.maxDeep {
		out[fieldType.Field().Name()] = nil
		return
	}

	shouldNil, err = j.fieldGen.ShouldBeNil(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.Field().HybridType(),
			OmniqlType: fieldType.Id(),
			Package:    fieldType.Package(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	if shouldNil {
		out[fieldType.Field().Name()] = nil
		return
	}

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.Field().HybridType(),
			OmniqlType: fieldType.Id(),
			Package:    fieldType.Package(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	r := make([]map[string]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		unionObject := map[string]interface{}{}

		childFieldType, err = j.fieldGen.Union.PickTable(path, fieldType)

		if err != nil {
			err = &Error{
				Path:       path + fmt.Sprintf("[%d].%s", i, childFieldType.Field().Name()),
				HybridType: fieldType.Field().HybridType(),
				OmniqlType: fieldType.Id(),
				Package:    fieldType.Package(),
				ErrorMsg:   err.Error(),
			}
			return
		}

		table := childFieldType.Field().Type().Table()

		err = j.fakeTable(path+fmt.Sprintf("[%d].%s", i, childFieldType.Field().Name()), unionObject, childFieldType, table)
		if err != nil {
			return
		}

		r = append(r, unionObject)
	}

	out[fieldType.Field().Name()] = r

	return
}
