package xfile

import (
	"os"
	"testing"
)

func TestTypeByExtension(t *testing.T) {
	mime := TypeByExtension("ABC.mp4")
	t.Log(mime)
}
func BenchmarkTypeByExtension(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TypeByExtension("ABC.mp4")
	}
}
func TestX(t *testing.T) {
	os.MkdirAll("C:\\Users\\Tangerg\\Desktop\\testFile\\testDir1", os.ModePerm)
}
