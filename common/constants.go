package common

import (
	"embed"
	"github.com/google/uuid"
)

var FS embed.FS

var Version = "v0.0.0"
var SQLitePath = "microblog.db"
var Theme = ""

var Username string
var Password string

var SessionSecret = uuid.New().String()
