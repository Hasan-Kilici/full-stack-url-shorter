package Database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"urlShorter/Utils"
	"fmt"
    b64 "encoding/base64"
    "errors"
)

func Login(Email, Password string) (string, error) {
	query := "SELECT token, password FROM users WHERE Email = $1"
	row := db.QueryRow(query, Email)

	var token, hashedPassword string
	row.Scan(&token, &hashedPassword)
	RealPassword, _ := b64.URLEncoding.DecodeString(hashedPassword)

	if(string(RealPassword) == Password){
		return token, nil
	}

	return "Kullanıcı Bulunamadı", errors.New("Email ya da şifre yanlış")
}

func Register(Username, Email, Password string) bool {	
	if !FindUserByEmail(Email){
		Token := Utils.GenerateToken(31)
		NewPassword := b64.URLEncoding.EncodeToString([]byte(Password))
		query := "INSERT INTO users (Token, Username, Email, Password, Perm) VALUES ($1, $2, $3, $4, $5)"
		_, err := db.Exec(query,Token,Username,Email,NewPassword,"User")
		if err != nil {
			fmt.Println(err)
			return false
		}
		
		return true
	} else {
		fmt.Println("Bu Email zaten Kullanılıyor")
		return false
	}
}

func FindUserByEmail(Email string) bool {
    query := "SELECT COUNT(*) FROM users WHERE Email = $1"
    row := db.QueryRow(query, Email)

    var count int
    err := row.Scan(&count)
    if err != nil {
        return false
    }

    if count == 0 {
        return false
    }

    return true
}

func FindUserByToken(Token string) (User, error){
    query := "SELECT * FROM users WHERE token = $1"
    row := db.QueryRow(query, Token)

    var user User
    err := row.Scan(&user.Token, &user.Username, &user.Email ,&user.Password, &user.Perm)
    if err != nil {
        if err == sql.ErrNoRows {
            return User{}, fmt.Errorf("kullanıcı bulunamadı")
        }
        return User{}, err
    }
    return user, nil
}