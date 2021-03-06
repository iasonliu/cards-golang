package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	username     string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{} // user ID, user
var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time

const sessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/login", login)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	showSessions()
	tpl.ExecuteTemplate(w, "index.tpl", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !isLogged(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "bar.tpl", u)
}
func list(w http.ResponseWriter, req *http.Request) {
	if isLogged(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "listuser.tpl", dbUsers)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if isLogged(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		username := req.FormValue("username")
		password, err := bcrypt.GenerateFromPassword([]byte(req.FormValue("password")), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "error", http.StatusServiceUnavailable)
			return
		}
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		role := req.FormValue("role")
		// username taken?
		if _, ok := dbUsers[username]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.Must(uuid.NewRandom())
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		})
		dbSessions[sID.String()] = session{username, time.Now()}

		// store user in dbUsers
		u := user{username, password, firstname, lastname, role}
		dbUsers[username] = u
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "signup.tpl", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if isLogged(w, req) {
		// clearup cookie
		c := &http.Cookie{
			Name:   "session",
			Value:  "",
			MaxAge: -1,
		}
		http.SetCookie(w, c)
		// delete session id from DB
		if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
			go cleanSessions()
		}
	}
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

func login(w http.ResponseWriter, req *http.Request) {
	if isLogged(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.Must(uuid.NewRandom())
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "login.tpl", nil)
}
