package releaze

import (
	"encoding/json"
	"net/http"
)

func HttpHandler(resp http.ResponseWriter, req *http.Request) {
	info := Get()
	bytes, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
	}
	resp.Write(bytes)
}
