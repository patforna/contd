package models

import (
	"time"
	"code.google.com/p/go-uuid/uuid"
)

type Pipeline struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Image     string
	Command   string
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		Id:uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Image:"java:8",
		Command:"javac -verbose Hello.java",
	}
}
