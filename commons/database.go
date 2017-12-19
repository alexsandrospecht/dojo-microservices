package commons

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func (ds *DataStore) FindAll(database, collectionName string, query, result interface{}) {
	collection := ds.GetCollection(database, collectionName)
	collection.Find(query).All(result)
}

func (ds *DataStore) FindOne(database, collectionName string, query, result interface{}) {
	collection := ds.GetCollection(database, collectionName)
	collection.FindId(query).One(result)
}

func (ds *DataStore) Delete(database, collectionName string, id bson.ObjectId) {
	collection := ds.GetCollection(database, collectionName)
	collection.RemoveId(id)
}