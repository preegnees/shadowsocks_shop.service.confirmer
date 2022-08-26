package mongo

import (
	models "confirmer/pkg/models"
)

type Mongo struct {
	config models.CnfDB
}

func NewMongo(cnf models.CnfDB) *Mongo {
	var mongo = Mongo{
		config: cnf,
	}
	// тут нужно получить клиент и тд
	return &mongo
}

var _ models.IStorage = (*Mongo)(nil)

func (m *Mongo) Put(id models.Id, label models.Label) error 

func (m *Mongo) Delete(id models.Id, label models.Label) error

func (m *Mongo) Update(id models.Id, label models.Label) error