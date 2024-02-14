package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "HeathChecker",
		Usage: "Samll tool to check websit is running or down",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "domain which need be checked",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "ip",
				Usage:    "ip which need be checked",
				Aliases:  []string{"p"},
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			port := ctx.String("port")
			if ctx.String("port") == "" {
				port = "80"
			}
			status := Check(ctx.String("domain"), port)
			fmt.Print(status)
			return nil
		},
	}
	error := app.Run(os.Args)
	if error != nil {
		log.Fatal(error)
	}
}
