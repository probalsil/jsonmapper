package jsonmapper_test

import (
	"testing"

	"jsonmapper"
)

func TestMapInputToOutput(t *testing.T) {
	input := jsonmapper.InputJSON{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	expectedOutput := jsonmapper.OutputJSON{
		Name: "John Doe",
		Age:  30,
	}
	output, err := jsonmapper.MapInputToOutput(input)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if output != expectedOutput {
		t.Errorf("got %v, want %v", output, expectedOutput)
	}
}
