package mdw

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/possawang/backoffice-persist-service/types"
	"github.com/possawang/go-service-lib-common/errorutils"
)

func GetUserData(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := parsingJson[types.GetUserDataReq](w, r)
		if 0 < req.Id {
			errorutils.Handle4xx(w, "ID must be more than 0", 3, "400")
		}
		h.ServeHTTP(w, r)
	})
}

func parsingJson[V any](w http.ResponseWriter, r *http.Request) V {
	var req V
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorutils.Handle4xx(w, err.Error(), 1, "400")
	}
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		errorutils.Handle4xx(w, err.Error(), 2, "400")
	}
	return req
}
