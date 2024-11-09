package administrationservice

import (
	"desabiller/helpers"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func (svc AdministrationService) CekJwt(ctx echo.Context) error {
	data := helpers.TokenJWTDecode(ctx)
	byte, _ := json.Marshal(data)
	fmt.Println(string(byte))
	return ctx.JSON(http.StatusOK, nil)
}
