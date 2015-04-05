# stitch

Stitch shell scripts via `source` directives into a single combined script.

## Installation

  Binary releases are available, or run:

```
$ go get github.com/tj/stitch/cmd/stitch
```

## Usage

#### func  File

```go
func File(file string) (out string, err error)
```
File reads `file` and stitches on `source <file>` directives.

## License

MIT