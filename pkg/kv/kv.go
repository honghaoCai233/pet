package kv

import (
	"encoding/json"
	"fmt"
)

type Key interface {
	int | uint |
		int8 | uint8 |
		int16 | uint16 |
		int32 | uint32 |
		int64 | uint64 |
		float32 | float64 |
		string
}
type Value interface {
	any
}
type KV[K Key, V Value] map[K]V

// KSVA key string value any
type KSVA = KV[string, any]

func New[K Key, V Value]() KV[K, V] {
	return make(map[K]V)
}
func NewKSVA() KSVA {
	return New[string, any]()
}

func (kv KV[K, V]) Len() int {
	return len(kv)
}

func (kv KV[K, V]) Clone() KV[K, V] {
	return Join(kv)
}

func (kv KV[K, V]) Keys() []K {
	rv := make([]K, 0, kv.Len())
	for k, _ := range kv {
		rv = append(rv, k)
	}
	return rv
}
func (kv KV[K, V]) Values() []V {
	rv := make([]V, 0, kv.Len())
	for _, v := range kv {
		rv = append(rv, v)
	}
	return rv
}

func (kv KV[K, V]) Put(k K, v V) KV[K, V] {
	kv[k] = v
	return kv
}
func (kv KV[K, V]) Get(k K) *Reply[K, V] {
	v, ok := kv[k]
	return newReply[K, V](k, v, ok)
}

func (kv KV[K, V]) Value(k K) (V, bool) {
	v, ok := kv[k]
	return v, ok
}

func (kv KV[K, V]) MustValue(k K) V {
	v, _ := kv[k]
	return v
}

func (kv KV[K, V]) Has(k K) bool {
	_, ok := kv[k]
	return ok
}

func (kv KV[K, V]) Marshal() ([]byte, error) {
	return json.Marshal(kv)
}

func (kv KV[K, V]) String() string {
	marshal, err := json.Marshal(kv)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (kv KV[K, V]) Unmarshal(v interface{}) error {
	marshal, err := kv.Marshal()
	if err != nil {
		return err
	}
	return json.Unmarshal(marshal, v)
}

func Join[K Key, V Value](kvs ...KV[K, V]) KV[K, V] {
	out := New[K, V]()
	for _, kv := range kvs {
		for k, v := range kv {
			out[k] = v
		}
	}
	return out
}

func Pairs(kvs ...interface{}) (KSVA, error) {
	if len(kvs)%2 == 1 {
		return nil, fmt.Errorf("KV: Pairs got the odd number of input pairs for KV: %d type", len(kvs))
	}
	kv := NewKSVA()
	var key string
	for i, s := range kvs {
		if i%2 == 0 {
			k, ok := s.(string)
			if !ok {
				return nil, fmt.Errorf("KV: the key need string type but got %T", s)
			}
			key = k
			continue
		}
		kv[key] = s
	}
	return kv, nil
}
