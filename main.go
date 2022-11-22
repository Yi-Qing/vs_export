package main

import (
    "encoding/json"
	"flag"
	"fmt"
	"os"
	"vs_export/sln"
	"io/ioutil"
)

func main() {
	path := flag.String("s", "", "sln file path")
	configuration := flag.String("c", "Debug|x64",
		"Configuration, [configuration|platform], default Debug|x64")
	flag.Parse()

	if *path == "" {
		usage()
		os.Exit(1)
	}

	solution, err := sln.NewSln(*path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	cmdList, err := solution.CompileCommandsJson(*configuration)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	js, err := json.Marshal(cmdList)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", js[:])
	ioutil.WriteFile("compile_commands.json", js[:], 0644)
}


func usage() {
	echo := `usage:sln_export_compile_commands options
			 -s   path                        sln filename
           -c   configuration               project configuration,eg Debug|x64.
                                            default Debug|x64
	`
	fmt.Println(echo)
}
