// main package for the gogpt
package main

import (
	"gogpt/cmd"
)

var (
	version = "local"
	commit  = "n/a"
)

func main() {
	cmd.Execute(version, commit)
}
