package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"note_title"`
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("your note titled %v has the following content:\n\n%v\n", note.Title, note.Content)
}
func (note Note) Save() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)

}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedOn: time.Now(),
	}, nil
}
