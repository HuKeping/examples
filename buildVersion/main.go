package main

import (
	"flag"
	"fmt"

	"github.com/HuKeping/examples/buildVersion/version"
)

func main() {
	v := flag.Bool("v", false, "print version and exit")
	flag.Parse()

	// Print version and exit
	if *v {
		fmt.Println("APIVerison: ", version.APIVersion)
		fmt.Println("GitCommit : ", version.GitCommit)
		fmt.Println("BuildDate : ", version.BuildDate)

		return
	}

	// We only get here when flag "-v" was not provided
	fmt.Println("Hola")
	return
}
