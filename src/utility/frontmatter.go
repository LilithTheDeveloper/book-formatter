package utility

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	types "book-formatter/types"
)

type SortedFile struct {
	FilePath string
	ChapterOrder int
}

func ParseFrontmatter(content string) (map[string]string, error) {
	frontmatter := make(map[string]string)
	lines := strings.Split(content, "\n")
	inFrontmatter := false

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "---" {
			inFrontmatter = !inFrontmatter
			continue
		}
		if inFrontmatter {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				frontmatter[key] = value
			}
		}
	}

	return frontmatter, nil
}

func GetChapterOrder(frontmatter map[string]string) int {
	chapterOrderStr, ok := frontmatter["chapter_order"]
	if !ok {
		return 0 
	}

	order, err := strconv.Atoi(chapterOrderStr)
	if err != nil {
		fmt.Printf("Error converting chapter_order to int: %v\n", err)
		return 0 
	}

	return order
}

func SortFilesByChapterOrder(file_paths []string, markdownFiles []types.MarkdownFile) []types.MarkdownFile {
	sortedFiles := make([]types.MarkdownFile, len(markdownFiles))
	copy(sortedFiles, markdownFiles)

	sort.Slice(sortedFiles, func(i, j int) bool {
		return sortedFiles[i].ChapterOrder < sortedFiles[j].ChapterOrder
	})

	return sortedFiles
}	