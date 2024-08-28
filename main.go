package main

import (
	"fmt"
	"net/http"
	"os"
	"task-app/config"
	"task-app/docs"
	"task-app/handlers"
	"task-app/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Configures system wide Logger object
	zlog.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	server := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowAllOrigins = true

	server.Use(cors.New(corsConfig), gin.Recovery())

	//load configurations
	configVariables := config.GetConfig()

	handler := handlers.NewHandler(configVariables)

	// register routes
	r := routes.NewRoutes(handler)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Task App APIs"
	docs.SwaggerInfo.Description = "This is the API docs for the Task application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	url := ginSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", configVariables.AppHost))
	server.GET("/swagger/*any", func(c *gin.Context) {
		if c.Param("any") == "/" || c.Param("any") == "" {
			c.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
		} else {
			ginSwagger.WrapHandler(swaggerFiles.Handler, url)(c)
		}
	})

	r.RegisterRoutes(server, handler)
	//run server
	if err := server.Run(fmt.Sprintf("%s:%s", configVariables.AppHost, configVariables.Port)); err != nil && err != http.ErrServerClosed {
		zlog.Fatal().Msgf("listen: %s", err)
	}
}
