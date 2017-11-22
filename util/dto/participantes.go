package dto

type Participante struct {
	Nome string `json:"nome"`
	Cpf  string `json:"cpf"`
}

type Participantes []Participante
