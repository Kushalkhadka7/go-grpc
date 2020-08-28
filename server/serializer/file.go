package file

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// WriteToFile store binary data in a file.
func WriteToFile(message proto.Message, fileName string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("Cannot marshal the file %w", err)
	}

	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("Cannot write to a file %w", err)
	}

	return nil
}
