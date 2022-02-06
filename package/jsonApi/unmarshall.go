package jsonApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func handleUnmarshalJSONError(err error) error {
	thrownError := fmt.Sprintf("Invald JSON in request body: %s", err)
	return errors.New(thrownError)
}

func unmarshalJSON(val interface{}, r io.Reader) error {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	if err := dec.Decode(val); err != nil {
		return handleUnmarshalJSONError(err)
	}
	return nil
}

func UnmarshalJSONRequest(val interface{}, r *http.Request) error {
	if sk, ok := r.Body.(io.Seeker); ok {
		_, err := sk.Seek(0, io.SeekStart)
		if err != nil {
			panic(err)
		}
	}
	return unmarshalJSON(val, r.Body)
}
