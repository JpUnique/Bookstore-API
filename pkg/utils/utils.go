package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) (body []byte, err error) {
	body, err = io.ReadAll(r.Body)
	if err != nil {
		err = errors.New("failed to parse body")
		return body, err
	}
	err = json.Unmarshal(body, x)
	if err != nil {
		err = errors.New("failed to unmarshal the  body")
		return body, err
	}
	return body, err
}
