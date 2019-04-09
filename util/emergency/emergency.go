package emergency

import (
	"runtime/debug"

	"github.com/angadthandi/goapp/log"
)

func Handle() {
	// recover from panic if one occured
	if err := recover(); err != nil {
		log.Errorf("Recovered from panic: %v", err)
		stackTrace := string(debug.Stack())
		// even though the stacktrace is verbose. We will want to log it at all
		// log levels considering how disruptive a panic is
		log.Errorf("Trace: %v", stackTrace)
	}
}
