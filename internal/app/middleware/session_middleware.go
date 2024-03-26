package middleware

import (
)

type SessionMiddlewareI interface {
}

type SessionMiddleware struct {
}

func NewSessionMiddleware() *SessionMiddleware {
	return &SessionMiddleware{}
}


