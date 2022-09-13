package chg

import (
	"fmt"
	"io"
	"strings"
)

// Item holds the change itself
type Item struct {
	Description string  `json:"description"`
	Items       []*Item `json:"items"`
}

// Render renders the change as a list item
func (i *Item) Render(w io.Writer, indent int) {
	bullet := fmt.Sprintf("%s- %s\n", strings.Repeat(" ", indent), i.Description)
	io.WriteString(w, bullet)
	for _, c := range i.Items {
		c.Render(w, indent+2)
	}
}
