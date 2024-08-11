package main

import (
	"Sentinel/lib"
	"fmt"
	"os"
)

func main() {
	args, err := lib.CliParser()
	if err != nil {
		fmt.Println(err)
		return
	}
	httpClient := lib.HttpClientInit(args.Timeout)
	localVersion := lib.GetCurrentLocalVersion()
	repoVersion := lib.GetCurrentRepoVersion(httpClient)
	fmt.Printf(" ===[ Sentinel, Version: %s ]===\n\n", localVersion)
	lib.VersionCompare(repoVersion, localVersion)
	if err := lib.CreateOutputDir(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	lib.DisplayCount = 0
	fmt.Print("[*] Method: ")
	if len(args.WordlistPath) == 0 {
		fmt.Println("PASSIVE")
		lib.PassiveEnum(&args, httpClient)
	} else {
		fmt.Println("DIRECT")
		if err := lib.DirectEnum(&args, httpClient); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}
}
