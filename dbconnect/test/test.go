package test

import (
	"fmt"
	"net/http"

	"github.com/angadthandi/goapp/log"
	"github.com/jinzhu/gorm"
)

// // Test Routes ----------------------------------------

// test handler
func TestHandler(
	w http.ResponseWriter,
	r *http.Request,
	dbRef *gorm.DB,
) {
	data := "Test Data!"
	// data := Echo(dbRef)
	log.Debugf("Test Page! %s", data)
	fmt.Fprintf(w, "Test Page! %s", data)
}
