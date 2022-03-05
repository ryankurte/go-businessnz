package businessnz

import (
	"github.com/ryankurte/go-businessnz/lib/base"
	"github.com/ryankurte/go-businessnz/lib/nzbn"
)

// Top level BusinessNZ API object
type BusinessNzApi struct {
	*base.Base

	/// NZBN Api provides search by name or lookups by NZBN
	Nzbn *nzbn.NzbnApi
}

// Create a BusinessNZ API instance
func NewBusinessNzApi(apiKey string, apiSecret string) BusinessNzApi {

	base := base.NewBase(apiKey, apiSecret, false, false)

	// Bind API keys
	a := BusinessNzApi{Base: &base}

	// Attach sub APIs
	a.Nzbn = &nzbn.NzbnApi{Base: &base}

	return a
}
