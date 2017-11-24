package commons

import (
	"log"

	"gopkg.in/mgo.v2"
)

type DataStore struct {
	Session *mgo.Session
}

func GetDataStore(url string) *DataStore {
	session, err := mgo.Dial(url)

	if err != nil {
		log.Fatal(err)
	}

	return &DataStore{Session:session}
}

func (ds *DataStore) GetCollection(database, collection string) *mgo.Collection {
	return ds.Session.DB(database).C(collection)
}
