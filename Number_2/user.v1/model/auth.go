package model

// Auth - Send token to user when login
type Auth struct{
	Message		string		`json:"message"`
	Token		string		`json:"token"`
}