package main

import (
	"github.com/zidoshare/go-store/controller"
	"github.com/zidoshare/go-store/service"
	"github.com/zidoshare/go-store/utils"
)

func main() {
	utils.Load()
	service.Connect()
	controller.Load()
}
