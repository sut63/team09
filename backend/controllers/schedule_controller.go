package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team09/app/ent"
	"github.com/team09/app/ent/department"
	"github.com/team09/app/ent/doctor"
	"github.com/team09/app/ent/office"
	"github.com/team09/app/ent/schedule"
)

// ScheduleController defines the struct for the user controller
type ScheduleController struct {
	client *ent.Client
	router gin.IRouter
}

type Schedule struct {
	Department int
	Office     int
	Doctor     int
	Activity   string
	Added      string
}

// CreateSchedule handles POST requests for adding schedule entities
// @Summary Create schedule
// @Description Create schedule
// @ID create-schedule
// @Accept   json
// @Produce  json
// @Param schedule body ent.Schedule true "Schedule entity"
// @Success 200 {object} ent.Schedule
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schedules [post]
func (ctl *ScheduleController) CreateSchedule(c *gin.Context) {
	obj := Schedule{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Schedule binding failed",
		})
		return
	}

	de, err := ctl.client.Department.
		Query().
		Where(department.IDEQ(int(obj.Department))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "department not found",
		})
		return
	}

	of, err := ctl.client.Office.
		Query().
		Where(office.IDEQ(int(obj.Office))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "office not found",
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

	time, err := time.Parse(time.RFC3339, obj.Added+":00+07:00")
	sh, err := ctl.client.Schedule.
		Create().
		SetDepartment(de).
		SetOffice(of).
		SetDocter(d).
		SetActivity(obj.Activity).
		SetAddedTime(time).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   sh,
	})
}

// GetSchedule handles GET requests to retrieve a schedule entity
// @Summary Get a schedule entity by ID
// @Description get schedule by ID
// @ID get-schedule
// @Produce  json
// @Param id path int true "Schedule ID"
// @Success 200 {object} ent.Schedule
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schedules/{id} [get]
func (ctl *ScheduleController) GetSchedule(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	sh, err := ctl.client.Schedule.
		Query().
		Where(schedule.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, sh)
}

// ListSchedule handles request to get a list of Schedule entities
// @Summary List schedule entities
// @Description list schedule entities
// @ID list-schedule
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Schedule
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schedules [get]
func (ctl *ScheduleController) ListSchedule(c *gin.Context) {
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

	Schedules, err := ctl.client.Schedule.
		Query().
		WithDepartment().
		WithOffice().
		WithDocter().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, Schedules)
}

// DeleteSchedule handles DELETE requests to delete a schedule entity
// @Summary Delete a schedule entity by ID
// @Description get schedule by ID
// @ID delete-schedule
// @Produce  json
// @Param id path int true "Schedule ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schedules/{id} [delete]
func (ctl *ScheduleController) DeleteSchedule(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Schedule.
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

// UpdateSchedule handles PUT requests to update a schedule entity
// @Summary Update a schedule entity by ID
// @Description update schedule by ID
// @ID update-schedule
// @Accept   json
// @Produce  json
// @Param id path int true "Schedule ID"
// @Param schedule body ent.Schedule true "Schedule entity"
// @Success 200 {object} ent.Schedule
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schedules/{id} [put]
func (ctl *ScheduleController) UpdateSchedule(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Schedule{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Schedule binding failed",
		})
		return
	}
	obj.ID = int(id)
	uq, err := ctl.client.Schedule.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, uq)
}

// NewScheduleController creates and registers handles for the schedule controller
func NewScheduleController(router gin.IRouter, client *ent.Client) *ScheduleController {
	rc := &ScheduleController{
		client: client,
		router: router,
	}
	rc.register()
	return rc
}

// InitScheduleController registers routes to the main engine
func (ctl *ScheduleController) register() {
	schedules := ctl.router.Group("/schedules")

	schedules.GET("", ctl.ListSchedule)

	// CRUD
	schedules.POST("", ctl.CreateSchedule)
	schedules.GET(":id", ctl.GetSchedule)
	schedules.DELETE(":id", ctl.DeleteSchedule)

}
