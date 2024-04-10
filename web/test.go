package web

import (
	res "LESGO/reservations"
	"html/template"
	"log"
	"net/http"
	"strconv"
	//"os"
	"context"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("executing homeHandler")
	renderTemplate(w, r, "home", nil)
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("executing byeHandler")
	renderTemplate(w, r, "byebye", nil)
}

func roomsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("executing roomsHandler")
	data := res.ListRoomsReturn()
	renderTemplate(w, r, "rooms", data)
}

func availableRoomsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("executing availableRoomsHandler")
	hasDate := r.URL.Query().Has("date")
	date := r.URL.Query().Get("date")
	var data any
	if hasDate && date != ""{
		log.Printf("got POST request. date(%t)=%s\n",
			hasDate, date)
		data = res.AreFreeReturn(date)
	} else {
		data = nil
	}
	renderTemplate(w, r, "avrooms", data)
}

func cancelReservationHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("executing cancelReservationHandler")
	data := res.ListReservationsReturn()
	renderTemplate(w, r, "canres", data)
}

func createReservationHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("executing createReservationHandler")

	hasDate := r.URL.Query().Has("date")
	date := r.URL.Query().Get("date")
	hasRoom := r.URL.Query().Has("room")
	room := r.URL.Query().Get("room")

	if hasDate && hasRoom {
		e, _ := strconv.Atoi(room)
		log.Printf("got POST request. date(%t)=%s, room(%t)=%s=%d\n",
			hasDate, date,
			hasRoom, room, e)
		res.CreateReservation(e, date)
		http.Redirect(w, r, "/list_reservations", http.StatusFound)
	} else {
		data := res.ListRoomsReturn()
		renderTemplate(w, r, "createres", data)
	}
}

func listReservationsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("executing listReservationsHandler")
	data := res.ListReservationsReturn()
	renderTemplate(w, r, "listres", data)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("executing notFoundHandler")
	renderTemplate(w, r, "notfound", nil)
}

var templates = template.Must(template.ParseFiles(
	"web/home.html", "web/avrooms.html",
	"web/byebye.html", "web/canres.html",
	"web/createres.html", "web/rooms.html",
	"web/listres.html", "web/notfound.html"))

func renderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data any) {
	// str := "web/" + tmpl + ".html"
	htmlstr := tmpl + ".html"
	// if _, err := os.Stat(str); err != nil {
	//http.Redirect(w, r, "/notfound", http.StatusNotFound)
	//	return
	//}
	err := templates.ExecuteTemplate(w, htmlstr, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Main() {
	var handler http.Handler
	srv := &http.Server{Addr: ":9000", Handler: handler}
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/die", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bye bye", http.StatusNotFound)
		context := context.Background()
		srv.Shutdown(context)
	})
	http.HandleFunc("/list_salles", roomsHandler)
	http.HandleFunc("/available_salles", availableRoomsHandler)
	http.HandleFunc("/notfound", notFoundHandler)
	http.HandleFunc("/create_reservation", createReservationHandler)
	http.HandleFunc("/cancel_reservation", cancelReservationHandler)
	http.HandleFunc("/list_reservations", listReservationsHandler)
	http.HandleFunc("/byebye", byeHandler)
	log.Println("Listening on :8000")

	go func() {
		httpError := srv.ListenAndServe()
		if httpError != nil {
			log.Println("While serving HTTP: ", httpError)
		}
	}()

	// go log.Fatal(http.ListenAndServe(":9000", nil))
}
