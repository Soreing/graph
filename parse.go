package graft

import (
	"github.com/Soreing/graft/resource"
	"github.com/Soreing/parsley"
)

// ParseResource parses JSON byte array into a msft graph resource
func ParseResource(
	dat []byte,
	res resource.Resource,
	cfgs ...parsley.Config,
) (err error) {
	err = parsley.Decode(dat, res, cfgs...)

	if err == nil && !res.Validate() {
		err = resource.NewValidationError()
	}

	return
}

func addODTField(f *[]parsley.Filter) {
	included := false
	for i := range *f {
		if (*f)[i].Field == resource.ODataTypeKey {
			included = true
		}
		if (*f)[i].Filter != nil {
			addODTField(&(*f)[i].Filter)
		}
	}

	if !included {
		*f = append(*f, parsley.Filter{
			Field:  resource.ODataTypeKey,
			Filter: nil,
		})
	}
}

func UseResourceFilter(f []parsley.Filter) parsley.Config {
	addODTField(&f)
	return parsley.UseFilter(f)
}
