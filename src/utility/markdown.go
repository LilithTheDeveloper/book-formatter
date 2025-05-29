package utility

import (
	// "os"
	types "book-formatter/types"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func PreprocessMarkdown(content string) string {
	content = RemoveFrontmatter(content)
	content = AppendCustomDiv(content)
	content = ReplaceHorizontalRules(content) 
	content = ReplaceLinks(content)
	content = ReplaceAuthorQuoteBlocks(content)

	return content
}

func MergeMarkdownFiles(files []types.MarkdownFile, generateStatistics bool) string {
	var mergedContent strings.Builder

	for _, file := range files {
		content, err := ReadFile(file.FilePath)
		if err != nil {
			continue // Handle error as needed
		}

		// Extract the file name without the path
		fileName := filepath.Base(file.FilePath)

		// Remove file ending 
		fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

		// Preprocess the content
		preprocessedContent := PreprocessMarkdown(content)

		// Append the preprocessed content to the merged content
		mergedContent.WriteString("\n# " + fileName + "\n\n")
		mergedContent.WriteString(preprocessedContent)
		mergedContent.WriteString("\n")
		mergedContent.WriteString(`\newpage`) 
		mergedContent.WriteString("\n")
	}

	if generateStatistics {
		stats := GenerateStatistics(mergedContent.String())
		mergedContent.WriteString("\n# Statistics\n\n")
		mergedContent.WriteString("Lines: " + strconv.Itoa(stats.Lines) + "\n\n")
		mergedContent.WriteString("Words: " + strconv.Itoa(stats.Words) + "\n\n")
		mergedContent.WriteString("Characters: " + strconv.Itoa(stats.Characters) + "\n\n")
	}

	return mergedContent.String()
}

func RemoveFrontmatter(content string) string {
	lines := strings.Split(content, "\n")
	inFrontmatter := false
	var result []string

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Start of frontmatter
		if i == 0 && trimmed == "---" {
			inFrontmatter = true
			continue
		}

		// End of frontmatter
		if trimmed == "---" && inFrontmatter {
			inFrontmatter = false
			continue
		}

		if !inFrontmatter {
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}

func ReplaceHorizontalRules(content string) string {
	content = strings.ReplaceAll(content, "---", `\rule{\linewidth}{0.4pt}`)
	return content
}


// Replace special quote blocks with a custom format.
func ReplaceAuthorQuoteBlocks(content string) string {
	re := regexp.MustCompile(`(?m)^>\s*\[!\s*quote\s*\|\s*([^\]]+)\]\s*(.*)$`)
	content = re.ReplaceAllStringFunc(content, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) < 3 {
			return match // No match, return original
		}
		title := strings.TrimSpace(parts[2])
		return "> **" + title + "**\n\n"
	})
	return content
}

// Removes links from the content and replaces them with their text content.
// e.g., 
// "[link text](http://example.com)" becomes "link text".
// "[[link text]]" becomes "link text".
// "[[link text|display text]]" becomes "display text".
func ReplaceLinks(content string) string {
	reMarkdown := regexp.MustCompile(`\[(.*?)\]\(.*?\)`)
	content = reMarkdown.ReplaceAllString(content, `$1`)

	reDoubleBrackets := regexp.MustCompile(`\[\[(.*?)(\|(.*?))?\]\]`)
	content = reDoubleBrackets.ReplaceAllStringFunc(content, func(match string) string {
		parts := strings.SplitN(match[2:len(match)-2], "|", 2)
		if len(parts) == 2 {
			return parts[1]
		}
		return parts[0]
	})

	return content
}

// Appends pandoc custom divs to the content when the pattern is matched.
// Custom divs are defined in the markdown content as follows:
// ---
// > [custom_div]
// > content
// Regular Text
func AppendCustomDiv(content string) string {
	lines := strings.Split(content, "\n")
	var result []string
	inCustomDiv := false

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		// Detect start of custom div: "---" followed by ">" line
		if !inCustomDiv && line == "---" && i+1 < len(lines) && strings.HasPrefix(strings.TrimSpace(lines[i+1]), ">") {
			result = append(result, "::: {.infobox}")
			inCustomDiv = true
			continue
		}

		if inCustomDiv {
			// End custom div on "---"
			if line == "---" {
				result = append(result, ":::")
				inCustomDiv = false
				continue
			}
			// Only process lines starting with ">"
			if strings.HasPrefix(line, ">") {
				contentLine := strings.TrimSpace(line[1:])
				if contentLine != "" {
					result = append(result, contentLine)
				} else if line == ">" {
					result = append(result, "") 
				}
				continue
			}
			// If line doesn't start with ">", close div and process line normally
			result = append(result, ":::")
			inCustomDiv = false
		}
		result = append(result, lines[i])
	}

	// Close any unclosed div at EOF
	if inCustomDiv {
		result = append(result, ":::")
	}

	return strings.Join(result, "\n")
}

func GenerateStatistics(content string) types.MarkdownFileStatistics { 
	var stats = types.MarkdownFileStatistics{
		Lines:       0,
		Words:      0,
		Characters: 0,
	}

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		stats.Lines++
		stats.Characters += len(line)
		words := strings.Fields(line)
		stats.Words += len(words)
	}

	return stats
}