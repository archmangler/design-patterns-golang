package main

import (
	"net/http"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Hello Looniverse!")

	app.render(w, "home.page.gohtml", nil)
}
