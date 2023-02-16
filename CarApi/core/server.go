package core

import (
	"car/core/controller"
)

var server = controller.Server{}

func Run() {

	server.Initialization()

	server.Run(":8084")
}
