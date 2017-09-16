package faker

import (
	"github.com/nebtex/omniql/commons/golang/oreflection"
	"github.com/nebtex/hybrids/golang/hybrids"
	"strings"
	"fmt"
)

func (j *Json) fakeTable(path string, out map[string]interface{}, fieldType oreflection.OType, tableType oreflection.Table) (err error) {
	var shouldNil bool
	var i hybrids.FieldNumber

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

	tableObject := map[string]interface{}{}

	for i = 0; i < tableType.FieldCount(); i++ {
		shouldGenerateField, err := j.fieldGen.ShouldGenerateField(path, fieldType, i)
		if err != nil {
			err = &Error{
				Path:       path,
				HybridType: fieldType.Field().HybridType(),
				OmniqlType: fieldType.Id(),
				Package:    fieldType.Package(),
				ErrorMsg:   err.Error(),
			}
			return err
		}
		if !shouldGenerateField {
			continue
		}

		childFieldName, childFieldType, _ := tableType.LookupFields().ByNumber(i)

		if childFieldType.Field().IsEnumeration() {
			enumType := childFieldType.Field().Type().Enumeration()
			err = j.fakeEnum(path+"."+childFieldName, tableObject, childFieldType, enumType)
			if err != nil {
				return err
			}
		}
		if childFieldType.Field().HybridType().IsScalar() {

			err = j.fakeScalar(path+"."+childFieldName, tableObject, childFieldType)
			if err != nil {
				return err
			}
		}
		if childFieldType.Field().HybridType().IsVectorScalar() {

			err = j.fakeVectorScalar(path+"."+childFieldName, tableObject, childFieldType)
			if err != nil {
				return err
			}
		}

		switch childFieldType.Field().HybridType() {
		case hybrids.Struct:
			structType := childFieldType.Field().Type().Struct()
			err = j.fakeStruct(path+"."+childFieldName, tableObject, childFieldType, structType)
		case hybrids.VectorStruct:
			structType := childFieldType.Field().Items().Type().Struct()
			err = j.fakeStruct(path+"."+childFieldName, tableObject, childFieldType, structType)
		case hybrids.String:
			err = j.fakeString(path+"."+childFieldName, tableObject, childFieldType)
		case hybrids.Byte:
			err = j.fakeByte(path+"."+childFieldName, tableObject, childFieldType)
		case hybrids.ResourceID:
			resource := childFieldType.Field().Type().Resource()
			err = j.fakeTheResourceID(path+"."+childFieldName, tableObject, childFieldType, resource)
		case hybrids.VectorString:
			err = j.fakeVectorString(path+"."+childFieldName, tableObject, childFieldType)
		case hybrids.VectorByte:
			err = j.fakeVectorByte(path+"."+childFieldName, tableObject, childFieldType)
		case hybrids.VectorResourceID:
			resource := childFieldType.Field().Items().Type().Resource()
			err = j.fakeVectorResourceID(path+"."+childFieldName, tableObject, childFieldType, resource)
		case hybrids.Table:
			table := childFieldType.Field().Type().Table()
			err = j.fakeTable(path+"."+childFieldName, tableObject, childFieldType, table)
		case hybrids.Union:
			union := childFieldType.Field().Type().Union()
			err = j.fakeUnion(path+"."+childFieldName, tableObject, childFieldType, union)
		case hybrids.VectorUnion:
			union := childFieldType.Field().Items().Type().Union()
			err = j.fakeVectorUnion(path+"."+childFieldName, tableObject, childFieldType, union)
		case hybrids.VectorTable:
			table := childFieldType.Field().Items().Type().Table()
			err = j.fakeVectorTable(path+"."+childFieldName, tableObject, childFieldType, table)
		}
		if err != nil {
			return err
		}
	}

	out[fieldType.Field().Name()] = tableObject
	return
}

func (j *Json) fakeVectorTable(path string, out map[string]interface{}, ft oreflection.OType, tableType oreflection.Table) (err error) {
	var shouldNil bool
	var vLen int
	var field hybrids.FieldNumber

	deep := len(strings.Split(path, "."))

	if deep > j.maxDeep {
		out[ft.Field().Name()] = nil
		return
	}

	shouldNil, err = j.fieldGen.ShouldBeNil(path, ft)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: ft.Field().HybridType(),
			OmniqlType: ft.Id(),
			Package:    ft.Package(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	if shouldNil {
		out[ft.Field().Name()] = nil
		return
	}

	//generate vector len
	vLen, err = j.fieldGen.VectorLen(path, ft)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: ft.Field().HybridType(),
			OmniqlType: ft.Id(),
			Package:    ft.Package(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	r := make([]map[string]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		tableObject := map[string]interface{}{}

		for field = 0; field < tableType.FieldCount(); field++ {
			childFieldName, childFieldType, _ := tableType.LookupFields().ByNumber(field)
			table := childFieldType.Field().Items().Type().Table()
			err = j.fakeTable(path+fmt.Sprintf("[%d].%s", i, childFieldName), tableObject, childFieldType, table)
			if err != nil {
				return
			}
		}
		r = append(r, tableObject)
	}

	out[ft.Field().Name()] = r

	return
}
