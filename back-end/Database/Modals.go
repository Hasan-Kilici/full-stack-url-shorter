package Database

type User struct{
	Token		string
	Username	string
	Email		string
	Password	string
	Perm		string
}

type Url struct {
	Token		string
	OwnerToken	string
	Redirect	string
	Click		int
}

type Clicked struct{
	Token 					string
	UrlToken 				string
	IP 						string
	Country 				string
	City 					string
	Latitude 				string
 	Longitude				string
}