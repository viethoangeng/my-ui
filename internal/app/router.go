/*******************************************************************************
 * Copyright 2017 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/
package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//LoadRestRoutes : init router
func LoadRestRoutes() http.Handler {
	r := mux.NewRouter()

	// Ping Resource
	r.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	// Configuration
	r.HandleFunc("/config", configHandler).Methods(http.MethodGet)

	// Get Readings by name in the time range
	r.HandleFunc(AppApiMonitorByDeviceName, RestGetReadingByDeviceNameInTimeRange).Methods(http.MethodGet)
	r.HandleFunc(AppApiMonitorByReadingName, RestGetReadingByReadingNameInTimeRange).Methods(http.MethodGet)
	r.HandleFunc(AppApiMonitorValueDescriptorByName, RestGetValueDescriptorByName).Methods(http.MethodGet)
	r.HandleFunc(AppApiMonitorValueDescriptorByDeviceName, RestGetValueDescriptorByDeviceName).Methods(http.MethodGet)
	return r
}

// Respond with PINGRESPONSE to see if the service is alive
func pingHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set(CONTENTTYPE, TEXTPLAIN)
	w.Write([]byte(PINGRESPONSE))
}

func configHandler(w http.ResponseWriter, _ *http.Request) {
	encode(Configuration, w)
}

// Helper function for encoding things for returning from REST calls
func encode(i interface{}, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	err := enc.Encode(i)
	// Problems encoding
	if err != nil {
		// LoggingClient.Error("Error encoding the data: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
