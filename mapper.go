package jsonmapper

import (
	"encoding/json"
	"errors"
)

// Struct for JSON input
type InputJSON struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

// Struct for JSON output
type OutputJSON struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// MapInputToOutput maps an InputJSON object to an OutputJSON object
func MapInputToOutput(input InputJSON) (OutputJSON, error) {
	if input.FirstName == "" || input.LastName == "" {
		return OutputJSON{}, errors.New("both first name and last name are required")
	}
	return OutputJSON{
		Name: input.FirstName + " " + input.LastName,
		Age:  input.Age,
	}, nil
}

// ConvertJSONToInput converts a JSON string to an InputJSON object
func ConvertJSONToInput(jsonString string) (InputJSON, error) {
	var input InputJSON
	err := json.Unmarshal([]byte(jsonString), &input)
	if err != nil {
		return InputJSON{}, err
	}
	return input, nil
}

// ConvertOutputToJSON converts an OutputJSON object to a JSON string
func ConvertOutputToJSON(output OutputJSON) (string, error) {
	jsonBytes, err := json.Marshal(output)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
