package Utils

import(
	"regexp"
	"strings"
	"fmt"
)

func ValidateEmail(mail string) bool {
	regex := regexp.MustCompile(`^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$`)
	return regex.MatchString(mail)
}

func ValidatePassword(password string) bool {
	var error string

	if strings.Contains(password, " ") {
		error = "Şifre boşluk içeremez"
	}

	if len(password) < 8 {
		error = "Şifre en az 8 karakter olmak zorundadır."
	}

	if len(password) > 32 {
		error = "Şifre en fazla 32 karakter olmalıdır."
	}

	if !regexp.MustCompile(`\d`).MatchString(password) {
		error = "Şifrede sayılar yer almalıdır."
	}

	if !regexp.MustCompile(`[a-zA-Z]`).MatchString(password) {
		error = "Şifrede metin bulunmak zorundadır."
	}

	if !regexp.MustCompile(`[!@#$%^&*(),.?":{}\|_<>]`).MatchString(password) {
		error = "Şifrede özel karakter bulunmak zorundadır."
	}

	if error != "" {
		fmt.Println(error)
	}

	return error == ""
}

func ValidateUsername(username string) bool {
	return !strings.Contains(username, " ")
}