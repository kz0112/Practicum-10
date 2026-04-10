package models

import "time"

type Appointment struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	UserID   uint      `json:"user_id"`
	DoctorID uint      `json:"doctor_id" binding:"required"`
	Date     time.Time `json:"date" binding:"required"`

	// 🔗 Relations
	User   User   `json:"user" gorm:"foreignKey:UserID" binding:"-"`
	Doctor Doctor `json:"doctor" gorm:"foreignKey:DoctorID" binding:"-"`
}
