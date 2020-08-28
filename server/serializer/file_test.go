package file_test

import (
	"appserver/sample"
	file "appserver/serializer"
	"fmt"
	"testing"
)

func TestFile(t *testing.T) {
	t.Parallel()

	fileName := "../temp/laptop.bin"
	laptop := sample.NewLaptop()

	err := file.WriteToFile(laptop, fileName)
	if err != nil {
		fmt.Println("test failed")

		return
	}

	fmt.Println("Test passed")

}
