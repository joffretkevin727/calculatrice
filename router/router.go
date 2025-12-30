package router

import (
	"net/http"
	"proj2/controller"
)

// CETTE FONCTION INITIALISE UN SERVEUR MUX, CONFIGURE LES ROUTES ET LES FICHIERS STATIQUES ET LE RETOURNE
func New() *http.ServeMux {
	mux := http.NewServeMux()

	//------------------- ROUTES -----------------------
	mux.HandleFunc("/home", controller.Home)
	mux.HandleFunc("/calculer", controller.Calculer)
	//--------------------------------------------------
	//-------------------- CSS -------------------------
	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	//--------------------------------------------------
	return mux
}
