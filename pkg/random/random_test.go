package random

import "testing"

func TestUUIDV4WithTimeStamp(t *testing.T) {
	id := UUIDV4WithTimeStamp()
	t.Log(id)
}

func TestRandIntStr(t *testing.T) {
	str := RandIntStr(6)
	t.Log(str)
}
func BenchmarkRandIntStr(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		RandIntStr(6)
	}
	b.StopTimer()
}

func TestUUIDV4(t *testing.T) {
	uuidv4 := UUIDV4()
	t.Log(uuidv4)
}

func TestRandAllString(t *testing.T) {
	allString := RandAllString(4)
	t.Log(allString)
}

func BenchmarkRandAllString(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		RandAllString(4)
	}
	b.StopTimer()
}

func TestRandString(t *testing.T) {
	t.Log(RandString(4))
}

func BenchmarkRandString(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		RandString(4)
	}
	b.StopTimer()
}
