package faker

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	fmocks "github.com/omniql/tools/fieldgen/mocks"
	rmocks "github.com/omniql/reflect/mocks"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
)

func Test_FakeEnumeration(t *testing.T) {
	Convey("Test_FakeEnumeration", t, func() {

		Convey("Should Fail the the random choice fail", func() {
			fg := &fmocks.Generator{}
			f := &Json{}
			fieldMock := &rmocks.FieldContainer{}
			enumType := &rmocks.EnumerationContainer{}

			fieldMock.On("Name").Return("field")
			fieldMock.On("ID").Return("dummy/test")
			fieldMock.On("HybridType").Return(hybrids.Uint16)

			enumGen := &fmocks.EnumerationGenerator{}
			enumGen.On("ShouldGenerateString", "test.field", fieldMock).Return(false, fmt.Errorf("entropy error"))
			fg.On("Enumeration").Return(enumGen)

			out := map[string]interface{}{}

			err := f.fakeEnum(fg, "test.field", out, fieldMock, enumType)

			So(err, ShouldNotBeNil)
			enumGen.AssertCalled(t, "ShouldGenerateString", "test.field", fieldMock)

		})

		Convey("if a string is selected, it should write a string in the json", func() {
			fg := &fmocks.Generator{}
			f := &Json{}
			fieldMock := &rmocks.FieldContainer{}
			enumType := &rmocks.EnumerationContainer{}

			fieldMock.On("Name").Return("field")
			fieldMock.On("ID").Return("dummy/test")
			fieldMock.On("HybridType").Return(hybrids.Uint16)

			enumGen := &fmocks.EnumerationGenerator{}
			enumGen.On("ShouldGenerateString", "test.field", fieldMock).Return(true, nil)
			enumGen.On("StringEnumeration", "test.field", fieldMock).Return("PASS", nil)

			fg.On("Enumeration").Return(enumGen)

			out := map[string]interface{}{}

			err := f.fakeEnum(fg, "test.field", out, fieldMock, enumType)

			So(err, ShouldBeNil)
			enumGen.AssertCalled(t, "ShouldGenerateString", "test.field", fieldMock)
			enumGen.AssertCalled(t, "StringEnumeration", "test.field", fieldMock)
			So(out["field"], ShouldEqual, "PASS")

		})

		Convey("if a string generation fails should fail", func() {
			fg := &fmocks.Generator{}
			f := &Json{}
			fieldMock := &rmocks.FieldContainer{}
			enumType := &rmocks.EnumerationContainer{}

			fieldMock.On("Name").Return("field")
			fieldMock.On("ID").Return("dummy/test")
			fieldMock.On("HybridType").Return(hybrids.Uint16)

			enumGen := &fmocks.EnumerationGenerator{}
			enumGen.On("ShouldGenerateString", "test.field", fieldMock).Return(true, nil)
			enumGen.On("StringEnumeration", "test.field", fieldMock).Return("", fmt.Errorf("entropy error"))
			fg.On("Enumeration").Return(enumGen)
			out := map[string]interface{}{}
			err := f.fakeEnum(fg, "test.field", out, fieldMock, enumType)
			So(err, ShouldNotBeNil)

		})

		Convey("should write a float64 in other cases", func() {
			Convey("Uint16", func() {

				fg := &fmocks.Generator{}
				f := &Json{}
				fieldMock := &rmocks.FieldContainer{}
				enumType := &rmocks.EnumerationContainer{}

				enumType.On("HybridType").Return(hybrids.Uint16)

				fieldMock.On("Name").Return("field")
				fieldMock.On("ID").Return("dummy/test")
				fieldMock.On("HybridType").Return(hybrids.Uint16)

				enumGen := &fmocks.EnumerationGenerator{}
				enumGen.On("ShouldGenerateString", "test.field", fieldMock).Return(false, nil)
				enumGen.On("Uint16Enumeration", "test.field", fieldMock).Return(uint16(30), nil)
				fg.On("Enumeration").Return(enumGen)

				out := map[string]interface{}{}

				err := f.fakeEnum(fg, "test.field", out, fieldMock, enumType)

				So(err, ShouldBeNil)
				So(out["field"], ShouldEqual, float64(30))

				Convey("Generator fails", func() {

					fg := &fmocks.Generator{}
					f := &Json{}
					fieldMock := &rmocks.FieldContainer{}
					enumType := &rmocks.EnumerationContainer{}

					enumType.On("HybridType").Return(hybrids.Uint16)

					fieldMock.On("Name").Return("field")
					fieldMock.On("ID").Return("dummy/test")
					fieldMock.On("HybridType").Return(hybrids.Uint16)

					enumGen := &fmocks.EnumerationGenerator{}
					enumGen.On("ShouldGenerateString", "test.field", fieldMock).Return(false, nil)
					enumGen.On("Uint16Enumeration", "test.field", fieldMock).Return(uint16(0), fmt.Errorf("entropy error"))
					fg.On("Enumeration").Return(enumGen)

					out := map[string]interface{}{}
					err := f.fakeEnum(fg, "test.field", out, fieldMock, enumType)
					So(err, ShouldNotBeNil)

				})
			})

			Convey("Uint8", func() {

				fg := &fmocks.Generator{}
				f := &Json{}
				fieldMock := &rmocks.FieldContainer{}
				enumType := &rmocks.EnumerationContainer{}

				enumType.On("HybridType").Return(hybrids.Uint8)

				fieldMock.On("Name").Return("field")
				fieldMock.On("ID").Return("dummy/test")
				fieldMock.On("HybridType").Return(hybrids.Uint8)

				enumGen := &fmocks.EnumerationGenerator{}
				enumGen.On("ShouldGenerateString", "test.field", fieldMock).Return(false, nil)
				enumGen.On("Uint8Enumeration", "test.field", fieldMock).Return(uint8(120), nil)
				fg.On("Enumeration").Return(enumGen)

				out := map[string]interface{}{}

				err := f.fakeEnum(fg, "test.field", out, fieldMock, enumType)

				So(err, ShouldBeNil)
				So(out["field"], ShouldEqual, float64(120))

				Convey("Generator fails", func() {

					fg := &fmocks.Generator{}
					f := &Json{}
					fieldMock := &rmocks.FieldContainer{}
					enumType := &rmocks.EnumerationContainer{}

					enumType.On("HybridType").Return(hybrids.Uint8)

					fieldMock.On("Name").Return("field")
					fieldMock.On("ID").Return("dummy/test")
					fieldMock.On("HybridType").Return(hybrids.Uint8)

					enumGen := &fmocks.EnumerationGenerator{}
					enumGen.On("ShouldGenerateString", "test.field", fieldMock).Return(false, nil)
					enumGen.On("Uint8Enumeration", "test.field", fieldMock).Return(uint8(0), fmt.Errorf("entropy error"))
					fg.On("Enumeration").Return(enumGen)

					out := map[string]interface{}{}
					err := f.fakeEnum(fg, "test.field", out, fieldMock, enumType)
					So(err, ShouldNotBeNil)

				})
			})

		})
	})
}

/*
func Test_FakeVectorEnumeration(t *testing.T) {
	Convey("Test_FakeVectorEnumeration", t, func() {

		Convey("Should Fail the the random choice fail", func() {
			fg := &fmocks.Generator{}
			f := &Json{}
			fieldMock := &rmocks.FieldContainer{}
			enumType := &rmocks.EnumerationContainer{}

			fieldMock.On("Name").Return("field")
			fieldMock.On("ID").Return("dummy/test")
			fieldMock.On("HybridType").Return(hybrids.Uint16)

			enumGen := &fmocks.EnumerationGenerator{}
			enumGen.On("ShouldGenerateString", "test.field", fieldMock).Return(false, fmt.Errorf("entropy error"))
			fg.On("Enumeration").Return(enumGen)

			out := map[string]interface{}{}

			err := f.fakeEnum(fg, "test.field", out, fieldMock, enumType)

			So(err, ShouldNotBeNil)
			enumGen.AssertCalled(t, "ShouldGenerateString", "test.field", fieldMock)

		})

		Convey("if a string is selected, it should write a string in the json", func() {
			fg := &fmocks.Generator{}
			f := &Json{}
			fieldMock := &rmocks.FieldContainer{}
			enumType := &rmocks.EnumerationContainer{}

			fieldMock.On("Name").Return("field")
			fieldMock.On("ID").Return("dummy/test")
			fieldMock.On("HybridType").Return(hybrids.Uint16)

			enumGen := &fmocks.EnumerationGenerator{}
			enumGen.On("ShouldGenerateString", "test.field", fieldMock).Return(true, nil)
			enumGen.On("StringEnumeration", "test.field", fieldMock).Return("PASS", nil)

			fg.On("Enumeration").Return(enumGen)

			out := map[string]interface{}{}

			err := f.fakeEnum(fg, "test.field", out, fieldMock, enumType)

			So(err, ShouldBeNil)
			enumGen.AssertCalled(t, "ShouldGenerateString", "test.field", fieldMock)
			enumGen.AssertCalled(t, "StringEnumeration", "test.field", fieldMock)
			So(out["field"], ShouldEqual, "PASS")

		})

		Convey("if a string generation fails should fail", func() {
			fg := &fmocks.Generator{}
			f := &Json{}
			fieldMock := &rmocks.FieldContainer{}
			enumType := &rmocks.EnumerationContainer{}

			fieldMock.On("Name").Return("field")
			fieldMock.On("ID").Return("dummy/test")
			fieldMock.On("HybridType").Return(hybrids.Uint16)

			enumGen := &fmocks.EnumerationGenerator{}
			enumGen.On("ShouldGenerateString", "test.field", fieldMock).Return(true, nil)
			enumGen.On("StringEnumeration", "test.field", fieldMock).Return("", fmt.Errorf("entropy error"))
			fg.On("Enumeration").Return(enumGen)
			out := map[string]interface{}{}
			err := f.fakeEnum(fg, "test.field", out, fieldMock, enumType)
			So(err, ShouldNotBeNil)

		})

		Convey("should write a float64 in other cases", func() {
			Convey("Uint16", func() {

				fg := &fmocks.Generator{}
				f := &Json{}
				fieldMock := &rmocks.FieldContainer{}
				enumType := &rmocks.EnumerationContainer{}

				enumType.On("HybridType").Return(hybrids.Uint16)

				fieldMock.On("Name").Return("field")
				fieldMock.On("ID").Return("dummy/test")
				fieldMock.On("HybridType").Return(hybrids.Uint16)

				enumGen := &fmocks.EnumerationGenerator{}
				enumGen.On("ShouldGenerateString", "test.field", fieldMock).Return(false, nil)
				enumGen.On("Uint16Enumeration", "test.field", fieldMock).Return(uint16(30), nil)
				fg.On("Enumeration").Return(enumGen)

				out := map[string]interface{}{}

				err := f.fakeEnum(fg, "test.field", out, fieldMock, enumType)

				So(err, ShouldBeNil)
				So(out["field"], ShouldEqual, float64(30))

				Convey("Generator fails", func() {

					fg := &fmocks.Generator{}
					f := &Json{}
					fieldMock := &rmocks.FieldContainer{}
					enumType := &rmocks.EnumerationContainer{}

					enumType.On("HybridType").Return(hybrids.Uint16)

					fieldMock.On("Name").Return("field")
					fieldMock.On("ID").Return("dummy/test")
					fieldMock.On("HybridType").Return(hybrids.Uint16)

					enumGen := &fmocks.EnumerationGenerator{}
					enumGen.On("ShouldGenerateString", "test.field", fieldMock).Return(false, nil)
					enumGen.On("Uint16Enumeration", "test.field", fieldMock).Return(uint16(0), fmt.Errorf("entropy error"))
					fg.On("Enumeration").Return(enumGen)

					out := map[string]interface{}{}
					err := f.fakeEnum(fg, "test.field", out, fieldMock, enumType)
					So(err, ShouldNotBeNil)

				})
			})

			Convey("Uint8", func() {

				fg := &fmocks.Generator{}
				f := &Json{}
				fieldMock := &rmocks.FieldContainer{}
				enumType := &rmocks.EnumerationContainer{}

				enumType.On("HybridType").Return(hybrids.Uint8)

				fieldMock.On("Name").Return("field")
				fieldMock.On("ID").Return("dummy/test")
				fieldMock.On("HybridType").Return(hybrids.Uint8)

				enumGen := &fmocks.EnumerationGenerator{}
				enumGen.On("ShouldGenerateString", "test.field", fieldMock).Return(false, nil)
				enumGen.On("Uint8Enumeration", "test.field", fieldMock).Return(uint8(120), nil)
				fg.On("Enumeration").Return(enumGen)

				out := map[string]interface{}{}

				err := f.fakeEnum(fg, "test.field", out, fieldMock, enumType)

				So(err, ShouldBeNil)
				So(out["field"], ShouldEqual, float64(120))

				Convey("Generator fails", func() {

					fg := &fmocks.Generator{}
					f := &Json{}
					fieldMock := &rmocks.FieldContainer{}
					enumType := &rmocks.EnumerationContainer{}

					enumType.On("HybridType").Return(hybrids.Uint8)

					fieldMock.On("Name").Return("field")
					fieldMock.On("ID").Return("dummy/test")
					fieldMock.On("HybridType").Return(hybrids.Uint8)

					enumGen := &fmocks.EnumerationGenerator{}
					enumGen.On("ShouldGenerateString", "test.field", fieldMock).Return(false, nil)
					enumGen.On("Uint8Enumeration", "test.field", fieldMock).Return(uint8(0), fmt.Errorf("entropy error"))
					fg.On("Enumeration").Return(enumGen)

					out := map[string]interface{}{}
					err := f.fakeEnum(fg, "test.field", out, fieldMock, enumType)
					So(err, ShouldNotBeNil)

				})
			})

		})
	})
}*/
