package models

import "time"

type Note struct {
	NoteID    string    `gorm:"primarykey; type:uuid; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	Title     string    `json:"title,omitempty" valid:"required~Title cannot be empty"`
	Note      string    `gorm:"default:text" json:"note,omitempty" valid:"-"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

type Notes []Note

func (Note) TableName() string {
	return "notes"
}
