package api

import (
	"movie/api/controller"
)

var server = controller.Server{}

func Run() {

	server.Initialization()

	server.Run(":1234")
}
