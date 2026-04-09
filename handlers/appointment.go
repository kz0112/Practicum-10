package handlers

import (
	"net/http"

	"Private-medical-clinic.backend/config"
	"Private-medical-clinic.backend/models"

	"github.com/gin-gonic/gin"
)

// =========================
// 📅 GET /appointments
// =========================
// GetAppointments godoc
// @Summary Get all appointments
// @Tags appointments
// @Produce json
// @Success 200 {array} models.Appointment
// @Router /appointments [get]
func GetAppointments(c *gin.Context) {
	var appointments []models.Appointment

	if err := config.DB.
		Preload("User").
		Preload("Doctor").
		Find(&appointments).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch appointments",
		})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

// =========================
// ➕ POST /appointments
// =========================
// CreateAppointment godoc
// @Summary Create appointment
// @Tags appointments
// @Accept json
// @Produce json
// @Param appointment body models.Appointment true "Appointment"
// @Success 201 {object} models.Appointment
// @Router /appointments [post]
func CreateAppointment(c *gin.Context) {
	var appointment models.Appointment

	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if appointment.UserID == 0 || appointment.DoctorID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "UserID and DoctorID are required",
		})
		return
	}

	if err := config.DB.Create(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create appointment",
		})
		return
	}

	// 🔥 preload response
	config.DB.Preload("User").Preload("Doctor").First(&appointment, appointment.ID)

	c.JSON(http.StatusCreated, appointment)
}

// =========================
// 🔍 GET /appointments/:id
// =========================
// GetAppointmentByID godoc
// @Summary Get appointment by ID
// @Tags appointments
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} models.Appointment
// @Router /appointments/{id} [get]
func GetAppointmentByID(c *gin.Context) {
	var appointment models.Appointment
	id := c.Param("id")

	if err := config.DB.
		Preload("User").
		Preload("Doctor").
		First(&appointment, id).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Appointment not found",
		})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

// =========================
// ❌ DELETE /appointments/:id
// =========================
// DeleteAppointment godoc
// @Summary Delete appointment
// @Tags appointments
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} map[string]string
// @Router /appointments/{id} [delete]
func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Appointment{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete appointment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Appointment deleted",
	})
}
