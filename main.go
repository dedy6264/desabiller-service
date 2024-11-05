package main

import (
	"context"
	"desabiller/apps"
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/repositories"
	"desabiller/routes"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo/middleware"
	id_translations "gopkg.in/go-playground/validator.v9/translations/id"
)

type CustomValidator struct {
	validator  *validator.Validate
	translator ut.Translator
}

// Passing Variable
var (
	uni         *ut.UniversalTranslator
	echoHandler echo.Echo
)

var ctx = context.Background()

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, row := range errs {
			return errors.New(row.Translate(cv.translator))
		}
	}

	return cv.validator.Struct(i)
}

func main() {

	if err := configs.OpenConnection(); err != nil {
		panic(fmt.Sprintf("Open Connection Faild: %s", err.Error()))
	}
	defer configs.CloseConnectionDB()

	// Connection database pgsql
	DB := configs.DBConnection()

	// Configuration Repository
	repo := repositories.NewRepositories(DB, ctx)

	// Configuration Repository and Services
	services := apps.SetupApp(DB, repo)
	// Routing API
	routes.RouteApi(echoHandler, services)
	echoHandler.Use(middleware.Recover())
	echoHandler.Use(middleware.Secure())
	echoHandler.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: configs.TRUE_VALUE,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	// echoHandler.Use(middleware.Logger())
	port := fmt.Sprintf(":%s", configs.GetEnv("APP_PORT", "8080"))
	echoHandler.Logger.Fatal(echoHandler.Start(port))
}

func init() {
	// BoardService()
	e := echo.New()
	echoHandler = *e
	validateCustom := validator.New()

	id := id.New()
	uni = ut.New(id, id)
	trans, _ := uni.GetTranslator("id")
	id_translations.RegisterDefaultTranslations(validateCustom, trans)
	e.Validator = &CustomValidator{validator: validateCustom, translator: trans}
	e.Static("/img/*", "assets/img")
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: configs.TRUE_VALUE,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		result := helpers.ResponseJSON(configs.FALSE_VALUE, strconv.Itoa(report.Code), err.Error(), nil)
		c.Logger().Error(report)
		c.JSON(report.Code, result)
	}

	// CEK HEADER DATA
	// var resultnya models.ReqHeaderLogin
	// var req string = "NSQMo/6kc8G/McFRSHdCDue70NTUwLjdvzsxG2orA7t5yUUw4q2AknY5/OEANMtEy5lxn4eng1UhJqG636ZnIz6rmQvAFRIb2jHPtCj6RYaFPRQyT8Fh5Yob3wOxunHXx8rD4HKjVqH41UHjdOJ98hra/CwmiVxT/mD0kTGtvBH9FMHgKr6MfJz56UMTPhpMosp4Y0RM4TMANSgAAtH4trRVF3FKe8kADeC7MSHzdgWUDwCgLBSk9tcnsbPgjoousdgG4esvhcosx5KNp6FJTrLvCLznr2OKNRUC4a4/h/1QFT7WOC9wWevIVRqtxovPfklSXgh/K86sTpusSsfI7bP0Zkorhk3byJo4p81EDoUDUknyl450vkGUMDlVUBGtTHWYJjRBOCjNsR+PLm3h11w6nnEBf0YBYLdosP7nwVwMJ72fWxtHMW9bVvLQFhz/lW1zYfX5LBR9hdz5n7kGWRayAMbSr4n3K6hnXIuRbbyanWeuAbNwgR6J3j3cuHBV"
	// resDec, err := helpers.AES256Decrypt(req, configs.AES256KEY)
	// if err != nil {
	// 	panic(err)
	// }

	// byte := []byte(resDec)
	// err = json.Unmarshal(byte, &resultnya)
	// if err != nil {
	// 	panic(err)
	// }
	// jssoonn, _ := json.Marshal(resultnya)
	// log.Panic(string(jssoonn))
}
