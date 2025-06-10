package utility

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)


func GenerateSinglePDF(filePath, outputPath, outputFileName, outputFormat, preambleFile, luaFilterFile string, pdfEngine string) error {
	return RunPandoc(filePath, outputPath, outputFileName, outputFormat, preambleFile, luaFilterFile, pdfEngine)
}

func RunPandoc(inputFile, outputPath, outputFileName, outputFormat, preambleFile string, luaFilterFile string, pdfEngine string) error {
	if inputFile == "" || outputPath == "" || outputFileName == "" || outputFormat == "" {
		return fmt.Errorf("input file, output path, output file name, and output format must be provided")
	}

	if preambleFile == "" {
		return fmt.Errorf("preamble file must be provided")
	}

	if luaFilterFile == "" {
		return fmt.Errorf("lua filter file must be provided")
	}

	
	// Build the full command before running it
	var cmdArgs []string

	absInput, err := filepath.Abs(inputFile)
	if err != nil {
		return fmt.Errorf("error getting absolute path of input file: %w", err)
	}
	absHeader, err := filepath.Abs(preambleFile)
	if err != nil {
		return fmt.Errorf("error getting absolute path of preamble file: %w", err)
	}
	absLuaFilter, err := filepath.Abs(luaFilterFile)
	if err != nil {
		return fmt.Errorf("error getting absolute path of lua filter file: %w", err)
	}
	absOut := filepath.Join(outputPath, outputFileName+"."+outputFormat)

	cmdArgs = append(cmdArgs, absInput) 
	cmdArgs = append(cmdArgs, "-o", absOut) // Use absolute path for output
	cmdArgs = append(cmdArgs, "--pdf-engine="+pdfEngine)
	cmdArgs = append(cmdArgs, "--include-in-header="+absHeader)
	cmdArgs = append(cmdArgs, "--lua-filter="+absLuaFilter)
	cmdArgs = append(cmdArgs, "--from", "markdown+native_divs", "--verbose")

	log.Print("Running pandoc with arguments: ", cmdArgs)

	// Construct the pandoc command
	cmd := exec.Command("pandoc", cmdArgs...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()


	fmt.Println("Pandoc command output:", stdout.String())
	

	if err != nil {
		fmt.Println("Pandoc command failed with error:", err)
		return fmt.Errorf("error running pandoc: %w", err)
	}
	return nil
}