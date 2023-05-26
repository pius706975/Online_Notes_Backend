package models

import "time"

type History struct {
	HistoryID string `gorm:"primarykey; type:uuid; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`

	Note_ID string `json:"note_id" valid:"uuidv4"`
	Note    Note   `gorm:"foreignkey:Note_ID; references:NoteID; constraint:OnUpdate:Cascade,OnDelete:CASCADE" json:"note_data" valid:"-"`

	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

type Histories []History

func (History) TableName() string {
	return "Histories"
}

