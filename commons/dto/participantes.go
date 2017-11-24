package dto

import "gopkg.in/mgo.v2/bson"

type Participante struct {
	ID   	bson.ObjectId 	`bson:"_id" json:"id"`
	Nome 	string 			`json:"nome"`
	Cpf  	string 			`json:"cpf"`
}

type Participantes []Participante
