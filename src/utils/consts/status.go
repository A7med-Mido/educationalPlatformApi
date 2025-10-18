package consts

import "net/http"


var(
	OK = http.StatusOK
	created = http.StatusCreated
	noContent = http.StatusNoContent
	badRequest = http.StatusBadRequest
	unauthorized = http.StatusUnauthorized
	forbidden = http.StatusForbidden
	notFound = http.StatusNotFound
	conflict = http.StatusConflict
	unprocessableEntity = http.StatusUnprocessableEntity
	internalServerError = http.StatusInternalServerError
)