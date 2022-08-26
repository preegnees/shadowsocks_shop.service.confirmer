package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"

	models "confirmer/pkg/models"
)

type Mongo struct {
	config models.CnfDB
}

var collection *mongo.Collection
var ctxWork = context.Background()

func NewMongo(cnf models.CnfDB) (*Mongo, error) {
	ctxInit, cancel := context.WithCancel(context.Background())
	log.Printf("Вызов функции NewMongo с параметрами конфигурации: %v\n", cnf)
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://localhost:%s/", string(cnf.Port)))
	chOk := make(chan struct{})
	go func(ch chan struct{}) {
		for {
			select {
			case <- time.After(time.Duration(cnf.TimeConnect) * time.Second):
				cancel()
			case <- ch:
				break
			}
		}
	}(chOk)
	client, err := mongo.Connect(ctxInit, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("$Ошибка при подключении к MonogDB, err:=%v", err)
	}
	log.Println("Успешное создание клиента Монго")
	err = client.Ping(ctxInit, nil)
  	if err != nil {
    	return nil, fmt.Errorf("$Ошибка при пинге MongoDB, err:=%v", err)
  	}
	log.Println("Успешный пинг монго")
	collection = client.Database(string(cnf.NameDB)).Collection(string(cnf.NameDB))
	chOk <- struct{}{}
	return &Mongo{}, nil
}

var _ models.IStorage = (*Mongo)(nil)

func (m *Mongo) Put(doc models.SaveToStorage) error {
	_, err := collection.InsertOne(ctxWork, doc)
	if err != nil {
		return fmt.Errorf("$Ошибка при добавилении данных, doc=%v, err:=%v", doc, err)
	}
	return nil
}

func (m *Mongo) Delete(s models.SaveToStorage) error {
	return nil
}

func (m *Mongo) Update(s models.SaveToStorage) error {
	return nil
}