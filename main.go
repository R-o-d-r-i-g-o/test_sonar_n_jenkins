package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "test_sonar_n_jenkins/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title						TODO APIs
// @version					1.0
// @description				Testing Swagger APIs.
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.url				http://www.swagger.io/support
// @contact.email				support@swagger.io
// @securityDefinitions.apiKey	JWT
// @in							header
// @name						token
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:8080
// @schemes					http
func main() {
	r := gin.Default()

	r.GET("/test", pedroGay)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// GetBookByISBN                godoc
//
//	@Summary		Get info from a helth route
//	@Description	get string by ID
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"ok"
//	@Router			/test [get]
func pedroGay(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}
