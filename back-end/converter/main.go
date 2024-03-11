package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

func main() {

	dir, _ := os.Getwd()
	inputFolder := filepath.Join(dir, "..", "..", "front-end", "src", "models")
	outputFolder := filepath.Join(dir, "..", "models")

	err := filepath.Walk(inputFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return nil
		}
		if !info.IsDir() && filepath.Ext(path) == ".ts" {
			outputFilename := filepath.Join(outputFolder, strings.TrimSuffix(info.Name(), ".ts")+".go")
			convertFile(path, outputFilename)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error scanning input folder: %v\n", err)
		return
	}

	fmt.Println("Conversion completed successfully.")
}

func convertFile(inputFilename, outputFilename string) {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		fmt.Printf("Error opening input file %s: %v\n", inputFilename, err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		fmt.Printf("Error creating output file %s: %v\n", outputFilename, err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var lines []string

	lines = append(lines, "package models\n")

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "interface") {
			lines = append(lines, processInterface(line))
		} else if strings.Contains(line, "import") {
			continue
		} else if strings.Contains(line, "}") {
			lines = append(lines, line)
		} else {
			pattern := `\b\d+\b|:|number|Record<string,|Set<string>|>`
			re := regexp.MustCompile(pattern)
			result := re.ReplaceAllStringFunc(line, func(match string) string {
				if strings.Contains(match, ":") {
					return ""
				}
				if strings.ToLower(match) == "number" {
					return "float64"
				}
				if match == "Record<string," {
					return "map[string]"
				}
				if match == "Set<string>" {
					return "map[string]struct{}"
				}

				return ""
			})

			lines = append(lines, capitalizeFirstCharacter(result))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning input file %s: %v\n", inputFilename, err)
		return
	}

	output := strings.Join(lines, "\n")
	_, err = outputFile.WriteString(output)
	if err != nil {
		fmt.Printf("Error writing to output file %s: %v\n", outputFilename, err)
		return
	}
}

func processInterface(line string) string {
	parts := strings.Fields(line)
	interfaceName := parts[2]

	structDeclaration := "type " + interfaceName + " struct {"

	return structDeclaration
}

func capitalizeFirstCharacter(s string) string {
	for i, char := range s {
		if !unicode.IsSpace(char) {

			spaces := ""
			for range i {
				spaces += " "
			}
			return spaces + string(unicode.ToUpper(char)) + s[i+1:]
		}
	}
	return s
}
