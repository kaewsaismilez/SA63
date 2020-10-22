package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaews/app/ent"
	"github.com/kaews/app/ent/patient"
)

// PatientController defines the struct for the patient controller
type PatientController struct {
	client *ent.Client
	router gin.IRouter
}

type Patient struct {
	name  string
	age   int
}

// PatientCreate handles POST requests for adding Patient entities
// @Summary Create Patient
// @Description Create Patient
// @ID create-Patient
// @Accept   json
// @Produce  json
// @Param Patient body ent.Patient true "Patient entity"
// @Success 200 {object} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients [post]
func (ctl *PatientController) PatientCreate(c *gin.Context) {
	obj := ent.Patient{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "patient binding failed",
		})
		return
	}

	p, err := ctl.client.Patient.
		Create().
		SetName(obj.Name).
		SetAge(obj.Age).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, p)
}

// GetPatient handles GET requests to retrieve a Patient entity
// @Summary Get a Patient entity by ID
// @Description get Patient by ID
// @ID get-Patient
// @Produce  json
// @Param id path int true "Patient ID"
// @Success 200 {object} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients/{id} [get]
func (ctl *PatientController) GetPatient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// ListPatient handles request to get a list of Patient entities
// @Summary List Patient entities
// @Description list Patient entities
// @ID list-Patient
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients [get]
func (ctl *PatientController) ListPatient(c *gin.Context) {
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

	patients, err := ctl.client.Patient.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, patients)
}

// NewPatientController creates and registers handles for the Patient controller
func NewPatientController(router gin.IRouter, client *ent.Client) *PatientController {
	pc := &PatientController{
		client: client,
		router: router,
	}
	pc.register()
	return pc
}

// InitPatientController registers routes to the main engine
func (ctl *PatientController) register() {
	patients := ctl.router.Group("/patients")
	patients.GET("", ctl.ListPatient)
	patients.POST("", ctl.PatientCreate)
	patients.GET(":id", ctl.GetPatient)

}
