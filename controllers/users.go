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
	Template struct {
		New Template
	}

	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n\n")
	fmt.Println("Hellow this is from New controller")
	fmt.Printf("\n\n")
	u.Template.New.Excute(w, nil)

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
