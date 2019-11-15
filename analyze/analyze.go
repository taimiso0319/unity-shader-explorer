package analyze

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func GetShaderDetails(paths []string) [][]string {
	var details [][]string
	details = make([][]string, len(paths))

	for index, p := range paths {
		var detail []string = analyze(p)
		details[index] = make([]string, len(detail))
		details[index] = detail
	}

	return details
}

func analyze(path string) []string {
	var detail []string
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	detail = append(detail, path)

	for scanner.Scan() {
		var line string = string(scanner.Text())
		fmt.Println(line)
		fmt.Println(utf8.RuneCountInString(line))
		strings.TrimLeft(line, string("0x20"))
		fmt.Println(line)
		fmt.Println(utf8.RuneCountInString(line))
		if strings.HasPrefix(line, "Shader") {
			detail = append(detail, line[strings.Index(line, "\""):strings.LastIndex(line, "\"")+1])
			break
		}
	}
	return detail
}
