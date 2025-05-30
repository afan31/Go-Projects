package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}
	return value, nil
}

func writeFloatToFile(value float64, filename string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}
