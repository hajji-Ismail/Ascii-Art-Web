package functions

import (
	"fmt"
	"os"
	"strings"
)

// this func just to read the banner file
func ReadFile(filename string) []string {
	data, err := os.ReadFile(filename)
	// handle err
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(0)
	}
	// handling if the banner file was writenf by windows
	stringdata := string(data)
	stringdata = strings.ReplaceAll(stringdata, "\r\n", "\n")

	result := strings.Split(stringdata, "\n")
	return result
}

// this is the the traitment functions
func TraitmentData(text []string, inputText string) string {
	// Normalize newlines
	inputText = strings.ReplaceAll(inputText, "\r\n", "\r")

	for _, char := range inputText {
		if (char < 32 && char != 13) || char > 126 {

			fmt.Println(char)
			return " Error : one of this charachter not in range "
		}
	}
	inputText = strings.ReplaceAll(inputText, "\r", "\r\n")
	result := ""

	words := strings.Split(inputText, "\r\n")

	result = Final_result(text, words)

	return result
}

func Final_result(arrData, words []string) string {
	result := ""
	for k := 0; k < len(words); k++ {
		
		for i := 0; i < 8; i++ {
			for j := 0; j < len(words[k]); j++ {
				Ascii := (int(words[k][j] - 32))

				start := Ascii*8 + Ascii + 1 + i

				result += arrData[start]

			}
			result += "\n"
		}
	}

	return result
}
