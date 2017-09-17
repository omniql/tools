package faker

import "github.com/nebtex/omniql/commons/golang/oreflection"
import (
	"github.com/nebtex/omniql/tools/fieldgen"
	"github.com/nebtex/hybrids/golang/hybrids"
)

func (j *Json) fakeStack(fg *fieldgen.Generator, application oreflection.Application, stackId string, desiredResourceCount int) (stack map[string]interface{}, err error) {
	var vLen int
	var i int
	stack = map[string]interface{}{}

	for i = 0; i < desiredResourceCount; i++ {
		tableObject := map[string]interface{}{}

		rid, resourceType, err := fg.Application().GenerateResource(application)
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
		j.fakeTable("", tableObject, rtype, rtype.Table())
		stack[string(rid)] = tableObject
	}

	return
}
