package tests

import (
	"log"
	"testing"

	assert "github.com/stretchr/testify/assert"

	models "confirmer/pkg/models"
	db "confirmer/pkg/service/db"
)

// Проверка на подключение с поднятым сервером и с верной конфигурацией
func TestStorage_TrueConf_and_UpServer(t *testing.T) {
	cnf := models.CnfDB{
		Port: "27017",
		NameDB: "TestConfirmer",
		TimeConnect: 5,
	}
	_, err := db.NewStorage(cnf)
	assert.True(t, err == nil)
}

// Проверка на подключение с поднятым сервером и с неверной конфигурацией
func TestStorage_FalseConf_and_UpServer(t *testing.T) {
	cnf := models.CnfDB{
		Port: "27015",
		NameDB: "TestConfirmer",
		TimeConnect: 5,
	}
	_, err := db.NewStorage(cnf)
	assert.True(t, err != nil)
	log.Println(err)
}

