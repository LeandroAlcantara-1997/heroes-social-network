package util

import (
	"strings"
	"time"
)

func ChunkTextByComma(text string) []string {
	if strings.Contains(text, ",") {
		return strings.Split(text, ",")
	}

	return []string{
		text,
	}
}

func GerPointer[T time.Time](value T) *T {
	return &value
}
