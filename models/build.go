package models

import (
	"fmt"
	"time"
	"code.google.com/p/go-uuid/uuid"
)

type Build struct {
	Id        string    `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	LogUrl    string    `json:"log_url"`
}

func NewBuild() *Build {
	id := uuid.New()
	return &Build{
		Id:id,
		Status:"CREATED",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		LogUrl: fmt.Sprintf("http://localhost:8080/log/%s", id),
	}
}
