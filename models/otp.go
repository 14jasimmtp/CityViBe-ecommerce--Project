package models

type OTP struct{
	Phone string `json:"phone" binding:"required"`
	Otp string `json:"otp" binding:"required"`
}