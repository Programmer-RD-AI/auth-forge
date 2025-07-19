package store

type BaseRepository interface {
	Create()
	Read()
	Update()
	Delete()
}
