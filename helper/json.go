package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromBody(request *http.Request, result interface{}) {

		decode := json.NewDecoder(request.Body)
		err :=	decode.Decode(result)
		PanicIfError(err)
}

func WriteFromResponse(writter http.ResponseWriter , result interface{}) {

	writter.Header().Add("Content-type","application/json")
	encode :=	json.NewEncoder(writter)
	err :=	encode.Encode(result)
	PanicIfError(err)
}