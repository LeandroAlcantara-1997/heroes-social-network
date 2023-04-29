package util

import "strings"

func ChunkTextByComma(text string) []string {
	if strings.Contains(text, ",") {
		return strings.Split(text, ",")
	}

	return []string{
		text,
	}
}
