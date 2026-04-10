package handlers

// GetAppointments godoc
// @Summary Get all appointments
// @Tags appointments
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Appointment
// @Router /appointments [get]
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
// @Security BearerAuth
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
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param appointment body models.Appointment true "Appointment"
// @Success 201 {object} models.Appointment
// @Router /appointments [post]
func CreateAppointment(c *gin.Context) {
	var appointment models.Appointment

	// 1️⃣ JSON оқу
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 🔥 2️⃣ ОСЫ ЖЕРГЕ ҚОСАСЫҢ
	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	userIDFloat, ok := userIDValue.(float64)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid user_id"})
		return
	}

	appointment.UserID = uint(userIDFloat)

	// 3️⃣ OPTIONAL VALIDATION
	if appointment.DoctorID == 0 {
		c.JSON(400, gin.H{"error": "DoctorID is required"})
		return
	}

	// 4️⃣ DB SAVE
	if err := config.DB.Create(&appointment).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 5️⃣ PRELOAD
	config.DB.Preload("Doctor").First(&appointment, appointment.ID)

	c.JSON(201, appointment)
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
