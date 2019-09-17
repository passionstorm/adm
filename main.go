package main

import (
	"adm/model"
	"adm/pkg/setting"
	"adm/routers"
	"net/http"
)

func main() {
	setting.Load()
	r := routers.Load()
	model.ConnectDB()
	model.AutoMigrate()

	http.ListenAndServe(":3000", r)
}
