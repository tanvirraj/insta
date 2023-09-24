package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

// error wrapper
func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFs(fs fs.FS, partterns ...string) (Template, error) {
	tpl := template.New(partterns[0])

	tpl = tpl.Funcs(template.FuncMap{
		"csrfField": func() (template.HTML, error) {
			return "", fmt.Errorf("csrfField not implemented")
		},
	})

	tpl, err := tpl.ParseFS(fs, partterns...)

	if err != nil {
		return Template{}, fmt.Errorf("parsing Error %v", tpl)
	}

	return Template{
		htmlTpl: tpl,
	}, nil
}

// func Parse(filepath string) (Template, error) {
// 	tpl, err := template.ParseFiles(filepath)
// 	if err != nil {
// 		return Template{}, fmt.Errorf("parsing Error %v", tpl)
// 	}
// 	return Template{
// 		htmlTpl: tpl,
// 	}, nil
// }

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	tpl, err := t.htmlTpl.Clone()
	if err != nil {
		log.Printf("cloning template: %v", err)
		http.Error(w, "There was an error rendering the page.", http.StatusInternalServerError)
		return
	}

	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
		},
	)

	w.Header().Set("Content-Type", "text/html; charset-utf8")
	err = tpl.Execute(w, data)
	if err != nil {
		log.Printf("Execute Template %v", err)
		http.Error(w, "Error in Excuting template", http.StatusInternalServerError)
		return
	}

}
