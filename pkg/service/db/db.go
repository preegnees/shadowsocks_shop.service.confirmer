package db

import (
	models "confirmer/pkg/models"
	mongo "confirmer/pkg/service/db/mongo"
)

func NewStorage(cnf models.CnfDB) models.IStorage {
	return mongo.NewMongo(cnf)
}