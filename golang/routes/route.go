package routes

import (
	"desabiller/services"
	hierarchyservicego "desabiller/services/hierarchyService"
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
	// admSvc := administrationservice.ApiAdministration(service)
	hierachySvc := hierarchyservicego.ApiHierarchy(service)
	aa := e.Group("/client")
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
