package main

import (
	"user.v1/deliver/restapi"
	"user.v1/repository/mysql"
)

func main() {
	mysql.Init()

	e := restapi.SetupRouter()
	e.Logger.Fatal(e.Start(":8081"))
}