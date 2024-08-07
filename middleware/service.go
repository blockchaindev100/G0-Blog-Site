package middleware

import "github.com/sirupsen/logrus"

type Middleware struct {
	Logger *logrus.Logger
}

func AcquireMiddleware(log *logrus.Logger) *Middleware {
	return &Middleware{Logger: log}
}
