package main

import (
	"log"

	"github.com/geoairng/Finalyear/startup"
)

// @title           Golang gin for E-commerce
// @version         1.0
// @description     An E-commerce app with gin framework
// @contact.email  lawalafeez820@gmail.com
// @securityDefinitions.apikey bearerToken
// @in header
// @name Authorization
// @tokenUrl http://localhost:8080/login

func main() {

	router := startup.StartApp()

	log.Fatal(router.Run())

}
