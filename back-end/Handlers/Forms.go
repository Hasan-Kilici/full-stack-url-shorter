package Handlers

type RegisterForm struct {
	Username 	string 	`json:"username"`
	Email 		string 	`json:"email"`
	Pass 		string 	`json:"pass"`
}

type LoginForm struct {
	Email 	string 	`json:"email"`
	Pass 	string 	`json:"pass"`
}

type CreateUrlForm struct {
	Redirect	string		`json:"redirect"`
}

type ClickedForm struct {
	IP			string		`json:"ip"`
	Country		string		`json:"country"`
	City		string		`json:"city"`
	Latitude	string		`json:"latitude"`
	Longitude	string		`json:"longitude"`
}