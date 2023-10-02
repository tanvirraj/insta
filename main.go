package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"

	"github.com/tanvirraj/insta/controllers"
	"github.com/tanvirraj/insta/models"
	"github.com/tanvirraj/insta/templates"
	"github.com/tanvirraj/insta/views"
)

func main() {
	r := chi.NewRouter()

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

	sessionService := models.SessionService{
		DB: db,
	}

	userC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}

	userC.Templates.New = views.Must(views.ParseFs(templates.FS, "signup.gohtml", "tailwind.gohtml"))

	userC.Templates.SignIn = views.Must(views.ParseFs(templates.FS, "signin.gohtml", "tailwind.gohtml"))

	r.Get("/signup", TimerMiddleware(userC.New))
	r.Get("/signin", userC.SignIn)
	r.Post("/signin", userC.ProcessSignIn)
	r.Post("/signout", userC.ProcessSignOut)
	r.Post("/users", userC.Create)
	r.Get("/users/me", userC.CurrentUser)

	CsrfMw := csrf.Protect([]byte("qazxswedctgbyhnujmiklopQAZWSXEDC"), csrf.Secure(false))

	fmt.Printf("Starting to server at :3030....")
	log.Fatal(http.ListenAndServe(":3030", CsrfMw(r)))
}

// custom middleware
func TimerMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h(w, r)
		fmt.Println("Total response tiem", time.Since(start))
	}
}
