package query

import (
	gorsk "github.com/git-ari/get-ready-GO/pkg/utl/model"
	"github.com/labstack/echo"
)

// List prepares data for list queries
func List(u *gorsk.AuthUser) (*gorsk.ListQuery, error) {
	switch true {
	case u.Role <= gorsk.AdminRole: // user is SuperAdmin or Admin
		return nil, nil
	case u.Role == gorsk.CompanyAdminRole:
		return &gorsk.ListQuery{Query: "company_id = ?", ID: u.CompanyID}, nil
	case u.Role == gorsk.LocationAdminRole:
		return &gorsk.ListQuery{Query: "location_id = ?", ID: u.LocationID}, nil
	default:
		return nil, echo.ErrForbidden
	}
}
