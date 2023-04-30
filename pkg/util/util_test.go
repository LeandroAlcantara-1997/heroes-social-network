package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sliceUseCase = "IronMan,Thor,Hulk,Capitain America"
)

func TestChunkTextByCommaWithComma(t *testing.T) {
	out := ChunkTextByComma(sliceUseCase)
	assert.Equal(t, []string{
		"IronMan",
		"Thor",
		"Hulk",
		"Capitain America",
	}, out)
}

func TestChunkTextByCommaWithoutComma(t *testing.T) {
	out := ChunkTextByComma("IronMan")
	assert.Equal(t, []string{
		"IronMan",
	}, out)

}
