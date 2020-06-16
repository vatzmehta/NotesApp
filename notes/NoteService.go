package Note

type NoteService struct {
	notesRepository NotesRepository
}

func ProvideNoteService(n NotesRepository) NoteService {
	return NoteService{n}
}

func (n *NoteService) FindAll() []Note {
	return n.notesRepository.FindAll()
}

func (n *NoteService) FindById(id uint) Note {
	return n.notesRepository.FindById(id)
}

func (n *NoteService) GetById(id uint) Note {
	return n.notesRepository.GetById(id)
}

func (n *NoteService) Create(note Note) Note {
	return n.notesRepository.Save(note)
}

func (n *NoteService) Delete(note Note) {
	n.notesRepository.Delete(note)
}

func (n *NoteService) DeleteAll() {
	n.notesRepository.DeleteAll()
}
