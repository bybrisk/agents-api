package handlers

import (
	"log"
)

type Agents struct {
 l *log.Logger
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

func NewAgents(l *log.Logger) *Agents{
	return &Agents{l}
}