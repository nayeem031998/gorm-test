package core

import (
	"gorm/core/controller"
)

var server = controller.Server{}

func Run() {

	server.Initialization()

	server.Run(":8082")
}
