package main

import (
	Actor "github.com/MoriMokata/project/backend/controller/Actor"
	Disease "github.com/MoriMokata/project/backend/controller/Disease"
	DrugAllergy "github.com/MoriMokata/project/backend/controller/DrugAllergy"
	LabResult "github.com/MoriMokata/project/backend/controller/LabResult"
	MedicalRecord "github.com/MoriMokata/project/backend/controller/MedicalRecord"
	MedicalHistory "github.com/MoriMokata/project/backend/controller/MedicalHistory"
	Refer "github.com/MoriMokata/project/backend/controller/Refer"
	Screening "github.com/MoriMokata/project/backend/controller/Screening"
	"github.com/MoriMokata/project/backend/entity"
	"github.com/MoriMokata/project/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())
	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			//api Diseases
			protected.GET("/api/ListDiseases", Disease.ListDiseases)

			//api DrugAllergy
			protected.GET("/api/CreateDrugAllergy", DrugAllergy.CreateDrugAllergy)
			protected.GET("/api/ListDrugAllergy", DrugAllergy.ListDrugAllergy)
			protected.GET("/api/ListDrug", DrugAllergy.ListDrug)

			//api MedicalRecord
			protected.GET("/api/ListMedicalRecord", MedicalRecord.ListMedicalRecord)
			protected.GET("/api/ListHealthInsurance", MedicalRecord.ListHealthInsurance)
			protected.GET("/api/ListNameTitle", MedicalRecord.ListNameTitle)
			protected.GET("/api/CreateMedicalRecord", MedicalRecord.CreateMedicalRecord)

			//api MedicalHistory
			protected.GET("/api/ListDepartments", MedicalHistory.ListDepartments)
			protected.GET("/api/CreateMedicalHistory", MedicalHistory.CreateMedicalHistory)
			protected.GET("/api/ListMedicalHistories", MedicalHistory.ListMedicalHistories)

			//api Refer
			protected.GET("/api/ListHospitals", Refer.ListHospitals)
			protected.GET("/api/CreateRefer", Refer.CreateRefer)
			protected.GET("/api/ListRefer", Refer.ListRefer)

			//api Screening
			protected.GET("/api/CreateScreening", Screening.CreateScreening)
			protected.GET("/api/ListScreening", Screening.ListScreening)
			protected.GET("/api/PreloadScreenings", Screening.PreloadScreenings)
			
			//api ListLabResult
			protected.GET("/api/ListLabType", LabResult.ListLabType)
			protected.GET("/api/ListLabRoom", LabResult.ListLabRoom)
			protected.POST("/api/CreateLabResult", LabResult.CreateLabResult)
			protected.GET("/api/ListLabResult", LabResult.ListLabResult)
			
			

		}
	}
	//Get func login/Actor
	r.POST("/api/LoginDoctor", Actor.LoginDoctor)
	r.POST("/api/LoginMedicalRecordOfficer", Actor.LoginMedicalRecordOfficer)
	r.POST("/api/LoginMedicalTech", Actor.LoginMedicalTech)
	r.POST("/api/LoginNurse", Actor.LoginNurse)
	

	// Run the server
	r.Run()
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
