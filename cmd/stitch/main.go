package main

import "github.com/tj/stitch"
import "github.com/tj/docopt"
import "log"
import "os"

var Version = "0.0.1"

const Usage = `
  Usage:
    stitch <file>
    stitch -h | --help
    stitch --version

  Options:
    -h, --help       output help information
    -v, --version    output version

`

func main() {
	args, err := docopt.Parse(Usage, nil, true, Version, false)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	s, err := stitch.File(args["<file>"].(string))
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	_, err = os.Stdout.WriteString(s)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
