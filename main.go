package main

import (
	"adm/model"
	"adm/routers"
	"flag"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"runtime"
)

func main() {
	addr := flag.String("port", ":3000", "Http Server Port")
	flag.Parse()

	runtime.GOMAXPROCS(1)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := routers.Load()
	model.Connect()
	http.ListenAndServe(*addr, r)
}
