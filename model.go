package api

type IModel interface {
    Create ()
    Delete ()
    Update (Columns Dict)
}
