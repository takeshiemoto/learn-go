package main

import "flag"

func parseArgs() (string, string, []string) {
	logDir := flag.String(
		"logDir", "", "Log output directory (default=stderr)")
	flag.Parse()
	return *logDir, flag.Arg(0), flag.Args()[1:]
}
