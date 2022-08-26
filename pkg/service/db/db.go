package db

import (
	"fmt"

	models "confirmer/pkg/models"
	mongo "confirmer/pkg/service/db/mongo"
)

func NewStorage(cnf models.CnfDB) (models.IStorage, error) {
	storage, err := mongo.NewMongo(cnf)
	if err != nil {
		return nil, fmt.Errorf("$Ошибка при получении хранилища, err:=%v", err)
	}
	return storage, nil
}