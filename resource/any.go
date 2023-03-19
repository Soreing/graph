package resource

import (
	"encoding/json"
	"errors"

	"github.com/Soreing/fastjson/reader"
)

//fastjson:skip
type AnyResource struct {
	value   any
	rType   ResourceType
	isSlice bool
}

func (r *AnyResource) SetValue(val any) (err error) {
	r.value = val
	if arr, ok := val.([]Resource); ok {
		r.isSlice = true
		if len(arr) > 0 {
			t, s := arr[0].GetType()
			for i := range arr {
				if nt, ns := arr[i].GetType(); nt != t || ns != s {
					t = MixedResourceType
					break
				}
			}
			r.rType = t
		} else {
			r.rType = UnknownResourceType
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
		for i := range slc {
			t, _ := slc[i].GetType()
			if t != r.rType || !slc[i].Validate() {
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

func peekODataType(r *reader.Reader) (tp []byte, err error) {
	var key []byte
	pos := r.GetPosition()
	err = r.OpenObject()
	for err == nil {
		if key, err = r.GetKey(); err == nil {
			switch string(key) {
			case "@odata.type":
				tp, err = r.GetByteArray()
				r.SetPosition(pos)
				return
			default:
				err = r.Skip()
			}

			if err == nil && !r.Next() {
				err = r.CloseObject()
				break
			}
		}
	}
	r.SetPosition(pos)
	return
}

func (o *AnyResource) sequenceResource(r *reader.Reader, idx int) (res []Resource, err error) {
	var odt []byte
	if odt, err = peekODataType(r); err == nil {
		obj := NewResource(odt)
		if err = obj.UnmarshalFastJSON(r); err == nil {
			if !r.Next() {
				res = make([]Resource, idx+1)
				res[idx] = obj
				return
			} else if res, err = o.sequenceResource(r, idx+1); err == nil {
				res[idx] = obj
			}
		}
	}
	return
}

func (o *AnyResource) resourceSlice(r *reader.Reader) (res []Resource, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequenceResource(r, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *AnyResource) UnmarshalFastJSON(r *reader.Reader) (err error) {
	switch r.GetType() {
	case reader.ObjectToken:
		var odt []byte
		if odt, err = peekODataType(r); err == nil {
			obj := NewResource(odt)
			if err = obj.UnmarshalFastJSON(r); err == nil {
				err = o.SetValue(obj)
			}
		}
	case reader.ArrayToken:
		var slc []Resource
		if slc, err = o.resourceSlice(r); err == nil {
			err = o.SetValue(slc)
		}
	default:
		err = errors.New("not a resource")
	}
	return
}

func (o *AnyResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}
