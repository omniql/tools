package faker

import (
	"github.com/omniql/reflect"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
	"github.com/nebtex/omniql/commons/golang/oreflection"
	"github.com/omniql/tools/fieldgen"
)

func (j *Json) fakeEnum(fg fieldgen.Generator, path string, out map[string]interface{}, ft reflect.FieldContainer, enumType reflect.EnumerationContainer) (err error) {
	var enumName string
	var enumUint8 uint8
	var enumUint16 uint16
	var enumUint32 uint32

	choiceString, err := fg.Enumeration().ShouldGenerateString(path, ft)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: ft.HybridType(),
			OmniqlType: ft.ID(),
			Package:    ft.Application().Path(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	if choiceString {
		enumName, err = fg.Enumeration().StringEnumeration(path, ft)
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

		out[ft.Field().Name()] = enumName
		return
	}

	switch enumType.HybridType() {

	case hybrids.Uint8:
		enumUint8, err = j.fieldGen.Enumeration().Uint8Enumeration(path, ft)
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
		//write the enum as json number
		out[ft.Field().Name()] = float64(enumUint8)

	case hybrids.Uint16:

		enumUint16, err = j.fieldGen.Enumeration().Uint16Enumeration(path, ft)
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
		//write the enum as json number
		out[ft.Field().Name()] = float64(enumUint16)

	case hybrids.Uint32:
		enumUint32, err = j.fieldGen.Enumeration().Uint32Enumeration(path, ft)
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
		//write the enum as json number
		out[ft.Field().Name()] = float64(enumUint32)
	}
	return
}

func (j *Json) fakeVectorEnum(fg fieldgen.Generator,path string, out map[string]interface{}, ft oreflection.OType, enumType oreflection.Enumeration) (err error) {
	var vLen int
	var shouldNil bool
	var enumName string
	var enumUint8 uint8
	var enumUint16 uint16
	var enumUint32 uint32
	var choiceString bool

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
		return err
	}

	r := make([]interface{}, 0, vLen)

	switch ft.Field().Items().HybridType() {

	case hybrids.Uint8:

		for i := 0; i < vLen; i++ {
			choiceString, err = j.fieldGen.Enumeration().ShouldGenerateString(path, ft)

			if err != nil {
				err = &Error{
					Path:       path + fmt.Sprintf("[%d]", i),
					HybridType: ft.Field().Items().HybridType(),
					OmniqlType: ft.Id(),
					Package:    ft.Package(),
					ErrorMsg:   err.Error(),
				}
				return
			}

			if choiceString {
				enumName, err = j.fieldGen.Enumeration().StringEnumeration(path, ft)
				if err != nil {
					err = &Error{
						Path:       path + fmt.Sprintf("[%d]", i),
						HybridType: ft.Field().Items().HybridType(),
						OmniqlType: ft.Id(),
						Package:    ft.Package(),
						ErrorMsg:   err.Error(),
					}
					return
				}

				r = append(r, enumName)
			} else {
				enumUint8, err = j.fieldGen.Enumeration().Uint8Enumeration(path, ft)
				if err != nil {
					err = &Error{
						Path:       path + fmt.Sprintf("[%d]", i),
						HybridType: ft.Field().Items().HybridType(),
						OmniqlType: ft.Id(),
						Package:    ft.Package(),
						ErrorMsg:   err.Error(),
					}
					return
				}
				r = append(r, float64(enumUint8))
			}
		}

	case hybrids.Uint16:
		for i := 0; i < vLen; i++ {
			choiceString, err = j.fieldGen.Enumeration().ShouldGenerateString(path, ft)

			if err != nil {
				err = &Error{
					Path:       path + fmt.Sprintf("[%d]", i),
					HybridType: ft.Field().Items().HybridType(),
					OmniqlType: ft.Id(),
					Package:    ft.Package(),
					ErrorMsg:   err.Error(),
				}
				return
			}

			if choiceString {
				enumName, err = j.fieldGen.Enumeration().StringEnumeration(path, ft)
				if err != nil {
					err = &Error{
						Path:       path + fmt.Sprintf("[%d]", i),
						HybridType: ft.Field().Items().HybridType(),
						OmniqlType: ft.Id(),
						Package:    ft.Package(),
						ErrorMsg:   err.Error(),
					}
					return
				}

				r = append(r, enumName)
			} else {
				enumUint16, err = j.fieldGen.Enumeration().Uint16Enumeration(path, ft)
				if err != nil {
					err = &Error{
						Path:       path + fmt.Sprintf("[%d]", i),
						HybridType: ft.Field().Items().HybridType(),
						OmniqlType: ft.Id(),
						Package:    ft.Package(),
						ErrorMsg:   err.Error(),
					}
					return
				}
				r = append(r, float64(enumUint16))
			}
		}
	case hybrids.Uint32:
		for i := 0; i < vLen; i++ {
			choiceString, err = j.fieldGen.Enumeration().ShouldGenerateString(path, ft)

			if err != nil {
				err = &Error{
					Path:       path + fmt.Sprintf("[%d]", i),
					HybridType: ft.Field().Items().HybridType(),
					OmniqlType: ft.Id(),
					Package:    ft.Package(),
					ErrorMsg:   err.Error(),
				}
				return
			}

			if choiceString {
				enumName, err = j.fieldGen.Enumeration().StringEnumeration(path, ft)
				if err != nil {
					err = &Error{
						Path:       path + fmt.Sprintf("[%d]", i),
						HybridType: ft.Field().Items().HybridType(),
						OmniqlType: ft.Id(),
						Package:    ft.Package(),
						ErrorMsg:   err.Error(),
					}
					return
				}

				r = append(r, enumName)
			} else {
				enumUint32, err = j.fieldGen.Enumeration().Uint32Enumeration(path, ft)
				if err != nil {
					err = &Error{
						Path:       path + fmt.Sprintf("[%d]", i),
						HybridType: ft.Field().Items().HybridType(),
						OmniqlType: ft.Id(),
						Package:    ft.Package(),
						ErrorMsg:   err.Error(),
					}
					return
				}
				r = append(r, float64(enumUint32))
			}
		}

	}

	out[ft.Field().Name()] = r
	return
}
