package main

import (
	"backend-demo/initialzers"
	"backend-demo/models"
)

func init() {
	initialzers.LoadEnv()
	initialzers.ConnectToDB()
}

func main() {
	initialzers.DB.AutoMigrate(&models.Post{})
}