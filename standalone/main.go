//
// Created by taimiso0319 on 2019/11/07.
// Copyright (c) 2019 taimiso0319. All rights reserved.
//

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/taimiso0319/unity-shader-explorer/analyze"
	"github.com/taimiso0319/unity-shader-explorer/collect"
	"github.com/taimiso0319/unity-shader-explorer/modify"
)

func main() {
	// Start setting flags
	var (
		limit   int
		json    bool
		m       bool
		a       bool
		orToAll bool
	)

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.BoolVar(&m, "modify", false, "Modify specific file. requires path work on.")
	f.BoolVar(&orToAll, "onlyrenderer-to-all", false, "Change only renderer to all.")
	f.BoolVar(&a, "analyze", false, "Analyzing files. requires path walk from.")
	f.IntVar(&limit, "limit", 5, "Max depth for limit explorering. work on the analyze option.")
	f.BoolVar(&json, "json", false, "Print result as Json.")
	f.Parse(os.Args[1:])

	for 0 < f.NArg() {
		f.Parse(f.Args()[1:])
	}
	// End setting flags

	if len(os.Args) <= 1 {
		fmt.Println("You have to set the path working on.")
		return
	}

	var workPath = os.Args[1]
	if _, err := os.Stat(workPath); os.IsNotExist(err) {
		// path does not exist
		fmt.Printf("%s does not exist. Please make sure the path exists.\n", workPath)
		return
	}
	// modify
	if m {
		modifyShader(workPath, orToAll)
		return
	}

	// analyze
	if a {
		if limit <= 0 {
			fmt.Printf("limit can not be less than 0\n")
			return
		}
		collectShaderData(workPath, limit, json)
		return
	}

	fmt.Printf("Please set options what work for. -h to see options.")
}

func collectShaderData(path string, limit int, json bool) {
	collect.SetDepthLimit(limit)
	var shaderPaths []string = collect.GetShaderPaths(os.Args[1])
	if json {
		shaders := analyze.ConvertToJson(analyze.GetShaderDetails(shaderPaths))
		fmt.Println(shaders)
	} else {
		for _, line := range analyze.GetShaderDetails(shaderPaths) {
			fmt.Println(line)
		}
	}
}

func modifyShader(path string, toAll bool) {
	if toAll {
		modify.ToAll(path)
	} else {
		modify.AddMetal(path)
	}
}
