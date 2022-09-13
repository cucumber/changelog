package chg

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemRender(t *testing.T) {
	i := Item{Description: "Item 1"}
	expected := "- Item 1\n"

	var buf bytes.Buffer
	i.Render(&buf, 0)
	result := buf.String()

	assert.Equal(t, expected, result)
}

func TestItemRenderWithChildren(t *testing.T) {
	c := &Item{Description: "a detail"}
	i := Item{Description: "Item 1", Items: []*Item{c}}
	expected := "- Item 1\n  - a detail\n"

	var buf bytes.Buffer
	i.Render(&buf, 0)
	result := buf.String()

	assert.Equal(t, expected, result)
}
