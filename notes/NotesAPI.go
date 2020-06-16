package Note

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type NotesAPI struct {
	noteService NoteService
}

func ProvideNoteAPI(n NoteService) NotesAPI {
	return NotesAPI{n}
}

func (n *NotesAPI) FindAll(c *gin.Context) {
	notes := n.noteService.FindAll()
	notesDto := ToNoteDtos(notes)
	c.JSON(http.StatusOK, gin.H{"notes": notesDto})
}

func (n *NotesAPI) FindById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	note := n.noteService.FindById(uint(id))
	if note == (Note{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	noteDto := ToNoteDto(note)
	c.JSON(http.StatusOK, gin.H{"note": noteDto})
}

func (n *NotesAPI) Create(c *gin.Context) {
	var noteDto NoteDTO
	err := c.BindJSON(&noteDto)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	note := n.noteService.Create(ToNote(noteDto))
	c.JSON(http.StatusOK, gin.H{"Note": ToNoteDto(note)})
}

func (n *NotesAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	note := n.noteService.GetById(uint(id))
	if note == (Note{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	n.noteService.Delete(note)
	c.Status(http.StatusOK)
}

func (n *NotesAPI) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	originalNote := n.noteService.GetById(uint(id))
	if originalNote == (Note{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	var noteDto NoteDTO
	err := c.BindJSON(&noteDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if noteDto.Details != "" {
		originalNote.Details = noteDto.Details
	}
	if noteDto.Name != "" {
		originalNote.Name = noteDto.Name
	}
	n.noteService.Create(originalNote)

	c.Status(http.StatusOK)
}

func (n *NotesAPI) DeleteAll(context *gin.Context) {
	n.noteService.DeleteAll()

}
