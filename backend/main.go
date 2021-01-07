package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team09/app/controllers"
	_ "github.com/team09/app/docs"
	"github.com/team09/app/ent"
)

type Doctor struct {
	Doctor []Doctor
}
type Doctor struct {
	Name        string
	Age         int
	Email       string
	Pnumber     int
	Address     string
	Educational string
}
type Gender struct {
	Gender []Gender
}

type Gender struct {
	Gender string
}
type Position struct {
	Position []Position
}

type Position struct {
	Position string
}
type Title struct {
	Title []Title
}

type Title struct {
	Title string
}
type Disease struct {
	Disease []Disease
}

type Disease struct {
	Disease string
}
type Course struct {
	Course []Course
}

type Course struct {
	Namecourse string
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

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewWoringtimeController(v1, client)
	controllers.NewOfficeController(v1, client)
	controllers.NewDoctorController(v1, client)
	controllers.NewTrainingController(v1, client)
	controllers.NewCourseController(v1, client)
	controllers.NewScheduleController(v1, client)
	controllers.NewDiseaseController(v1, client)
	controllers.NewTitleController(v1, client)
	controllers.NewGenderController(v1, client)
	controllers.NewPositionController(v1, client)

	//setcourse
	courses := Courses{
		Course: []Course{
			Course{"ลักสูตรฝึกอบรมแพทย์ประจำบ้าน"},
			Course{"หลักสูตรฝึกอบรมแพทย์ประจำบ้านต่อยอด"},
			Course{"หลูกสูตร Ph.D. in Clinical Sciences"},
		},
	}

	for _, co := range courses.Course {
		client.Course.
			Create().
			SetNamecourse(co.Namecourse).
			Save(context.Background())
	}

	// Set Titles Data
	titles := Titles{
		Title: []Title{
			Title{"นายแพทย์"},
			Title{"แพทย์หญิง"},
		},
	}

	for _, t := range titles.Title {
		client.Title.
			Create().
			SetTitle(t.Title).
			Save(context.Background())
	}

	// Set Genders Data
	genders := Genders{
		Gender: []Gender{
			Gender{"ชาย"},
			Gender{"หญิง"},
		},
	}

	for _, g := range genders.Gender {
		client.Gender.
			Create().
			SetGender(g.Gender).
			Save(context.Background())
	}

	// Set Positions Data
	positions := Positions{
		Position: []Position{
			Position{"นายแพทย์ปฏิบัติการ"},
			Position{"นายแพทย์ชำนาญการ"},
			Position{"นายแพทย์ชำนาญการพิเศษ"},
			Position{"นายแพทย์เชี่ยวชาญ"},
			Position{"นายแพทย์ทรงคุณวุฒิ"},
			Position{"อื่นๆ"},
		},
	}

	for _, p := range positions.Position {
		client.Position.
			Create().
			SetPosition(p.Position).
			Save(context.Background())
	}

	// Set Diseases Data
	diseases := Diseases{
		Disease: []Disease{
			Disease{"โรคมะเร็ง"},
			Disease{"โรคหลอดเลือดหัวใจ"},
			Disease{"โรคเบาหวาน"},
			Disease{"โรคความดันโลหิตสูง"},
			Disease{"วัณโรคที่มากับอากาศ"},
			Disease{"โรคปอดเรื้อรัง"},
			Disease{"โรคภูมิแพ้"},
			Disease{"อื่นๆ"},
		},
	}

	for _, d := range diseases.Disease {
		client.Disease.
			Create().
			SetDisease(d.Disease).
			Save(context.Background())
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
