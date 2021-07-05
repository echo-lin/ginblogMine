package main

import (
	"ginblogMine/model"
	"ginblogMine/routes"
)

func main(){
	model.InitDb()
	routes.InitRouter()
}