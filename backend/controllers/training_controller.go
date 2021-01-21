package controllers

import (
	"context"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team09/app/ent"
	"github.com/team09/app/ent/course"
	"github.com/team09/app/ent/department"
	"github.com/team09/app/ent/doctor"
)

// TrainingController defines the struct for the training controller
type TrainingController struct {
	client *ent.Client
	router gin.IRouter
}

type Training struct {
	Course       int
	Department   int
	Doctor       int
	Branch       string
	Dateone      string
	Datetwo      string
	Doctoridcard string
	Hour        string
}

// CreateTraining handles POST requests for adding training entities
// @Summary Create training
// @Description Create training
// @ID create-training
// @Accept   json
// @Produce  json
// @Param training body ent.Training true "Training entity"
// @Success 200 {object} ent.Training
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /trainings [post]
func (ctl *TrainingController) CreateTraining(c *gin.Context) {
	obj := Training{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "training binding failed",
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
			"error": "Doctor not found",
		})
		return
	}

	time1, err := time.Parse(time.RFC3339, obj.Dateone+"T00:00:00Z")
	time2, err := time.Parse(time.RFC3339, obj.Datetwo+"T00:00:00Z")
	t, err := ctl.client.Training.
		Create().
		SetCourse(co).
		SetDepartment(de).
		SetDoctor(d).
		SetBranch(obj.Branch).
		SetDateone(time1).
		SetDatetwo(time2).
		SetDoctoridcard(obj.Doctoridcard).
		SetHour(obj.Hour).
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
		"data":   t,
	})
}

// ListTraining handles request to get a list of training entities
// @Summary List training entities
// @Description list training entities
// @ID list-training
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Training
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /trainings [get]
func (ctl *TrainingController) ListTraining(c *gin.Context) {
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

	trainings, err := ctl.client.Training.
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

	c.JSON(200, trainings)
}

// NewTrainingController creates and registers handles for the training controller
func NewTrainingController(router gin.IRouter, client *ent.Client) *TrainingController {
	tr := &TrainingController{
		client: client,
		router: router,
	}
	tr.register()
	return tr
}

// InitTrainingController registers routes to the main engine
func (ctl *TrainingController) register() {
	trainings := ctl.router.Group("/trainings")

	trainings.GET("", ctl.ListTraining)
	trainings.POST("", ctl.CreateTraining)

}
