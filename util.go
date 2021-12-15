package main

type Repository struct {
	Name    string
	Content interface{}
}

func (r *Repository) ContainsBytes() bool {
	_, ok := r.Content.([]byte)
	return ok
}
