package oxyKit

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"net/http"
)

// StatusClientClosedRequest non-standard HTTP status code for client disconnection.
const StatusClientClosedRequest = 499

// StatusClientClosedRequestText non-standard HTTP status for client disconnection.
const StatusClientClosedRequestText = "Client Closed Request"

// errorHandler Standard error handler.
type errorHandler struct {
	logger *logrus.Logger
}

func (handler *errorHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request, err error) {
	statusCode := http.StatusInternalServerError

	//if netErr, ok := err.(net.Error); ok {
	//	if netErr.Timeout() {
	//		statusCode = http.StatusGatewayTimeout
	//	} else {
	//		statusCode = http.StatusBadGateway
	//	}
	//} else if errors.Is(err, io.EOF) {
	//	statusCode = http.StatusBadGateway
	//} else if errors.Is(err, context.Canceled) {
	//	statusCode = StatusClientClosedRequest
	//}
	var netErr net.Error
	if errors.As(err, &netErr) {
		if netErr.Timeout() {
			statusCode = http.StatusGatewayTimeout
		} else {
			statusCode = http.StatusBadGateway
		}
	} else if errors.Is(err, io.EOF) {
		statusCode = http.StatusBadGateway
	} else if errors.Is(err, context.Canceled) {
		statusCode = StatusClientClosedRequest
	}

	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(statusText(statusCode)))

	//netErr.log.Debug("'%d %s' caused by: %v", statusCode, statusText(statusCode), err)
	handler.logger.Errorf("'%d %s' caused by: %v", statusCode, statusText(statusCode), err)
}

func statusText(statusCode int) string {
	if statusCode == StatusClientClosedRequest {
		return StatusClientClosedRequestText
	}
	return http.StatusText(statusCode)
}
