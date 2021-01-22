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
	"github.com/team09/app/ent/extradoctor"
)

// OfficeController defines the struct for the office controller
type OfficeController struct {
	client *ent.Client
	router gin.IRouter
}

type Office struct {
	Officename  	string
	Added1			string
	Added2 			string
	Doctor      	int
	Department  	int
	Extradoctor  	int
	Roomnumber		string
	Doctoridcard    string
}

// CreateOffice handles POST requests for adding office entities
// @Summary Create office
// @Description Create office
// @ID create-office
// @Accept   json
// @Produce  json
// @Param office body ent.Office true "Office entity"
// @Success 200 {object} ent.Office
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /offices [post]
func (ctl *OfficeController) CreateOffice(c *gin.Context) {
	obj := Office{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "office binding failed",
		})
		return
	}

	et, err := ctl.client.Extradoctor.
		Query().
		Where(extradoctor.IDEQ(int(obj.Extradoctor))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "specialdoctor not found",
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

	time1, err := time.Parse(time.RFC3339, obj.Added1+"T00:00:00Z")
	time2, err := time.Parse(time.RFC3339, obj.Added2+"T00:00:00Z")
	of, err := ctl.client.Office.
		Create().
		SetOfficename(obj.Officename).
		SetDoctoridcard(obj.Doctoridcard).
		SetRoomnumber(obj.Roomnumber).
		SetDoctor(d).
		SetDepartment(de). 
		SetExtradoctor(et).
		SetAddedTime1(time1).
		SetAddedTime2(time2).
		Save(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error" : err,
		})
		return
	}
	
	c.JSON(200, gin.H{
		"status": true,
		"data":   of,
	})
}

// GetOffice handles GET requests to retrieve a office entity
// @Summary Get a office entity by ID
// @Description get office by ID
// @ID get-office
// @Produce  json
// @Param id path int true "Office ID"
// @Success 200 {object} ent.Office
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /offices/{id} [get]
func (ctl *OfficeController) GetOffice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	of, err := ctl.client.Office.
		Query().
		Where(office.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, of)
}

// ListOffice handles request to get a list of office entities
// @Summary List office entities
// @Description list office entities
// @ID list-office
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Office
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /offices [get]
func (ctl *OfficeController) ListOffice(c *gin.Context) {
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

	offices, err := ctl.client.Office.
		Query().
		WithDoctor().
		WithDepartment().
		WithExtradoctor().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, offices)
}

// DeleteOffice handles DELETE requests to delete a office entity
// @Summary Delete a office entity by ID
// @Description get office by ID
// @ID delete-office
// @Produce  json
// @Param id path int true "Office ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /offices/{id} [delete]
func (ctl *OfficeController) DeleteOffice(c *gin.Context) {
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

// NewOfficeController creates and registers handles for the office controller
func NewOfficeController(router gin.IRouter, client *ent.Client) *OfficeController {
	ofc := &OfficeController{
		client: client,
		router: router,
	}
	ofc.register()
	return ofc
}

// InitOfficeController registers routes to the main engine
func (ctl *OfficeController) register() {
	offices := ctl.router.Group("/offices")
	offices.GET("", ctl.ListOffice)
	// CRUD
	offices.POST("", ctl.CreateOffice)
	offices.GET(":id", ctl.GetOffice)
	offices.DELETE(":id", ctl.DeleteOffice)
}
