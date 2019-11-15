package analyze

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetShaderDetails(paths []string) map[int]map[string]string {
	var details = make(map[int]map[string]string)

	for index, p := range paths {
		details[index] = make(map[string]string)
		details[index] = analyze(p)
	}

	return details
}

func analyze(path string) map[string]string {
	var detail = make(map[string]string)
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	//var passCount int = 0

	// Initialize mapped params
	detail["path"] = path
	detail["renderers"] = "all"
	detail["isSurface"] = "false"
	//detail["isMultiPass"] = "false"

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var line string = string(scanner.Text())
		line = strings.TrimSpace(line)
		//fmt.Println(strings.Index(line, "Shader"))
		//fmt.Println(len(line))
		if strings.HasPrefix(line, "//") {
			continue
		}

		if strings.Contains(line, "//") {
			line = line[:strings.Index(line, "//")]
		}
		if strings.HasPrefix(line, "Shader") {
			detail["name"] = line[strings.Index(line, "\"")+1 : strings.LastIndex(line, "\"")]
			continue
		}
		if strings.HasPrefix(line, "#pragma") {
			if strings.Contains(line, "surface") {
				detail["isSurface"] = "true"
			}

			if strings.Contains(line, "only_renderers") {
				detail["renderers"] = line[strings.LastIndex(line, "only_renderers")+15:]
			}
		}
		//if strings.HasPrefix(line, "Pass") {
		//	passCount++
		//}
	}
	//if passCount > 1 {
	//	detail["isMultiPass"] = "true"
	//}
	return detail
}
