package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team09/app/ent"
	"github.com/team09/app/ent/specialist"
)

// SpecialistController defines the struct for the specialist controller
type SpecialistController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateSpecialist handles POST requests for adding specialist entities
// @Summary Create specialist
// @Description Create specialist
// @ID create-specialist
// @Accept   json
// @Produce  json
// @Param specialist body ent.Specialist true "Specialist entity"
// @Success 200 {object} ent.Specialist
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specialists [post]
func (ctl *SpecialistController) CreateSpecialist(c *gin.Context) {
	obj := ent.Specialist{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "specialist binding failed",
		})
		return
	}

	d, err := ctl.client.Specialist.
		Create().
		SetSpecialist(obj.Specialist).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, d)
}

// GetSpecialist handles GET requests to retrieve a specialist entity
// @Summary Get a specialist entity by ID
// @Description get specialist by ID
// @ID get-specialist
// @Produce  json
// @Param id path int true "Specialist ID"
// @Success 200 {object} ent.Specialist
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specialists/{id} [get]
func (ctl *SpecialistController) GetSpecialist(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	d, err := ctl.client.Specialist.
		Query().
		Where(specialist.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, d)
}

// ListSpecialist handles request to get a list of specialist entities
// @Summary List specialist entities
// @Description list specialist entities
// @ID list-specialist
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Specialist
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specialists [get]
func (ctl *SpecialistController) ListSpecialist(c *gin.Context) {
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

	specialists, err := ctl.client.Specialist.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, specialists)
}

// DeleteSpecialist handles DELETE requests to delete a specialist entity
// @Summary Delete a specialist entity by ID
// @Description get specialist by ID
// @ID delete-specialist
// @Produce  json
// @Param id path int true "Specialist ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specialists/{id} [delete]
func (ctl *SpecialistController) DeleteSpecialist(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Specialist.
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

// NewSpecialistController creates and registers handles for the specialist controller
func NewSpecialistController(router gin.IRouter, client *ent.Client) *SpecialistController {
	uc := &SpecialistController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitSpecialistController registers routes to the main engine
func (ctl *SpecialistController) register() {
	specialists := ctl.router.Group("/specialists")

	specialists.GET("", ctl.ListSpecialist)
	// CRUD
	specialists.POST("", ctl.CreateSpecialist)
	specialists.GET(":id", ctl.GetSpecialist)
	specialists.DELETE(":id", ctl.DeleteSpecialist)
}
