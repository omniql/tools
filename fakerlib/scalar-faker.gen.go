// Code generated, DO NOT EDIT.
package faker

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"strconv"
	"github.com/omniql/reflect"
	"fmt"
	"github.com/omniql/tools/fieldgen"
)

func (j *Json) fakeBoolean(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var v bool

	v, err = fg.Boolean(path, fieldType)

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
	
func (j *Json) fakeVectorBoolean(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var vLen int
	var v bool

	//generate vector len
	vLen, err = fg.VectorLen(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.HybridType(),
			OmniID:      fieldType.ID(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  fg.Boolean(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Items().HybridType(),
				OmniID:  fieldType.ID(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, v)

	}

	out[fieldType.Name()] = r

	return
}


func (j *Json) fakeInt8(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var v int8

	v, err = fg.Int8(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	out[fieldType.Name()] = float64(v)

	return
}
	
func (j *Json) fakeVectorInt8(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var vLen int
	var v int8

	//generate vector len
	vLen, err = fg.VectorLen(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.HybridType(),
			OmniID:      fieldType.ID(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  fg.Int8(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Items().HybridType(),
				OmniID:  fieldType.ID(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Name()] = r

	return
}


func (j *Json) fakeUint8(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var v uint8

	v, err = fg.Uint8(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	out[fieldType.Name()] = float64(v)

	return
}
	
func (j *Json) fakeVectorUint8(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var vLen int
	var v uint8

	//generate vector len
	vLen, err = fg.VectorLen(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.HybridType(),
			OmniID:      fieldType.ID(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  fg.Uint8(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Items().HybridType(),
				OmniID:  fieldType.ID(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Name()] = r

	return
}


func (j *Json) fakeInt16(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var v int16

	v, err = fg.Int16(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	out[fieldType.Name()] = float64(v)

	return
}
	
func (j *Json) fakeVectorInt16(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var vLen int
	var v int16

	//generate vector len
	vLen, err = fg.VectorLen(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.HybridType(),
			OmniID:      fieldType.ID(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  fg.Int16(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Items().HybridType(),
				OmniID:  fieldType.ID(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Name()] = r

	return
}


func (j *Json) fakeUint16(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var v uint16

	v, err = fg.Uint16(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	out[fieldType.Name()] = float64(v)

	return
}
	
func (j *Json) fakeVectorUint16(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var vLen int
	var v uint16

	//generate vector len
	vLen, err = fg.VectorLen(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.HybridType(),
			OmniID:      fieldType.ID(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  fg.Uint16(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Items().HybridType(),
				OmniID:  fieldType.ID(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Name()] = r

	return
}


func (j *Json) fakeInt32(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var v int32

	v, err = fg.Int32(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	out[fieldType.Name()] = float64(v)

	return
}
	
func (j *Json) fakeVectorInt32(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var vLen int
	var v int32

	//generate vector len
	vLen, err = fg.VectorLen(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.HybridType(),
			OmniID:      fieldType.ID(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  fg.Int32(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Items().HybridType(),
				OmniID:  fieldType.ID(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Name()] = r

	return
}


func (j *Json) fakeUint32(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var v uint32

	v, err = fg.Uint32(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	out[fieldType.Name()] = float64(v)

	return
}
	
func (j *Json) fakeVectorUint32(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var vLen int
	var v uint32

	//generate vector len
	vLen, err = fg.VectorLen(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.HybridType(),
			OmniID:      fieldType.ID(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  fg.Uint32(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Items().HybridType(),
				OmniID:  fieldType.ID(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Name()] = r

	return
}


func (j *Json) fakeInt64(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var v int64

	v, err = fg.Int64(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	out[fieldType.Name()] = strconv.FormatInt(int64(v), 10)

	return
}
	
func (j *Json) fakeVectorInt64(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var vLen int
	var v int64

	//generate vector len
	vLen, err = fg.VectorLen(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.HybridType(),
			OmniID:      fieldType.ID(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  fg.Int64(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Items().HybridType(),
				OmniID:  fieldType.ID(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, strconv.FormatInt(int64(v), 10))

	}

	out[fieldType.Name()] = r

	return
}


func (j *Json) fakeUint64(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var v uint64

	v, err = fg.Uint64(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	out[fieldType.Name()] = strconv.FormatUint(uint64(v), 10)

	return
}
	
func (j *Json) fakeVectorUint64(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var vLen int
	var v uint64

	//generate vector len
	vLen, err = fg.VectorLen(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.HybridType(),
			OmniID:      fieldType.ID(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  fg.Uint64(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Items().HybridType(),
				OmniID:  fieldType.ID(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, strconv.FormatUint(uint64(v), 10))

	}

	out[fieldType.Name()] = r

	return
}


func (j *Json) fakeFloat32(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var v float32

	v, err = fg.Float32(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	out[fieldType.Name()] = float64(v)

	return
}
	
func (j *Json) fakeVectorFloat32(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var vLen int
	var v float32

	//generate vector len
	vLen, err = fg.VectorLen(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.HybridType(),
			OmniID:      fieldType.ID(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  fg.Float32(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Items().HybridType(),
				OmniID:  fieldType.ID(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Name()] = r

	return
}


func (j *Json) fakeFloat64(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var v float64

	v, err = fg.Float64(path, fieldType)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: fieldType.HybridType(),
			OmniID:     fieldType.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	out[fieldType.Name()] = float64(v)

	return
}
	
func (j *Json) fakeVectorFloat64(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer) (err error) {
	var vLen int
	var v float64

	//generate vector len
	vLen, err = fg.VectorLen(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.HybridType(),
			OmniID:      fieldType.ID(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  fg.Float64(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Items().HybridType(),
				OmniID:  fieldType.ID(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, float64(v))

	}

	out[fieldType.Name()] = r

	return
}



func (j *Json) fakeScalar(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer)(err error){

	switch fieldType.HybridType(){

	case hybrids.Boolean:
		err = j.fakeBoolean(fg, path, out, fieldType)

	case hybrids.Int8:
		err = j.fakeInt8(fg, path, out, fieldType)

	case hybrids.Uint8:
		err = j.fakeUint8(fg, path, out, fieldType)

	case hybrids.Int16:
		err = j.fakeInt16(fg, path, out, fieldType)

	case hybrids.Uint16:
		err = j.fakeUint16(fg, path, out, fieldType)

	case hybrids.Int32:
		err = j.fakeInt32(fg, path, out, fieldType)

	case hybrids.Uint32:
		err = j.fakeUint32(fg, path, out, fieldType)

	case hybrids.Int64:
		err = j.fakeInt64(fg, path, out, fieldType)

	case hybrids.Uint64:
		err = j.fakeUint64(fg, path, out, fieldType)

	case hybrids.Float32:
		err = j.fakeFloat32(fg, path, out, fieldType)

	case hybrids.Float64:
		err = j.fakeFloat64(fg, path, out, fieldType)

	default:
		err = &Error{
			Path:        path,
			HybridType:  fieldType.HybridType(),
			OmniID:      fieldType.ID(),
			ErrorMsg:    "path not recognized as  scalar",
		}
    }
	return
}


func (j *Json) fakeVectorScalar(fg fieldgen.Generator, path string, out map[string]interface{}, fieldType reflect.FieldContainer)(err error){

	switch fieldType.HybridType(){

	case hybrids.VectorBoolean:
		err = j.fakeVectorBoolean(fg, path, out, fieldType)

	case hybrids.VectorInt8:
		err = j.fakeVectorInt8(fg, path, out, fieldType)

	case hybrids.VectorUint8:
		err = j.fakeVectorUint8(fg, path, out, fieldType)

	case hybrids.VectorInt16:
		err = j.fakeVectorInt16(fg, path, out, fieldType)

	case hybrids.VectorUint16:
		err = j.fakeVectorUint16(fg, path, out, fieldType)

	case hybrids.VectorInt32:
		err = j.fakeVectorInt32(fg, path, out, fieldType)

	case hybrids.VectorUint32:
		err = j.fakeVectorUint32(fg, path, out, fieldType)

	case hybrids.VectorInt64:
		err = j.fakeVectorInt64(fg, path, out, fieldType)

	case hybrids.VectorUint64:
		err = j.fakeVectorUint64(fg, path, out, fieldType)

	case hybrids.VectorFloat32:
		err = j.fakeVectorFloat32(fg, path, out, fieldType)

	case hybrids.VectorFloat64:
		err = j.fakeVectorFloat64(fg, path, out, fieldType)

	default:
		err = &Error{
			Path:        path,
			HybridType:  fieldType.HybridType(),
			OmniID:      fieldType.ID(),
			ErrorMsg:    "path not recognized as vector of scalar",
		}
	}
	return
}


