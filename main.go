package main

import (
	"adm/app/models"
	"adm/app/pkg/view"
	"adm/cmd"
	"runtime"
)

func init() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
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
