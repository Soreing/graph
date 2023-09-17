package resource

import (
	"errors"
	"fmt"

	"github.com/Soreing/parsley"
	parse "github.com/Soreing/parsley"
	reader "github.com/Soreing/parsley/reader"
)

//parsley:skip
type AnyResource struct {
	value   any
	rType   ResourceType
	isSlice bool
}

func (r *AnyResource) set(v any, rt ResourceType, slc bool) {
	r.value = v
	r.rType = rt
	r.isSlice = slc
}

func (r *AnyResource) SetValue(val any) (err error) {
	r.value = val
	if arr, ok := val.([]Resource); ok {
		r.isSlice = true
		if len(arr) == 0 {
			r.rType = UnknownResourceType
		} else {
			t, s := arr[0].GetType()
			for _, rs := range arr {
				if nt, ns := rs.GetType(); nt != t || ns != s {
					t = MixedResourceType
					break
				}
			}
			r.rType = t
		}
	} else if obj, ok := val.(Resource); ok {
		r.rType, _ = obj.GetType()
		r.isSlice = false
	} else {
		err = errors.New("not a resource")
	}
	return
}

func (r *AnyResource) GetValue() any {
	return r.value
}

func (r *AnyResource) GetType() (ResourceType, bool) {
	return r.rType, r.isSlice
}

func (r *AnyResource) Validate() bool {
	if r.isSlice {
		slc := r.value.([]Resource)
		for _, rs := range slc {
			t, _ := rs.GetType()
			if t != r.rType || !rs.Validate() {
				return false
			}
		}
	} else {
		obj := r.value.(Resource)
		t, _ := obj.GetType()
		if t != r.rType || !obj.Validate() {
			return false
		}
	}
	return true
}

func (r *AnyResource) UnmarshalJSON(data []byte) (err error) {
	if err = parsley.Unmarshal(data, r); err == nil {
		if ok := r.Validate(); !ok {
			err = NewValidationError()
		}
	}
	return
}

func pickFilter(odt []byte, filters []parse.Filter) []parse.Filter {
	for i := range filters {
		if string(odt) == filters[i].Field {
			return filters[i].Filter
		}
	}
	return nil
}

func peekODataType(r *reader.Reader) (tp []byte, err error) {
	var key []byte
	pos := r.GetPosition()
	err = r.OpenObject()

	for err == nil {
		key, err = r.Key()
		if err == nil {
			switch string(key) {
			case "@odata.type":
				tp, err = r.Bytes()
				r.SetPosition(pos)
				return
			default:
				err = r.Skip()
			}

			if err == nil && !r.Next() {
				break
			}
		}
	}

	err = fmt.Errorf("data type field not found")
	r.SetPosition(pos)
	return
}

func (o *AnyResource) sequence(
	r *reader.Reader,
	f []parse.Filter,
	idx int,
) (res []Resource, err error) {
	var odt []byte
	odt, err = peekODataType(r)
	if err == nil {
		obj := NewResource(odt)
		flt := pickFilter(odt, f)
		err = obj.DecodeObjectPJSON(r, flt)
		if err == nil {
			if !r.Next() {
				res = make([]Resource, idx+1)
				res[idx] = obj
			} else {
				res, err = o.sequence(r, f, idx+1)
				if err == nil {
					res[idx] = obj
				}
			}
		}
	}
	return
}

func (o *AnyResource) slice(
	r *reader.Reader,
	f []parse.Filter,
) (res []Resource, err error) {
	err = r.OpenArray()
	if err == nil {
		res, err = o.sequence(r, f, 0)
		if err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *AnyResource) DecodeObjectPJSON(
	r *reader.Reader,
	f []parse.Filter,
) (err error) {
	switch r.Token() {
	case reader.NullToken:
		o.set(nil, NullResourceType, false)

	case reader.ObjectToken:
		var odt []byte
		odt, err = peekODataType(r)
		if err == nil {
			obj := NewResource(odt)
			flt := pickFilter(odt, f)
			err = obj.DecodeObjectPJSON(r, flt)
			if err == nil {
				err = o.SetValue(obj)
			}
		}

	case reader.ArrayToken:
		var slc []Resource
		slc, err = o.slice(r, f)
		if err == nil {
			err = o.SetValue(slc)
		}

	default:
		err = errors.New("not a resource")
	}

	return
}
