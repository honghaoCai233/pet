package kv

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	kv := New[string, any]()
	t.Log(kv)
}
func TestJoin(t *testing.T) {
	kv1 := New[string, string]()
	kv1.Put("a", "a")
	kv2 := New[string, string]()
	kv2.Put("b", "b")
	join := Join(kv1, kv2)
	t.Log(join)
}
func TestPairs(t *testing.T) {
	pairs, err := Pairs("a", "a", "b", "b")
	t.Log(err)
	t.Log(pairs)
}

func TestKV_Put(t *testing.T) {
	kv := New[string, string]()
	kv.Put("a", "a")
	kv.Put("b", "b")
	kv.Put("c", "c")
	kv.Put("d", "d").
		Put("e", "e").
		Put("f", "f")
	t.Log(kv)
}
func TestKV_Get(t *testing.T) {
	kv := New[string, string]()
	kv.Put("a", "a")
	kv.Put("b", "b")
	kv.Put("c", "c")
	kv.Put("d", "d")
	v, err := kv.Get("a").String()
	t.Log(err)
	t.Log(v)
	v2, err := kv.Get("b").Int64()
	t.Log(err)
	t.Log(v2)
	v3, err := kv.Get("e").String()
	t.Log(err)
	t.Log(v3)
}

func TestKV_Has(t *testing.T) {
	kv := New[string, string]()
	kv.Put("a", "a")
	kv.Put("b", "b")
	kv.Put("c", "c")
	kv.Put("d", "d")
	t.Log(kv.Has("a"))
	t.Log(kv.Has("b"))
	t.Log(kv.Has("e"))
}

func TestKV_Len(t *testing.T) {
	kv := New[string, string]()
	kv.Put("a", "a")
	kv.Put("b", "b")
	kv.Put("c", "c")
	kv.Put("d", "d")
	t.Log(kv.Len())
}
func TestKV_Clone(t *testing.T) {
	kv := New[string, string]()
	kv.Put("a", "a")
	kv.Put("b", "b")
	kv.Put("c", "c")
	kv.Put("d", "d")
	clone := kv.Clone()
	fmt.Printf("%p", kv)
	t.Log("----------")
	fmt.Printf("%p", clone)
}
func TestKV_Clone2(t *testing.T) {
	kv := New[string, string]()
	kv.Put("a", "a")
	kv.Put("b", "b")
	kv.Put("c", "c")
	kv.Put("d", "d")
	clone := kv.Clone()
	clone.Put("e", "e")
	t.Log(kv.Has("e"))
	t.Log(clone.Has("e"))
}

func TestKV_Marshal(t *testing.T) {
	kv := New[string, string]()
	kv.Put("a", "a")
	kv.Put("b", "b")
	kv.Put("c", "c")
	kv.Put("d", "d")
	marshal, err := kv.Marshal()
	t.Log(err)
	t.Log(marshal)
}

type Asd struct {
	Asd string `json:"asd"`
	Dsa int    `json:"dsa"`
}

func TestKV_Unmarshal(t *testing.T) {
	k := New[string, any]()
	k.Put("asd", "sra")
	k.Put("dsa", 1)
	a := Asd{}
	err := k.Unmarshal(&a)
	t.Log(err)
	t.Log(a)
}

func TestKV_Keys(t *testing.T) {
	k := New[string, int]()
	k.Put("a", 1)
	k.Put("b", 2)
	k.Put("c", 3)
	keys := k.Keys()
	t.Log(keys)
}

func TestKV_Values(t *testing.T) {
	k := New[string, int]()
	k.Put("a", 1)
	k.Put("b", 2)
	k.Put("c", 3)
	vals := k.Values()
	t.Log(vals)
}

func TestNewKSVA(t *testing.T) {
	ksva := NewKSVA()
	ksva.Put("a", "a").
		Put("b", "b").
		Put("c", "c").
		Put("1", 1).
		Put("2", 2).
		Put("3", 3).
		Put("1.111", 1.111)
	t.Log(ksva)
}
