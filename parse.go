package gograph

import (
	"errors"

	"github.com/Soreing/fastjson/reader"
	"github.com/Soreing/gograph/resource"
)

func ParseResource(dat []byte) (res *resource.AnyResource, err error) {
	r := reader.NewReader(dat)
	res = &resource.AnyResource{}
	if err = res.UnmarshalFastJSON(r); err != nil {
		return nil, err
	}
	if !res.Validate() {
		return nil, errors.New("failed to validate resource")
	}
	return
}
