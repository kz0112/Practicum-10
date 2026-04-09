package main

import (
	"log"

	"Private-medical-clinic.backend/config"
	"Private-medical-clinic.backend/handlers"
	"Private-medical-clinic.backend/middleware"
	"Private-medical-clinic.backend/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "Private-medical-clinic.backend/docs"
)

func main() {

	// 🔌 DB
	config.ConnectDB()

	// 🔄 Migration
	config.DB.AutoMigrate(
		&models.User{},
		&models.Doctor{},
		&models.Appointment{},
	)

	// 🚀 Router
	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// =========================
	// 🔐 AUTH
	// =========================
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// =========================
	// 👨‍⚕️ DOCTORS (public)
	// =========================
	r.GET("/doctors", handlers.GetDoctors)
	r.POST("/doctors", handlers.CreateDoctor)

	// =========================
	// 🔒 PROTECTED ROUTES
	// =========================
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/users", handlers.GetUsers)
	auth.GET("/users/:id", handlers.GetUserByID)

	auth.GET("/appointments", handlers.GetAppointments)
	auth.POST("/appointments", handlers.CreateAppointment)

	// =========================
	// ▶️ START
	// =========================
	log.Println("🚀 Server started at http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
