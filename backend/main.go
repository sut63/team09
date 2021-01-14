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

type Doctors struct {
	Doctor []Doctor
}
type Doctor struct {
	Name        string
	Age         int
	Email       string
	Password    string
	Address     string
	Educational string
	Phone		string
}
type Genders struct {
	Gender []Gender
}

type Gender struct {
	Gender string
}
type Positions struct {
	Position []Position
}

type Position struct {
	Position string
}
type Titles struct {
	Title []Title
}

type Title struct {
	Title string
}
type Diseases struct {
	Disease []Disease
}

type Disease struct {
	Disease string
}
type Courses struct {
	Course []Course
}

type Course struct {
	Namecourse string
}

type Departments struct {
	Department []Department
}

type Department struct {
	Name string
	// Detail	string
}

type Missions struct {
	Mission []Mission
}

type Mission struct {
	MissionType string
}

type Offices struct {
	Office []Office
}

type Office struct {
	Officename string
}

type Specialists struct {
	Specialist []Specialist
}

type Specialist struct {
	Specialist string
}

// @title SUT SA Example API Playlist Vidoe
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

    client, err := ent.Open("sqlite3", "file:doctor.db?cache=shared&_fk=1")
    if err != nil {
        log.Fatalf("fail to open sqlite3: %v", err)
    }
    defer client.Close()

    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

	v1 := router.Group("/api/v1")
	controllers.NewOfficeController(v1, client)
	controllers.NewDoctorController(v1, client)
	controllers.NewTrainingController(v1, client)
	controllers.NewCourseController(v1, client)
	controllers.NewScheduleController(v1, client)
	controllers.NewDiseaseController(v1, client)
	controllers.NewTitleController(v1, client)
	controllers.NewGenderController(v1, client)
	controllers.NewPositionController(v1, client)
	controllers.NewDepartmentController(v1, client)
	controllers.NewMissionController(v1, client)
	controllers.NewSpecialistController(v1, client)


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
			Position{"แพทย์ปฏิบัติการ"},
			Position{"แพทย์ชำนาญการ"},
			Position{"แพทย์ชำนาญการพิเศษ"},
			Position{"แพทย์เชี่ยวชาญ"},
			Position{"แพทย์ทรงคุณวุฒิ"},
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
			Disease{"โรคหลอดเลือดหัวใจ"},
			Disease{"โรคเบาหวาน"},
			Disease{"โรคความดันโลหิตสูง"},
			Disease{"วัณโรคที่มากับอากาศ"},
			Disease{"โรคปอดเรื้อรัง"},
			Disease{"โรคภูมิแพ้"},
			Disease{"ไม่มีโรคประจำตัว"},
			Disease{"อื่นๆ"},
		},
	}

	for _, di := range diseases.Disease {
		client.Disease.
			Create().
			SetDisease(d.Disease).
			Save(context.Background())
	}

	// Set Department Data
	departments := Departments{
		Department: []Department{
			Department{"แผนกรังสี"},
			Department{"แผนกห้องปฏิบัติการทางการแพทย์"},
			Department{"แผนกศัลยกรรม"},
			Department{"แผนกวิสัญญี"},
			Department{"แผนกกุมารเวช"},
			Department{"แผนกสูตินรีเวช"},
			Department{"แผนกเวชศาสตร์ฟื้นฟู"},
			Department{"แผนกอายุรกรรม"},
			Department{"แผนกจักษุ"},
			Department{"แผนกหู คอ จมูก"},
			Department{"แผนกจิตเวช"},
		},
	}

	for _, de := range departments.Department {
		client.Department.
			Create().
			SetName(de.Name).
			// SetDetail(de.Detail).
			Save(context.Background())
	}

	// Set Mission Data
	missions := Missions{
		Mission: []Mission{
			Mission{"หัวหน้าแผนก"},
			Mission{"รองหัวหน้าแผนก"},
			// Mission{"แผนกผู้ป่วยหนัก"},
			// Mission{"แผนกรังสี"},
			// Mission{"แผนกห้องปฏิบัติการทางการแพทย์"},
			// Mission{"แผนกศัลยกรรม"},
			// Mission{"แผนกวิสัญญี"},
			// Mission{"แผนกกุมารเวช"},
			// Mission{"แผนกสูตินรีเวช"},
			// Mission{"แผนกเวชศาสตร์ฟื้นฟู"},
			// Mission{"แผนกอายุรกรรม"},
			// Mission{"แผนกจักษุ"},
			// Mission{"แผนกหู คอ จมูก"},
			// Mission{"แผนกเภสัชกรรม"},
			// Mission{"แผนกจิตเวช"},
		},
	}

	for _, m := range missions.Mission {
		client.Mission.
			Create().
			SetMissionType(m.MissionType).
			Save(context.Background())
	}

	// Set Officename
	offices := Offices{
		Office: []Office{
			Office{"โรงพยาบาลกรุงเทพ"},
			Office{"โรงพยาบาลบำเน็จนราดูล"},
			Office{"โรงพยาบาลมหาราช"},
			Office{"โรงพยาบาลเซนแมรี่"},
			Office{"โรงพยาบาลหนองคาย"},
			Office{"ศิริราชพยาบาล มหาวิทยาลัยมหิดล"},
			Office{"โรงพยาบาลเจ้าพระยา"},
			Office{"โรงพยาบาลยันฮี"},
			Office{"โรงพยาบาลบำราศนราดูร"},
		},
	}

	for _, of := range offices.Office {
		client.Mission.
			Create().
			SetMissionType(of.Officename).
			Save(context.Background())
	}

	// Set Specialist Data
	specialists := Specialists{
		Specialist: []Specialist{
			//Specialist{"แผนกรังสี"},
			Specialist{"รังสีร่วมรักษาระบบประสาท"},
			Specialist{"รังสีรักษาและมะเร็งวิทยา"},
			Specialist{"รังสีวิทยาวินิจฉัย"},
			//Specialist{"แผนกห้องปฏิบัติการทางการแพทย์"},
			Specialist{"จุลกายวิภาคศาสตร์"},
			Specialist{"ปรสิตวิทยา"},
			Specialist{"เวชศาสตร์การบริการโลหิต"},
			//Specialist{"แผนกศัลยกรรม"},
			Specialist{"ศัลยศาสตร์ตกแต่งและเสริมสร้าง"},
			Specialist{"ศัลยศาสตร์หลอดเลือด"},
			Specialist{"ศัลยศาสตร์ลำไส้ใหญ่และทวารหนัก"},
			//Specialist{"แผนกวิสัญญี"},
			Specialist{"วิสัญญีวิทยาเพื่อการผ่าตัดหัวใจ หลอดเลือดใหญ่และทรวงอก"},
			Specialist{"วิสัญญีวิทยาสำหรับผู้ป่วยโรคทางระบบประสาท"},
			Specialist{"วิสัญญีวิทยาเพื่อการระงับปวด"},
			//Specialist{"แผนกกุมารเวช"},
			Specialist{"กุมารเวชศาสตร์โรคติดเชื้อ"},
			Specialist{"กุมารเวชศาสตร์พัฒนาการและพฤติกรรม"},
			Specialist{"กุมารเวชศาสตร์โรคภูมิแพ้และภูมิคุ้มกัน"},
			//Specialist{"แผนกสูตินรีเวช"},
			Specialist{"เวชศาสตร์มารดาและทารกในครรภ์"},
			Specialist{"มะเร็งนรีเวชวิทยา"},
			Specialist{"เวชศาสตร์การเจริญพันธุ์"},
			//Specialist{"แผนกเวชศาสตร์ป้องกัน"},
			Specialist{"ระบาดวิทยา"},
			Specialist{"อาชีวเวชศาสตร์"},
			Specialist{"เวชศาสตร์การบิน"},
			//Specialist{"แผนกอายุรกรรม"},
			Specialist{"อายุรศาสตร์โรคข้อและรูมาติสซั่ม"},
			Specialist{"อายุรศาสตร์โรคภูมิแพ้และอิมมูโนวิทยาคลินิก"},
			Specialist{"อายุรศาสตร์โรคระบบทางเดินอาหาร"},
			//Specialist{"แผนกจักษุ"},
			Specialist{"จักษุวิทยาโรคต้อหิน"},
			Specialist{"จักษุวิทยาการตรวจคลื่นไฟฟ้า"},
			Specialist{"ศัลยกรรมจักษุตกแต่งและเสริมสร้าง"},
			//Specialist{"แผนกหู คอ จมูก"},
			Specialist{"โสต ศอ นาสิกวิทยา"},
			Specialist{"ศัลยศาสตร์ตกแต่งและเสริมสร้างใบหน้า"},
			//Specialist{"แผนกจิตเวช"},
			Specialist{"จิตเวชศาสตร์ทั่วไป"},
			Specialist{"จิตเวชศาสตร์เด็กและวัยรุ่น"},
		},
	}

	for _, sl := range specialists.Specialist {
		client.Specialist.
			Create().
			SetSpecialist(sl.Specialist).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()

}