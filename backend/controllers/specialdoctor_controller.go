package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team09/app/ent"
	"github.com/team09/app/ent/specialdoctor"
	"github.com/team09/app/ent/doctor"
	"github.com/team09/app/ent/department"
	"github.com/team09/app/ent/extradoctor"
)

// SpecialdoctorController defines the struct for the specialdoctor controller
type SpecialdoctorController struct {
	client *ent.Client
	router gin.IRouter
}
type Specialdoctor struct {
	Doctorid			string
	Roomnumber			string
	Other				string 
	Extradoctor			int 
	Doctor 				int
	Department 			int
}

// CreateSpecialdoctor handles POST requests for adding specialdoctor entities
// @Summary Create specialdoctor
// @Description Create specialdoctor
// @ID create-specialdoctor
// @Accept   json
// @Produce  json
// @Param specialdoctor body ent.Specialdoctor true "Specialdoctor entity"
// @Success 200 {object} ent.Specialdoctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specialdoctors [post]
func (ctl *SpecialdoctorController) CreateSpecialdoctor(c *gin.Context) {
	obj := Specialdoctor{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "specialdoctor binding failed",
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
	s, err := ctl.client.Extradoctor.
		Query().
		Where(extradoctor.IDEQ(int(obj.Extradoctor))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Doctor not found",
		})
		return
	}

	sl, err := ctl.client.Specialdoctor.
		Create().
		SetExtradoctor(s).
		SetDoctor(d).
		SetDepartment(de).
		SetDoctorid(obj.Doctorid).
		SetRoomnumber(obj.Roomnumber).
		SetOther(obj.Other).
		Save(context.Background())
		
		if err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{
				"status": false,
				"error":  err,
			})
			return
		}
	
		c.JSON(200, gin.H{
			"status": true,
			"data":   sl,
		})
}

// GetSpecialdoctor handles GET requests to retrieve a specialdoctor entity
// @Summary Get a specialdoctor entity by ID
// @Description get specialdoctor by ID
// @ID get-specialdoctor
// @Produce  json
// @Param id path int true "Specialdoctor ID"
// @Success 200 {object} ent.Specialdoctor
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specialdoctors/{id} [get]
func (ctl *SpecialdoctorController) GetSpecialdoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	sl, err := ctl.client.Specialdoctor.
		Query().
		Where(specialdoctor.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, sl)
}

// ListSpecialdoctor handles request to get a list of specialdoctorentities
// @Summary List specialdoctor entities
// @Description list specialdoctor entities
// @ID list-specialdoctor
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Specialdoctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specialdoctors [get]
func (ctl *SpecialdoctorController) ListSpecialdoctor(c *gin.Context) {
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

	specialdoctors, err := ctl.client.Specialdoctor.
		Query().
		Limit(limit).
		Offset(offset).
		WithDepartment().
		WithDoctor().
		WithExtradoctor().
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, specialdoctors)
}

// DeleteSpecialdoctor handles DELETE requests to delete a specialdoctor entity
// @Summary Delete a specialdoctor entity by ID
// @Description get specialdoctor by ID
// @ID delete-specialdoctor
// @Produce  json
// @Param id path int true "Specialdoctor ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specialdoctors/{id} [delete]
func (ctl *SpecialdoctorController) DeleteSpecialdoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Specialdoctor.
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

// NewSpecialdoctorController creates and registers handles for the specialdoctor controller
func NewSpecialdoctorController(router gin.IRouter, client *ent.Client) *SpecialdoctorController {
	slc := &SpecialdoctorController{
		client: client,
		router: router,
	}
	slc.register()
	return slc
}

// InitSpecialdoctorController registers routes to the main engine
func (ctl *SpecialdoctorController) register() {
	specialdoctors := ctl.router.Group("/specialdoctors")

	specialdoctors.GET("", ctl.ListSpecialdoctor)
	// CRUD
	specialdoctors.POST("", ctl.CreateSpecialdoctor)
	specialdoctors.GET(":id", ctl.GetSpecialdoctor)
	specialdoctors.DELETE(":id", ctl.DeleteSpecialdoctor)
}
