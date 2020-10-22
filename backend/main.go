package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kaews/app/controllers"
	_ "github.com/kaews/app/docs"
	"github.com/kaews/app/ent"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Dentists struct {
	Dentist []Dentist
}

type Dentist struct {
	Name string
}

type Employees struct {
	Employee []Employee
}

type Employee struct {
	Email    string
	Password string
	Name     string
}

type Patients struct {
	Patient []Patient
}

type Patient struct {
	Name string
	Age  int
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewDentistController(v1, client)
	controllers.NewEmployeeController(v1, client)
	controllers.NewMedicalfileController(v1, client)
	controllers.NewPatientController(v1, client)

	// Set dentists Data
	dentists := Dentists{
		Dentist: []Dentist{
			Dentist{"ทพญ.นัดดา มนูญศิลป์"},
			Dentist{"ทพ.ภูวดล โกศลอิทธิกุล"},
			Dentist{"ทพญ.นัดดา มนูญศิลป์"},
			Dentist{"ทพ.วัลลภ จันทร์สว่าง"},
		},
	}

	for _, d := range dentists.Dentist {
		client.Dentist.
			Create().
			SetName(d.Name).
			Save(context.Background())
	}

	// Set Employees Data
	employees := Employees{
		Employee: []Employee{
			Employee{"b6009168@g.sut.ac.th", "055555", "รัตน์ตวัน ขาวฉลาด"},
			Employee{"chanyeol@gmail.com", "12345cy", "ใจรัก นาดาว"},
			Employee{"baekhyun@gmail.com", "baekycute92", "นีรนาท ถิระศุภะ"},
			Employee{"somsri@gmail.com", "02811111", "สมศรี สุขใจ"},
		},
	}

	for _, e := range employees.Employee {
		client.Employee.
			Create().
			SetName(e.Name).
			SetEmail(e.Email).
			SetPassword(e.Password).
			Save(context.Background())
	}

	// Set Patient Data
	patients := Patients{
		Patient: []Patient{
			Patient{"พายุ นาคราช", 23},
			Patient{"พีรพงศ์ สันติวงศ์", 25},
			Patient{"เยาวลักษณ์ สุขทวี", 19},
			Patient{"ศักดินา ไธสง", 50},
		},
	}

	for _, p := range patients.Patient {
		client.Patient.
			Create().
			SetName(p.Name).
			SetAge(p.Age).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
