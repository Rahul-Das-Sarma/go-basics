package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (n Note) Display() {
	fmt.Printf("Title: %s\n Content: %s\n\n", n.title, n.content)
}
func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("title and content are required")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
