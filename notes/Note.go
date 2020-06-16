package Note

import (
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	Name      string
	Details   string
}
