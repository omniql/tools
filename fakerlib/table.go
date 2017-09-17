package faker

import (
	"github.com/omniql/reflect"
	"github.com/nebtex/hybrids/golang/hybrids"
	"strings"
	"fmt"
	"github.com/omniql/tools/fieldgen"
)

func (j *Json) fakeTable(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer, tableType reflect.TableContainer) (err error) {
	var shouldNil bool
	var i int
	var shouldGenerateField bool

	deep := len(strings.Split(path, "."))

	if deep > j.maxDeep {
		out[fieldType.Name()] = nil
		return
	}

	tableObject := map[string]interface{}{}

	for i = 0; i < tableType.LookupFields().FieldCount(); i++ {
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

		childFieldType, _ := tableType.LookupFields().FieldByPosition(hybrids.FieldNumber(i))

		if childFieldType.ValueType().Kind() == reflect.Enumeration {
			enumType := childFieldType.ValueType().Enumeration()
			err = j.fakeEnum(fg, path+"."+childFieldType.Name(), tableObject, childFieldType, enumType)
			if err != nil {
				return err
			}
			continue
		}

		if childFieldType.HybridType().IsScalar() {

			err = j.fakeScalar(fg, path+"."+childFieldType.Name(), tableObject, childFieldType)
			if err != nil {
				return err
			}
			continue
		}

		if childFieldType.HybridType() == hybrids.Struct {
			structType := childFieldType.ValueType().Struct()
			err = j.fakeStruct(fg, path+"."+childFieldType.Name(), tableObject, childFieldType, structType)
			if err != nil {
				return err
			}
			continue
		}

		shouldNil, err = fg.ShouldBeNil(path, childFieldType)

		if err != nil {
			err = &Error{
				Path:       path + "." + childFieldType.Name(),
				HybridType: childFieldType.HybridType(),
				OmniID:     childFieldType.ID(),
				ErrorMsg:   err.Error(),
			}
			return
		}

		if shouldNil {
			out[fieldType.Name()] = nil
			return
		}

		if childFieldType.HybridType().IsVectorScalar() {

			err = j.fakeVectorScalar(fg, path+"."+childFieldType.Name(), tableObject, childFieldType)
			if err != nil {
				return err
			}
		}

		switch childFieldType.HybridType() {
		case hybrids.VectorStruct:
			structType := childFieldType.Items().ValueType().Struct()
			err = j.fakeStruct(fg, path+"."+childFieldType.Name(), tableObject, childFieldType, structType)
		case hybrids.String:
			err = j.fakeString(fg, path+"."+childFieldType.Name(), tableObject, childFieldType)
		case hybrids.Byte:
			err = j.fakeByte(fg, path+"."+childFieldType.Name(), tableObject, childFieldType)
		case hybrids.ResourceID:
			resource := childFieldType.ValueType().Resource()
			err = j.fakeTheResourceID(fg, path+"."+childFieldType.Name(), tableObject, childFieldType, resource)
		case hybrids.VectorString:
			err = j.fakeVectorString(fg, path+"."+childFieldType.Name(), tableObject, childFieldType)
		case hybrids.VectorByte:
			err = j.fakeVectorByte(fg, path+"."+childFieldType.Name(), tableObject, childFieldType)
		case hybrids.VectorResourceID:
			resource := childFieldType.Items().ValueType().Resource()
			err = j.fakeVectorResourceID(fg, path+"."+childFieldType.Name(), tableObject, childFieldType, resource)
		case hybrids.Table:
			table := childFieldType.ValueType().Table()
			err = j.fakeTable(fg, path+"."+childFieldType.Name(), tableObject, childFieldType, table)
		case hybrids.Union:
			union := childFieldType.ValueType().Union()
			err = j.fakeUnion(fg, path+"."+childFieldType.Name(), tableObject, childFieldType, union)
		case hybrids.VectorUnion:
			union := childFieldType.Items().ValueType().Union()
			err = j.fakeVectorUnion(fg, path+"."+childFieldType.Name(), tableObject, childFieldType, union)
		case hybrids.VectorTable:
			table := childFieldType.Items().ValueType().Table()
			err = j.fakeVectorTable(fg, path+"."+childFieldType.Name(), tableObject, childFieldType, table)
		}
		if err != nil {
			return err
		}
	}

	out[fieldType.Name()] = tableObject
	return
}

func (j *Json) fakeVectorTable(fg fieldgen.Generator, path string, out map[string]interface{}, ft reflect.FieldContainer, tableType reflect.TableContainer) (err error) {
	var vLen int

	deep := len(strings.Split(path, "."))

	if deep > j.maxDeep {
		out[ft.Name()] = nil
		return
	}

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

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		tableObject := map[string]interface{}{}
		err = j.fakeTable(fg, path+fmt.Sprintf("[%d]", i), tableObject, ft, tableType)
		if err != nil {
			return
		}
		r = append(r, tableObject[ft.Name()])
	}

	out[ft.Name()] = r

	return
}
