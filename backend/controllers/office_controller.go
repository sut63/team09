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
	"github.com/team09/app/ent/extradoctor"
	"github.com/team09/app/ent/office"
)

// OfficeController defines the struct for the office controller
type OfficeController struct {
	client *ent.Client
	router gin.IRouter
}

type Office struct {
	Officename   string
	Firsttime    string
	Lasttime  	 string
	Doctor       int
	Department   int
	Extradoctor  int
	Roomnumber   string
	Doctoridcard string
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

	firsttime, err := time.Parse(time.RFC3339, obj.Firsttime+"T00:00:00Z")
	Lasttime, err := time.Parse(time.RFC3339, obj.Lasttime+"T00:00:00Z")
	of, err := ctl.client.Office.
		Create().
		SetOfficename(obj.Officename).
		SetDoctoridcard(obj.Doctoridcard).
		SetRoomnumber(obj.Roomnumber).
		SetDoctor(d).
		SetDepartment(de).
		SetExtradoctor(et).
		SetFirsttime(firsttime).
		SetFinallytime(Lasttime).
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
		WithDepartment().
		WithDoctor().
		WithExtradoctor().
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

// GetSearchWorkhistory handles GET requests to retrieve a Office entity
// @Summary Get a Office entity by Search
// @Description get Office by Search
// @ID get-Office-by-search
// @Produce  json
// @Param Office query string false "Search Office"
// @Success 200 {object} ent.Office
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /searchworkhistorys [get]
func (ctl *OfficeController) GetSearchWorkhistory(c *gin.Context) {
	officesearch := c.Query("office")

	of, err := ctl.client.Office.
		Query().
		WithDoctor().
		WithDepartment().
		WithExtradoctor().
		Where(office.DoctoridcardContains(officesearch)).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": of,
	})
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
	// Search
	searchworkhistory := ctl.router.Group("/searchworkhistorys")
	searchworkhistory.GET("", ctl.GetSearchWorkhistory)
	// CRUD
	offices.POST("", ctl.CreateOffice)
	offices.GET(":id", ctl.GetOffice)

}
