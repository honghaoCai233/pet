package location

import "testing"

func TestCode2ZhLocation(t *testing.T) {
	t.Log(Code2ZhLocation("CN"))
	t.Log(Code2ZhLocation("US"))
}
func TestZHLocation2Code(t *testing.T) {
	t.Log(ZHLocation2Code("火星"))
	t.Log(ZHLocation2Code("   "))
	t.Log(ZHLocation2Code("中国"))
	t.Log(ZHLocation2Code("美国"))
	t.Log(ZHLocation2Code("巴西"))
	t.Log(ZHLocation2Code("中国香港"))
}
