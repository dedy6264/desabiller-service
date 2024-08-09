package administrationservice

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
)

func (svc AdministrationService) CekJwt(ctx echo.Context) error {
	a := ctx.Get("user").(*jwt.Token)
	// claim := a.Claims.(jwt.MapClaims)
	fmt.Println("::::", a)
	return ctx.JSON(http.StatusOK, nil)
}
