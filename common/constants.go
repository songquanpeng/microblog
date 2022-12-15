package common

import (
	"embed"
	"github.com/google/uuid"
)

//go:embed public
var FS embed.FS

var Version = "v0.0.0"
var SQLitePath = "microblog.db"

var Username string
var Password string

var SessionSecret = uuid.New().String()
