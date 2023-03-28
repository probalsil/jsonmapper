package jsonmapper

import (
	"fmt"
	//"jsonmapper"
)

func main() {
	input := InputJSON{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	output, err := MapInputToOutput(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonString, err := ConvertOutputToJSON(output)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output JSON:", jsonString)
}
