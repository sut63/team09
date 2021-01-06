package controllers

import (
	"context"
   	"fmt"
	"strconv"

	"github.com/team09/app/ent"
	"github.com/gin-gonic/gin"
	"github.com/team09/app/ent/mission"
)

// MissionController defines the struct for the mission controller
type MissionController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateMission handles POST requests for adding mission entities
// @Summary Create mission
// @Description Create mission
// @ID create-mission
// @Accept   json
// @Produce  json
// @Param doctor body ent.Mission true "Mission entity"
// @Success 200 {object} ent.Mission
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /missions [post]
func (ctl *MissionController) CreateMission(c *gin.Context) {
	obj := ent.Mission{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "mission binding failed",
		})
		return
	}
	m, err := ctl.client.Mission.
		Create().
		SetMissionname(obj.Missionname).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	c.JSON(200, m)
}

// GetMission handles GET requests to retrieve a mission entity
// @Summary Get a mission entity by ID
// @Description get mission by ID
// @ID get-mission
// @Produce  json
// @Param id path int true "Mission ID"
// @Success 200 {object} ent.Mission
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /missions/{id} [get]
func (ctl *MissionController) GetMission(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	m, err := ctl.client.Mission.
		Query().
		Where(mission.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, m)
}

// ListMission handles request to get a list m mission entities
// @Summary List mission entities
// @Description list mission entities
// @ID list-mission
// @Produce json
// @Param limit  query int false "Limit"
// @Param mfset query int false "Offset"
// @Success 200 {array} ent.Mission
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /missions [get]
func (ctl *MissionController) ListMission(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	mfsetQuery := c.Query("mfset")
	mfset := 0
	if mfsetQuery != "" {
		mfset64, err := strconv.ParseInt(mfsetQuery, 10, 64)
		if err == nil {
			mfset = int(mfset64)
		}
	}

	missions, err := ctl.client.Mission.
		Query().
		Limit(limit).
		Offset(mfset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, missions)
}

// DeleteMission handles DELETE requests to delete a mission entity
// @Summary Delete a mission entity by ID
// @Description get mission by ID
// @ID delete-mission
// @Produce  json
// @Param id path int true "Mission ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /missions/{id} [delete]
func (ctl *MissionController) DeleteMission(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Mission.
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

// NewMissionController creates and registers handles for the mission controller
func NewMissionController(router gin.IRouter, client *ent.Client) *MissionController {
	mi := &MissionController{
		client: client,
		router: router,
	}
	mi.register()
	return mi
}

// InitMissionController registers routes to the main engine
func (ctl *MissionController) register() {
	missions := ctl.router.Group("/missions")
	missions.GET("", ctl.ListMission)
	// CRUD
	missions.POST("", ctl.CreateMission)
	missions.GET(":id", ctl.GetMission)
	missions.DELETE(":id", ctl.DeleteMission)
}