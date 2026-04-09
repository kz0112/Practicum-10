package models

type Doctor struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" binding:"required"`
	Specialty string `json:"specialty" binding:"required"`

	// 🔗 Relations
	Appointments []Appointment `json:"appointments"`
}
