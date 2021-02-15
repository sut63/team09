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
	Password	string
	Address     string
	Educational string
	Phone    	string
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
}

type Specialdoctor struct {
	Specialdoctor []Specialdoctor
}

type Specialdoctors struct {
	Other string
	Roomnumber string
	Doctorid string
}

type Missions struct {
	Mission []Mission
}

type Mission struct {
	Mission string
}

type Offices struct {
	Office []Office
}

type Office struct {
	Officename string
}

type Extradoctors struct {
	Extradoctor []Extradoctor
}

type Extradoctor struct {
	Specialname string
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

    client, err := ent.Open("sqlite3", "file:dinformation.db?cache=shared&_fk=1")
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
	controllers.NewExtradoctorController(v1, client)
	controllers.NewDetailController(v1, client)
	controllers.NewSpecialdoctorController(v1, client)

	//setcourse
	courses := Courses{
		Course: []Course{
			Course{"หลักสูตรฝึกอบรมแพทย์ประจำบ้าน"},
			Course{"หลักสูตรฝึกอบรมแพทย์ประจำบ้านต่อยอด"},
			Course{"หลักสูตร Ph.D. in Clinical Sciences"},
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
			Disease{"โรคมะเร็ง"},
			Disease{"โรคหลอดเลือดหัวใจ"},
			Disease{"โรคเบาหวาน"},
			Disease{"โรคความดันโลหิตสูง"},
			Disease{"วัณโรคที่มากับอากาศ"},
			Disease{"โรคปอดเรื้อรัง"},
			Disease{"โรคภูมิแพ้"},
			Disease{"อื่นๆ"},
			Disease{"ไม่มีโรค"},
		},
	}

	for _, d := range diseases.Disease {
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
			Save(context.Background())
	}

	// Set Mission Data
	missions := Missions{
		Mission: []Mission{
			Mission{"หัวหน้าแผนก,รองหัวหน้าแผนก,พยาบาลวิชาชีพ,ผู้ช่วยพยาบาล,เจ้าหน้าที่ธุรการ"},
            Mission{"หัวหน้าแผนก,รองหัวหน้าแผนก,พยาบาลวิชาชีพ,ผู้ช่วยพยาบาล,เจ้าพนักงานงานการแพทย์"},
            Mission{"หัวหน้าแผนก,รองหัวหน้าแผนก,พยาบาลชำนาญการพิเศษ,ผู้ช่วยพยาบาล,เจ้าพนักงานงานการแพทย์"},
		},
	}

	for _, m := range missions.Mission {
		client.Mission.
			Create().
			SetMission(m.Mission).
			Save(context.Background())
	}

	// Set Special Data
	extradoctors := Extradoctors{
		Extradoctor: []Extradoctor{
			//Special{"แผนกรังสี"},
			Extradoctor{"รังสีร่วมรักษาระบบประสาท"},
			Extradoctor{"รังสีรักษาและมะเร็งวิทยา"},
			Extradoctor{"รังสีวิทยาวินิจฉัย"},
			//Special{"แผนกห้องปฏิบัติการทางการแพทย์"},
			Extradoctor{"จุลกายวิภาคศาสตร์"},
			Extradoctor{"ปรสิตวิทยา"},
			Extradoctor{"เวชศาสตร์การบริการโลหิต"},
			//Special{"แผนกศัลยกรรม"},
			Extradoctor{"ศัลยศาสตร์ตกแต่งและเสริมสร้าง"},
			Extradoctor{"ศัลยศาสตร์หลอดเลือด"},
			Extradoctor{"ศัลยศาสตร์ลำไส้ใหญ่และทวารหนัก"},
			//Special{"แผนกวิสัญญี"},
			Extradoctor{"วิสัญญีวิทยาเพื่อการผ่าตัดหัวใจ หลอดเลือดใหญ่และทรวงอก"},
			Extradoctor{"วิสัญญีวิทยาสำหรับผู้ป่วยโรคทางระบบประสาท"},
			Extradoctor{"วิสัญญีวิทยาเพื่อการระงับปวด"},
			//Special{"แผนกกุมารเวช"},
			Extradoctor{"กุมารเวชศาสตร์โรคติดเชื้อ"},
			Extradoctor{"กุมารเวชศาสตร์พัฒนาการและพฤติกรรม"},
			Extradoctor{"กุมารเวชศาสตร์โรคภูมิแพ้และภูมิคุ้มกัน"},
			//Special{"แผนกสูตินรีเวช"},
			Extradoctor{"เวชศาสตร์มารดาและทารกในครรภ์"},
			Extradoctor{"มะเร็งนรีเวชวิทยา"},
			Extradoctor{"เวชศาสตร์การเจริญพันธุ์"},
			//Special{"แผนกเวชศาสตร์ป้องกัน"},
			Extradoctor{"ระบาดวิทยา"},
			Extradoctor{"อาชีวเวชศาสตร์"},
			Extradoctor{"เวชศาสตร์การบิน"},
			//Special{"แผนกอายุรกรรม"},
			Extradoctor{"อายุรศาสตร์โรคข้อและรูมาติสซั่ม"},
			Extradoctor{"อายุรศาสตร์โรคภูมิแพ้และอิมมูโนวิทยาคลินิก"},
			Extradoctor{"อายุรศาสตร์โรคระบบทางเดินอาหาร"},
			//Special{"แผนกจักษุ"},
			Extradoctor{"จักษุวิทยาโรคต้อหิน"},
			Extradoctor{"จักษุวิทยาการตรวจคลื่นไฟฟ้า"},
			Extradoctor{"ศัลยกรรมจักษุตกแต่งและเสริมสร้าง"},
			//Specialist{"แผนกหู คอ จมูก"},
			Extradoctor{"โสต ศอ นาสิกวิทยา"},
			Extradoctor{"ศัลยศาสตร์ตกแต่งและเสริมสร้างใบหน้า"},
			//Special{"แผนกจิตเวช"},
			Extradoctor{"จิตเวชศาสตร์ทั่วไป"},
			Extradoctor{"จิตเวชศาสตร์เด็กและวัยรุ่น"},
		},
	}

	for _, sl := range extradoctors.Extradoctor {
		client.Extradoctor.
			Create().
			SetSpecialname(sl.Specialname).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()

}