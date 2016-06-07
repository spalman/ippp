package main

import (
	"math"
	"strconv"
)

func F1() {
	ai, _ := strconv.Atoi(returnIndexInRP("a"))
	Si, _ := strconv.Atoi(returnIndexInRP("S"))
	Hai, _ := strconv.Atoi(returnIndexInRP("Ha"))
	RP[Si].value = RP[ai].value * RP[Hai].value

}
func F2() {
	ai, _ := strconv.Atoi(returnIndexInRP("a"))
	Si, _ := strconv.Atoi(returnIndexInRP("S"))
	Hai, _ := strconv.Atoi(returnIndexInRP("Ha"))
	RP[ai].value = RP[Si].value / RP[Hai].value

}
func F3() {
	ai, _ := strconv.Atoi(returnIndexInRP("a"))
	Si, _ := strconv.Atoi(returnIndexInRP("S"))
	alphai, _ := strconv.Atoi(returnIndexInRP("alpha"))
	RP[ai].value = math.Sqrt(RP[Si].value) / math.Sqrt(math.Sin(RP[alphai].value))
}
func F4() {
	ai, _ := strconv.Atoi(returnIndexInRP("a"))
	Si, _ := strconv.Atoi(returnIndexInRP("S"))
	betai, _ := strconv.Atoi(returnIndexInRP("beta"))
	RP[ai].value = math.Sqrt(RP[Si].value) / math.Sqrt(math.Sin(RP[betai].value))
}
func F5() {
	d1i, _ := strconv.Atoi(returnIndexInRP("d1"))
	d2i, _ := strconv.Atoi(returnIndexInRP("d2"))
	ai, _ := strconv.Atoi(returnIndexInRP("a"))
	RP[ai].value = math.Sqrt(RP[d1i].value*RP[d1i].value+RP[d2i].value*RP[d2i].value) / 2
}
func F6() {
	d1i, _ := strconv.Atoi(returnIndexInRP("d1"))
	d2i, _ := strconv.Atoi(returnIndexInRP("d2"))
	ai, _ := strconv.Atoi(returnIndexInRP("a"))
	RP[d1i].value = math.Sqrt(4*RP[ai].value*RP[ai].value - RP[d2i].value*RP[d2i].value)
}
func F7() {
	d1i, _ := strconv.Atoi(returnIndexInRP("d1"))
	d2i, _ := strconv.Atoi(returnIndexInRP("d2"))
	ai, _ := strconv.Atoi(returnIndexInRP("a"))
	RP[d2i].value = math.Sqrt(4*RP[ai].value*RP[ai].value - RP[d1i].value*RP[d1i].value)
}
func F8() {
	d1i, _ := strconv.Atoi(returnIndexInRP("d1"))
	d2i, _ := strconv.Atoi(returnIndexInRP("d2"))
	Si, _ := strconv.Atoi(returnIndexInRP("S"))
	RP[d1i].value = 2 * RP[Si].value / RP[d2i].value
}
func F9() {
	d1i, _ := strconv.Atoi(returnIndexInRP("d1"))
	d2i, _ := strconv.Atoi(returnIndexInRP("d2"))
	Si, _ := strconv.Atoi(returnIndexInRP("S"))
	RP[d2i].value = 2 * RP[Si].value / RP[d1i].value
}
func F10() {
	ai, _ := strconv.Atoi(returnIndexInRP("a"))
	alphai, _ := strconv.Atoi(returnIndexInRP("alpha"))
	Si, _ := strconv.Atoi(returnIndexInRP("S"))
	RP[Si].value = RP[ai].value * RP[ai].value * math.Sin(RP[alphai].value)
}
func F11() {
	d1i, _ := strconv.Atoi(returnIndexInRP("d1"))
	d2i, _ := strconv.Atoi(returnIndexInRP("d2"))
	Si, _ := strconv.Atoi(returnIndexInRP("S"))
	RP[Si].value = 0.5 * RP[d1i].value * RP[d2i].value
}
func F12() {
	ai, _ := strconv.Atoi(returnIndexInRP("a"))
	Si, _ := strconv.Atoi(returnIndexInRP("S"))
	Hai, _ := strconv.Atoi(returnIndexInRP("Ha"))
	RP[Hai].value = RP[Si].value / RP[ai].value
}
