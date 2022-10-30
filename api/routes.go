package api

import "github.com/gin-gonic/gin"

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	app.Router.GET("/users", app.getUser())
	app.Router.POST("/users", app.createUser())
}
