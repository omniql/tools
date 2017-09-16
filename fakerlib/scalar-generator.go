// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates scalar-decoder.gen.go It can be invoked by running
// go generate

package main

import "os"
import "text/template"
import "log"

import (
	"io"
	"github.com/nebtex/hybrids/golang/hybrids"
)

func RenderScalarFakers(f io.Writer) {
	var err error
	var tmp *template.Template

	tmp, err = template.New("ScalarFakers").Parse(`
func (j *Json) fake{{.scalar.String}}(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var v {{.scalar.NativeType}}

	v, err = j.fieldGen.{{.scalar.String}}(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	out[fieldType.Field().Name()] = v

	return
}
	
func (j *Json) fakeVector{{.scalar.String}}(path string, out map[string]interface{}, fieldType oreflection.OType) (err error) {
	var shouldNil bool
	var vLen int
	var v {{.scalar.NativeType}}

	shouldNil, err = j.fieldGen.ShouldBeNil(path, fieldType)

	if err != nil {
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
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
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    err.Error(),
		}
		return
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.fieldGen.{{.scalar.String}}(path, fieldType)
		if err != nil {
			err = &Error{
				Path:        path+fmt.Sprintf("[%d]", i),
				HybridType:  fieldType.Field().Items().HybridType(),
				OmniqlType:  fieldType.Id(),
				Package:     fieldType.Package(),
				ErrorMsg:    err.Error(),
			}
			return
	    }
		r = append(r, {{.jsTransform}})

	}

	out[fieldType.Field().Name()] = r

	return
}

`)
	die(err)

	scalars := []hybrids.Types{
		hybrids.Boolean,
		hybrids.Int8,
		hybrids.Uint8,
		hybrids.Int16,
		hybrids.Uint16,
		hybrids.Int32,
		hybrids.Uint32,
		hybrids.Int64,
		hybrids.Uint64,
		hybrids.Float32,
		hybrids.Float64,
	}
	jsonFunction := []string{
		"v",
		"float64(v)",
		"float64(v)",
		"float64(v)",
		"float64(v)",
		"float64(v)",
		"float64(v)",
		"strconv.FormatInt(int64(v), 10)",
		"strconv.FormatUint(uint64(v), 10)",
		"float64(v)",
		"float64(v)",
	}

	jsonType := []string{
		"bool",
		"float64",
		"float64",
		"float64",
		"float64",
		"float64",
		"float64",
		"string",
		"string",
		"float64",
		"float64",
	}

	min := []string{
		"hybrids.MinBoolean",
		"hybrids.MinInt8",
		"hybrids.MinUint8",
		"hybrids.MinInt16",
		"hybrids.MinUint16",
		"hybrids.MinInt32",
		"hybrids.MinUint32",
		"-200000000",
		"0",
		"-158",
		"-3500",
	}

	max := []string{
		"hybrids.MaxBoolean",
		"hybrids.MaxInt8",
		"hybrids.MaxUint8",
		"hybrids.MaxInt16",
		"hybrids.MaxUint16",
		"hybrids.MaxInt32",
		"hybrids.MaxUint32",
		"-200000000",
		"0",
		"-158",
		"-3500",
	}

	for index, v := range scalars {
		err = tmp.Execute(f, map[string]interface{}{
			"scalar":      v,
			"jsTransform": jsonFunction[index],
			"jsonType":    jsonType[index],
			"max":         max[index],
			"min":         min[index],
		})
		die(err)
	}

	decodeScalarTemplate, err := template.New("decodeScalarTemplate").Parse(`

func (j *Json) fakeScalar(path string, out map[string]interface{}, fieldType oreflection.OType)(err error){

	switch fieldType.Field().HybridType(){
{{ range $index, $value := .scalars }}
	case hybrids.{{$value.String}}:
		err = j.fake{{$value.String}}(path, out, fieldType)
{{ end }}
	default:
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package:     fieldType.Package(),
			ErrorMsg:    "path not recognized as  scalar",
		}
    }
	return
}


func (j *Json) fakeVectorScalar(path string, out map[string]interface{}, fieldType oreflection.OType)(err error){

	switch fieldType.Field().HybridType(){
{{ range $index, $value := .scalars }}
	case hybrids.Vector{{$value.String}}:
		err = j.fakeVector{{$value.String}}(path, out, fieldType)
{{ end }}
	default:
		err = &Error{
			Path:        path,
			HybridType:  fieldType.Field().HybridType(),
			OmniqlType:  fieldType.Id(),
			Package: 	 fieldType.Package(),
			ErrorMsg:    "path not recognized as vector of scalar",
		}
	}
	return
}


`)
	die(err)

	err = decodeScalarTemplate.Execute(f, map[string]interface{}{
		"scalars": scalars,
	})
	die(err)

}

func RenderTestScalarFaker(f io.Writer) {
	var err error
	var tmp *template.Template

	tmp, err = template.New("ScalarFakersTest").Parse(`

func Test_Fake{{.scalar.String}}(t *testing.T) {
	Convey("Test_Fake{{.scalar.String}}", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}

			otype:= &rmocks.OType{}
			fieldMock := &rmocks.Field{}
			otype.On("Field").Return(fieldMock)
			fieldMock.On("Name").Return("field")

			fg.On("{{.scalar.String}}", "test.field", otype).Return({{.value}}, nil)
			out := map[string]interface{}{}

			err := f.fake{{.scalar.String}}("test.field", out, otype)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, {{.jsonValue}})
			fg.AssertCalled(t, "{{.scalar.String}}", "test.field", otype)

		})

		Convey("Should return error if the generator returns an error", func() {

			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}

			otype:= &rmocks.OType{}
	        fieldMock := &rmocks.Field{}
			otype.On("Field").Return(fieldMock)
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.{{.scalar.String}})

			otype.On("Id").Return("Table/Test")
			otype.On("Package").Return("io.test.app")

			fg.On("{{.scalar.String}}", "test.field", otype).Return({{.value}}, fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fake{{.scalar.String}}("test.field", out, otype)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.Package, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.{{.scalar.String}})
			So(ef.OmniqlType, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVector{{.scalar.String}}(t *testing.T) {
	Convey("Test_FakeVector{{.scalar.String}}", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fieldMock := &rmocks.Field{}
			otype.On("Field").Return(fieldMock)
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Vector{{.scalar.String}})

			otype.On("Id").Return("Struct/Test")
			otype.On("Package").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVector{{.scalar.String}}("test.field", out, otype)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype)
			ef := err.(*Error)
			So(ef.Package, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Vector{{.scalar.String}})
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fieldMock := &rmocks.Field{}
			otype.On("Field").Return(fieldMock)
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Vector{{.scalar.String}})

			otype.On("Id").Return("Struct/Test")
			otype.On("Package").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVector{{.scalar.String}}("test.field", out, otype)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fieldMock := &rmocks.Field{}
			otype.On("Field").Return(fieldMock)
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Vector{{.scalar.String}})

			otype.On("Id").Return("Struct/Test")
			otype.On("Package").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype).Return(false, nil)
			fg.On("VectorLen", "test.field", otype).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVector{{.scalar.String}}("test.field", out, otype)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype)
			fg.AssertCalled(t, "VectorLen", "test.field", otype)
			So(ef.Package, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.Vector{{.scalar.String}})
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fieldMock := &rmocks.Field{}
			otype.On("Field").Return(fieldMock)
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Vector{{.scalar.String}})

			itemsMock:= &rmocks.Items{}
			fieldMock.On("Items").Return(itemsMock)
			itemsMock.On("HybridType").Return(hybrids.{{.scalar.String}})

			otype.On("Id").Return("Struct/Test")
			otype.On("Package").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype).Return(false, nil)
			fg.On("VectorLen", "test.field", otype).Return(10, nil)
			fg.On("{{.scalar.String}}", "test.field", otype).Return({{.value}}, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVector{{.scalar.String}}("test.field", out, otype)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", otype)
			fg.AssertCalled(t, "VectorLen", "test.field", otype)
			So(ef.Package, ShouldEqual, "io.test.app")
			So(ef.HybridType, ShouldEqual, hybrids.{{.scalar.String}})
			So(ef.OmniqlType, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fg := &fmocks.FieldGenerator{}
			f := &Json{fieldGen: fg}
			otype := &rmocks.OType{}

			fieldMock := &rmocks.Field{}
			otype.On("Field").Return(fieldMock)
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Vector{{.scalar.String}})

			otype.On("Id").Return("Struct/Test")
			otype.On("Package").Return("io.test.app")
			fg.On("ShouldBeNil", "test.field", otype).Return(false, nil)
			fg.On("VectorLen", "test.field", otype).Return(10, nil)
			fg.On("{{.scalar.String}}", "test.field", otype).Return({{.value}}, nil)

			out := map[string]interface{}{}

			err := f.fakeVector{{.scalar.String}}("test.field", out, otype)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, {{.jsonValue}})
		})
	})
}

`)
	die(err)
	scalars := []hybrids.Types{
		hybrids.Boolean,
		hybrids.Int8,
		hybrids.Uint8,
		hybrids.Int16,
		hybrids.Uint16,
		hybrids.Int32,
		hybrids.Uint32,
		hybrids.Int64,
		hybrids.Uint64,
		hybrids.Float32,
		hybrids.Float64,
	}
	testValue := []string{
		"true",
		"int8(-60)",
		"uint8(87)",
		"int16(-500)",
		"uint16(458)",
		"int32(-60000)",
		"uint32(23568778)",
		"int64(-54545788)",
		"uint64(123587)",
		"float32(25.50)",
		"float64(-356.4545)",
	}

	jsonValue := []string{
		"true",
		"float64(-60)",
		"float64(87)",
		"float64(-500)",
		"float64(458)",
		"float64(-60000)",
		"float64(23568778)",
		"\"-54545788\"",
		"\"123587\"",
		"float64(25.50)",
		"float64(-356.4545)",
	}

	for index, v := range scalars {
		err = tmp.Execute(f, map[string]interface{}{
			"scalar":    v,
			"value":     testValue[index],
			"jsonValue": jsonValue[index],
		})
		die(err)
	}

}

func main() {
	f, err := os.Create("scalar-faker.gen.go")
	die(err)
	defer f.Close()
	_, err = f.Write([]byte("// Code generated, DO NOT EDIT.\n"))
	die(err)
	_, err = f.Write([]byte(`package faker
`))
	die(err)
	_, err = f.Write([]byte(`
import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"strconv"
	"github.com/nebtex/omniql/commons/golang/oreflection"
	"fmt"
)
`))
	die(err)
	RenderScalarFakers(f)

	test, err := os.Create("scalar-faker_test.go")
	die(err)
	defer test.Close()
	_, err = test.Write([]byte("// Code generated, DO NOT EDIT.\n"))
	die(err)
	_, err = test.Write([]byte(`package faker
`))
	die(err)

	_, err = test.Write([]byte(`
import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	fmocks "github.com/nebtex/omniql/tools/faker/fieldgen/mocks"
	rmocks "github.com/nebtex/omniql/commons/golang/oreflection/mocks"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
)
`))
	die(err)
	RenderTestScalarFaker(test)

	/*RenderVectorScalarDecoders(f)

	test, err := os.Create("scalar-decoder_test.go")
	die(err)
	defer test.Close()
	_, err = test.Write([]byte("// Code generated. DO NOT EDIT.\n"))
	die(err)
	_, err = test.Write([]byte(`package omarshaller
`))
	die(err)
	_, err = test.Write([]byte(`
import (
	"testing"
	"github.com/nebtex/hybrids/golang/hybrids/mocks"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
)
`))
	die(err)
	RenderTestScalarDecoders(test)
	RenderVectorScalar_TEST_Decoders(test)*/
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
