package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "Usage:")
		fmt.Fprintln(flag.CommandLine.Output(), "")
		fmt.Fprintln(flag.CommandLine.Output(), "  stamp -[s|m|n] timestamp")
		fmt.Fprintln(flag.CommandLine.Output(), "")
		fmt.Fprintln(flag.CommandLine.Output(), "  timestamp   An integer to parse as a date")
		flag.PrintDefaults()
	}

	secs := flag.Bool("s", true, "Parse input as seconds")
	millis := flag.Bool( "m", false, "Parse input as milliseconds")
	nanos := flag.Bool( "n", false, "Parse input as nanoseconds")

	flag.Parse()

	arg := flag.Arg(0)
	if arg == "" {
		fmt.Println("Missing argument")
		fmt.Println("")
		flag.Usage()
		os.Exit(2)
	}

	val, err := strconv.ParseInt(arg, 10, 64)
	switch {
	case err != nil:
		fmt.Println("'%s' could not be parsed as a number: %v", arg, err)
		flag.Usage()
		os.Exit(2)
	}

	var t time.Time

	switch {
	case *millis:
		t = time.UnixMilli(val)
	case *nanos:
		t = time.UnixMicro(val)
	case *secs:
		t = time.Unix(val, 0)
	}

	fmt.Println(t.Format(time.RFC3339))
}
