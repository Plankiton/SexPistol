package api
import "gorm.io/gorm"

type IModel interface {
    Create ()
    Delete ()
    Update (Columns Dict)
}

type Model struct {
    gorm.Model
    ModelType string
}
