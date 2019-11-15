//
// Created by taimiso0319 on 2019/11/07.
// Copyright (c) 2019 taimiso0319. All rights reserved.
//

package collect

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var depthLimit int = 5
var currentDepth int = 0

// Test shows useless words as a first step.
func Test() {
	fmt.Print("test")
}

// WalkPath explorer directries recusively from the root path.
func WalkPath(root string) {
	var paths []string = dirwalk(root)
	fmt.Println(paths)
	fmt.Println(len(paths))
}

// SetDepthLimit changes the limit of directries the WalkPath explorers.
func SetDepthLimit(depth int) {
	depthLimit = depth
}

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			currentDepth++
			if currentDepth > depthLimit {
				fmt.Println("Too deep. To avoid this error if necessarily, please run the change limit method first.")
				os.Exit(0) // I think I should change exit code to notice what happends.
			}
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			currentDepth--
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}
