package controllers
 
import (
   "context"
   "fmt"
   "strconv"

   "github.com/team09/app/ent"
   "github.com/team09/app/ent/special_doctor"
   "github.com/gin-gonic/gin"
)
 
// Special_DoctorController defines the struct for the special_doctor controller
type Special_DoctorController struct {
   client *ent.Client
   router gin.IRouter
}

// CreateSpecial_Doctor handles POST requests for adding special_doctor entities
// @Summary Create special_doctor
// @Description Create special_doctor
// @ID create-special_doctor
// @Accept   json
// @Produce  json
// @Param special_doctor body ent.Special_Doctor true "Special_Doctor entity"
// @Success 200 {object} ent.Special_Doctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /special_doctors [post]
func (ctl *Special_DoctorController) CreateSpecial_Doctor(c *gin.Context) {
	obj := ent.Special_Doctor{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "special_doctor binding failed",
		})
		return
	}
  
	d, err := ctl.client.Special_Doctor.
		Create().
		SetSpecial_Doctor(obj.Special_Doctor).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, d)
 }
 
// GetSpecial_Doctor handles GET requests to retrieve a special_doctor entity
// @Summary Get a special_doctor entity by ID
// @Description get special_doctor by ID
// @ID get-special_doctor
// @Produce  json
// @Param id path int true "Special_Doctor ID"
// @Success 200 {object} ent.Special_Doctor
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /special_doctors/{id} [get]
func (ctl *Special_DoctorController) GetSpecial_Doctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	d, err := ctl.client.Special_Doctor.
		Query().
		Where(special_doctor.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, d)
 }

 // ListSpecial_Doctor handles request to get a list of special_doctor entities
// @Summary List special_doctor entities
// @Description list special_doctor entities
// @ID list-special_doctor
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Special_Doctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /special_doctors [get]
func (ctl *Special_DoctorController) ListSpecial_Doctor(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {limit = int(limit64)}
	}
  
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {offset = int(offset64)}
	}
  
	special_doctors, err := ctl.client.Special_Doctor.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, special_doctors)
 }
 
 // DeleteSpecial_Doctor handles DELETE requests to delete a special_doctor entity
// @Summary Delete a special_doctor entity by ID
// @Description get special_doctor by ID
// @ID delete-special_doctor
// @Produce  json
// @Param id path int true "Special_Doctor ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /special_doctors/{id} [delete]
func (ctl *Special_DoctorController) DeleteSpecial_Doctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Special_Doctor.
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
 
 // NewSpecial_DoctorController creates and registers handles for the special_doctor controller
func NewSpecial_DoctorController(router gin.IRouter, client *ent.Client) *Special_DoctorUserController {
	uc := &Special_DoctorController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitSpecial_DoctorController registers routes to the main engine
 func (ctl *Special_DoctorController) register() {
	special_doctors := ctl.router.Group("/special_doctors")
  
	special_doctors.GET("", ctl.ListSpecial_Doctor)
	// CRUD
	special_doctors.POST("", ctl.CreateSpecial_Doctor)
	special_doctors.GET(":id", ctl.GetSpecial_Doctor)
	special_doctors.DELETE(":id", ctl.DeleteSpecial_Doctor)
 }
 