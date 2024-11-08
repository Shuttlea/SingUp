/*
 * Sing Up / Sing In
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"singup/token"
)

var DB *gorm.DB

func AuthActivateGet(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello from Activate handler!")
}

func AuthRegisterPost(w http.ResponseWriter, r *http.Request) {
	// read email and password from form
	err := r.ParseForm()
	if err != nil {
		slog.Error("AuthRegisterPost", "error", err)
	}
	email := r.PostFormValue("email")
	password := r.PostFormValue(("text"))

	// check email for existing in database
	var user1 User
	res := DB.Where("email = ?", email).Find(&user1)
	if res.RowsAffected != 0 {
		w.WriteHeader(409)
		tmpl, err := template.ParseFiles("./ui/html/singup.html")
		if err != nil {
			slog.Error("Parse template file", "error", err)
			return
		}
		tmpl.Execute(w, "This email is already registered")
		return
	}

	// make token with email and expired time 24 hours
	tokenString, err := token.MakeToken(email)
	if err != nil {
		slog.Error("Make Token", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	slog.Debug("Generated","token",tokenString)

	// DEBUG 
	// a,b := token.MakeToken("adsff")
	em,ex,_ := token.VerifyToken(tokenString)
	slog.Debug("Decode token","email",em,"exp",ex, "err", err)

	// encode password for storage
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		slog.Error("AuthRegister bcrypt", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// save user to database
	user := User{Email: email, Password: string(hash)}
	result := DB.Create(&user)
	if result.Error != nil {
		slog.Error("Error create new user:", "error", result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, "./ui/html/check_email.html")
}
