package app

import (
	"encoding/json"
	"net/http"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"gopkg.in/resty.v1"
)

func UpdateMapName() error {
	var sliceDevice []DeviceShort
	url := GetEndpoint(ClientMetadata) + "device"
	resp, err := resty.R().Get(url)
	if err == nil {
		json.Unmarshal(resp.Body(), &sliceDevice)

		for _, d := range sliceDevice {
			MapObjectName[d.NameID] = d.UserName["UserName"]["UserName"]
		}
		// fmt.Println("test map_name: %v", MapObjectName)
	}
	return err
}

func getFromMapName(nameID string) string {
	return MapObjectName[nameID]
}
func getHTTPStatus(err error) int {
	if err != nil {
		chk, ok := err.(*types.ErrServiceClient)
		if ok {
			return chk.StatusCode
		}
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func reponseHTTPrequest(w http.ResponseWriter, body []byte, err error) {
	status := getHTTPStatus(err)
	if status != http.StatusOK {
		http.Error(w, err.Error(), status)
	} else {
		if len(body) > 0 {
			w.Header().Set(CONTENTTYPE, APPLICATIONJSON)
		}
		w.Write(body)
	}
}
