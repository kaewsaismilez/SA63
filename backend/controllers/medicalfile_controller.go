package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaews/app/ent"
	"github.com/kaews/app/ent/dentist"
	"github.com/kaews/app/ent/employee"
	"github.com/kaews/app/ent/medicalfile"
	"github.com/kaews/app/ent/patient"
)

// MedicalfileController defines the struct for the medicalfile controller
type MedicalfileController struct {
	client *ent.Client
	router gin.IRouter
}

type Medicalfile struct { 
	Dentist  int
	Employee int
	Patient  int
	Detial   string
	Added    string
}

// MedicalfileCreate handles POST requests for adding medicalfile entities
// @Summary Create medicalfile
// @Description Create medicalfile
// @ID create-medicalfile
// @Accept   json
// @Produce  json
// @Param Medicalfile body ent.Medicalfile true "Medicalfile entity"
// @Success 200 {object} ent.Medicalfile
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalfiles [post]
func (ctl *MedicalfileController) MedicalfileCreate(c *gin.Context) {
	obj := Medicalfile{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "medicalfile binding failed",
		})
		return 
	}

	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
		return
	}

	d, err := ctl.client.Dentist.
		Query().
		Where(dentist.IDEQ(int(obj.Dentist))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "dentist not found",
		})
		return
	}

	e, err := ctl.client.Employee.
		Query().
		Where(employee.IDEQ(int(obj.Employee))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "employee not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Added)

	m, err := ctl.client.Medicalfile.
		Create().
		SetPatient(p).
		SetDentist(d).
		SetEmployee(e).
		SetDetail(obj.Detial).
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
		"data":   m,
	})

}

// GetMedicalfile handles GET requests to retrieve a Medicalfile entity
// @Summary Get a Medicalfile entity by ID
// @Description get Medicalfile by ID
// @ID get-Medicalfile
// @Produce  json
// @Param id path int true "Medicalfile ID"
// @Success 200 {object} ent.Medicalfile
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalfiles/{id} [get]
func (ctl *MedicalfileController) GetMedicalfile(c *gin.Context) { //ดึงข้อมูลตาม pk มาใส่ใน list
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	m, err := ctl.client.Medicalfile.
		Query().
		Where(medicalfile.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, m)
}

// ListMedicalfile handles request to get a list of Medicalfile entities
// @Summary List Medicalfile entities
// @Description list Medicalfile entities
// @ID list-Medicalfile
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Medicalfile
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalfiles [get]
func (ctl *MedicalfileController) ListMedicalfile(c *gin.Context) { //ดึง combobox 
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
	medicalfiles, err := ctl.client.Medicalfile.
		Query().
		WithDentist().
		WithEmployee().
		WithPatient().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, medicalfiles)
}

// NewMedicalfileController creates and registers handles for the Medicalfile controller
func NewMedicalfileController(router gin.IRouter, client *ent.Client) *MedicalfileController {
	mc := &MedicalfileController{
		client: client,
		router: router,
	}
	mc.register()
	return mc
}

// InitMedicalfileController registers routes to the main engine
func (ctl *MedicalfileController) register() {
	medicalfiles := ctl.router.Group("/medicalfiles")
	medicalfiles.GET("", ctl.ListMedicalfile)
	medicalfiles.POST("", ctl.MedicalfileCreate)
	medicalfiles.GET(":id", ctl.GetMedicalfile)
}
