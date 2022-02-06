package jsonApi

import (
	"context"
	"encoding/json"
	"net/http"
)

type HttpCommand func(ctx context.Context, r *http.Request) (interface{}, error)

func JSON(fn HttpCommand) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var marshalled []byte
		var err error
		var res interface{}

		writer.Header().Add("Content-Type", "application/json")
		res, err = fn(request.Context(), request)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			marshalled, err = json.Marshal(err.Error())
			if err != nil {
				_, _ = writer.Write([]byte(err.Error()))
				return
			}
			_, _ = writer.Write(marshalled)
			return
		}
		writer.WriteHeader(http.StatusOK)
		marshalled, err = json.Marshal(res)
		if err != nil {
			_, _ = writer.Write([]byte(err.Error()))
			return
		}
		_, _ = writer.Write(marshalled)

	}
}
