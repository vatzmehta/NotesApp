package Note

import "time"

type NoteDTO struct {
	Id uint
	Name string
	Details string
	CreatedAt time.Time
	UpdatedAt time.Time
}
