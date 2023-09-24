package controllers

import (
	"fmt"
	"net/http"

	"github.com/tanvirraj/insta/models"
)

// we are parsing templates and putting it in views.templates file
// but we also don't want to show user comes controller and get parse error
// when we parsing templates we putting users templates into ` Users Struct````
// So here, Template is a struct inside Users Template which is creating new template struct
// then In the Main file first we create new instance Users Struct
//userC := controllers.Users{}
//And put parsed tempalte into Users struct
//userC.Template.New = views.Must(views.ParseFs(templates.FS, "signup.gohtml", "tailwind.gohtml"))

//userC := controllers.Users{}
//userC.Template.New = views.Must(views.ParseFs(templates.FS, "signup.gohtml", "tailwind.gohtml"))

//follwoing struct can be written this way too

// type Users struct {
// 	Templates UserTemplates
// }

// type UserTemplates struct {
// 	New views.Template
// }

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}

	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {

	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, r, data)

}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	fmt.Print(err)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "user created : %w", user)

}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, r, data)

}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.UserService.Authenticate(data.Email, data.Password)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:  "email",
		Value: user.Email,
		Path:  "/",
	}

	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "user details : %+v", user)

}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	email, err := r.Cookie("email")
	if err != nil {
		fmt.Fprint(w, "The email cookkie could not be read")
		return
	}
	fmt.Fprintf(w, "Email Cokkie: %s\n", email.Value)
	fmt.Fprintf(w, "Heder  %+v\n", r.Header)
}
