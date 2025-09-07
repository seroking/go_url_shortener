package helpers

import (
	"github.com/deatil/go-encoding/encoding"
)

func GenerateCode(id string) string {
	base62Data := encoding.FromString(id).Base62Encode().ToString()
	// fmt.Printf("starting url : %v\nencoded value: %v", id, base62Data)
	return base62Data

}
