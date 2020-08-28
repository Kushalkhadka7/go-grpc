package sample

import (
	laptop "appserver/pb"
	"math/rand"
)

func randomKeyboardLayout() laptop.Keyboard_Layout {
	switch rand.Intn(2) {
	case 1:
		return laptop.Keyboard_QWERTY

	case 2:
		return laptop.Keyboard_AZWETY

	default:
		return laptop.Keyboard_QWERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomInt() uint32 {
	return 1
}

func randomString() string {
	return "hello"
}

func randomFloat() float64 {
	return 2.6
}

func randomEnum() laptop.Memory_Unit {
	return 1
}

func randomDriver() laptop.Storage_Driver {
	return 1
}

func randomScreenPanel() laptop.Screen_Panel {
	switch rand.Intn(2) {
	case 1:
		return laptop.Screen_IPS

	case 2:
		return laptop.Screen_OLED

	default:
		return laptop.Screen_IPS
	}
}

func randomWeight() *laptop.Laptop_WeightKg {
	return &laptop.Laptop_WeightKg{}
}
