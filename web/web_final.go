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

func home_handler(w http.ResponseWriter, r *http.Request) {
	/*
		Function to handle requests to the home page.

		Parameters:
			- w (http.ResponseWriter): HTTP response writer.
			- r (*http.Request): HTTP request.

		Returns:
			Nothing.
	*/
	V := os.Getenv("verbose")
	// Checks if the user has enabled logs
	if V == "true" {
		log.Println("executing home_handler")
	}
	render_template(w, "home", nil)
}

func bye_handler(w http.ResponseWriter, r *http.Request) {
	/*
		Function to handle requests to terminate the server.

		Parameters:
			- w (http.ResponseWriter): HTTP response writer.
			- r (*http.Request): HTTP request.

		Returns:
			Nothing.
	*/
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("executing bye_handler")
	}
	render_template(w, "byebye", nil)
}

func rooms_handler(w http.ResponseWriter, r *http.Request) {
	/*
		Function to handle requests to list rooms.

		Parameters:
			- w (http.ResponseWriter): HTTP response writer.
			- r (*http.Request): HTTP request.

		Returns:
			Nothing.
	*/
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("executing rooms_handler")
	}
	data := res.List_rooms()
	render_template(w, "rooms", data)
}

func available_rooms_handler(w http.ResponseWriter, r *http.Request) {
	/*
		Function to handle requests to list available rooms.

		Parameters:
			- w (http.ResponseWriter): HTTP response writer.
			- r (*http.Request): HTTP request.

		Returns:
			Nothing.
	*/
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("executing available_rooms_handler")
	}
	// Retrieves the parameters/filters of the query
	hasDate := r.URL.Query().Has("date")
	date := r.URL.Query().Get("date")
	hasRoom := r.URL.Query().Has("room")
	room := r.URL.Query().Get("room")
	var data = []any{}
	data = append(data, res.List_rooms())
	if (hasDate && date != "") || (hasRoom && room != "") {
		log.Printf("got POST request. date(%t)=%s, room(%t)=%s\n",
			hasDate, date, hasRoom, room)
		// If a date filter is entered, we retrieve the reservations by date, otherwise by room
		if date != "" {
			data = append(data, res.List_reservations_by_date(date))
		} else {
			i, _ := strconv.Atoi(room)
			data = append(data, res.List_reservations_by_room(i))
		}
	} else {
		data = append(data, res.List_rooms())
	}
	fmt.Println(data)
	render_template(w, "avrooms", data)
}

func cancel_reservation_handler(w http.ResponseWriter, r *http.Request) {
	/*
		Function to handle requests to cancel a reservation.

		Parameters:
			- w (http.ResponseWriter): HTTP response writer.
			- r (*http.Request): HTTP request.

		Returns:
			Nothing.
	*/
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("executing cancel_reservation_handler")
	}
	hasroom := r.URL.Query().Has("room")
	room := r.URL.Query().Get("room")
	if hasroom && room != "" {
		log.Printf("got POST request. room(%t)=%s\n",
			hasroom, room)
		i, _ := strconv.Atoi(room)
		res.Delete_reservation(i)
		// Returns to the reservations list page (content in the constant)
		http.Redirect(w, r, LR, http.StatusFound)
	} else {
		data := res.List_reservations()
		render_template(w, "canres", data)
	}
}

func create_reservation_handler(w http.ResponseWriter, r *http.Request) {
	/*
		Function to handle requests to create a reservation.

		Parameters:
			- w (http.ResponseWriter): HTTP response writer.
			- r (*http.Request): HTTP request.

		Returns:
			Nothing.
	*/
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("executing create_reservation_handler")
	}

	hasDate := r.URL.Query().Has("date")
	date := r.URL.Query().Get("date")
	hasDate2 := r.URL.Query().Has("date2")
	date2 := r.URL.Query().Get("date2")
	hasRoom := r.URL.Query().Has("room")
	room := r.URL.Query().Get("room")

	if hasDate && hasRoom && hasDate2 && date != "" && date2 != "" && room != "" {
		e, _ := strconv.Atoi(room)
		log.Printf("got POST request. date(%t)=%s, date2(%t)=%s, room(%t)=%s=%d\n",
			hasDate, date,
			hasDate2, date2,
			hasRoom, room, e)
		res.Create_reservation(e, res.Convert_string_to_datetime(date), res.Convert_string_to_datetime(date2))
		http.Redirect(w, r, LR, http.StatusFound)
	} else {
		data := res.List_rooms()
		render_template(w, "createres", data)
	}
}

func list_reservations_handler(w http.ResponseWriter, r *http.Request) {
	/*
		Function to handle requests to list reservations.

		Parameters:
			- w (http.ResponseWriter): HTTP response writer.
			- r (*http.Request): HTTP request.

		Returns:
			Nothing.
	*/
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("executing list_reservations_handler")
	}
	data := res.List_reservations()
	render_template(w, "listres", data)
}

func not_found_handler(w http.ResponseWriter, r *http.Request) {
	/*
		Function to handle requests for non-existent routes.

		Parameters:
			- w (http.ResponseWriter): HTTP response writer.
			- r (*http.Request): HTTP request.

		Returns:
			Nothing.
	*/
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("executing notFound_handler")
	}
	render_template(w, "notfound", nil)
}

func render_template(w http.ResponseWriter, tmpl string, data any) {
	/*
		Function to render an HTML template.

		Parameters:
			- w (http.ResponseWriter): HTTP response writer.
			- tmpl (string): Name of the HTML template.
			- data (any): Data to pass to the template.

		Returns:
			Nothing.
	*/
	// str := "web/" + tmpl + ".html"
	htmlstr := tmpl + ".html"
	// _, err := os.Stat(str); 
	// if err != nil {
	// 	http.Redirect(w, r, "/notfound", http.StatusNotFound)
	// 		return
	// }
	err := templates.ExecuteTemplate(w, htmlstr, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Main() {
	/*
		Function serving as the entry point for the web application.
	*/
	var handler http.Handler
	port := os.Getenv("WEB_PORT")
	if port == ""{
		port = "80"
	}
	srv := &http.Server{Addr: ":"+port, Handler: handler}
	// http.HandleFunc("/", home_handler)
	http.HandleFunc("/home", home_handler)
	http.HandleFunc("/die", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bye bye", http.StatusNotFound)
		context := context.Background()
		srv.Shutdown(context)
	})
	http.HandleFunc("/list_salles", rooms_handler)
	http.HandleFunc("/available_salles", available_rooms_handler)
	http.HandleFunc("/notfound", not_found_handler)
	http.HandleFunc("/create_reservation", create_reservation_handler)
	http.HandleFunc("/cancel_reservation", cancel_reservation_handler)
	http.HandleFunc(LR, list_reservations_handler)
	http.HandleFunc("/byebye", bye_handler)
	V := os.Getenv("verbose")
	if V == "true" {
		log.Println("Listening on :"+port)
	}

	go func() {
		http_error := srv.ListenAndServe()
		if http_error != nil {
			if V == "true" {
				log.Println(ROUGE, "While serving HTTP: ", END, http_error)
			}
		}
	}()

	// go log.Fatal(http.ListenAndServe(":9000", nil))
}
