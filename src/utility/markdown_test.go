package utility

import (
	"os"
	"testing"
)

func TestPreprocessMarkdown(t *testing.T) {
	content := `---
title: Test Document
author: John Doe
chapter_order: 2
---
This is a test document.
It has some content.
---
This is a horizontal rule:
---
This is a link: [link text](http://example.com)
This is a double bracket link: [[link text|display text]]
This is another double bracket link: [[link text]]
---
>
>This section is a special block for custom divs.
>
>Still with some content and [[link text|display text]].
Now more regular content.
`

	expected := `This is a test document.
It has some content.
\rule{\linewidth}{0.4pt}
This is a horizontal rule:
\rule{\linewidth}{0.4pt}
This is a link: link text
This is a double bracket link: display text
This is another double bracket link: link text
::: {.infobox}

This section is a special block for custom divs.

Still with some content and display text.
:::
Now more regular content.
`

	preprocessed := PreprocessMarkdown(content)
	if preprocessed != expected {
		t.Errorf("Expected:\n%q\nGot:\n%q", expected, preprocessed)
		t := t // capture range variable
		err := os.WriteFile("expected.txt", []byte(expected), 0644)
		if err != nil {
			t.Errorf("Failed to write expected.txt: %v", err)
		}
		err = os.WriteFile("got.txt", []byte(preprocessed), 0644)
		if err != nil {
			t.Errorf("Failed to write got.txt: %v", err)
		}
	}
}


func TestAppendCustomDiv(t *testing.T){
	content := `---
> Some text
> This is a custom div content.
>This is regular content.
This is another line of regular content.` 

	expected := `::: {.infobox}
Some text
This is a custom div content.
This is regular content.
:::
This is another line of regular content.`

	preprocessed := AppendCustomDiv(content)
	if preprocessed != expected {
		t.Errorf("Expected:\n%q\nGot:\n%q", expected, preprocessed)
		t := t // capture range variable
		err := os.WriteFile("expected_custom_div.txt", []byte(expected), 0644)
		if err != nil {
			t.Errorf("Failed to write expected_custom_div.txt: %v", err)
		}
		err = os.WriteFile("got_custom_div.txt", []byte(preprocessed), 0644)
		if err != nil {
			t.Errorf("Failed to write got_custom_div.txt: %v", err)
		}
	}
}