package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/tanvirraj/insta/controllers"
	"github.com/tanvirraj/insta/models"
	"github.com/tanvirraj/insta/templates"
	"github.com/tanvirraj/insta/views"
)

func main() {
	r := chi.NewRouter()
	// r.Use(middleware.Logger)
	tpl := views.Must(views.ParseFs(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFs(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFs(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{
		DB: db,
	}

	userC := controllers.Users{
		UserService: &userService,
	}
	userC.Template.New = views.Must(views.ParseFs(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", userC.New)

	r.Post("/users", userC.Create)

	// tpl =
	// 	r.Get("/signup", controllers.StaticHandler(tpl))

	fmt.Printf("Starting to server at :3030....")
	// http.ListenAndServe(":3000", r)

	log.Fatal(http.ListenAndServe(":3030", r))
}
