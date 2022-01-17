package main

import (
	"external.v1/deliver/restapi"
	"external.v1/repository/mysql"
)

func main() {
	mysql.Init()

	e := restapi.SetupRouter()
	e.Logger.Fatal(e.Start(":8082"))
}