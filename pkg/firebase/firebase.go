package firebase

import "google.golang.org/api/option"

type App struct {
	authOverride     map[string]interface{}
	dbURL            string
	projectID        string
	serviceAccountID string
	storageBucket    string
	opts             []option.ClientOption
}
