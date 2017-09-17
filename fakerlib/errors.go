package faker

import (
	"fmt"
	"github.com/nebtex/hybrids/golang/hybrids"
)

//DecodeError contains the data to fully identify  an error produced by the decoder
type Error struct {
	// field full path
	Path string `json:"path,omitempty"`
	// field underlying type
	HybridType hybrids.Types `json:"hybrid_type,omitempty"`
	OmniID     string        `json:"omniql_type,omitempty"`
	ErrorMsg   string        `json:"error,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("@Faker #Path: %s, #HybridType: %s, #OmniID: %s, Error: %s", e.Path, e.HybridType.String(), e.OmniID, e.ErrorMsg)
}
