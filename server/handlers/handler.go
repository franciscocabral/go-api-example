package handlers

type Handler interface {
	Get()
	Post()
	Put()
	Delete()
	Default()
}
