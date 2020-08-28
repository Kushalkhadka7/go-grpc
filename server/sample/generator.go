package sample

import (
	laptop "appserver/pb"

	"github.com/golang/protobuf/ptypes"
)

// NewKeyboard generates sample data for keyboard.
func NewKeyboard() *laptop.Keyboard {
	keyboard := &laptop.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}

// NewCPU generates sample data for CPU.
func NewCPU() *laptop.CPU {
	cpu := &laptop.CPU{
		Name:        randomString(),
		Brand:       randomString(),
		NoOfCores:   randomInt(),
		NoOfThreads: randomInt(),
		MinGhz:      randomFloat(),
		MazGhz:      randomFloat(),
	}

	return cpu
}

// NewGPU generates sample data fro GPU.
func NewGPU() *laptop.GPU {
	gpu := &laptop.GPU{
		Brand:  randomString(),
		Name:   randomString(),
		MinGhz: randomFloat(),
		MazGhz: randomFloat(),
		Memory: &laptop.Memory{
			Value: randomInt(),
			Unit:  randomEnum(),
		},
	}

	return gpu
}

// NewStorage generates sample data for storage.
func NewStorage() *laptop.Storage {
	storage := &laptop.Storage{
		Driver: randomDriver(),
		Memory: &laptop.Memory{
			Value: randomInt(),
			Unit:  randomEnum(),
		},
	}

	return storage
}

// NewScreen generates sample data for the screen.
func NewScreen() *laptop.Screen {
	screen := &laptop.Screen{
		Size: 2,
		Resolution: &laptop.Screen_Resolution{
			Width:  2,
			Height: 3,
		},
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

// NewLaptop generates sampled data for laptop.
func NewLaptop() *laptop.Laptop {
	laptop := &laptop.Laptop{
		Id:    randomString(),
		Brand: randomString(),
		Name:  randomString(),
		Cpu:   NewCPU(),
		Ram: &laptop.Memory{
			Value: randomInt(),
			Unit:  randomEnum(),
		},
		Gpus: []*laptop.GPU{
			NewGPU(),
			NewGPU(),
		},
		Storage: []*laptop.Storage{
			NewStorage(),
			NewStorage(),
		},
		Keyboard:    NewKeyboard(),
		Weight:      randomWeight(),
		PriceUsd:    2.2,
		ReleaseYear: 2020,
		UpdatedAt:   ptypes.TimestampNow(),
	}

	return laptop
}
