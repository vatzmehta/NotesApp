package Note

func ToNote(dto NoteDTO) Note{
	return Note{Name: dto.Name,Details: dto.Details}
}

func ToNoteDto(note Note) NoteDTO{
	return NoteDTO{note.ID,note.Name,note.Details,note.CreatedAt,note.UpdatedAt}
}

func ToNoteDtos(notes []Note) []NoteDTO{
	notesDto := make([]NoteDTO, len(notes))

	for i,v := range notes {
		notesDto[i] = ToNoteDto(v)
	}
	return notesDto
}