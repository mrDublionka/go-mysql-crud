package utils

import (
	"cloud.google.com/go/storage"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var StorageClient *storage.Client

type FirebaseConfig struct {
	ApiKey            string
	AuthDomain        string
	ProjectId         string
	StorageBucket     string
	MessagingSenderId string
	AppId             string
	MeasurementId     string
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
