package models

// Id
type Id string

// Label. Строка, которая генерируется при выставления щета, чтобы можно было потом узнать был ли платеж
type Label string

// IStorage. Интерфейс для хранилища
type IStorage interface {
	Put(Id, Label) error
	Delete(Id, Label) error
	Update(Id, Label) error
}

// Port. Порт для базы данных
type Port string

// NameDB. Имя базы данных.
type NameDB string

// CnfDB. Конфигурация для базы данных.
type CnfDB struct {
	Port
	NameDB
}