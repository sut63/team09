package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team09/app/ent"
	"github.com/team09/app/ent/extradoctor"

)

// ExtradoctorController defines the struct for the extradoctor controller
type ExtradoctorController struct {
	client *ent.Client
	router gin.IRouter
}

// GetExtradoctor handles GET requests to retrieve a extradoctor entity
// @Summary Get a extradoctor entity by ID
// @Description get extradoctor by ID
// @ID get-extradoctor
// @Produce  json
// @Param id path int true "Special ID"
// @Success 200 {object} ent.Extradoctor
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /extradoctors/{id} [get]
func (ctl *ExtradoctorController) GetExtradoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	s, err := ctl.client.Extradoctor.
		Query().
		Where(extradoctor.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListExtradoctor handles request to get a list of extradoctor entities
// @Summary Listextradoctor entities
// @Description list extradoctor entities
// @ID list-extradoctor
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Extradoctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /extradoctors [get]
func (ctl *ExtradoctorController) ListExtradoctor(c *gin.Context) {
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

	extradoctors, err := ctl.client.Extradoctor.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, extradoctors)
}

// NewExtradoctorController creates and registers handles for the Extradoctor controller
func NewExtradoctorController(router gin.IRouter, client *ent.Client) *ExtradoctorController {
	dep := &ExtradoctorController{
		client: client,
		router: router,
	}
	dep.register()
	return dep
}

// InitExtradoctorController registers routes to the main engine
func (ctl *ExtradoctorController) register() {
	extradoctors := ctl.router.Group("/extradoctors")

	extradoctors.GET("", ctl.ListExtradoctor)
	// CRUD
	extradoctors.GET(":id", ctl.GetExtradoctor)
}
