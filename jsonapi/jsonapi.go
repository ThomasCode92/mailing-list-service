package jsonapi

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func setJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func fromJson[T any](body io.Reader, target T) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	json.Unmarshal(buf.Bytes(), &target)
}

func returnJson[T any](w http.ResponseWriter, withData func() (T, error)) {
	setJsonHeader(w)

	data, serverErr := withData()

	if serverErr != nil {
		w.WriteHeader(500)
		serverErrJson, err := json.Marshal(&serverErr)
		if err != nil {
			log.Print(err)
			return
		}
		w.Write(serverErrJson)
		return
	}

	dataJson, err := json.Marshal(&data)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	w.Write(dataJson)
}

func returnErr(w http.ResponseWriter, err error, code int) {
	// We use returnJson so it can modify `w` and generate a JSON response.
	// `returnJson` will call `handleErr` which also modifies `w`.
	// `returnJson` will check the return value of `handleErr` to confirm if it should return an error or valid data.
	// Nothing is actually returned from `returnErr` because the return value is written into `w`.
	handleErr := func() (interface{}, error) {
		var errorMessage struct{ Err string }
		errorMessage.Err = err.Error()
		w.WriteHeader(code)
		return errorMessage, nil
	}
	returnJson(w, handleErr)
}
