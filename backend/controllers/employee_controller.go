package controllers

import (
   "context"
   "strconv"
   "github.com/kaews/app/ent"
   "github.com/kaews/app/ent/employee"
   "github.com/gin-gonic/gin"
)

// EmployeeController defines the struct for the employee controller
type EmployeeController struct {
   client *ent.Client
   router gin.IRouter
}

type Employee struct {
	name   string
	email  string
	password string
}


// EmployeeCreate handles POST requests for adding Employee entities
// @Summary Create Employee
// @Description Create Employee
// @ID create-Employee
// @Accept   json
// @Produce  json
// @Param Employee body ent.Employee true "Employee entity"
// @Success 200 {object} ent.Employee
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /employees [post]
func (ctl *EmployeeController) EmployeeCreate(c *gin.Context) {
	obj := ent.Employee{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "employee binding failed",
		})
		return
	}
  
	e, err := ctl.client.Employee.
		Create().
		SetName(obj.Name).
		SetEmail(obj.Email).
		SetPassword(obj.Password).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed", //400 คือหา path ปลายทางไม่เจอ เชื่อมไปไม่ได้ 
		})
		return
	}
  
	c.JSON(200, e)
 }
 

// GetEmployee handles GET requests to retrieve a Employee entity
// @Summary Get a Employee entity by ID
// @Description get Employee by ID
// @ID get-Employee
// @Produce  json
// @Param id path int true "Employee ID"
// @Success 200 {object} ent.Employee
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /employees/{id} [get]
func (ctl *EmployeeController) GetEmployee(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	e, err := ctl.client.Employee.
		Query().
		Where(employee.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(), //404 คือ หาไม่เจอเลย
		})
		return
	}

	c.JSON(200, e)
}

// ListEmployee handles request to get a list of Employee entities
// @Summary List Employee entities
// @Description list Employee entities
// @ID list-Employee
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Employee
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /employees [get]
func (ctl *EmployeeController) ListEmployee(c *gin.Context) {
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

	employees, err := ctl.client.Employee.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, employees)
}

// NewEmployeeController creates and registers handles for the Employee controller
func NewEmployeeController(router gin.IRouter, client *ent.Client) *EmployeeController {
	ec := &EmployeeController{
		client: client,
		router: router,
	}
	ec.register()
	return ec
 }
 
 // InitEmployeeController registers routes to the main engine
 func (ctl *EmployeeController) register() {
	employees := ctl.router.Group("/employees")
	employees.GET("", ctl.ListEmployee)
	employees.POST("", ctl.EmployeeCreate)
	employees.GET(":id", ctl.GetEmployee)
 
 }