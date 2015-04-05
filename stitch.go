// Stitch shell scripts via `source` directives into a single combined script.
package stitch

import "path/filepath"
import "io/ioutil"
import "regexp"
import "fmt"

var (
	sources    = regexp.MustCompile(`(?m)^[ \t]*source[ \t]*([^;\n]+)`)
	whitespace = regexp.MustCompile(`\s+`)
)

// File reads `file` and stitches on `source <file>` directives.
func File(file string) (out string, err error) {
	defer capture(&err)

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	dir := filepath.Dir(file)
	in := string(b)

	out = sources.ReplaceAllStringFunc(in, func(s string) string {
		file := whitespace.Split(s, 2)[1]
		file = filepath.Join(dir, file)

		out, err := File(file)
		if err != nil {
			panic(err)
		}

		return fmt.Sprintf("\n# source %s:\n%s", file, out)
	})

	return
}

func capture(err *error) {
	if e := recover(); e != nil {
		switch e.(type) {
		case error:
			*err = e.(error)
		default:
			panic(e)
		}
	}
}
