package providerservices

import (
	"desabiller/utils"
	"fmt"
	"log"

	"github.com/labstack/echo"
)

func (svc providerServices) GetToken(ctx echo.Context) error {
	respByte, _, err := utils.WorkerPostWithBearerGetToken("http://localhost:8080/api/getToken", "", nil, "json")
	if err != nil {
		log.Println("FAILED GET TOKEN")
	}
	fmt.Println(string(respByte))
	return nil
}
