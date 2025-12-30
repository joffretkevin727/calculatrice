package controller

import (
	"bytes"
	"net/http"
	"proj2/structure"
	"strconv"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	// Analyse et rend le template spécifié
	tmpl := template.Must(template.ParseFiles("template/" + filename))
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, data); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}
	w.Write(buf.Bytes())
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := structure.PageData{
		Operation: []string{"+", "-", "*", "/"},
	}
	RenderTemplate(w, "home.html", data)
}
func Calculer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	valeur1str := r.FormValue("valeur1")
	valeur2str := r.FormValue("valeur2")
	operation := r.FormValue("operation")

	valeur1, err1 := strconv.ParseFloat(valeur1str, 64)
	valeur2, err2 := strconv.ParseFloat(valeur2str, 64)

	var res float64
	var error string = ""
	if err1 != nil || err2 != nil {
		error = "Veuillez saisir des nombres valides."
	} else {
		switch operation {
		case "+":
			res = valeur1 + valeur2
		case "-":
			res = valeur1 - valeur2
		case "*":
			res = valeur1 * valeur2
		case "/":
			if valeur2 != 0 {
				res = valeur1 / valeur2
			} else {
				error = "division par zero"
			}
		default:
			if operation == "" {
				error = "aucune operation n'est selectionner"
			}
		}
	}

	data := structure.PageData{
		Operation: []string{"+", "-", "*", "/"},
		Resultat:  res,
		Error:     error,
		HasResult: error == "",
	}
	RenderTemplate(w, "home.html", data)
}
