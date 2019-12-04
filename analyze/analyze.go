package analyze

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func GetShaderDetails(paths []string) []map[string]interface{} {
	details := make([]map[string]interface{}, len(paths))

	for index, p := range paths {
		details[index] = make(map[string]interface{})
		details[index] = analyze(p)
	}

	return details
}

func ConvertToJson(datas []map[string]interface{}) string {
	jsonString, err := json.Marshal(datas)
	if err != nil {
		fmt.Println(err)
	}
	return string(jsonString)
}

func analyze(path string) map[string]interface{} {
	detail := make(map[string]interface{})
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	//var passCount int = 0
	//foundShaderName := false
	// Initialize mapped params
	detail["path"] = path
	detail["renderers"] = "all"
	detail["isSurface"] = false
	//detail["isMultiPass"] = "false"
	renderersVal := 16383

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := string(scanner.Text())
		line = strings.TrimSpace(line)
		line = strings.Trim(line, string([]byte{239, 187, 191}))
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
				detail["isSurface"] = true
			}

			if strings.Contains(line, "only_renderers") {
				renderersString := line[strings.LastIndex(line, "only_renderers")+15:]
				tempRenderersVal := calcRenderers(renderersString)
				if renderersVal != tempRenderersVal {
					if tempRenderersVal < renderersVal {
						detail["renderers"] = renderersString
						renderersVal = tempRenderersVal
					}
				}
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

const (
	ord3d9    = 1 << iota // 1 	// Direct3D 9
	ord3d119x             // 2 	// Direct3D 11 9.x
	ord3d11               // 4 	// Direct3D 11/12
	orglcore              // 8 	// OpenGL 3.x / 4.x
	orgles                // 16 	// OpenGL ES 2.0
	orgles3               // 32 	// OpenGL ES 3.0
	ormetal               // 64 	// iOS/Mac Metal
	orvulkan              // 128 	// Vulkan
	orxbox360             // 256 	// Xbox 360
	orxboxone             // 512 	// xbox One
	orps4                 // 1024 // PlayStation 4
	orpsp2                // 2048 // PlayStation Vita
	orn3ds                // 4096 // Nintendo 3DS
	orwiiu                // 8192 // Nintendo Wii U
)

func calcRenderers(renderers string) int {
	val := 0
	if strings.Contains(renderers, "d3d9") {
		val += ord3d9
	}
	d3d11num := strings.Count(renderers, "d3d11")
	if d3d11num == 1 {
		if strings.Contains(renderers, "d3d11_9x") {
			val += ord3d119x
		} else {
			val += ord3d11
		}
	} else if d3d11num == 2 {
		val += ord3d11
		val += ord3d119x
	}
	if strings.Contains(renderers, "glcore") {
		val += orglcore
	}
	if strings.Contains(renderers, "gles") {
		val += orgles
	}
	if strings.Contains(renderers, "gles3") {
		val += orgles3
	}
	if strings.Contains(renderers, "metal") {
		val += ormetal
	}
	if strings.Contains(renderers, "vulkan") {
		val += orvulkan
	}
	if strings.Contains(renderers, "xbox360") {
		val += orxbox360
	}
	if strings.Contains(renderers, "xboxone") {
		val += orxboxone
	}
	if strings.Contains(renderers, "ps4") {
		val += orps4
	}
	if strings.Contains(renderers, "psp2") {
		val += orpsp2
	}
	if strings.Contains(renderers, "n3ds") {
		val += orn3ds
	}
	if strings.Contains(renderers, "wiiu") {
		val += orwiiu
	}
	return val
}
