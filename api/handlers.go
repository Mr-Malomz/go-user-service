package api

import (
	"context"
	"go-user-service/data"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const appTimeout = time.Second * 10

func (app *Config) getUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		data, err := app.getUserService()
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusOK, data)
	}
}

func (app *Config) createUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.User
		defer cancel()

		app.validateBody(ctx, &payload)

		newData := data.User{
			FirstName:   payload.FirstName,
			LastName:    payload.LastName,
			PhoneNumber: payload.PhoneNumber,
		}

		data, err := app.createUserService(newData)
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusAccepted, data)
	}
}
