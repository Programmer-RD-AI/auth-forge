package model

type BaseStore interface {
	Connect()
	Create()
	Read()
	Update()
	Delete()
	Close()
}
