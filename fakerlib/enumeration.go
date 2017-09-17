package faker

import (
	"github.com/omniql/reflect"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
	"github.com/omniql/tools/fieldgen"
)

func (j *Json) fakeEnum(fg fieldgen.Generator, path string, out map[string]interface{}, ft reflect.FieldContainer, enumType reflect.EnumerationContainer) (err error) {
	var enumName string
	var enumUint8 uint8
	var enumUint16 uint16

	choiceString, err := fg.Enumeration().ShouldGenerateString(path, ft)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: ft.HybridType(),
			OmniID:     ft.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	if choiceString {
		enumName, err = fg.Enumeration().StringEnumeration(path, ft)
		if err != nil {
			err = &Error{
				Path:       path,
				HybridType: ft.HybridType(),
				OmniID:     ft.ID(),
				ErrorMsg:   err.Error(),
			}
			return
		}

		out[ft.Name()] = enumName
		return
	}

	switch enumType.HybridType() {

	case hybrids.Uint8:
		enumUint8, err = fg.Enumeration().Uint8Enumeration(path, ft)
		if err != nil {
			err = &Error{
				Path:       path,
				HybridType: ft.HybridType(),
				OmniID:     ft.ID(),
				ErrorMsg:   err.Error(),
			}
			return
		}
		//write the enum as json number
		out[ft.Name()] = float64(enumUint8)

	case hybrids.Uint16:

		enumUint16, err = fg.Enumeration().Uint16Enumeration(path, ft)
		if err != nil {
			err = &Error{
				Path:       path,
				HybridType: ft.HybridType(),
				OmniID:     ft.ID(),
				ErrorMsg:   err.Error(),
			}
			return
		}
		//write the enum as json number
		out[ft.Name()] = float64(enumUint16)
	}
	return
}

func (j *Json) fakeVectorEnum(fg fieldgen.Generator, path string, out map[string]interface{}, ft reflect.FieldContainer, enumType reflect.EnumerationContainer) (err error) {
	var vLen int
	var shouldNil bool
	var enumName string
	var enumUint8 uint8
	var enumUint16 uint16
	var choiceString bool

	shouldNil, err = fg.ShouldBeNil(path, ft)

	if err != nil {
		err = &Error{
			Path:       path,
			HybridType: ft.HybridType(),
			OmniID:     ft.ID(),
			ErrorMsg:   err.Error(),
		}
		return
	}

	if shouldNil {
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
		return err
	}

	r := make([]interface{}, 0, vLen)

	switch ft.Items().HybridType() {

	case hybrids.Uint8:

		for i := 0; i < vLen; i++ {
			choiceString, err = fg.Enumeration().ShouldGenerateString(path, ft)

			if err != nil {
				err = &Error{
					Path:       path + fmt.Sprintf("[%d]", i),
					HybridType: ft.Items().HybridType(),
					OmniID:     ft.ID(),
					ErrorMsg:   err.Error(),
				}
				return
			}

			if choiceString {
				enumName, err = fg.Enumeration().StringEnumeration(path, ft)
				if err != nil {
					err = &Error{
						Path:       path + fmt.Sprintf("[%d]", i),
						HybridType: ft.Items().HybridType(),
						OmniID:     ft.ID(),
						ErrorMsg:   err.Error(),
					}
					return
				}

				r = append(r, enumName)
			} else {
				enumUint8, err = fg.Enumeration().Uint8Enumeration(path, ft)
				if err != nil {
					err = &Error{
						Path:       path + fmt.Sprintf("[%d]", i),
						HybridType: ft.Items().HybridType(),
						OmniID:     ft.ID(),
						ErrorMsg:   err.Error(),
					}
					return
				}
				r = append(r, float64(enumUint8))
			}
		}

	case hybrids.Uint16:
		for i := 0; i < vLen; i++ {
			choiceString, err = fg.Enumeration().ShouldGenerateString(path, ft)

			if err != nil {
				err = &Error{
					Path:       path + fmt.Sprintf("[%d]", i),
					HybridType: ft.Items().HybridType(),
					OmniID:     ft.ID(),
					ErrorMsg:   err.Error(),
				}
				return
			}

			if choiceString {
				enumName, err = fg.Enumeration().StringEnumeration(path, ft)
				if err != nil {
					err = &Error{
						Path:       path + fmt.Sprintf("[%d]", i),
						HybridType: ft.Items().HybridType(),
						OmniID:     ft.ID(),
						ErrorMsg:   err.Error(),
					}
					return
				}

				r = append(r, enumName)
			} else {
				enumUint16, err = fg.Enumeration().Uint16Enumeration(path, ft)
				if err != nil {
					err = &Error{
						Path:       path + fmt.Sprintf("[%d]", i),
						HybridType: ft.Items().HybridType(),
						OmniID:     ft.ID(),
						ErrorMsg:   err.Error(),
					}
					return
				}
				r = append(r, float64(enumUint16))
			}
		}

	}

	out[ft.Name()] = r
	return
}
