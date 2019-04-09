package rest

import (
	"encoding/json"
	"net/http"

	"github.com/angadthandi/goapp/registry"
	"github.com/angadthandi/goapp/util/emergency"
	"github.com/jinzhu/gorm"

	log "github.com/angadthandi/goapp/log"
	"github.com/angadthandi/goapp/util/config"
)

type GenericAPIRecieve struct {
	Api     string          `json:"api"`
	Message json.RawMessage `json:"message"`
}

type GenericAPIResponse struct {
	Api     string      `json:"api"`
	Message interface{} `json:"message"`
}

type ErrorResponse struct {
	ErrorCode    int    `json:"code"`
	ErrorMessage string `json:"message"`
}

type SuccessResponse struct {
	SuccessMessage interface{} `json:"message"`
}

// handler for rest/API
func API(
	w http.ResponseWriter,
	dbRef *gorm.DB,
	reg *registry.Registry,
	configData config.ConfigStruct,
	jsonMsg json.RawMessage,
) {
	var (
		resp    GenericAPIResponse
		recieve GenericAPIRecieve
	)

	defer emergency.Handle()

	err := json.Unmarshal(jsonMsg, &recieve)
	if err != nil {
		log.Errorf("rest/API JSON unmarshal error: %v", err)
		return
	}

	// TODO validate token/input, might be done before this gets called

	var respErr error
	switch recieve.Api {
	// case "auth":
	// 	resp.Message, respErr = auth.Authenticate(
	// 		configData.JWTAuthSecret, recieve.Message)
	// case "register":
	// 	user, respErr := user.Register(dbRef, w, recieve.Message)
	// 	if respErr == nil {
	// 		// TODO can respond from register with new UserID
	// 		// and create a new func in auth pkg,
	// 		// which directly creates token from a UserID

	// 		// success register
	// 		// create token & send it
	// 		authData := auth.AuthRecieve{
	// 			Username: user.Username.String,
	// 			Password: user.Password.String,
	// 		}
	// 		b, err := json.Marshal(authData)
	// 		if err != nil {
	// 			respErr = err
	// 			break
	// 		}

	// 		resp.Message, respErr = auth.Authenticate(
	// 			configData.JWTAuthSecret, b)
	// 	}

	default:
		unknownApi(recieve.Api, w)
		log.Errorf("unknown api string: %v", recieve.Api)
		return
	}

	if respErr != nil {
		log.Errorf("rest/API response error: %v", respErr)
		return
	}

	// Response
	setDefaultHeader(w)

	resp.Api = recieve.Api
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

	// log.Debugf("rest/API JSON Response: %v", string(b))
	// fmt.Fprintf(w, "rest/API JSON Response: %v", string(b))
	respondWithSuccess(b, w)
}

func unknownApi(
	apiString string,
	w http.ResponseWriter,
) {
	respondWithError(
		http.StatusMethodNotAllowed,
		"unknown api call: "+apiString,
		w,
	)
}

func respondWithError(
	errorCode int,
	errorMessage string,
	w http.ResponseWriter,
) {
	errJson := GenericAPIResponse{
		Api: "error",
		Message: ErrorResponse{
			ErrorCode:    errorCode,
			ErrorMessage: errorMessage,
		},
	}
	err := json.NewEncoder(w).Encode(errJson)
	if err != nil {
		log.Errorf("Was unable to respond with error: %v because: %v",
			errorMessage, err)
	}
}

func setDefaultHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func respondWithSuccess(
	message interface{},
	w http.ResponseWriter,
) {
	respJson := GenericAPIResponse{
		Api: "success",
		Message: SuccessResponse{
			SuccessMessage: message,
		},
	}
	log.Debugf("API Response: %v", respJson)
	err := json.NewEncoder(w).Encode(respJson)
	if err != nil {
		log.Errorf("Was unable to respond with success: %v because: %v",
			message, err)
	}
}
