//go:build prod

package environment

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var _ EnvContext = (*context)(nil)

type context struct {
}

func (p *context) Init(app *fiber.App, db *gorm.DB) {
}

func (p *context) GetWorkingDirectory() (string, error) {
	return os.Executable()
}

func init() {
	Mode = Production
	Context = &context{}
}
