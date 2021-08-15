package controllers

import (
	"net/http"
	"go-todo/config"
)

func StartMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":" + config.Config.Port, nil)
}

