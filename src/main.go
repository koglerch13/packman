package main

import (
	"os"
	"packman/common"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "pack-man",
		Usage:   "Generic package loader",
		Version: "0.1.0",
		Commands: []*cli.Command{
			{
				Name:      "load",
				Aliases:   []string{"l"},
				Usage:     "Loads and extracts the packages in the config file. A path to a config file can be supplied. If there is no explicit path, it defaults to './packman-config.xml'",
				UsageText: "pack-man load [config-file]",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "verbose", Usage: "Adds extra output"},
				},
				Action: loadAction,
			},
			{
				Name:    "cache",
				Aliases: []string{"c"},
				Usage:   "Provides cache operations (cache clear)",
				Subcommands: []*cli.Command{
					{
						Name:   "clear",
						Usage:  "Clears the cache directory.",
						Action: cacheClearAction,
					},
				},
			},
		},
	}

	error := app.Run(os.Args)
	common.ExitOnError("Cannot start pack-man", error)
}

func loadAction(cliContext *cli.Context) error {
	const defaultConfigFilePath = "./packman-config.xml"

	args := cliContext.Args()
	configFilePath := args.First()
	if configFilePath == "" {
		configFilePath = defaultConfigFilePath
	}

	common.Verbose = cliContext.Bool("verbose")

	Load(configFilePath)

	return nil
}

func cacheClearAction(cliContext *cli.Context) error {
	ClearCache()
	return nil
}
