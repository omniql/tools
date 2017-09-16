package fieldgen

import (
	"github.com/omniql/reflect"
	"github.com/nebtex/hybrids/golang/hybrids"
	"github.com/nebtex/omniql/commons/golang/oreflection"
)

//go:generate mockery -name=EnumerationGenerator
//EnumerationGenerator ...
type EnumerationGenerator interface {
	StringEnumeration(path string, fieldType reflect.FieldContainer) (string, error)
	Uint8Enumeration(path string, fieldType reflect.FieldContainer) (uint8, error)
	Uint16Enumeration(path string, fieldType reflect.FieldContainer) (uint16, error)
	ShouldGenerateString(path string, fieldType reflect.FieldContainer) (bool, error)
}

//go:generate mockery -name=EnumerationGenerator
//UnionGenerator ...
type UnionGenerator interface {
	SelectTable(path string, fieldType oreflection.OType) (childFieldType oreflection.OType, err error)
}

//go:generate mockery -name=FieldGenerator
//FieldGenerator generate random data using the reflection interface
type Generator interface {
	//generate random vector len
	VectorLen(path string, ot reflect.FieldContainer, maxlen int) (int, error)

	//return if a field should be nil or not (don't use this on scalar or structs)
	ShouldBeNil(path string, ot reflect.FieldContainer) (bool, error)

	//return a random Boolean tha follow the spec of th field
	Boolean(path string, field reflect.FieldContainer) (bool, error)

	//return a random Int8 tha follow the spec of th field
	Int8(path string, field reflect.FieldContainer) (int8, error)

	//return a random Uint8 tha follow the spec of th field
	Uint8(path string, field reflect.FieldContainer) (uint8, error)

	//return a random Int16 tha follow the spec of th field
	Int16(path string, field reflect.FieldContainer) (int16, error)

	//return a random Uint16 tha follow the spec of th field
	Uint16(path string, field reflect.FieldContainer) (uint16, error)

	//return a random Int32 tha follow the spec of th field
	Int32(path string, field reflect.FieldContainer) (int32, error)

	//return a random Uint32 tha follow the spec of th field
	Uint32(path string, field reflect.FieldContainer) (uint32, error)

	//return a random Int64 tha follow the spec of th field
	Int64(path string, field reflect.FieldContainer) (int64, error)

	//return a random Uint64 tha follow the spec of th field
	Uint64(path string, field reflect.FieldContainer) (uint64, error)

	//return a random Float32 tha follow the spec of th field
	Float32(path string, field reflect.FieldContainer) (float32, error)

	//return a random Float64 tha follow the spec of th field
	Float64(path string, field reflect.FieldContainer) (float64, error)

	//return a random String tha follow the spec of th field
	String(path string, field reflect.FieldContainer) (string, error)

	//return a random String tha follow the spec of th field
	Byte(path string, field reflect.FieldContainer) ([]byte, error)

	//return a random String tha follow the spec of th field
	ResourceID(path string, field reflect.FieldContainer, resourceIdType hybrids.ResourceIDType) ([]byte, error)

	//return a random Enumeration,
	Enumeration() EnumerationGenerator

	//Decide if a field should be generated or not
	ShouldGenerateField(path string, table oreflection.OType, fn hybrids.FieldNumber) (bool, error)

	//Union random generator helpers
	Union() UnionGenerator
}
