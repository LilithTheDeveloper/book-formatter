package main

import (
	"context"
	"log"
	"os"
	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v3"
	commands "book-formatter/commands"
	types "book-formatter/types"
)


func main() {}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var defaultConfig = types.Config{
		InputDir:          "./input",
		OutputDir:         "./output",
		OutputFormat:      "pdf",
		GenerateSingleFile: true,
		GenerateStatistics: true,
		IgnoredFiles: 	[]string{"Chapters.md"},
		IgnoredDirs: 	[]string{"Generated", "output", "input", "build"},
		PandocPdfEngine: "pdflatex",
	}

	cfg , err := load_config("../config.yaml")
	if err != nil {
		log.Printf("Error loading config file: \"%v\"\nLoading default configuration...", err)
		cfg = &defaultConfig
	} else {
		log.Printf("Configuration loaded successfully: %+v", cfg)
	}

	load_cli(cfg)
}

func load_config(path string) (*types.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg types.Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func load_cli(config *types.Config) {
	// Initialize CLI command with flags
	cmd := &cli.Command{
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{
				Name:  "convert",
				Description: "Convert files from input directory to output format",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "input",
						Value: config.InputDir,
						Usage: "Iput dir",
						Aliases: []string{"i"},
						Destination: &config.InputDir,
					},
					&cli.StringFlag{
						Name:  "output",
						Value: config.OutputDir,
						Usage: "Output dir",
						Aliases: []string{"o"},
						Destination: &config.OutputDir,

					},
					&cli.StringFlag{
						Name:  "format",
						Value: config.OutputFormat,
						Usage: "Output format",
						Aliases: []string{"f"},
						Destination: &config.OutputFormat,
					},
					&cli.BoolFlag{
						Name:  "single_file",
						Value: config.GenerateSingleFile,
						Usage: "Gen single file?",
						Aliases: []string{"s"},
						Destination: &config.GenerateSingleFile,
					},
					&cli.BoolFlag{
						Name:  "generate_stats",
						Value: config.GenerateStatistics,
						Usage: "Gen statistics?",
						Aliases: []string{"g"},
						Destination: &config.GenerateStatistics,
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if config.InputDir == "" || config.OutputDir == "" {
						return cli.Exit("Input and output directories must be specified", 1)
					}

					commands.Convert(
						config.InputDir,
						config.OutputDir,
						config.OutputFormat,
						config.GenerateSingleFile,
						config.GenerateStatistics,
						config.IgnoredFiles,
						config.IgnoredDirs,
						config.LuaFilePath,
						config.PreambleFilePath,
						config.PandocPdfEngine,
					)
					return nil
				},
			},	
		},
    }

    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}
