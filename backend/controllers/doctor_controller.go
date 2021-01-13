package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team09/app/ent"
	"github.com/team09/app/ent/disease"
	"github.com/team09/app/ent/doctor"
	"github.com/team09/app/ent/gender"
	"github.com/team09/app/ent/position"
	"github.com/team09/app/ent/title"
)

// DoctorController defines the struct for the doctor controller
type DoctorController struct {
	client *ent.Client
	router gin.IRouter
}
type Doctor struct {
	Name 		string
	Age			int
	Email		string
	Password 	string
	Address  	string
	Educational string
	Phone 		int
	Title   	int
	Position 	int
	Gender 		int 
	Disease 	int
}

// CreateDoctor handles POST requests for adding doctor entities
// @Summary Create doctor
// @Description Create doctor
// @ID create-doctor
// @Accept   json
// @Produce  json
// @Param doctor body ent.Doctor true "Doctor entity"
// @Success 200 {object} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctors [post]
func (ctl *DoctorController) CreateDoctor(c *gin.Context) {
	obj := Doctor{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "doctor binding failed",
		})
		return
	}

	t, err := ctl.client.Title.
		Query().
		Where(title.IDEQ(int(obj.Title))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "title not found",
		})
		return
	}

	p, err := ctl.client.Position.
		Query().
		Where(position.IDEQ(int(obj.Position))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "position not found",
		})
		return
	}

	di, err := ctl.client.Disease.
		Query().
		Where(disease.IDEQ(int(obj.Disease))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "disease not found",
		})
		return
	}

	g, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(int(obj.Gender))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "gender not found",
		})
		return
	}

	d, err := ctl.client.Doctor.
		Create().
		SetName(obj.Name).
		SetAddress(obj.Address).
		SetEducational(obj.Educational).
		SetPassword(obj.Password).
		SetAge(obj.Age).
		SetPhone(obj.Phone).
		SetEmail(obj.Email).
		SetTitle(t).
		SetPosition(p).
		SetGender(g).
		SetDisease(di).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	
	c.JSON(200, gin.H{
		"status": true,
		"data":   d,
	})
}

// GetDoctor handles GET requests to retrieve a doctor entity
// @Summary Get a doctor entity by ID
// @Description get doctor by ID
// @ID get-doctor
// @Produce  json
// @Param id path int true "Doctor ID"
// @Success 200 {object} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctors/{id} [get]
func (ctl *DoctorController) GetDoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	d, err := ctl.client.Doctor.
		Query().
		Where(doctor.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, d)
}

// ListDoctor handles request to get a list of doctor entities
// @Summary List doctor entities
// @Description list doctor entities
// @ID list-doctor
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctors [get]
func (ctl *DoctorController) ListDoctor(c *gin.Context) {
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

	doctors, err := ctl.client.Doctor.
		Query().
		WithTitle().
		WithGender().
		WithPosition().
		WithDisease().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, doctors)
}

// DeleteDoctor handles DELETE requests to delete a doctor entity
// @Summary Delete a doctor entity by ID
// @Description get doctor by ID
// @ID delete-doctor
// @Produce  json
// @Param id path int true "Doctor ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctors/{id} [delete]
func (ctl *DoctorController) DeleteDoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Doctor.
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

// NewDoctorController creates and registers handles for the doctor controller
func NewDoctorController(router gin.IRouter, client *ent.Client) *DoctorController {
	dc := &DoctorController{
		client: client,
		router: router,
	}
	dc.register()
	return dc
}

// InitDoctorController registers routes to the main engine
func (ctl *DoctorController) register() {
	doctors := ctl.router.Group("/doctors")
	doctors.GET("", ctl.ListDoctor)
	// CRUD
	doctors.POST("", ctl.CreateDoctor)
	doctors.GET(":id", ctl.GetDoctor)
	doctors.DELETE(":id", ctl.DeleteDoctor)
}