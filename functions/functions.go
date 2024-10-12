package functions

import (
	"os"
	"strings"
)

// this func just to read the banner file
func ReadFile(filename string) ([]string, error, bool) {
	data, err := os.ReadFile(filename)
	// handle err
	if err != nil {
		return nil, err, true
	}
	// handling if the banner file was writenf by windows
	stringdata := string(data)
	stringdata = strings.ReplaceAll(stringdata, "\r\n", "\n")

	result := strings.Split(stringdata, "\n")
	return result, nil, false
}

// this is the the traitment functions
func TraitmentData(text []string, inputText string) string {
	// Normalize newlines
	inputText = strings.ReplaceAll(inputText, "\r\n", "\r")

	for _, char := range inputText {
		if (char < 32 && char != 13) || char > 126 {
			return "our ascii do not suport speciale caracters"
		}
	}

	result := ""

	words := strings.Split(inputText, "\r")

	result = Final_result(text, words)

	return result
}

func Final_result(arrData, words []string) string {
	result := ""
	for k, word := range words {
		if word == "" && len(words) > 2 {
			result += "\r\n"
		} else {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(words[k]); j++ {
					Ascii := (int(words[k][j] - 32))

					start := Ascii*8 + Ascii + 1 + i

					result += arrData[start]

				}
				result += "\r\n"
			}
		}
	}

	return result
}
