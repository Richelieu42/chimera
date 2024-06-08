package forwardKit

import (
	"github.com/richelieu-yang/chimera/v3/src/contextKit"
	"log"
	"net/http"
)

func DefaultErrorHandler(errLogger *log.Logger) func(w http.ResponseWriter, r *http.Request, err error) {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		if err != nil {
			errLogger.Printf("Fail to forward request, error: %s.", err.Error())
			contextKit.Set(r, "", err)
		}
	}
}
