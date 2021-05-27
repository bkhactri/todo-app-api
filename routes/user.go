package routes

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/tobias/todo-app/config"
	db "github.com/tobias/todo-app/database"
	"github.com/tobias/todo-app/models"
	res "github.com/tobias/todo-app/utils"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if govalidator.IsNull(email) || govalidator.IsNull(password) {
		res.JSON(w, 400, "Field can not empty")
		return
	}

	email = models.Santize(email)
	password = models.Santize(password)

	collection := db.ConnectUser()
	var user models.User

	if err := collection.Where("email = ?", email).First(&user).Error; err != nil || !govalidator.IsEmail(email) {
		res.JSON(w, 400, "Email not found")
		return
	}

	if err := models.CompareHashedPassword(user.Password, password); err != nil {
		res.JSON(w, 401, " Password incorrect")
		return
	}
	config.Store(w, r, email)
	res.JSON(w, 200, "Sign in successfully")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	if govalidator.IsNull(name) || govalidator.IsNull(email) || govalidator.IsNull(password) || govalidator.IsNull(confirmPassword) {
		res.JSON(w, 400, "Field can not empty")
		return
	}

	if !govalidator.IsEmail(email) {
		res.JSON(w, 400, "Email is incorrect")
		return
	}

	if password != confirmPassword {
		res.JSON(w, 400, "Confirm Password is incorrect")
		return
	}

	name = models.Santize(name)
	email = models.Santize(email)
	password = models.Santize(password)

	collection := db.ConnectUser()
	var user models.User

	if err := collection.Where("email = ?", email).First(&user).Error; err == nil {
		res.JSON(w, 409, "Email has been used")
		return
	}

	hashedPassword, err := models.HashPassword(password)

	if err != nil {
		res.JSON(w, 500, "Register has failed")
		return
	}

	newUser := models.User{Name: name, Email: email, Password: hashedPassword}
	if err := collection.Create(&newUser).Error; err != nil {
		res.JSON(w, 500, "Can not create User")
		return
	}

	res.JSON(w, 200, "Register Successfully")
}
