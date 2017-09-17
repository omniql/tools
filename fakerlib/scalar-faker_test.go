// Code generated, DO NOT EDIT.
package faker

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	fmocks "github.com/omniql/tools/fieldgen/mocks"
	rmocks "github.com/omniql/reflect/mocks"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
)


func Test_FakeBoolean(t *testing.T) {
	Convey("Test_FakeBoolean", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")

			fg.On("Boolean", "test.field", fieldMock).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeBoolean(fg, "test.field", out, fieldMock)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, true)
			fg.AssertCalled(t, "Boolean", "test.field", fieldMock)

		})

		Convey("Should return error if the generator returns an error", func() {

			fg := &fmocks.Generator{}
			f := &Json{}

	        fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Boolean)

			fieldMock.On("ID").Return("Table/Test")

			fg.On("Boolean", "test.field", fieldMock).Return(true, fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeBoolean(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.Boolean)
			So(ef.OmniID, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorBoolean(t *testing.T) {
	Convey("Test_FakeVectorBoolean", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorBoolean)
			fieldMock.On("ID").Return("Struct/Test")


			fieldMock.On("Id").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorBoolean(fg, "test.field", out, fieldMock)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.VectorBoolean)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorBoolean)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorBoolean(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorBoolean)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorBoolean(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.VectorBoolean)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorBoolean)

			itemsMock:= &rmocks.ItemsContainer{}
			fieldMock.On("Items").Return(itemsMock)
			itemsMock.On("HybridType").Return(hybrids.Boolean)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Boolean", "test.field", fieldMock).Return(true, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorBoolean(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.Boolean)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorBoolean)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Boolean", "test.field", fieldMock).Return(true, nil)

			out := map[string]interface{}{}

			err := f.fakeVectorBoolean(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, true)
		})
	})
}



func Test_FakeInt8(t *testing.T) {
	Convey("Test_FakeInt8", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")

			fg.On("Int8", "test.field", fieldMock).Return(int8(-60), nil)
			out := map[string]interface{}{}

			err := f.fakeInt8(fg, "test.field", out, fieldMock)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(-60))
			fg.AssertCalled(t, "Int8", "test.field", fieldMock)

		})

		Convey("Should return error if the generator returns an error", func() {

			fg := &fmocks.Generator{}
			f := &Json{}

	        fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Int8)

			fieldMock.On("ID").Return("Table/Test")

			fg.On("Int8", "test.field", fieldMock).Return(int8(-60), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeInt8(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.Int8)
			So(ef.OmniID, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorInt8(t *testing.T) {
	Convey("Test_FakeVectorInt8", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt8)
			fieldMock.On("ID").Return("Struct/Test")


			fieldMock.On("Id").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorInt8(fg, "test.field", out, fieldMock)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt8)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt8)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorInt8(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt8)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt8(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt8)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt8)

			itemsMock:= &rmocks.ItemsContainer{}
			fieldMock.On("Items").Return(itemsMock)
			itemsMock.On("HybridType").Return(hybrids.Int8)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Int8", "test.field", fieldMock).Return(int8(-60), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt8(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.Int8)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt8)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Int8", "test.field", fieldMock).Return(int8(-60), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorInt8(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(-60))
		})
	})
}



func Test_FakeUint8(t *testing.T) {
	Convey("Test_FakeUint8", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")

			fg.On("Uint8", "test.field", fieldMock).Return(uint8(87), nil)
			out := map[string]interface{}{}

			err := f.fakeUint8(fg, "test.field", out, fieldMock)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(87))
			fg.AssertCalled(t, "Uint8", "test.field", fieldMock)

		})

		Convey("Should return error if the generator returns an error", func() {

			fg := &fmocks.Generator{}
			f := &Json{}

	        fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Uint8)

			fieldMock.On("ID").Return("Table/Test")

			fg.On("Uint8", "test.field", fieldMock).Return(uint8(87), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeUint8(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.Uint8)
			So(ef.OmniID, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorUint8(t *testing.T) {
	Convey("Test_FakeVectorUint8", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint8)
			fieldMock.On("ID").Return("Struct/Test")


			fieldMock.On("Id").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorUint8(fg, "test.field", out, fieldMock)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint8)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint8)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorUint8(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint8)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint8(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint8)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint8)

			itemsMock:= &rmocks.ItemsContainer{}
			fieldMock.On("Items").Return(itemsMock)
			itemsMock.On("HybridType").Return(hybrids.Uint8)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Uint8", "test.field", fieldMock).Return(uint8(87), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint8(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.Uint8)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint8)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Uint8", "test.field", fieldMock).Return(uint8(87), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorUint8(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(87))
		})
	})
}



func Test_FakeInt16(t *testing.T) {
	Convey("Test_FakeInt16", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")

			fg.On("Int16", "test.field", fieldMock).Return(int16(-500), nil)
			out := map[string]interface{}{}

			err := f.fakeInt16(fg, "test.field", out, fieldMock)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(-500))
			fg.AssertCalled(t, "Int16", "test.field", fieldMock)

		})

		Convey("Should return error if the generator returns an error", func() {

			fg := &fmocks.Generator{}
			f := &Json{}

	        fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Int16)

			fieldMock.On("ID").Return("Table/Test")

			fg.On("Int16", "test.field", fieldMock).Return(int16(-500), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeInt16(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.Int16)
			So(ef.OmniID, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorInt16(t *testing.T) {
	Convey("Test_FakeVectorInt16", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt16)
			fieldMock.On("ID").Return("Struct/Test")


			fieldMock.On("Id").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorInt16(fg, "test.field", out, fieldMock)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt16)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt16)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorInt16(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt16)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt16(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt16)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt16)

			itemsMock:= &rmocks.ItemsContainer{}
			fieldMock.On("Items").Return(itemsMock)
			itemsMock.On("HybridType").Return(hybrids.Int16)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Int16", "test.field", fieldMock).Return(int16(-500), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt16(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.Int16)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt16)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Int16", "test.field", fieldMock).Return(int16(-500), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorInt16(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(-500))
		})
	})
}



func Test_FakeUint16(t *testing.T) {
	Convey("Test_FakeUint16", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")

			fg.On("Uint16", "test.field", fieldMock).Return(uint16(458), nil)
			out := map[string]interface{}{}

			err := f.fakeUint16(fg, "test.field", out, fieldMock)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(458))
			fg.AssertCalled(t, "Uint16", "test.field", fieldMock)

		})

		Convey("Should return error if the generator returns an error", func() {

			fg := &fmocks.Generator{}
			f := &Json{}

	        fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Uint16)

			fieldMock.On("ID").Return("Table/Test")

			fg.On("Uint16", "test.field", fieldMock).Return(uint16(458), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeUint16(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.Uint16)
			So(ef.OmniID, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorUint16(t *testing.T) {
	Convey("Test_FakeVectorUint16", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint16)
			fieldMock.On("ID").Return("Struct/Test")


			fieldMock.On("Id").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorUint16(fg, "test.field", out, fieldMock)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint16)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint16)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorUint16(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint16)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint16(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint16)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint16)

			itemsMock:= &rmocks.ItemsContainer{}
			fieldMock.On("Items").Return(itemsMock)
			itemsMock.On("HybridType").Return(hybrids.Uint16)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Uint16", "test.field", fieldMock).Return(uint16(458), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint16(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.Uint16)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint16)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Uint16", "test.field", fieldMock).Return(uint16(458), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorUint16(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(458))
		})
	})
}



func Test_FakeInt32(t *testing.T) {
	Convey("Test_FakeInt32", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")

			fg.On("Int32", "test.field", fieldMock).Return(int32(-60000), nil)
			out := map[string]interface{}{}

			err := f.fakeInt32(fg, "test.field", out, fieldMock)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(-60000))
			fg.AssertCalled(t, "Int32", "test.field", fieldMock)

		})

		Convey("Should return error if the generator returns an error", func() {

			fg := &fmocks.Generator{}
			f := &Json{}

	        fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Int32)

			fieldMock.On("ID").Return("Table/Test")

			fg.On("Int32", "test.field", fieldMock).Return(int32(-60000), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeInt32(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.Int32)
			So(ef.OmniID, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorInt32(t *testing.T) {
	Convey("Test_FakeVectorInt32", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt32)
			fieldMock.On("ID").Return("Struct/Test")


			fieldMock.On("Id").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorInt32(fg, "test.field", out, fieldMock)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt32)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt32)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorInt32(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt32)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt32(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt32)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt32)

			itemsMock:= &rmocks.ItemsContainer{}
			fieldMock.On("Items").Return(itemsMock)
			itemsMock.On("HybridType").Return(hybrids.Int32)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Int32", "test.field", fieldMock).Return(int32(-60000), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt32(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.Int32)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt32)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Int32", "test.field", fieldMock).Return(int32(-60000), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorInt32(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(-60000))
		})
	})
}



func Test_FakeUint32(t *testing.T) {
	Convey("Test_FakeUint32", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")

			fg.On("Uint32", "test.field", fieldMock).Return(uint32(23568778), nil)
			out := map[string]interface{}{}

			err := f.fakeUint32(fg, "test.field", out, fieldMock)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(23568778))
			fg.AssertCalled(t, "Uint32", "test.field", fieldMock)

		})

		Convey("Should return error if the generator returns an error", func() {

			fg := &fmocks.Generator{}
			f := &Json{}

	        fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Uint32)

			fieldMock.On("ID").Return("Table/Test")

			fg.On("Uint32", "test.field", fieldMock).Return(uint32(23568778), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeUint32(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.Uint32)
			So(ef.OmniID, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorUint32(t *testing.T) {
	Convey("Test_FakeVectorUint32", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint32)
			fieldMock.On("ID").Return("Struct/Test")


			fieldMock.On("Id").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorUint32(fg, "test.field", out, fieldMock)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint32)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint32)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorUint32(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint32)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint32(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint32)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint32)

			itemsMock:= &rmocks.ItemsContainer{}
			fieldMock.On("Items").Return(itemsMock)
			itemsMock.On("HybridType").Return(hybrids.Uint32)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Uint32", "test.field", fieldMock).Return(uint32(23568778), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint32(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.Uint32)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint32)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Uint32", "test.field", fieldMock).Return(uint32(23568778), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorUint32(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(23568778))
		})
	})
}



func Test_FakeInt64(t *testing.T) {
	Convey("Test_FakeInt64", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")

			fg.On("Int64", "test.field", fieldMock).Return(int64(-54545788), nil)
			out := map[string]interface{}{}

			err := f.fakeInt64(fg, "test.field", out, fieldMock)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, "-54545788")
			fg.AssertCalled(t, "Int64", "test.field", fieldMock)

		})

		Convey("Should return error if the generator returns an error", func() {

			fg := &fmocks.Generator{}
			f := &Json{}

	        fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Int64)

			fieldMock.On("ID").Return("Table/Test")

			fg.On("Int64", "test.field", fieldMock).Return(int64(-54545788), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeInt64(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.Int64)
			So(ef.OmniID, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorInt64(t *testing.T) {
	Convey("Test_FakeVectorInt64", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt64)
			fieldMock.On("ID").Return("Struct/Test")


			fieldMock.On("Id").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorInt64(fg, "test.field", out, fieldMock)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt64)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt64)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorInt64(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt64)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt64(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.VectorInt64)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt64)

			itemsMock:= &rmocks.ItemsContainer{}
			fieldMock.On("Items").Return(itemsMock)
			itemsMock.On("HybridType").Return(hybrids.Int64)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Int64", "test.field", fieldMock).Return(int64(-54545788), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorInt64(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.Int64)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorInt64)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Int64", "test.field", fieldMock).Return(int64(-54545788), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorInt64(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, "-54545788")
		})
	})
}



func Test_FakeUint64(t *testing.T) {
	Convey("Test_FakeUint64", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")

			fg.On("Uint64", "test.field", fieldMock).Return(uint64(123587), nil)
			out := map[string]interface{}{}

			err := f.fakeUint64(fg, "test.field", out, fieldMock)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, "123587")
			fg.AssertCalled(t, "Uint64", "test.field", fieldMock)

		})

		Convey("Should return error if the generator returns an error", func() {

			fg := &fmocks.Generator{}
			f := &Json{}

	        fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Uint64)

			fieldMock.On("ID").Return("Table/Test")

			fg.On("Uint64", "test.field", fieldMock).Return(uint64(123587), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeUint64(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.Uint64)
			So(ef.OmniID, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorUint64(t *testing.T) {
	Convey("Test_FakeVectorUint64", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint64)
			fieldMock.On("ID").Return("Struct/Test")


			fieldMock.On("Id").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorUint64(fg, "test.field", out, fieldMock)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint64)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint64)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorUint64(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint64)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint64(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.VectorUint64)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint64)

			itemsMock:= &rmocks.ItemsContainer{}
			fieldMock.On("Items").Return(itemsMock)
			itemsMock.On("HybridType").Return(hybrids.Uint64)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Uint64", "test.field", fieldMock).Return(uint64(123587), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorUint64(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.Uint64)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorUint64)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Uint64", "test.field", fieldMock).Return(uint64(123587), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorUint64(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, "123587")
		})
	})
}



func Test_FakeFloat32(t *testing.T) {
	Convey("Test_FakeFloat32", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")

			fg.On("Float32", "test.field", fieldMock).Return(float32(25.50), nil)
			out := map[string]interface{}{}

			err := f.fakeFloat32(fg, "test.field", out, fieldMock)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(25.50))
			fg.AssertCalled(t, "Float32", "test.field", fieldMock)

		})

		Convey("Should return error if the generator returns an error", func() {

			fg := &fmocks.Generator{}
			f := &Json{}

	        fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Float32)

			fieldMock.On("ID").Return("Table/Test")

			fg.On("Float32", "test.field", fieldMock).Return(float32(25.50), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeFloat32(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.Float32)
			So(ef.OmniID, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorFloat32(t *testing.T) {
	Convey("Test_FakeVectorFloat32", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorFloat32)
			fieldMock.On("ID").Return("Struct/Test")


			fieldMock.On("Id").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorFloat32(fg, "test.field", out, fieldMock)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.VectorFloat32)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorFloat32)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorFloat32(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorFloat32)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorFloat32(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.VectorFloat32)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorFloat32)

			itemsMock:= &rmocks.ItemsContainer{}
			fieldMock.On("Items").Return(itemsMock)
			itemsMock.On("HybridType").Return(hybrids.Float32)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Float32", "test.field", fieldMock).Return(float32(25.50), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorFloat32(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.Float32)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorFloat32)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Float32", "test.field", fieldMock).Return(float32(25.50), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorFloat32(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(25.50))
		})
	})
}



func Test_FakeFloat64(t *testing.T) {
	Convey("Test_FakeFloat64", t, func() {

		Convey("Should populate field with the data of FieldGenerator", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")

			fg.On("Float64", "test.field", fieldMock).Return(float64(-356.4545), nil)
			out := map[string]interface{}{}

			err := f.fakeFloat64(fg, "test.field", out, fieldMock)

			So(err, ShouldBeNil)
			So(out["field"], ShouldEqual, float64(-356.4545))
			fg.AssertCalled(t, "Float64", "test.field", fieldMock)

		})

		Convey("Should return error if the generator returns an error", func() {

			fg := &fmocks.Generator{}
			f := &Json{}

	        fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.Float64)

			fieldMock.On("ID").Return("Table/Test")

			fg.On("Float64", "test.field", fieldMock).Return(float64(-356.4545), fmt.Errorf("failed entropy"))
			out := map[string]interface{}{}

			err := f.fakeFloat64(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.Float64)
			So(ef.OmniID, ShouldEqual, "Table/Test")
			So(ef.Path, ShouldEqual, "test.field")


		})
	})
}

	
func Test_FakeVectorFloat64(t *testing.T) {
	Convey("Test_FakeVectorFloat64", t, func() {

		Convey("should return error if ShouldBeNil selector returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorFloat64)
			fieldMock.On("ID").Return("Struct/Test")


			fieldMock.On("Id").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, fmt.Errorf("entropy error"))
			out := map[string]interface{}{}

			err := f.fakeVectorFloat64(fg, "test.field", out, fieldMock)

			So(err, ShouldNotBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			ef := err.(*Error)
			So(ef.HybridType, ShouldEqual, hybrids.VectorFloat64)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should put the field nil if  ShouldBeNil returns true", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorFloat64)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(true, nil)
			out := map[string]interface{}{}

			err := f.fakeVectorFloat64(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			value, ok := out["field"]
			So(value, ShouldEqual, nil)
			So(ok, ShouldBeTrue)

		})

		Convey("should return error if the random vector len generator returns errors", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorFloat64)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(0, fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorFloat64(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.VectorFloat64)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field")

		})

		Convey("should return error if the random value generator returns error", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorFloat64)

			itemsMock:= &rmocks.ItemsContainer{}
			fieldMock.On("Items").Return(itemsMock)
			itemsMock.On("HybridType").Return(hybrids.Float64)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Float64", "test.field", fieldMock).Return(float64(-356.4545), fmt.Errorf("entropy error"))

			out := map[string]interface{}{}

			err := f.fakeVectorFloat64(fg, "test.field", out, fieldMock)
			So(err, ShouldNotBeNil)
			ef := err.(*Error)

			fg.AssertCalled(t, "ShouldBeNil", "test.field", fieldMock)
			fg.AssertCalled(t, "VectorLen", "test.field", fieldMock)
			So(ef.HybridType, ShouldEqual, hybrids.Float64)
			So(ef.OmniID, ShouldEqual, "Struct/Test")
			So(ef.Path, ShouldEqual, "test.field[0]")

		})

		Convey("Should set a vector with the random len generated if all was ok", func() {
			fg := &fmocks.Generator{}
			f := &Json{}

			fieldMock := &rmocks.FieldContainer{}
			fieldMock.On("Name").Return("field")
			fieldMock.On("HybridType").Return(hybrids.VectorFloat64)

			fieldMock.On("ID").Return("Struct/Test")
			fg.On("ShouldBeNil", "test.field", fieldMock).Return(false, nil)
			fg.On("VectorLen", "test.field", fieldMock).Return(10, nil)
			fg.On("Float64", "test.field", fieldMock).Return(float64(-356.4545), nil)

			out := map[string]interface{}{}

			err := f.fakeVectorFloat64(fg, "test.field", out, fieldMock)
			So(err, ShouldBeNil)
			vector, ok := out["field"].([]interface{})
			So(ok, ShouldBeTrue)

			So(len(vector), ShouldEqual, 10)
			So(vector[0], ShouldEqual, float64(-356.4545))
		})
	})
}

