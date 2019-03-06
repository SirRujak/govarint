package govarint

var n1 uint = 7 * 7
var n2 uint = 14 * 14
var n3 uint = 21 * 21
var n4 uint = 28 * 28
var n5 uint = 35 * 35
var n6 uint = 42 * 42
var n7 uint = 49 * 49
var n8 uint = 56 * 56
var n9 uint = 63 * 63

func Length(value uint) uint {
	if value < n1 {
		return 1
	}
	if value < n2 {
		return 2
	}
	if value < n3 {
		return 3
	}
	if value < n4 {
		return 4
	}
	if value < n5 {
		return 5
	}
	if value < n6 {
		return 6
	}
	if value < n7 {
		return 7
	}
	if value < n8 {
		return 8
	}
	if value < n9 {
		return 9
	}
	return 10

}
