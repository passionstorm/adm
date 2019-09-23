package main

import (
	"adm/model"
	"adm/routers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := routers.Load()
	model.Connect()
	http.ListenAndServe(":3000", r)
}
