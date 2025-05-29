package commands

import (
	"log"
	util "book-formatter/utility"
	types "book-formatter/types"
)

func Convert(
	inputDir, outputDir, outputFormat string,
	singleFile, generateStats bool,
	ignoredFiles []string,
	ignoredDirs []string,
	luaFilePath, preambleFilePath string,
) error {
	log.Printf("Processing files in %s, outputting to %s in %s format", inputDir, outputDir, outputFormat)
	file_paths, err := util.GetFiles(inputDir, ignoredFiles)
	if err != nil {
		log.Printf("Error retrieving files from %s: %v", inputDir, err)
	}
	
	var markdownFiles []types.MarkdownFile

	for _, file_path := range file_paths {
		file, err := util.ReadFile(file_path)
		if err != nil {
			log.Printf("Error reading file %s: %v", file_path, err)
			continue
		}
		fm, err := util.ParseFrontmatter(file)
		if err != nil {
			log.Printf("Error parsing frontmatter for file %s: %v", file, err)
			continue
		}

		markdownFiles = append(markdownFiles, types.MarkdownFile{
			FilePath:    file_path,
			Frontmatter: fm,
			ChapterOrder: util.GetChapterOrder(fm),
		})
	}	

	markdownFiles = util.SortFilesByChapterOrder(file_paths, markdownFiles)
	
	if singleFile {
		content := util.MergeMarkdownFiles(markdownFiles, generateStats)

		path := "./cache/output.md"

		err = util.WriteFile(path, content)
		if err != nil {
			log.Printf("Error writing merged markdown file: %v", err)
		}

		err = util.GenerateSinglePDF(path, outputDir, "output", outputFormat, preambleFilePath, luaFilePath)
		if err != nil {
			log.Printf("Error generating PDF from merged markdown: %v", err)
		}

		return nil
	} else {
		log.Printf("Generating multiple output files in %s format", outputFormat)

		for _, file := range markdownFiles {
			content, err := util.ReadFile(file.FilePath)
			if err != nil {
				log.Printf("Error reading file %s: %v", file.FilePath, err)
				continue
			}

			preprocessedContent := util.PreprocessMarkdown(content)
			outputFilePath := outputDir + "/" + file.FilePath[len(inputDir):]
			err = util.WriteFile(outputFilePath, preprocessedContent)
			if err != nil {
				log.Printf("Error writing preprocessed file %s: %v", outputFilePath, err)
				continue
			}

			log.Printf("Processed and wrote file: %s", outputFilePath)
		}
	}


	if err != nil {
		log.Printf("Error preprocessing markdown files: %v", err)
	}

	return nil
}