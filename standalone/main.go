//
// Created by taimiso0319 on 2019/11/07.
// Copyright (c) 2019 taimiso0319. All rights reserved.
//

package main

import (
	"os"
	//"flag"
	"flag"
	"fmt"

	"github.com/taimiso0319/unity-shader-explorer/collect"
)

func main() {
	// Start setting flags
	var (
		limit int
	)

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.IntVar(&limit, "limit", 5, "Max depth for limit explorering.")
	f.Parse(os.Args[1:])
	for 0 < f.NArg() {
		f.Parse(f.Args()[1:])
	}
	if limit <= 0 {
		fmt.Printf("limit can not be less than 0\n")
		return
	}
	// End setting flags

	if len(os.Args) <= 1 {
		fmt.Println("You have to set the path start walk from.")
		return
	}
	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		// path does not exist
		fmt.Printf("%s does not exist. Please make sure the path exists.\n", os.Args[1])
		return
	}

	fmt.Println(limit)
	collect.SetDepthLimit(limit)

	collect.WalkPath(os.Args[1])
}
