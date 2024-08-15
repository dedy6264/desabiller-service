package hierarchyservice

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
)

func (svc HierarcyService) CheckUserStatus(ctx echo.Context) error {
	var (
	// svcName    = "CheckUserStatus"
	// respGlobal models.Response
	// dbTime     = time.Now().Format(time.RFC3339)
	)
	a := ctx.Get("user").(*jwt.Token)
	claim := a.Claims.(jwt.MapClaims)
	exp := claim["exp"].(float64)
	fmt.Println("::::", exp)
	return nil
}
