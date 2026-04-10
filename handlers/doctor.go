package handlers

import (
	"net/http"

	"Private-medical-clinic.backend/config"
	"Private-medical-clinic.backend/models"

	"github.com/gin-gonic/gin"
)

// =========================
// 👨‍⚕️ GET /doctors
// =========================
// GetDoctors godoc
// @Summary Get all doctors
// @Tags doctors
// @Produce json
// @Success 200 {array} models.Doctor
// @Router /doctors [get]
func GetDoctors(c *gin.Context) {
	var doctors []models.Doctor

	if err := config.DB.
		Preload("Appointments").
		Preload("Appointments.User").
		Find(&doctors).Error; err != nil {

		c.JSON(500, gin.H{
			"error": "Failed to fetch doctors",
		})
		return
	}

	c.JSON(200, doctors)
}

// =========================
// ➕ POST /doctors
// =========================
// CreateDoctor godoc
// @Summary Create doctor
// @Tags doctors
// @Accept json
// @Produce json
// @Param doctor body models.Doctor true "Doctor"
// @Success 201 {object} models.Doctor
// @Router /doctors [post]
func CreateDoctor(c *gin.Context) {
	var doctor models.Doctor

	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if doctor.Name == "" || doctor.Specialty == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name and Specialty are required",
		})
		return
	}

	if err := config.DB.Create(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create doctor",
		})
		return
	}

	c.JSON(http.StatusCreated, doctor)
}
