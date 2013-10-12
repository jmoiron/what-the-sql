package main

import (
	"encoding/json"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/pat"
	"github.com/gorilla/schema"
	"github.com/jmoiron/mandira"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const PRELOAD_TEMPLATES = false

var templates *mandira.Loader
var decoder = schema.NewDecoder()

type handler http.HandlerFunc
type Router struct {
	pat.Router
}

func (r *Router) Get(pattern string, handler func(w http.ResponseWriter, req *http.Request)) *mux.Route {
	return r.Router.Get(pattern, http.HandlerFunc(handler))
}

func (r *Router) Post(pattern string, handler func(w http.ResponseWriter, req *http.Request)) *mux.Route {
	return r.Router.Post(pattern, http.HandlerFunc(handler))
}

func abort(w http.ResponseWriter, code int, message string) {
	t := templates.MustGet("base.mnd")
	w.WriteHeader(code)
	w.Write([]byte(t.Render(map[string]interface{}{
		"content": `<h1>An error occured:</h1><p>` + message + `</p>`,
	})))
}

func render(template string, context map[string]interface{}) string {
	base := templates.MustGet("base.mnd")
	t := templates.MustGet(template)
	return t.RenderInLayout(base, context)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(render("index.mnd", map[string]interface{}{})))
}

func newTest(w http.ResponseWriter, req *http.Request) {
	_, name, err := CreateNewDatabase()
	if err != nil {
		abort(w, 500, "Something unexpected happened: "+err.Error())
		return
	}
	http.Redirect(w, req, "/"+name, 303)
}

func detail(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get(":id")
	dbm, err := OpenDatabase(id)
	if err != nil {
		abort(w, 500, "Could not open a database for that id: "+err.Error())
		return
	}
	status := Status{}
	err = dbm.Dbx.Get(&status, "SELECT * FROM _status;")
	if err != nil {
		abort(w, 500, "Could not open a database for that id: "+err.Error())
		return
	}
	w.Write([]byte(render("test.mnd", map[string]interface{}{"Id": id, "Question": Questions[status.CurrentQuestion-1]})))
}

type Sql struct {
	Id  string `schema:"id"`
	Sql string `schema:"sql"`
}

type JsonResponse struct {
	Ok           bool     `json:"ok"`
	Response     string   `json:"response"`
	NextQuestion Question `json:"next,omitempty"`
	Answered     bool     `json:"answered,omitempty"`
}

type JsonError struct {
	Error string `json:"error"`
}

func (j JsonResponse) Bytes() []byte {
	b, err := json.Marshal(j)
	if err != nil {
		return JsonError{"Could not unmarshal: " + err.Error()}.Bytes()
	}
	return b
}

func (e JsonError) Bytes() []byte {
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return b
}

func exec(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sql := &Sql{}
	req.ParseForm()
	decoder.Decode(sql, req.PostForm)

	dbm, err := OpenDatabase(sql.Id)
	if err != nil {
		w.Write(JsonError{"could not open database for that id: " + err.Error()}.Bytes())
		return
	}

	result, err := execSql(dbm, sql.Sql)

	if err != nil {
		w.Write(JsonError{err.Error()}.Bytes())
		return
	}

	response := JsonResponse{Ok: true, Response: result.String()}

	// if the current question was answered correctly, set this db to the next question.
	if result.Answered {
		// there's some weird maths below because sometimes CurrentQuestion is
		// being used as a number and sometimes it's being used as an index
		response.Answered = true
		if result.CurrentQuestion == len(Questions) {
			response.NextQuestion = Questions[result.CurrentQuestion-1]
		} else {
			response.NextQuestion = Questions[result.CurrentQuestion]
			_, err = dbm.Exec("UPDATE _status SET currentquestion=?;", result.CurrentQuestion+1)
			if err != nil {
				log.Println(err)
			}
		}
	}

	w.Write(response.Bytes())
}

func init() {
	r := &Router{}
	r.Get("/new", newTest)
	r.Get("/{id:[^/]+}", detail)
	r.Post("/{id:[^/]+}/exec", exec)
	r.Get("/", index)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	handler := handlers.LoggingHandler(os.Stdout, r)
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		now := time.Now().UTC()
		ts := now.Format(time.RFC1123)
		ts = strings.Replace(ts, "UTC", "GMT", 1)
		w.Header().Set("Server", "sql-test")
		w.Header().Set("Date", ts)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		handler.ServeHTTP(w, req)
	}))

	templates = mandira.NewLoader("./templates", PRELOAD_TEMPLATES)

}
