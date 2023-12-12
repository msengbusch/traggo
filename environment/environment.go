package environment

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type EnvMode string

const (
	Development EnvMode = "development"
	Production          = "production"
)

var Mode EnvMode       // Set via development/production.go init func
var Context EnvContext // Set via development/production.go init func

var BuildCommit string = "_"  // Set via ldflags
var BuildDate string = "_"    // Set via ldflags
var BuildVersion string = "_" // Set via ldflags

type EnvContext interface {
	Init(app *fiber.App, db *gorm.DB)

	GetWorkingDirectory() (string, error)
}
