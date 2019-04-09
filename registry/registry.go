package registry

import (
	"encoding/json"
	"sync"
)

type ClientID struct {
	TokenID  string
	UniqueID string
	UserID   int
}

type Registry struct {
	clientMapLock sync.RWMutex

	clientMap map[ClientID](chan []byte)
}

func NewRegistry() *Registry {
	return &Registry{
		clientMap: make(map[ClientID](chan []byte)),
	}
}

func (r *Registry) RegisterClient(
	clientID ClientID,
	ch chan []byte,
) {
	r.clientMapLock.Lock()
	defer r.clientMapLock.Unlock()

	r.clientMap[clientID] = ch
}

func (r *Registry) UnRegisterClient(
	clientID ClientID,
) {
	r.clientMapLock.Lock()
	defer r.clientMapLock.Unlock()

	delete(r.clientMap, clientID)
}

func (r *Registry) SendToAllClients(
	jsonMsg json.RawMessage,
) {
	r.clientMapLock.RLock()
	defer r.clientMapLock.RUnlock()

	for _, ch := range r.clientMap {
		ch <- jsonMsg
	}
}

func (r *Registry) SendToCaller(
	clientID ClientID,
	jsonMsg json.RawMessage,
) {
	r.clientMapLock.RLock()
	defer r.clientMapLock.RUnlock()

	for cID, ch := range r.clientMap {
		if clientID == cID {
			ch <- jsonMsg
			break
		}
	}
}
