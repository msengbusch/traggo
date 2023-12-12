//go:build !prod

package environment

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/traggo/server/server"
	"gorm.io/gorm"
)

var _ EnvContext = (*context)(nil)

type context struct {
}

func (d *context) Init(app *fiber.App, db *gorm.DB) {
	server.InitSwagger(app)
}

func (d *context) GetWorkingDirectory() (string, error) {
	return filepath.Dir("."), nil
}

func init() {
	Mode = Development
	Context = &context{}
}
