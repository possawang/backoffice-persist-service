package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/possawang/backoffice-persist-service/entity"
	"github.com/possawang/backoffice-persist-service/types"
	"github.com/possawang/go-persist-lib-common/connection"
	"github.com/possawang/go-service-lib-common/commonutils"
	"github.com/possawang/go-service-lib-common/domain"
	"github.com/possawang/go-service-lib-common/errorutils"
)

func GetUserData(w http.ResponseWriter, r *http.Request) {
	req := parsingJson[types.GetUserDataReq](w, r)
	var user entity.User
	trx := connection.DB.First(&user, req.Id)
	if trx.Error != nil {
		errorutils.Handle5xx(w, trx.Error, 3, "500")
	}
	if 0 >= trx.RowsAffected {
		errorutils.Handle5xx(w, fmt.Errorf("user with id %d not found", req.Id), 4, "500")
	}
	var alloweds []entity.Allowed
	whereAllowed := entity.Allowed{Role: user.Role}
	whereAllowed.Deleted = false
	connection.DB.Where(&whereAllowed).Find(&alloweds)
	res := types.GetUserDataRes{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role.Name,
		Allowed: commonutils.ArrayMap(alloweds, func(v entity.Allowed) types.UserAllowedAccess {
			return types.UserAllowedAccess{
				Endpoint: v.Endpoint,
				Method:   v.Method,
			}
		}),
	}
	givingResponse(w, res)
}

func givingResponse[V any](w http.ResponseWriter, data V) {
	res := domain.BaseResponse[V]{
		Data:       data,
		Msg:        "Sukses",
		StatusCode: "00",
	}
	bytes, err := json.Marshal(res)
	if err != nil {
		errorutils.Handle5xx(w, err, 5, "500")
	}
	w.Header().Set("status", "200")
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func parsingJson[V any](w http.ResponseWriter, r *http.Request) V {
	var req V
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorutils.Handle5xx(w, err, 1, "500")
	}
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		errorutils.Handle5xx(w, err, 2, "500")
	}
	return req
}
