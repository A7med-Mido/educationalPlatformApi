package constants


import "net/http"


const (
	OK              = http.StatusOK
	Created         = http.StatusCreated
	InternalError   = http.StatusInternalServerError
	Unauthorized    = http.StatusUnauthorized
	Forbidden       = http.StatusForbidden
	NotFound        = http.StatusNotFound
)