package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
)

func PrettyPrint(v interface{}) error {
	var buf bytes.Buffer

	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(v); err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}

	jsonStr := buf.String()

	var coloredOutput strings.Builder
	var inString bool
	var isKey bool
	var insideArray bool

	for i := 0; i < len(jsonStr); i++ {
		char := jsonStr[i]

		if char == '"' {
			inString = !inString

			if !inString {
				coloredOutput.WriteString(colorString(string(char), isKey))

				if isKey {
					isKey = false
				}
			} else {
				if !insideArray {
					for j := i - 1; j >= 0; j-- {
						if jsonStr[j] == '{' || jsonStr[j] == ',' {
							isKey = true
							break
						}
						if !unicode.IsSpace(rune(jsonStr[j])) {
							break
						}
					}
				}

				coloredOutput.WriteString(colorString(string(char), isKey))
			}
		} else if !inString {
			switch char {
			case '{', '}', '[', ']', ':', ',':
				if char == '[' {
					insideArray = true
				} else if char == ']' {
					insideArray = false
				}
				coloredOutput.WriteString(colorString(string(char), false))
			default:
				coloredOutput.WriteString(string(char))
			}
		} else {
			coloredOutput.WriteString(colorString(string(char), isKey))
		}
	}

	fmt.Println(coloredOutput.String())
	return nil
}

func colorString(str string, isKey bool) string {
	if isKey {
		return "\x1b[36m" + str + "\x1b[0m"
	} else {
		return "\x1b[33m" + str + "\x1b[0m"
	}
}

func Te() {
	fmt.Println("something")
}
