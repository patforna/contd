package models

import (
	"time"
	"code.google.com/p/go-uuid/uuid"
)

type Build struct {
	Id              string    `json:"id"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	OutputStreamUrl string    `json:"output_stream_url"`
}

func NewBuild() *Build {
	return &Build{
		Id:uuid.New(),
		Status:"CREATED",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		OutputStreamUrl:"http://...",
	}
}
