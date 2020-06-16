package Note

import "github.com/jinzhu/gorm"

type NotesRepository struct {
	DB *gorm.DB
}

func ProvideNotesRepository(DB *gorm.DB) NotesRepository {
	return NotesRepository{DB: DB}
}

func (n *NotesRepository ) FindAll() []Note{
	var notes []Note
	n.DB.Find(&notes)
	return notes
}

func (n *NotesRepository) FindById(id uint) Note{
	var note Note
	n.DB.First(&note,id)
	return note
}

func (n *NotesRepository) Save(note Note) Note{
	n.DB.Save(&note)
	return note
}

func (n *NotesRepository) Delete(note Note){
	n.DB.Delete(&note)
}

func (n *NotesRepository) DeleteAll() {
	n.DB.Delete(&Note{})
}

func (n *NotesRepository) GetById(id uint) Note {
	var note Note
	n.DB.First(&note,id)
	return note

}
