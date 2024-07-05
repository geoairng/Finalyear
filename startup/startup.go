package startup

import (
	"fmt"

	"github.com/geoairng/Finalyear/database"

	"github.com/geoairng/Finalyear/handlers"
	"github.com/gin-gonic/gin"
)

// @Summary Welcome page
// @Description Greet
// @Tags Default
// @Produce json
// @Success 200
// @Router / [get]
func Greet(c *gin.Context) {

	c.JSON(200, gin.H{"Greeting": "Welcome"})
}

func StartApp() *gin.Engine {

	router := gin.Default()

	db, err := database.SetUp()
	if err != nil {
		fmt.Println(err)

	}

	Server := handlers.Server{DB: db}

	router.GET("/", Greet)
	router.POST("/user", Server.Register)

	return router
}
