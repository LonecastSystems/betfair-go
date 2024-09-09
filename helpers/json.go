package helpers

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadJson[T any](res *http.Response, response *T) (err error) {
	defer res.Body.Close()

	if body, err := io.ReadAll(res.Body); err == nil {
		return json.Unmarshal(body, &response)
	}

	return nil
}
