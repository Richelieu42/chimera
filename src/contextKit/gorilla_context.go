package contextKit

import (
	"github.com/gorilla/context"
	"net/http"
)

var (
	Set func(r *http.Request, key, val interface{}) = context.Set

	Get      func(r *http.Request, key interface{}) interface{}         = context.Get
	GetOk    func(r *http.Request, key interface{}) (interface{}, bool) = context.GetOk
	GetAll   func(r *http.Request) map[interface{}]interface{}          = context.GetAll
	GetAllOk func(r *http.Request) (map[interface{}]interface{}, bool)  = context.GetAllOk

	Delete func(r *http.Request, key interface{}) = context.Delete
	Clear  func(r *http.Request)                  = context.Clear

	Purge func(maxAge int) int = context.Purge

	// ClearHandler !!!: 能有效避免内存泄露.
	ClearHandler func(h http.Handler) http.Handler = context.ClearHandler
)
