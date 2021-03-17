package controllers

import (
	"context"
	"fmt"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team09/app/ent"
	"github.com/team09/app/ent/course"
	"github.com/team09/app/ent/department"
	"github.com/team09/app/ent/doctor"
	"github.com/team09/app/ent/training"
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
	Trainingplace       string
	Firstday      string
	Lastday      string
	Doctoridcard string
	Hour         int
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

	time1, err := time.Parse(time.RFC3339, obj.Firstday+"T00:00:00Z")
	time2, err := time.Parse(time.RFC3339, obj.Lastday+"T00:00:00Z")
	t, err := ctl.client.Training.
		Create().
		SetCourse(co).
		SetDepartment(de).
		SetDoctor(d).
		SetTrainingplace(obj.Trainingplace).
		SetFirstday(time1).
		SetLastday(time2).
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

// GetTraining handles GET requests to retrieve a training entity
// @Summary Get a training entity by ID
// @Description get training by ID
// @ID get-training
// @Produce  json
// @Param id path int true "Training ID"
// @Success 200 {object} ent.Training
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /trainings/{id} [get]

func (ctl *TrainingController) GetTraining(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	t, err := ctl.client.Training.
		Query().
		Where(training.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, t)
}

// GetSearchTrainingTable handles GET requests to retrieve a Training entity
// @Summary Get a Training entity by Search
// @Description get Training by Search
// @ID get-Training-by-search
// @Produce  json
// @Param Training query string false "Search Training"
// @Success 200 {object} ent.Training
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /trainingtables [get]
func (ctl *TrainingController) GetSearchTrainingTable(c *gin.Context) {
	trainingsearch := c.Query("training")

	t, err := ctl.client.Training.
		Query().
		WithCourse().
		WithDepartment().
		WithDoctor().
		Where(training.DoctoridcardContains(trainingsearch)).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": t,
	})
}


// DeleteTraining handles DELETE requests to delete a training entity
// @Summary Delete a training entity by ID
// @Description get training by ID
// @ID delete-training
// @Produce  json
// @Param id path int true "Training ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /trainings/{id} [delete]
func (ctl *TrainingController) DeleteTraining(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Training.
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

	// Search
	searchTrainingTables := ctl.router.Group("/trainingtable")
	searchTrainingTables.GET("", ctl.GetSearchTrainingTable)

	trainings.GET("", ctl.ListTraining)
	trainings.POST("", ctl.CreateTraining)
	trainings.GET(":id", ctl.GetTraining)
	trainings.DELETE(":id", ctl.DeleteTraining)

}
