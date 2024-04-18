package web

import (
	res "LESGO/reservations"
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

var templates = template.Must(template.ParseFiles(
	"web/pages/home.html", "web/pages/avrooms.html",
	"web/pages/byebye.html", "web/pages/canres.html",
	"web/pages/createres.html", "web/pages/rooms.html",
	"web/pages/listres.html", "web/pages/notfound.html"))

const (
	END   = "\033[0m"
	ROUGE = "\033[31;01;51m"
	LR    = "/list_reservations"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("executing homeHandler")
	}
	renderTemplate(w, "home", nil)
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("executing byeHandler")
	}
	renderTemplate(w, "byebye", nil)
}

func roomsHandler(w http.ResponseWriter, r *http.Request) {
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("executing roomsHandler")
	}
	data := res.ListRooms()
	renderTemplate(w, "rooms", data)
}

func availableRoomsHandler(w http.ResponseWriter, r *http.Request) {
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("executing availableRoomsHandler")
	}
	hasDate := r.URL.Query().Has("date")
	date := r.URL.Query().Get("date")
	hasRoom := r.URL.Query().Has("room")
	room := r.URL.Query().Get("room")
	var data = []any{}
	data = append(data, res.ListRooms())
	if (hasDate && date != "") || (hasRoom && room != "") {
		log.Printf("got POST request. date(%t)=%s, room(%t)=%s\n",
			hasDate, date, hasRoom, room)
		if date != "" {
			data = append(data, res.ListReservationsByDate(date))
		} else {
			i, _ := strconv.Atoi(room)
			data = append(data, res.ListReservationsByRoom(i))
		}
	} else {
		data = append(data, res.ListRooms())
	}
	fmt.Println(data)
	renderTemplate(w, "avrooms", data)
}

func cancelReservationHandler(w http.ResponseWriter, r *http.Request) {
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("executing cancelReservationHandler")
	}
	hasroom := r.URL.Query().Has("room")
	room := r.URL.Query().Get("room")
	if hasroom && room != "" {
		log.Printf("got POST request. room(%t)=%s\n",
			hasroom, room)
		i, _ := strconv.Atoi(room)
		res.DeleteReservation(i)
		http.Redirect(w, r, LR, http.StatusFound)
	} else {
		data := res.ListReservations()
		renderTemplate(w, "canres", data)
	}
}

func createReservationHandler(w http.ResponseWriter, r *http.Request) {
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("executing createReservationHandler")
	}

	hasDate := r.URL.Query().Has("date")
	date := r.URL.Query().Get("date")
	hasDate2 := r.URL.Query().Has("date2")
	date2 := r.URL.Query().Get("date2")
	hasRoom := r.URL.Query().Has("room")
	room := r.URL.Query().Get("room")

	if hasDate && hasRoom && hasDate2 {
		e, _ := strconv.Atoi(room)
		log.Printf("got POST request. date(%t)=%s, date2(%t)=%s, room(%t)=%s=%d\n",
			hasDate, date,
			hasDate2, date2,
			hasRoom, room, e)
		res.CreateReservation(e, res.ConvertStringToDatetime(date), res.ConvertStringToDatetime(date2))
		http.Redirect(w, r, LR, http.StatusFound)
	} else {
		data := res.ListRooms()
		renderTemplate(w, "createres", data)
	}
}

func listReservationsHandler(w http.ResponseWriter, r *http.Request) {
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("executing listReservationsHandler")
	}
	data := res.ListReservations()
	renderTemplate(w, "listres", data)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("executing notFoundHandler")
	}
	renderTemplate(w, "notfound", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
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
	// http.HandleFunc("/", homeHandler)
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
	http.HandleFunc(LR, listReservationsHandler)
	http.HandleFunc("/byebye", byeHandler)
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("Listening on :9000")
	}

	go func() {
		httpError := srv.ListenAndServe()
		if httpError != nil {
			if V == "true" {
				log.Println(ROUGE, "While serving HTTP: ", END, httpError)
			}
		}
	}()

	// go log.Fatal(http.ListenAndServe(":9000", nil))
}
