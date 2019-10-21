package main

import (
	"adm/pkg/config"
	"adm/pkg/db"
	"adm/routers"
	"flag"
	"net/http"
	"runtime"
)

func main() {
	addr := flag.String("port", ":3000", "Http Server Port")
	flag.Parse()
	runtime.GOMAXPROCS(1)
	config.Load()
	db.GetConnection()
	r := routers.Load()
	//model.Connect()
	http.ListenAndServe(*addr, r)
}
