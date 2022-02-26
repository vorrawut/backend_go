package cors

import "os"

var Default = map[string]string{
	"Access-Control-Allow-Origin":  os.Getenv("CORS_DOMAINS"),
	"Access-Control-Allow-Methods": "GET,POST,PUT,PATCH,DELETE,OPTIONS",
	"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
}
