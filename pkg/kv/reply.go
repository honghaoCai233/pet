package kv

import (
	"errors"
	"github.com/spf13/cast"
)

var NotFound = errors.New("not found")

type Reply[K Key, V Value] struct {
	k  K
	v  V
	ok bool
}

func newReply[K Key, V Value](k K, v V, ok bool) *Reply[K, V] {
	return &Reply[K, V]{
		k:  k,
		v:  v,
		ok: ok,
	}
}

func (r *Reply[K, V]) String() (string, error) {
	if !r.ok {
		return "", NotFound
	}
	return cast.ToStringE(r.v)
}
func (r *Reply[K, V]) MustString() string {
	s, _ := r.String()
	return s
}
func (r *Reply[K, V]) Int64() (int64, error) {
	if !r.ok {
		return 0, NotFound
	}
	return cast.ToInt64E(r.v)
}
func (r *Reply[K, V]) Float64() (float64, error) {
	if !r.ok {
		return 0, NotFound
	}
	return cast.ToFloat64E(r.v)
}
func (r *Reply[K, V]) Bool() (bool, error) {
	if !r.ok {
		return false, NotFound
	}
	return cast.ToBoolE(r.v)
}
func (r *Reply[K, V]) Value() (interface{}, error) {
	if !r.ok {
		return nil, NotFound
	}
	return r.v, nil
}
