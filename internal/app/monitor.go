package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/resty.v1"
)

// RestGetReadingByDeviceNameInTimeRange :get all readings by nameID in the time range
func RestGetReadingByDeviceNameInTimeRange(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	devicename := vars[DEVICENAME]
	start := vars[START]
	end := vars[END]
	limit := vars[LIMIT]

	startInt, _ := strconv.ParseInt(start, 10, 64)
	endInt, _ := strconv.ParseInt(end, 10, 64)
	limitInt, _ := strconv.ParseInt(limit, 10, 64)

	body, err := getReadingByDeviceNameInTimeRange(devicename, startInt, endInt, limitInt)
	reponseHTTPrequest(w, body, err)
}

// RestGetReadingByReadingNameInTimeRange :get all readings by nameID in the time range
func RestGetReadingByReadingNameInTimeRange(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	readingname := vars[READINGNAME]
	start := vars[START]
	end := vars[END]
	limit := vars[LIMIT]

	startInt, _ := strconv.ParseInt(start, 10, 64)
	endInt, _ := strconv.ParseInt(end, 10, 64)
	limitInt, _ := strconv.ParseInt(limit, 10, 64)

	body, err := getReadingByReadingNameInTimeRange(readingname, startInt, endInt, limitInt)
	reponseHTTPrequest(w, body, err)
}

// getReadingByDeviceNameInTimeRange :get all readings by nameID in the time range
func getReadingByDeviceNameInTimeRange(nameID string, from int64, to int64, limit int64) ([]byte, error) {
	const batchSize = 100
	const maxRequest = 100 // max = maxRequest * batchSize = 10 000
	if limit > (maxRequest * batchSize) {
		limit = maxRequest * batchSize
	}

	result := make([]Reading, batchSize*maxRequest)
	var pos int64
	toStr := strconv.FormatInt(to, 10)
	batchStr := strconv.FormatInt(batchSize, 10)

	var count int

	for ok := true; ok; ok = (count == batchSize) && (limit > 0) {
		fromStr := strconv.FormatInt(from, 10)
		resp, err := resty.R().Get(GetEndpoint(ClientData) + "reading/" + fromStr + "/" + toStr + "/" + batchStr)

		var readings []Reading
		if err = json.Unmarshal(resp.Body(), &readings); err != nil {
			return nil, err
		}

		count = len(readings)
		if count > 0 {
			from = readings[count-1].Created + 1
		}

		for _, reading := range readings {
			if reading.Device != nameID {
				continue
			}
			reading.UserName = getFromMapName(reading.Device)
			result[pos] = reading
			pos++
			if pos == limit {
				break
			}
		}
		limit = limit - pos
	}

	return json.Marshal(result[:pos])
}

// getReadingByReadingNameInTimeRange :get all readings of all Devices by ReadingName in the time range
func getReadingByReadingNameInTimeRange(readingname string, from int64, to int64, limit int64) ([]byte, error) {
	const batchSize = 100
	const maxRequest = 100 // max = maxRequest * batchSize = 10 000
	if limit > (maxRequest * batchSize) {
		limit = maxRequest * batchSize
	}

	result := make([]Reading, batchSize*maxRequest)
	var pos int64
	toStr := strconv.FormatInt(to, 10)
	batchStr := strconv.FormatInt(batchSize, 10)

	var count int

	for ok := true; ok; ok = (count == batchSize) && (limit > 0) {
		fromStr := strconv.FormatInt(from, 10)
		resp, err := resty.R().Get(GetEndpoint(ClientData) + "reading/" + fromStr + "/" + toStr + "/" + batchStr)

		var readings []Reading
		if err = json.Unmarshal(resp.Body(), &readings); err != nil {
			return nil, err
		}

		count = len(readings)
		if count > 0 {
			from = readings[count-1].Created + 1
		}

		for _, reading := range readings {
			if reading.Name != readingname {
				continue
			}
			reading.UserName = getFromMapName(reading.Device)
			result[pos] = reading
			pos++
			if pos == limit {
				break
			}
		}
		limit = limit - pos
	}

	return json.Marshal(result[:pos])
}

func RestGetValueDescriptorByName(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)

	name := vars[NAME]
	body, err := getValueDescriptorByName(name)
	reponseHTTPrequest(w, body, err)
}

func RestGetValueDescriptorByDeviceName(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)

	devicename := vars[DEVICENAME]
	body, err := getValueDescriptorByDeviceName(devicename)
	reponseHTTPrequest(w, body, err)
}

//GetValueDescriptorByName : get ValueDescriptor by name
func getValueDescriptorByName(name string) ([]byte, error) {
	var vd ValueDescriptor
	resp, err := resty.R().Get(GetEndpoint(ClientData) + "valuedescriptor/name/" + name)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(resp.Body(), &vd); err != nil {
		return nil, err
	}
	return json.Marshal(vd)
}

//GetValueDescriptorByDeviceName : get ValueDescriptors by devicename
func getValueDescriptorByDeviceName(name string) ([]byte, error) {
	var vd []ValueDescriptor
	resp, err := resty.R().Get(GetEndpoint(ClientData) + "valuedescriptor/devicename/" + name)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(resp.Body(), &vd); err != nil {
		return nil, err
	}
	return json.Marshal(vd)
}
