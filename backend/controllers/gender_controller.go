package controllers
 
import (
   "context"
   "fmt"
   "strconv"

   "github.com/team09/app/ent"
   "github.com/team09/app/ent/gender"
   "github.com/gin-gonic/gin"
)

// GenderController defines the struct for the gender controller
type GenderController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateGender handles POST requests for adding gender entities
// @Summary Create gender
// @Description Create gender
// @ID create-gender
// @Accept   json
// @Produce  json
// @Param gender body ent.Gender true "Gender entity"
// @Success 200 {object} ent.Gender
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /genders [post]
func (ctl *GenderController) CreateGender(c *gin.Context) {
	obj := ent.Gender{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "gender binding failed",
		})
		return
	}
	d, err := ctl.client.Gender.
		Create().
		SetGender(obj.Gender).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	c.JSON(200, d)
}

// GetGender handles GET requests to retrieve a gender entity
// @Summary Get a gender entity by ID
// @Description get gender by ID
// @ID get-gender
// @Produce  json
// @Param id path int true "Gender ID"
// @Success 200 {object} ent.Gender
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /genders/{id} [get]
func (ctl *GenderController) GetGender(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	d, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, d)
}

// ListGender handles request to get a list of gender entities
// @Summary List gender entities
// @Description list gender entities
// @ID list-gender
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Gender
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /genders [get]
func (ctl *GenderController) ListGender(c *gin.Context) {
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
	genders, err := ctl.client.Gender.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
	c.JSON(200, genders)
}

// DeleteGender handles DELETE requests to delete a gender entity
// @Summary Delete a gender entity by ID
// @Description get gender by ID
// @ID delete-gender
// @Produce  json
// @Param id path int true "Gender ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /genders/{id} [delete]
func (ctl *GenderController) DeleteGender(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Gender.
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

// NewGenderController creates and registers handles for the gender controller
func NewGenderController(router gin.IRouter, client *ent.Client) *GenderController {
	dc := &GenderController{
		client: client,
		router: router,
	}
	dc.register()
	return dc
 }
  
 // InitUserController registers routes to the main engine
 func (ctl *GenderController) register() {
	genders := ctl.router.Group("/genders")
	genders.GET("", ctl.ListGender)
	// CRUD
	genders.POST("", ctl.CreateGender)
	genders.GET(":id", ctl.GetGender)
	genders.DELETE(":id", ctl.DeleteGender)
 }