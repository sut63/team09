package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team09/app/ent"
	"github.com/team09/app/ent/course"
)

// CourseController defines the struct for the course controller
type CourseController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateCourse handles POST requests for adding course entities
// @Summary Create course
// @Description Create course
// @ID create-course
// @Accept   json
// @Produce  json
// @Param course body ent.Course true "Course entity"
// @Success 200 {object} ent.Course
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courses [post]
func (ctl *CourseController) CreateCourse(c *gin.Context) {
	obj := ent.Course{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Course binding failed",
		})
		return
	}

	co, err := ctl.client.Course.
		Create().
		SetNamecourse(obj.Namecourse).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, co)
}

// GetCourse handles GET requests to retrieve a course entity
// @Summary Get a course entity by ID
// @Description get course by ID
// @ID get-course
// @Produce  json
// @Param id path int true "Course ID"
// @Success 200 {object} ent.Course
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courses/{id} [get]
func (ctl *CourseController) GetCourse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	co, err := ctl.client.Course.
		Query().
		Where(course.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, co)
}

// ListCourse handles request to get a list of course entities
// @Summary List course entities
// @Description list course entities
// @ID list-course
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Course
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courses [get]
func (ctl *CourseController) ListCourse(c *gin.Context) {
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

	courses, err := ctl.client.Course.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, courses)
}

// NewCourseController creates and registers handles for the course controller
func NewCourseController(router gin.IRouter, client *ent.Client) *CourseController {
	cc := &CourseController{
		client: client,
		router: router,
	}
	cc.register()
	return cc
}

// InitCourseController registers routes to the main engine
func (ctl *CourseController) register() {
	courses := ctl.router.Group("/courses")
	courses.POST("", ctl.CreateCourse)
	courses.GET("", ctl.ListCourse)
	courses.GET(":id", ctl.GetCourse)
}
