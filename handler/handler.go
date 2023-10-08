package handler

import (
	"main/logs"
	"net/http"
)

func OpenServer() {
	logs.Fatal(http.ListenAndServe(":8080", nil))
}
