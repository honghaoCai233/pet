package xphone

import "testing"

func TestHide(t *testing.T) {
	hide := HideSensitive("+15")
	t.Log(hide)
	hide = HideSensitive("+8615")
	t.Log(hide)
	hide = HideSensitive("+115700483629")
	t.Log(hide)
	hide = HideSensitive("+8615700483629")
	t.Log(hide)
	four := GetLastFour("12312")
	t.Log(four)
}

func BenchmarkHide(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HideSensitive("+8615700483629")
	}
}

func TestIsCNPhoneNumber(t *testing.T) {
	ok := IsCNPhoneNumber("+8615700483629")
	t.Log(ok)
}
