package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team09/app/ent"
	"github.com/team09/app/ent/workingtime"
)

// WorkingtimeController defines the struct for the office controller
type WorkingtimeController struct {
	client *ent.Client
	router gin.IRouter
}
type Workingtime struct {
	Added1 string
	Added2 string
}

// CreateWorkingtime handles POST requests for adding workingtime entities
// @Summary Create workingtime
// @Description Create workingtime
// @ID create-workingtime
// @Accept   json
// @Produce  json
// @Param doctor body ent.Workingtime true "Workingtime entity"
// @Success 200 {object} ent.Workingtime
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /workingtimes [post]
func (ctl *WorkingtimeController) CreateWorkingtime(c *gin.Context) {
	obj := Workingtime{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "workingtime binding failed",
		})
		return
	}

	time1, err := time.Parse(time.RFC3339, obj.Added1)
	time2, err := time.Parse(time.RFC3339, obj.Added2)
	wt, err := ctl.client.Workingtime.
		Create().
		SetAddedTime1(time1).
		SetAddedTime2(time2).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	c.JSON(200, wt)
}

// GetWorkingtime handles GET requests to retrieve a Workingtime entity
// @Summary Get a workingtime entity by ID
// @Description get workingtime by ID
// @ID get-workingtime
// @Produce  json
// @Param id path int true "Workingtime ID"
// @Success 200 {object} ent.Workingtime
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /workingtimes/{id} [get]
func (ctl *WorkingtimeController) GetWorkingtime(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	wt, err := ctl.client.Workingtime.
		Query().
		Where(workingtime.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, wt)
}

// ListWorkingtime handles request to get a list of workingtime entities
// @Summary List workingtime entities
// @Description list workingtime entities
// @ID list-office
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Workingtime
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /workingtimes [get]
func (ctl *WorkingtimeController) ListWorkingtime(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	workingtimes, err := ctl.client.Workingtime.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, workingtimes)
}

// DeleteWokingtime handles DELETE requests to delete a workingtime entity
// @Summary Delete a workingtime entity by ID
// @Description get workingtime by ID
// @ID delete-workingtime
// @Produce  json
// @Param id path int true "Workingtime ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /workingtimes/{id} [delete]
func (ctl *WorkingtimeController) DeleteWokingtime(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Office.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// NewWoringtimeController creates and registers handles for the workingtime controller
func NewWoringtimeController(router gin.IRouter, client *ent.Client) *WorkingtimeController {
	wt := &WorkingtimeController{
		client: client,
		router: router,
	}
	wt.register()
	return wt
}

// InitWorkingtimeController registers routes to the main engine
func (ctl *WorkingtimeController) register() {
	workingtimes := ctl.router.Group("/workingtimes")
	workingtimes.GET("", ctl.ListWorkingtime)
	// CRUD
	workingtimes.POST("", ctl.CreateWorkingtime)
	workingtimes.GET(":id", ctl.GetWorkingtime)
	workingtimes.DELETE(":id", ctl.DeleteWokingtime)
}
