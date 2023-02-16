package controller

import (
	
)

var server = Server{}

func Run() {

	server.DBInitialization()

	server.Run(":8084")
}
