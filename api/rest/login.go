package rest

import (
	"encoding/json"
	"net/http"

	"github.com/angadthandi/goapp/auth"
	log "github.com/angadthandi/goapp/log"
	"github.com/angadthandi/goapp/model/tbluser"
	"github.com/angadthandi/goapp/module/user"
	"github.com/angadthandi/goapp/util/config"
	"github.com/jinzhu/gorm"
)

// Login func validates user creds,
// if valid responds with usertoken for user to login to system
func Login(
	w http.ResponseWriter,
	dbRef *gorm.DB,
	configData config.ConfigStruct,
	jsonMsg json.RawMessage,
) {
	var (
		resp GenericAPIResponse
		msg  user.LoginStruct
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

	// validate creds from db
	user, err := tbluser.ValidateLogin(
		dbRef,
		msg.Username.String,
		msg.Password.String,
	)
	if err != nil {
		respondWithError(
			http.StatusInternalServerError,
			"invalid login",
			w,
		)
		return
	}

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

	resp.Api = "login"
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
