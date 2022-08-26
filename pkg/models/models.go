package models

// Id
type Id string

// Label. Строка, которая генерируется при выставления щета, чтобы можно было потом узнать был ли платеж
type Label string

// SaveToStorage. Структура для работы с базой данных
type SaveToStorage struct {
	Id `bson:"id"`
	Label `bson:"label"`
}

// IStorage. Интерфейс для хранилища
type IStorage interface {
	Put(SaveToStorage) error
	Delete(SaveToStorage) error
	Update(SaveToStorage) error
}

// Port. Порт для базы данных
type Port string

// NameDB. Имя базы данных.
type NameDB string

// TimeConnect. Время ожидания при подключении к базе данных.
type TimeConnect int8

// CnfDB. Конфигурация для базы данных.
type CnfDB struct {
	Port
	NameDB
	TimeConnect
}