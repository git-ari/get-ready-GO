package password

import (
	"time"

	"github.com/git-ari/get-ready-GO/pkg/api/password"
	gorsk "github.com/git-ari/get-ready-GO/pkg/utl/model"
	"github.com/labstack/echo"
)

// New creates new password logging service
func New(svc password.Service, logger gorsk.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents password logging service
type LogService struct {
	password.Service
	logger gorsk.Logger
}

const name = "password"

// Change logging
func (ls *LogService) Change(c echo.Context, id int, oldPass, newPass string) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Change password request", err,
			map[string]interface{}{
				"req":  id,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Change(c, id, oldPass, newPass)
}
