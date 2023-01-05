package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gitlab.engdb.com.br/apigin/domain/entities"
	"gitlab.engdb.com.br/apigin/interfaces"
	errorUtils "gitlab.engdb.com.br/apigin/utils/error"
)

type UserHandler struct {
	userUseCase interfaces.UserUseCase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewProjectHandler(r *gin.Engine, us interfaces.UserUseCase) *gin.Context {
	handler := &UserHandler{
		userUseCase: us,
	}
	r.GET("/getUsers", handler.getUsers)
	r.GET("/getUserById/:id", handler.getUserById)
	r.POST("/createUser", handler.createUser)
	r.DELETE("/:id", handler.deleteUser)
	r.PUT("/:id", handler.editUser)

	return nil
}

func (uh *UserHandler) getUsers(c *gin.Context) {
	ctx := context.Background()

	resp, code, errorResp := uh.userUseCase.GetUser(ctx)
	if errorResp != nil {
		c.JSON(code, errorResp)
		return
	}

	c.JSON(code, resp)
}

func (uh *UserHandler) getUserById(c *gin.Context) {
	ctx := context.Background()

	inputId := c.Param("id")

	resp, code, errorResp := uh.userUseCase.GetUserById(ctx, inputId)
	if errorResp != nil {
		c.JSON(code, errorResp)
		return
	}

	c.JSON(code, resp)
}

func (uh *UserHandler) createUser(c *gin.Context) {
	ctx := context.Background()

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

	code, errorResp := uh.userUseCase.CreateUser(ctx, body)
	if errorResp != nil {
		c.JSON(code, errorResp)
		return
	}

	c.String(200, "User Created")
}

func (uh *UserHandler) deleteUser(c *gin.Context) {
	ctx := context.Background()

	inputId := c.Param("id")

	code, errorResp := uh.userUseCase.DeleteUser(ctx, inputId)
	if errorResp != nil {
		c.JSON(code, errorResp)
		return
	}

	c.String(200, "User Deleted")
}

func (uh *UserHandler) editUser(c *gin.Context) {
	ctx := context.Background()

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

	code, errorResp := uh.userUseCase.EditUser(ctx, body, inputId)
	if errorResp != nil {
		c.JSON(code, errorResp)
		return
	}

	c.String(200, "User Updated")
}
