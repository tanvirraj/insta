package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

// error wrapper
func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFs(fs fs.FS, partterns ...string) (Template, error) {
	tpl, err := template.ParseFS(fs, partterns...)

	if err != nil {
		return Template{}, fmt.Errorf("parsing Error %v", tpl)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing Error %v", tpl)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Excute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset-utf8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("Execute Template %v", err)
		http.Error(w, "Error in Excuting template", http.StatusInternalServerError)
		return
	}

}
