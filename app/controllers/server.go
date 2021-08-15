package controllers

import (
	"net/http"
	"go-todo/config"
)

func StartMainServer() error {
	return http.ListenAndServe(":" + config.Config.Port, nil)
}

