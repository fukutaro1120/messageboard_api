package controller

import (
	"fmt"
	"message-board-api/model"
	"message-board-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ArticleController ...
type ArticleController struct {
	UseCase *usecase.ArticleUseCase
}

// List ...
func (c *ArticleController) List(ctx *gin.Context) {
	articles, err := c.UseCase.List(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewError(err))
		return
	}

	ctx.JSON(200, gin.H{
		"articles": articles,
	})
}

// Parms ...
type Parms struct {
	ID      int
	Name    string `json:"name" form:"name" binding:"required"`
	Message string `json:"message" form:"message" binding:"required"`
}

// DeleteParams ...
type DeleteParams struct {
	ID int
}

// UpParms ...
type UpParms struct {
	ID      int
	Message string `json:"message" form:"message" binding:"required"`
}

// Create ...
func (c *ArticleController) Create(ctx *gin.Context) {
	var p Parms
	if err := ctx.ShouldBind(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Printf("controller.createでエラー：%v", err)
	}
	c.UseCase.Create(ctx, &model.Article{Name: p.Name, Message: p.Message})
}

// Delete ...
func (c *ArticleController) Delete(ctx *gin.Context) {
	// var dp DeleteParams
	msgID := ctx.Query("id")
	messageID, err := strconv.Atoi(msgID)
	if err != nil {
		fmt.Printf("controller.deleteでエラー：%v", err)
	}
	dp := &DeleteParams{ID: messageID}
	c.UseCase.Delete(ctx, &model.Article{ID: dp.ID})
}

// Update   ...
func (c *ArticleController) Update(ctx *gin.Context) {
	var up UpParms
	if err := ctx.ShouldBind(&up); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Printf("controller.updateでエラー：%v", err)
	}
	c.UseCase.Update(ctx, &model.Article{ID: up.ID, Message: up.Message})
}
