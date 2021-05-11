package models

import "errors"

// Post -> models

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	Private bool   `json:"private"`
	Author  User   `json:"author"`
	Base
}

// TableName -> TableName
func (p Post) TableName() string {
	return "post"
}

// ToMap -> covert to map
func (p Post) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         p.ID,
		"title":      p.Title,
		"body":       p.Body,
		"private":    p.Body,
		"author":     p.Author,
		"created_at": p.CreatedAt,
		"updated_at": p.UpdatedAt,
	}
}

// Validate Post data
func (p *Post) Validate() error {
	var err error
	if p.Title == "" {
		err = errors.New("required title")
		return err
	}
	if p.Body == "" {
		err = errors.New("required body")
		return err
	}
	return nil
}
