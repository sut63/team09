package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team09/app/ent"
	"github.com/team09/app/ent/course"
	"github.com/team09/app/ent/department"
	"github.com/team09/app/ent/detail"
	"github.com/team09/app/ent/doctor"
)

// DetailController defines the struct for the detail controller
type DetailController struct {
	client *ent.Client
	router gin.IRouter
}
type Detail struct {
	Explain    string
	Course     int
	Department int
	Phone string
	Email string
	Doctor     int
	Departmentid string
}

// CreateDetail handles POST requests for adding detail entities
// @Summary Create detail
// @Description Create detail
// @ID create-detail
// @Accept   json
// @Produce  json
// @Param Detail body ent.Detail true "Detail entity"
// @Success 200 {object} ent.Detail
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /details [post]
func (ctl *DetailController) CreateDetail(c *gin.Context) {
	obj := Detail{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Detail binding failed",
		})
		return
	}

	co, err := ctl.client.Course.
		Query().
		Where(course.IDEQ(int(obj.Course))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Course not found",
		})
		return
	}

	de, err := ctl.client.Department.
		Query().
		Where(department.IDEQ(int(obj.Department))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Department not found",
		})
		return
	}

	d, err := ctl.client.Doctor.
		Query().
		Where(doctor.IDEQ(int(obj.Doctor))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "doctor not found",
		})
		return
	}

	dt, err := ctl.client.Detail.
		Create().
		SetCourse(co).
		SetDepartment(de).
		SetExplain(obj.Explain).
		SetPhone(obj.Phone).
		SetEmail(obj.Email).
		SetDoctor(d).
		SetDepartmentid(obj.Departmentid).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   dt,
	})
}

// ListDetail handles request to get a list of detail entities
// @Summary List detail entities
// @Description list detail entities
// @ID list-detail
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Detail
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /details [get]
func (ctl *DetailController) ListDetail(c *gin.Context) {
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

	details, err := ctl.client.Detail.
		Query().
		WithCourse().
		WithDepartment().
		WithDoctor().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, details)
}

// GetDetail handles GET requests to retrieve a detail entity
// @Summary Get a detail entity by ID
// @Description get detail by ID
// @ID get-detail
// @Produce  json
// @Param id path int true "detail ID"
// @Success 200 {object} ent.Detail
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /details/{id} [get]
func (ctl *DetailController) GetDetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	dt, err := ctl.client.Detail.
		Query().
		Where(detail.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, dt)
}

// GetSearchDetailTable handles GET requests to retrieve a Detail entity
// @Summary Get a Detail entity by Search
// @Description get Detail by Search
// @ID get-Detail-by-search
// @Produce  json
// @Param Detail query string false "Search Detail"
// @Success 200 {object} ent.Detail
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /detailtables [get]
func (ctl *DetailController) GetSearchDetailTable(c *gin.Context) {
	detailsearch := c.Query("detail")

	dt, err := ctl.client.Detail.
		Query().
		WithCourse().
		WithDepartment().
		WithDoctor().
		Where(detail.DepartmentidContains(detailsearch)).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": dt,
	})
}

// DeleteDetaile handles DELETE requests to delete a detail entity
// @Summary Delete a detail entity by ID
// @Description get detail by ID
// @ID delete-detail
// @Produce  json
// @Param id path int true "detail ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /details/{id} [delete]
func (ctl *DetailController) DeleteDetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Detail.
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

// NewDetailController creates and registers handles for the detail controller
func NewDetailController(router gin.IRouter, client *ent.Client) *DetailController {
	dta := &DetailController{
		client: client,
		router: router,
	}
	dta.register()
	return dta
}

// InitDetailController registers routes to the main engine
func (ctl *DetailController) register() {
	details := ctl.router.Group("/details")
	details.GET("", ctl.ListDetail)
	// Search
	searchdetail := ctl.router.Group("/detailtables")
	searchdetail.GET("", ctl.GetSearchDetailTable)
	// CRUD
	details.POST("", ctl.CreateDetail)
	details.GET(":id", ctl.GetDetail)
	details.DELETE(":id", ctl.DeleteDetail)
}
