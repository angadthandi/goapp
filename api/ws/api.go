package ws

import (
	"encoding/json"

	"github.com/angadthandi/goapp/log"
	"github.com/angadthandi/goapp/registry"
	"github.com/angadthandi/goapp/util/config"
	"github.com/angadthandi/goapp/util/emergency"
	"github.com/jinzhu/gorm"
)

type GenericWSRecieve struct {
	Api     string          `json:"api"`
	Message json.RawMessage `json:"message"`
}

type GenericWSResponse struct {
	Api     string      `json:"api"`
	Message interface{} `json:"message"`
}

// handler for ws/API
func API(
	dbRef *gorm.DB,
	reg *registry.Registry,
	configData config.ConfigStruct,
	caller registry.ClientID,
	jsonMsg json.RawMessage,
) {
	var (
		resp                GenericWSResponse
		recieve             GenericWSRecieve
		sendMsgToAllClients bool
	)

	defer emergency.Handle()

	err := json.Unmarshal(jsonMsg, &recieve)
	if err != nil {
		log.Errorf("ws/API JSON unmarshal error: %v", err)
		return
	}

	switch recieve.Api {
	case "test":
		resp.Message = recieve.Message
		sendMsgToAllClients = true

	default:
		resp.Message = "Default JSON Message!"
	}

	// Response
	resp.Api = recieve.Api
	b, err := json.Marshal(resp)
	if err != nil {
		log.Errorf("ws/API JSON Marshal error: %v", err)
		return
	}

	log.Debugf("ws/API JSON Response: %v", string(b))
	if sendMsgToAllClients {
		reg.SendToAllClients(b)
	} else {
		reg.SendToCaller(caller, b)
	}
}
