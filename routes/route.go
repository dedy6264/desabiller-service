package routes

import (
	"crypto/subtle"
	"desabiller/configs"
	"desabiller/services"
	auth "desabiller/services/authService"
	helperservices "desabiller/services/helperServices"
	trxservice "desabiller/services/trxService"

	// helperservices "desabiller/services/helperServices"
	hierarchyservicego "desabiller/services/hierarchyService"
	providerservice "desabiller/services/providerServices"
	savingservices "desabiller/services/savingServices"
	"log"

	// nfeatureassignmentsservice "desabiller/services/nFeaturesService/featureAssignmentService"
	// nfeaturesservice "desabiller/services/nFeaturesService/featuresService"
	// clientservicego "desabiller/services/nHierarchyService/clientService"
	// merchantoutlettservicego "desabiller/services/nHierarchyService/merchantOutletService"
	// merchantservicego "desabiller/services/nHierarchyService/merchantService"
	// useroutlettservicego "desabiller/services/nHierarchyService/userOutletService"
	// nuserdashboardservice "desabiller/services/nUserDashboardService"

	// trxservice "desabiller/services/trxService"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RouteApi(e echo.Echo, service services.UsecaseService) {
	providerSvc := providerservice.NewApiProviderServices(service)
	admSvc := auth.ApiAdministration(service)
	hierachySvc := hierarchyservicego.RepoHierarchy(service)

	// userAct := useractivityservice.NewApiUserActivityService(service)
	helperSvc := helperservices.NewApiHelperService(service)
	trxSvc := trxservice.NewRepoTrxService(service)
	savingSvc := savingservices.NewApiSavingServices(service)
	login := e.Group("/login")
	login.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Println("[Start]")
		log.Println("EndPoint :", c.Path())
		log.Println("Header :", c.Request().Header)
		log.Println("Body :", string(reqBody))
		log.Println("Response :", string(resBody))
		log.Println("[End]")
	}))
	login.POST("/", admSvc.Login)
	{
		userapp := e.Group("/user-app")
		userapp.Use(middleware.BasicAuth(func(pss, pwd string, ctx echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(pss), []byte("joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(pwd), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		}))
		userapp.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		userapp.POST("/get", hierachySvc.GetUserApps)
		userapp.POST("/add", hierachySvc.AddUserApp)
		userapp.POST("/drop", hierachySvc.DropUserApp)
		userapp.POST("/update", hierachySvc.UpdateUserApp)
		userapp.POST("/verivicationotp", hierachySvc.VerificationOTP)
	}
	{
		aa := e.Group("/cif")
		aa.Use(middleware.BasicAuth(func(pss, pwd string, ctx echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(pss), []byte("joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(pwd), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		}))
		aa.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		aa.POST("/add", savingSvc.AddCif)
		aa.POST("/update", savingSvc.UpdateCif)
		aa.POST("/drop", savingSvc.DropCif)
		aa.POST("/gets", savingSvc.GetCifs)

		bb := e.Group("/saving-type")
		bb.Use(middleware.BasicAuth(func(pss, pwd string, ctx echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(pss), []byte("joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(pwd), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		}))
		bb.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		bb.POST("/add", savingSvc.AddSavingType)
		bb.POST("/update", savingSvc.UpdateSavingType)
		bb.POST("/drop", savingSvc.DropSavingType)
		bb.POST("/gets", savingSvc.GetSavingTypes)

		cc := e.Group("/saving-segment")
		cc.Use(middleware.BasicAuth(func(pss, pwd string, ctx echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(pss), []byte("joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(pwd), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		}))
		cc.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		cc.POST("/add", savingSvc.AddSavingSegment)
		cc.POST("/update", savingSvc.UpdateSavingSegment)
		cc.POST("/drop", savingSvc.DropSavingSegment)
		cc.POST("/gets", savingSvc.GetSavingSegments)

		ee := e.Group("/saving-transaction")
		ee.Use(middleware.BasicAuth(func(pss, pwd string, ctx echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(pss), []byte("joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(pwd), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		}))
		ee.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		ee.POST("/gets", savingSvc.GetSavingTransactions)
		ee.POST("/update", savingSvc.UpdateSavingTransaction)

		dd := e.Group("/account")
		dd.Use(middleware.BasicAuth(func(pss, pwd string, ctx echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(pss), []byte("joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(pwd), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		}))
		dd.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		dd.POST("/add", savingSvc.AddAccount)
		dd.POST("/gets", savingSvc.GetAccounts)
		dd.POST("/drop", savingSvc.DropAccount)
		dd.POST("/update", savingSvc.UpdateAccount)
		dd.POST("/setpin", savingSvc.SetPin)
	}
	{
		proType := e.Group("/product-type")
		proType.Use(middleware.BasicAuth(func(pss, pwd string, ctx echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(pss), []byte("joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(pwd), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		}))
		proType.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		proType.POST("/get", providerSvc.GetProductTypes)
		proType.POST("/add", providerSvc.AddProductType)
		proType.POST("/drop", providerSvc.DropProductType)
		proType.POST("/update", providerSvc.UpdateProductType)

		proCtgr := e.Group("/product-category")
		proCtgr.Use(middleware.BasicAuth(func(pss, pwd string, ctx echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(pss), []byte("joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(pwd), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		}))
		proCtgr.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		proCtgr.POST("/get", providerSvc.GetProductCategories)
		proCtgr.POST("/add", providerSvc.AddProductCategory)
		proCtgr.POST("/drop", providerSvc.DropProductCategory)
		proCtgr.POST("/update", providerSvc.UpdateProductCategory)

		proRefnce := e.Group("/product-reference")
		proRefnce.Use(middleware.BasicAuth(func(pss, pwd string, ctx echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(pss), []byte("joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(pwd), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		}))
		proRefnce.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		proRefnce.POST("/get", providerSvc.GetProductReferences)
		proRefnce.POST("/add", providerSvc.AddProductReference)
		proRefnce.POST("/drop", providerSvc.DropProductReference)
		proRefnce.POST("/update", providerSvc.UpdateProductReference)

		product := e.Group("/product")
		// product.Use(middleware.BasicAuth(func(pss, pwd string, ctx echo.Context) (bool, error) {
		// 	if subtle.ConstantTimeCompare([]byte(pss), []byte("joe")) == 1 &&
		// 		subtle.ConstantTimeCompare([]byte(pwd), []byte("secret")) == 1 {
		// 		return true, nil
		// 	}
		// 	return false, nil
		// }))
		product.Use(middleware.JWT([]byte(configs.KEY)))
		product.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		product.POST("/get", providerSvc.GetProducts)
		product.POST("/add", providerSvc.AddProduct)
		product.POST("/drop", providerSvc.DropProduct)
		product.POST("/update", providerSvc.UpdateProduct)

	}

	// {
	// 	//eksternal callback
	// 	callback := e.Group("/callback")
	// 	callback.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	// 		log.Println("[Start]")
	// 		log.Println("EndPoint :", c.Path())
	// 		log.Println("Header :", c.Request().Header)
	// 		log.Println("Body :", string(reqBody))
	// 		log.Println("Response :", string(resBody))
	// 		log.Println("[End]")
	// 	}))
	// 	// callback.POST("/iak", trxSvc.IAKCallback)
	// }

	{ //user
		user := e.Group("/user")
		user.Use(middleware.JWT([]byte(configs.KEY)))
		user.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		user.POST("/", admSvc.CekJwt)
		user.POST("/getuser", hierachySvc.GetUser)
	}
	{
		//BILLER transaksi
		mm := e.Group("/biller")
		mm.Use(middleware.JWT([]byte(configs.KEY)))
		mm.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		mm.POST("/inquiry", trxSvc.InquiryBiller)
		mm.POST("/payment", trxSvc.PaymentBiller)
		mm.POST("/advice", trxSvc.Advice)
	}
	{
		//transaksi
		mobileTrx := e.Group("/trx")
		mobileTrx.Use(middleware.JWT([]byte(configs.KEY)))
		mobileTrx.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		mobileTrx.POST("/getHistory", trxSvc.HistoryTrxBillerReports)
		mobileTrx.POST("/getTrx", trxSvc.TrxBillerReport)
		mobileTrx.POST("/getTrxs", trxSvc.TrxBillerReports)
	}
	{
		//helper
		helper := e.Group("/helper")
		helper.Use(middleware.JWT([]byte(configs.KEY)))
		helper.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		helper.POST("/getReference", helperSvc.GetOperatorService)
	}
}
