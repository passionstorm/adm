package main

import (
	"adm/app/models"
	"adm/app/pkg/view"
	"adm/cmd"
	"log"
	"runtime"
)

func init() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func main() {
	// Connect to database
	db := models.InitDB()
	defer db.Close()
	// Setup templates
	view.LoadTemplates()
	// Get started
	cmd.Execute()
}
