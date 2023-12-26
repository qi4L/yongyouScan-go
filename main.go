package main

import (
	"flag"
	"fmt"
	"os"
	"yongyouScan/pkg/config"
	"yongyouScan/pkg/qi"
)

var (
	Url  string
	file string
)

func usage() {
	fmt.Println(`Usage of main.exe:
  -u url
      you target, example: 127.0.0.1
  -f targets.txt
	  Read the target from the file and test the vulnerabilities in batches
  `)
}

func main() {
	flag.StringVar(&Url, "u", "", "your target")
	flag.StringVar(&file, "f", "", "Specify batch target")
	flag.Usage = usage
	flag.Parse()

	if Url == "" && file == "" {
		usage()
		os.Exit(0)
	}

	qi.Logo()

	qi4l := config.WorkExp{
		Url: Url,
	}

	if file != "" {
		lines, err := qi.ReadLinesFromFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, line := range lines {
			qi4l.Url = line
			qi4l.YonYouScanRun()
		}
	}

	qi4l.YonYouScanRun()
}
