package main

import (
	"fmt"

	"github.com/alecthomas/kong"

	_ "github.com/digaxfr/go-sandbox/internal/dg"
)

type Globals struct {
	Debug bool `kong:"help='Enable debug logging',optional"`
}

var cli struct {
	Globals

	SSH SshCmd `cmd:"" help:"SSH subcommand"`
}

func main() {
	ctx := kong.Parse(&cli)

	fmt.Printf("%v\n", ctx)

	err := ctx.Run()
	if err != nil {
		fmt.Println(err)
	}
}
