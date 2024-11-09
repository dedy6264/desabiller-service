package routes

import (
	"crypto/subtle"
	"desabiller/configs"
	"desabiller/services"
	auth "desabiller/services/authService"
	hierarchyservicego "desabiller/services/hierarchyService"
	trxservice "desabiller/services/trxService"
	useractivityservice "desabiller/services/userActivityService"
	"log"

	// nfeatureassignmentsservice "desabiller/services/nFeaturesService/featureAssignmentService"
	// nfeaturesservice "desabiller/services/nFeaturesService/featuresService"
	// clientservicego "desabiller/services/nHierarchyService/clientService"
	// merchantoutlettservicego "desabiller/services/nHierarchyService/merchantOutletService"
	// merchantservicego "desabiller/services/nHierarchyService/merchantService"
	// useroutlettservicego "desabiller/services/nHierarchyService/userOutletService"
	// nuserdashboardservice "desabiller/services/nUserDashboardService"
	providerservice "desabiller/services/providerServices"

	// trxservice "desabiller/services/trxService"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RouteApi(e echo.Echo, service services.UsecaseService) {
	providerSvc := providerservice.NewApiProviderServices(service)
	admSvc := auth.ApiAdministration(service)
	hierachySvc := hierarchyservicego.ApiHierarchy(service)
	userAct := useractivityservice.NewApiUserActivityService(service)
	trxSvc := trxservice.NewApiTrxService(service)

	{ //dashboard

		aa := e.Group("/client")
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
		aa.POST("/add", hierachySvc.AddClient)
		aa.POST("/gets", hierachySvc.GetClients)
		aa.POST("/drop", hierachySvc.DropClient)
		aa.POST("/update", hierachySvc.UpdateClient)
		bb := e.Group("/group")
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
		bb.POST("/add", hierachySvc.AddGroup)
		bb.POST("/gets", hierachySvc.GetGroups)
		bb.POST("/drop", hierachySvc.DropGroup)
		bb.POST("/update", hierachySvc.UpdateGroup)
		cc := e.Group("/merchant")
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
		cc.POST("/add", hierachySvc.AddMerchant)
		cc.POST("/gets", hierachySvc.GetMerchants)
		cc.POST("/drop", hierachySvc.DropMerchant)
		cc.POST("/update", hierachySvc.UpdateMerchant)
		dd := e.Group("/merchantOutlet")
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
		dd.POST("/add", hierachySvc.AddMerchantOutlet)
		dd.POST("/gets", hierachySvc.GetMerchantOutlets)
		dd.POST("/drop", hierachySvc.DropMerchantOutlet)
		dd.POST("/update", hierachySvc.UpdateMerchantOutlet)
		ee := e.Group("/provider")
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
		ee.POST("/add", providerSvc.AddProvider)
		ee.POST("/gets", providerSvc.GetProviders)
		ee.POST("/drop", providerSvc.DropProvider)
		ee.POST("/update", providerSvc.UpdateProvider)
		ff := e.Group("/clan")
		ff.Use(middleware.BasicAuth(func(pss, pwd string, ctx echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(pss), []byte("joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(pwd), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		}))
		ff.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		ff.POST("/add", providerSvc.AddProductClan)
		ff.POST("/gets", providerSvc.GetProductClans)
		ff.POST("/drop", providerSvc.DropProductClan)
		ff.POST("/update", providerSvc.UpdateProductClan)
		gg := e.Group("/category")
		gg.Use(middleware.BasicAuth(func(pss, pwd string, ctx echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(pss), []byte("joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(pwd), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		}))
		gg.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		gg.POST("/add", providerSvc.AddProductCategory)
		gg.POST("/gets", providerSvc.GetProductCategories)
		gg.POST("/drop", providerSvc.DropProductCategory)
		gg.POST("/update", providerSvc.UpdateProductCategory)
		hh := e.Group("/type")
		hh.Use(middleware.BasicAuth(func(pss, pwd string, ctx echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(pss), []byte("joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(pwd), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		}))
		hh.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		hh.POST("/add", providerSvc.AddProductType)
		hh.POST("/gets", providerSvc.GetProductTypes)
		hh.POST("/drop", providerSvc.DropProductType)
		hh.POST("/update", providerSvc.UpdateProductType)
		ii := e.Group("/product-provider")
		ii.Use(middleware.BasicAuth(func(pss, pwd string, ctx echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(pss), []byte("joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(pwd), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		}))
		ii.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		ii.POST("/add", providerSvc.AddProductProvider)
		ii.POST("/gets", providerSvc.GetProductProviders)
		ii.POST("/drop", providerSvc.DropProductProvider)
		ii.POST("/update", providerSvc.UpdateProductProvider)
		jj := e.Group("/product")
		jj.Use(middleware.BasicAuth(func(pss, pwd string, ctx echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(pss), []byte("joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(pwd), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		}))
		jj.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		jj.POST("/add", providerSvc.AddProduct)
		jj.POST("/gets", providerSvc.GetProducts)
		jj.POST("/drop", providerSvc.DropProduct)
		jj.POST("/update", providerSvc.UpdateProduct)
	}
	{ //eksternal callback
		callback := e.Group("/callback")
		callback.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		callback.POST("/iak", trxSvc.IAKCallback)
	}
	{ //login
		kk := e.Group("/login")
		kk.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		kk.POST("/", admSvc.Login)
	}
	{ //user
		ll := e.Group("/user")
		ll.Use(middleware.JWT([]byte(configs.KEY)))
		ll.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		ll.POST("/", admSvc.CekJwt)
		ll.POST("/get", userAct.GetMerchantOutlets)
		ll.POST("/update", userAct.UpdateMerchantOutlet)
	}
	{ //transaksi
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
	// trxSvc := trxservice.NewApiTrxService(service)
	// // nHierachySvc := nhierarchyservice.NewApiNHierarchyServices(service)
	// nClientSvc := clientservicego.NewApiNHierarchyClientServices(service)
	// nMerchantSvc := merchantservicego.NewApiNHierarchyMerchantServices(service)
	// nMerchantOutletSvc := merchantoutlettservicego.NewApiNHierarchyMerchantOutletServices(service)
	// nUserOutletSvc := useroutlettservicego.NewApiNHierarchyUserOutletServices(service)
	// nUserDashboardSvc := nuserdashboardservice.NewApiNUserDashboardServices(service)
	// nFeaturesSvc := nfeaturesservice.NewApiNFeaturesServices(service)
	// nFeaturesAssignmentSvc := nfeatureassignmentsservice.NewApiNFeaturesAssignmentServices(service)

	// dvc := e.Group("/dvc")
	// // dvc.Use(middleware.JWT([]byte("PRODMKPMobileMyBills")))
	// // dvc.Use(middleware.JWT([]byte(configs.APP_KEY)))
	// dvc.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowCredentials: true,
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	// }))
	// dvc.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	// 	log.Println("[Start]")
	// 	log.Println("EndPoint :", c.Path())
	// 	log.Println("Header :", c.Request().Header)
	// 	log.Println("Body :", string(reqBody))
	// 	log.Println("Response :", string(resBody))
	// 	log.Println("[End]")
	// }))

}
