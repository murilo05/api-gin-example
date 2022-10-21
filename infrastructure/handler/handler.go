package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gitlab.engdb.com.br/apigin/domain/entities"
	"gitlab.engdb.com.br/apigin/domain/usecase"
	errorUtils "gitlab.engdb.com.br/apigin/utils/error"
)

func Routes(r *gin.Engine) *gin.Context {
	r.GET("/getUsers", getUsers)
	r.GET("/getUserById/:id", getUserById)
	r.POST("/createUser", createUser)
	r.DELETE("/deleteUser/:id", deleteUser)
	r.PUT("/editUser/:id", editUser)

	return nil
}

func getUsers(c *gin.Context) {
	var token = c.GetHeader("token")

	gctx := context.Background()
	gctx = context.WithValue(gctx, "token", token)
	ctx, cancel := context.WithTimeout(gctx, 30000)
	defer cancel()

	resp, code, errorResp := usecase.GetUser(ctx)
	if errorResp != nil {
		c.JSON(code, errorResp)
		return
	}

	c.JSON(code, resp)
}

func getUserById(c *gin.Context) {
	var token = c.GetHeader("token")

	gctx := context.Background()
	gctx = context.WithValue(gctx, "token", token)
	ctx, cancel := context.WithTimeout(gctx, 30000)
	defer cancel()

	inputId := c.Param("id")

	resp, code, errorResp := usecase.GetUserById(ctx, inputId)
	if errorResp != nil {
		c.JSON(code, errorResp)
		return
	}

	c.JSON(code, resp)
}

func createUser(c *gin.Context) {
	var token = c.GetHeader("token")

	gctx := context.Background()
	gctx = context.WithValue(gctx, "token", token)
	ctx, cancel := context.WithTimeout(gctx, 30000)
	defer cancel()

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errResp := errorUtils.CreateError(500, err.Error())
		c.JSON(500, errResp)
		return
	}

	body := entities.User{}

	json.Unmarshal(jsonData, &body)
	if err != nil {
		errResp := errorUtils.CreateError(500, err.Error())
		c.JSON(500, errResp)
		return
	}

	code, errorResp := usecase.CreateUser(ctx, body)
	if errorResp != nil {
		c.JSON(code, errorResp)
		return
	}

	c.String(200, "User Created")
}

func deleteUser(c *gin.Context) {
	var token = c.GetHeader("token")

	gctx := context.Background()
	gctx = context.WithValue(gctx, "token", token)
	ctx, cancel := context.WithTimeout(gctx, 30000)
	defer cancel()

	inputId := c.Param("id")

	code, errorResp := usecase.DeleteUser(ctx, inputId)
	if errorResp != nil {
		c.JSON(code, errorResp)
		return
	}

	c.String(200, "User Deleted")
}

func editUser(c *gin.Context) {
	var token = c.GetHeader("token")

	gctx := context.Background()
	gctx = context.WithValue(gctx, "token", token)
	ctx, cancel := context.WithTimeout(gctx, 30000)
	defer cancel()

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errResp := errorUtils.CreateError(500, err.Error())
		c.JSON(500, errResp)
		return
	}

	body := entities.User{}

	json.Unmarshal(jsonData, &body)
	if err != nil {
		errResp := errorUtils.CreateError(500, err.Error())
		c.JSON(500, errResp)
		return
	}

	inputId := c.Param("id")

	code, errorResp := usecase.EditUser(ctx, body, inputId)
	if errorResp != nil {
		c.JSON(code, errorResp)
		return
	}

	c.String(200, "User Updated")
}
