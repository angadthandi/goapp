package route

import (
	"io/ioutil"
	"net/http"

	"github.com/angadthandi/goapp/dbconnect/test"
	"github.com/angadthandi/goapp/gosocket"
	"github.com/angadthandi/goapp/registry"
	"github.com/jinzhu/gorm"

	"github.com/angadthandi/goapp/api/rest"
	"github.com/gorilla/mux"

	log "github.com/angadthandi/goapp/log"
	"github.com/angadthandi/goapp/util/config"
)

func Handle(
	dbRef *gorm.DB,
	hub *gosocket.Hub,
	reg *registry.Registry,
	configData config.ConfigStruct,
) {
	r := mux.NewRouter().StrictSlash(true)

	// register route, with no auth check for token
	r.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		b := postDataHandler(r)

		rest.Register(w, dbRef, configData, b)
	})

	// login route, with no auth check for token
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		b := postDataHandler(r)

		rest.Register(w, dbRef, configData, b)
	})

	// r.HandleFunc("/", rest.Home)

	// add token check for all api calls
	r.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {

		b := postDataHandler(r)

		// TODO validate token before accessing API

		rest.API(w, dbRef, reg, configData, b)
	})

	// generate websocket conn based on a valid token
	r.HandleFunc("/ws/{token}", func(w http.ResponseWriter, r *http.Request) {

		varsMap := mux.Vars(r)
		token, ok := varsMap["token"]
		if !ok {
			log.Error("Invalid token!")
			return
		}

		// TODO validate token before upgrading to ws

		gosocket.ServeWs(
			hub,
			w,
			r,
			dbRef,
			reg,
			configData,
			token,
		)
	})

	// Test Routes --------------------------------
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		test.TestHandler(w, r, dbRef)
	})
	// Test Routes --------------------------------

	// // static files
	// r.HandleFunc("/vendor/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.StripPrefix("/vendor/",
	// 		http.FileServer(http.Dir("./public")))
	// })
	//   r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.Handle("/", r)
}

func postDataHandler(
	r *http.Request,
) []byte {
	// w.Header().Set("Content-Type", "application/json")

	// var m Member
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Errorf("error reading post data: %v", err)
		return nil
	}

	// TODO parse body & extract token
	// TODO verify token

	//   json.Unmarshal(b, &m)

	//   members = append(members, m)

	//   j, _ := json.Marshal(m)
	//   w.Write(j)

	return b
}
