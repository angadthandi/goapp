package rest

import (
	"encoding/json"
	"net/http"

	"github.com/angadthandi/goapp/auth"
	log "github.com/angadthandi/goapp/log"
	"github.com/angadthandi/goapp/module/user"
	"github.com/angadthandi/goapp/util/config"
	"github.com/jinzhu/gorm"
)

// Register func creates new user if unique username/email.
// And responds with usertoken for user to login to system
func Register(
	w http.ResponseWriter,
	dbRef *gorm.DB,
	configData config.ConfigStruct,
	jsonMsg json.RawMessage,
) {
	var (
		resp GenericAPIResponse
		msg  user.SignupStruct
		err  error
	)

	err = json.Unmarshal(jsonMsg, &msg)
	if err != nil {
		respondWithError(
			http.StatusInternalServerError,
			"unable to unmarshal json",
			w,
		)
		log.Errorf("unable to unmarshal json: %v", err)
		return
	}

	// signup user
	user, err := user.Signup(dbRef, msg)
	if err != nil {
		respondWithError(
			http.StatusInternalServerError,
			"unable to signup",
			w,
		)
		log.Errorf("unable to signup: %v", err)
		return
	}

	// success register
	// create token & send it
	usertoken, err := auth.CreateToken(
		configData.JWTAuthSecret, int(user.UserID.Int64))
	if err != nil {
		respondWithError(
			http.StatusInternalServerError,
			"unable to create token",
			w,
		)
		log.Errorf("unable to create token: %v", err)
		return
	}

	// marshal token & send
	resp.Message, err = json.Marshal(usertoken)
	if err != nil {
		respondWithError(
			http.StatusInternalServerError,
			"unable to marshal json to sender",
			w,
		)
		log.Errorf("unable to marshal json: %v", err)
		return
	}

	// Response
	setDefaultHeader(w)

	resp.Api = "register"
	b, err := json.Marshal(resp)
	if err != nil {
		respondWithError(
			http.StatusInternalServerError,
			"unable to marshal json to sender",
			w,
		)
		log.Errorf("rest/API JSON Marshal error: %v", err)
		return
	}

	respondWithSuccess(b, w)
}
