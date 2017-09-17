package faker

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/omniql/tools/fieldgen"
	"github.com/omniql/reflect"
)

func (j *Json) fakeString(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var v string

	v, err = fg.String(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	out[fieldType.Name()] = v

	return
}

func (j *Json) fakeByte(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var v []byte

	v, err = fg.Byte(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	out[fieldType.Name()] = b64.StdEncoding.EncodeToString(v)

	return
}

func (j *Json) fakeTheResourceID(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer, resource reflect.ResourceContainer) (err error) {
	var v []byte

	v, err = fg.ResourceID(path, fieldType, resource.Application().ResourceIDType())

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	out[fieldType.Name()] = b64.StdEncoding.EncodeToString(v)

	return
}

func (j *Json) fakeVectorString(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var vLen int
	var v string

	//generate vector len
	vLen, err = fg.VectorLen(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err = fg.String(path, fieldType)
		if err != nil {
			err = &Error{
				Path:       path + fmt.Sprintf("[%d]", i),
				HybridType: fieldType.Items().HybridType(),
				OmniID:     fieldType.ID(),
				ErrorMsg:   err.Error(),
			}
			return
		}
		r = append(r, v)

	}

	out[fieldType.Name()] = r

	return
}

func (j *Json) fakeVectorByte(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var vLen int
	var v []byte

	//generate vector len
	vLen, err = fg.VectorLen(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err = fg.Byte(path, fieldType)
		if err != nil {
			err = &Error{
				Path:       path + fmt.Sprintf("[%d]", i),
				HybridType: fieldType.Items().HybridType(),
				OmniID:     fieldType.ID(),
				ErrorMsg:   err.Error(),
			}
			return
		}
		r = append(r, b64.StdEncoding.EncodeToString(v))

	}

	out[fieldType.Name()] = r

	return
}

func (j *Json) fakeVectorResourceID(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer, resource reflect.ResourceContainer) (err error) {
	var vLen int
	var v []byte

	//generate vector len
	vLen, err = fg.VectorLen(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err = fg.ResourceID(path, fieldType, resource.Application().ResourceIDType())
		if err != nil {
			err = &Error{
				Path:       path + fmt.Sprintf("[%d]", i),
				HybridType: fieldType.Items().HybridType(),
				OmniID:     fieldType.ID(),
				ErrorMsg:   err.Error(),
			}
			return
		}
		r = append(r, b64.StdEncoding.EncodeToString(v))

	}

	out[fieldType.Name()] = r

	return
}
