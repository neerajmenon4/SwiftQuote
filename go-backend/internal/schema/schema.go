package schema

import (
	"io/ioutil"
)

// ReadSchemaCSV reads the entire contents of a CSV file and returns it as a string.
func ReadSchemaCSV(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
