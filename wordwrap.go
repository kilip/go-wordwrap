package wordwrap

import (
	"strings"
)

// Wrap returns wrapped string to a given number of characters using "\n" character
func Wrap(text string, width uint) string {
	breakChar := "\n"
	return WrapF(text, width, breakChar, false)
}

// WrapF wraps a given text with full options
func WrapF(text string, width uint, breakChar string, cutLongWord bool) string {
	if len(text) == 0 {
		return ""
	}

	if len(breakChar) == 0 {
		panic("breakChar can't be an empty string")
	}

	if width == 0 && cutLongWord {
		panic("cutLongWords can't be true when width is 0")
	}

	if len(breakChar) == 1 && !cutLongWord {
		return doWrap(text, width, breakChar)
	}
	return doWrapWithCutLongWords(text, width, breakChar)
}

func doWrapWithCutLongWords(text string, width uint, char string) string {
	return ""
}

func doWrap(text string, width uint, breakChar string) string {
	var lastStart, lastSpace, current uint
	var lines []string
	var line string
	lastStart = 0
	lastSpace = 0

	for current = 0; current < uint(len(text)); current++ {
		cText := string(text[current])
		if cText != breakChar {
			line += cText
		}
		if cText == breakChar {
			lastStart = current + 1
			lastSpace = current + 1
			if line != "" && line != breakChar {
				lines = append(lines, line)
				line = ""
			}
			lines = append(lines, "")
		} else if cText == " " {
			if (current - lastStart) >= width {
				//sumLen := width - uint(len(line))
				if line != "" {
					lines = append(lines, text[lastStart:current+1])
					line = ""
				}
				lastStart = current + 1
			}
			lastSpace = current
		} else if current-lastStart >= width && lastStart != lastSpace {
			lines = append(lines, text[lastStart:lastSpace+1])
			line = text[lastSpace+1 : current+1]
			lastStart = lastSpace + 1
		} else if current-lastStart >= width {
			// long word cases
			lines = append(lines, text[lastStart:current-1])
			line = text[current-1 : current+1]
			lastStart = current
		}

	}

	if line != "" {
		lines = append(lines, line)
	}
	return strings.Join(lines, breakChar)
}
