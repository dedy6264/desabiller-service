package routes

import (
	"desabiller/configs"
	"desabiller/services"
	administrationservice "desabiller/services/administrationService"
	hierarchyservicego "desabiller/services/hierarchyService"
	nfeatureassignmentsservice "desabiller/services/nFeaturesService/featureAssignmentService"
	nfeaturesservice "desabiller/services/nFeaturesService/featuresService"
	clientservicego "desabiller/services/nHierarchyService/clientService"
	merchantoutlettservicego "desabiller/services/nHierarchyService/merchantOutletService"
	merchantservicego "desabiller/services/nHierarchyService/merchantService"
	useroutlettservicego "desabiller/services/nHierarchyService/userOutletService"
	nuserdashboardservice "desabiller/services/nUserDashboardService"
	productservice "desabiller/services/productService"
	trxservice "desabiller/services/trxService"
	"log"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RouteApi(e echo.Echo, service services.UsecaseService) {
	productSvc := productservice.ApiProduct(service)
	admSvc := administrationservice.ApiAdministration(service)
	hierachySvc := hierarchyservicego.ApiHierarchy(service)
	trxSvc := trxservice.NewApiTrxService(service)
	// nHierachySvc := nhierarchyservice.NewApiNHierarchyServices(service)
	nClientSvc := clientservicego.NewApiNHierarchyClientServices(service)
	nMerchantSvc := merchantservicego.NewApiNHierarchyMerchantServices(service)
	nMerchantOutletSvc := merchantoutlettservicego.NewApiNHierarchyMerchantOutletServices(service)
	nUserOutletSvc := useroutlettservicego.NewApiNHierarchyUserOutletServices(service)
	nUserDashboardSvc := nuserdashboardservice.NewApiNUserDashboardServices(service)
	nFeaturesSvc := nfeaturesservice.NewApiNFeaturesServices(service)
	nFeaturesAssignmentSvc := nfeatureassignmentsservice.NewApiNFeaturesAssignmentServices(service)

	//login
	open := e.Group("/open")
	{
		open.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		open.POST("/login", admSvc.DeviceLogin)
		// open.OPTIONS("/login", admSvc.DeviceLogin)
	}
	//device activity
	dvc := e.Group("/dvc")
	dvc.Use(middleware.JWT([]byte("PRODMKPMobileMyBills")))
	// dvc.Use(middleware.JWT([]byte(configs.APP_KEY)))
	dvc.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	dvc.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Println("[Start]")
		log.Println("EndPoint :", c.Path())
		log.Println("Header :", c.Request().Header)
		log.Println("Body :", string(reqBody))
		log.Println("Response :", string(resBody))
		log.Println("[End]")
	}))
	{ //product device
		pPo := dvc.Group("/productpos")
		pPo.POST("/list", productSvc.GetListProductPos)
		// pPo.OPTIONS("/list", productSvc.GetListProductPos)
		pPo.POST("/add", productSvc.AddProductPos)
		pPo.POST("/update", productSvc.UpdateProductPos)
		pPo.POST("/drop", productSvc.DropProductPos)
	}
	{ //product category device
		pCa := dvc.Group("/productcategory")
		pCa.POST("/list", productSvc.GetListProductCategory)
		pCa.POST("/add", productSvc.AddProductCategory)
		pCa.POST("/update", productSvc.UpdateProductCategory)
		pCa.POST("/drop", productSvc.DropProductCategory)
	}

	dash := e.Group("/dash")
	{
		dash.Use(middleware.JWT([]byte(configs.APP_KEY)))
		dash.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}))
		dash.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		dash.GET("/cek", hierachySvc.CheckUserStatus)
		{
			cli := dash.Group("/client")
			cli.POST("/list", hierachySvc.GetListClient)
			cli.POST("/add", hierachySvc.AddClient)
			cli.POST("/drop", hierachySvc.DropClient)
			cli.POST("/update", hierachySvc.UpdateClient)
		}
		{
			mer := dash.Group("/merchant")
			mer.POST("/list", hierachySvc.GetListMerchant)
			mer.POST("/add", hierachySvc.AddMerchant)
			mer.POST("/drop", hierachySvc.DropMerchant)
			mer.POST("/update", hierachySvc.UpdateMerchant)
		}
		{
			mOu := dash.Group("/merchantoutlet")
			mOu.POST("/list", hierachySvc.GetListMerchantOutlet)
			mOu.POST("/add", hierachySvc.AddMerchantOutlet)
			mOu.POST("/drop", hierachySvc.DropMerchantOutlet)
			mOu.POST("/update", hierachySvc.UpdateMerchantOutlet)
		}
		{
			uOu := dash.Group("/useroutlet")
			uOu.POST("/list", hierachySvc.GetListUserOutlet)
			uOu.POST("/add", hierachySvc.AddUserOutlet)
			uOu.POST("/drop", hierachySvc.DropUserOutlet)
			uOu.POST("/update", hierachySvc.UpdateUserOutlet)
		}
		{
			uDe := dash.Group("/outletdevice")
			uDe.POST("/list", hierachySvc.GetListOutletDevice)
			uDe.POST("/add", hierachySvc.AddOutletDevice)
			uDe.POST("/drop", hierachySvc.DropOutletDevice)
			uDe.POST("/update", hierachySvc.UpdateOutletDevice)
		}
		{
			pTy := dash.Group("/producttype")
			pTy.POST("/list", productSvc.GetListProductType)
		}
		{
			pCa := dash.Group("/productcategory")
			pCa.POST("/list", productSvc.GetListProductCategory)
			pCa.POST("/add", productSvc.AddProductCategory)
			pCa.POST("/update", productSvc.UpdateProductCategory)
			pCa.POST("/drop", productSvc.DropProductCategory)
		}
		{
			pBp := dash.Group("/productbillerprovider")
			pBp.POST("/list", productSvc.GetListProductBillerProvider)
			pBp.POST("/add", productSvc.AddProductBillerProvider)
			pBp.POST("/update", productSvc.UpdateProductBillerProvider)
			pBp.POST("/drop", productSvc.DropProductBillerProvider)
		}
		{
			pBi := dash.Group("/productbiller")
			pBi.POST("/list", productSvc.GetListProductBiller)
			pBi.POST("/add", productSvc.AddProductBiller)
			pBi.POST("/update", productSvc.UpdateProductBiller)
			pBi.POST("/drop", productSvc.DropProductBiller)
		}
		{
			pPo := dash.Group("/productpos")
			pPo.POST("/list", productSvc.GetListProductPos)
			pPo.OPTIONS("/list", productSvc.GetListProductPos)
			pPo.POST("/add", productSvc.AddProductPos)
			pPo.OPTIONS("/add", productSvc.AddProductPos)
			pPo.POST("/update", productSvc.UpdateProductPos)
			pPo.POST("/drop", productSvc.DropProductPos)
		}
		{
			seg := dash.Group("/segment")
			seg.POST("/list", productSvc.GetListSegment)
			seg.POST("/add", productSvc.AddSegment)
			seg.POST("/update", productSvc.UpdateSegment)
			seg.POST("/drop", productSvc.DropSegment)
		}
		{
			sPr := dash.Group("/segmentproduct")
			sPr.POST("/list", productSvc.GetListSegmentProduct)
			sPr.POST("/add", productSvc.AddSegmentProduct)
			sPr.POST("/update", productSvc.UpdateSegmentProduct)
			sPr.POST("/drop", productSvc.DropSegmentProduct)
		}
		{
			pMd := dash.Group("/paymentmethod")
			pMd.POST("/list", productSvc.GetListPaymentMethod)
			pMd.POST("/add", productSvc.AddPaymentMethod)
			pMd.POST("/update", productSvc.UpdatePaymentMethod)
			pMd.POST("/drop", productSvc.DropPaymentMethod)
		}
		{
			pMc := dash.Group("/paymentmethodcategory")
			pMc.POST("/list", productSvc.GetListPaymentMethodCategory)
			pMc.POST("/add", productSvc.AddPaymentMethodCategory)
			pMc.POST("/update", productSvc.UpdatePaymentMethodCategory)
			pMc.POST("/drop", productSvc.DropPaymentMethodCategory)
		}
	}
	//api khusus transaksi pos dan biller
	tRx := e.Group("/trx")
	{
		tRx.Use(middleware.JWT([]byte(configs.APP_KEY)))
		tRx.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}))
		tRx.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			log.Println("[Start]")
			log.Println("EndPoint :", c.Path())
			log.Println("Header :", c.Request().Header)
			log.Println("Body :", string(reqBody))
			log.Println("Response :", string(resBody))
			log.Println("[End]")
		}))
		{
			pos := tRx.Group("/pos")
			pos.POST("/inq", trxSvc.InquiryPos)
			pos.POST("/payment", trxSvc.PaymentPos)
			pos.POST("/report", trxSvc.GetListReportPos)
		}
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	v1 := e.Group("/v1") //nVersion
	v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	v1.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Println("[Start]")
		log.Println("EndPoint :", c.Path())
		log.Println("Header :", c.Request().Header)
		log.Println("Body :", string(reqBody))
		log.Println("Response :", string(resBody))
		log.Println("[End]")
	}))
	v1.Use(middleware.JWT([]byte(configs.APP_KEY)))
	// v1.Use(middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
	// 	if subtle.ConstantTimeCompare([]byte(username), []byte("mkpmobile")) == 1 &&
	// 		subtle.ConstantTimeCompare([]byte(password), []byte("mkpmobile123")) == 1 {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))
	{
		v1.POST("/addUserDash", nUserDashboardSvc.CreateUserDashboard)
		v1.POST("/getUserDash", nUserDashboardSvc.GetUserDashboard)
		v1.POST("/getListUserDash", nUserDashboardSvc.GetListUserDashboard)
		v1.POST("/dropUserDash", nUserDashboardSvc.DropUserDashboard)
		v1.POST("/updateUserDash", nUserDashboardSvc.UpdateUserDashboard)
	}
	{
		v1.POST("/addClient", nClientSvc.CreateClientService)
		v1.POST("/getSingleClient", nClientSvc.GetClientService)
		v1.POST("/getListClient", nClientSvc.GetListClientService)
		v1.POST("/dropClient", nClientSvc.DropClientService)
		v1.POST("/updateClient", nClientSvc.UpdateClientService)
	}
	{
		v1.POST("/addMerchant", nMerchantSvc.CreateMerchantService)
		v1.POST("/getSingleMerchant", nMerchantSvc.GetMerchantService)
		v1.POST("/getListMerchant", nMerchantSvc.GetListMerchantService)
		v1.POST("/dropMerchant", nMerchantSvc.DropMerchantService)
		v1.POST("/updateMerchant", nMerchantSvc.UpdateMerchantService)
	}
	{
		v1.POST("/addMerchantOutlet", nMerchantOutletSvc.CreateMerchantOutletService)
		v1.POST("/getSingleMerchantOutlet", nMerchantOutletSvc.GetMerchantOutletService)
		v1.POST("/getListMerchantOutlet", nMerchantOutletSvc.GetListMerchantOutletService)
		v1.POST("/dropMerchantOutlet", nMerchantOutletSvc.DropMerchantOutletService)
		v1.POST("/updateMerchantOutlet", nMerchantOutletSvc.UpdateMerchantOutletService)
	}
	{
		v1.POST("/addUserOutlet", nUserOutletSvc.CreateUserOutletService)
		v1.POST("/getSingleUserOutlet", nUserOutletSvc.GetUserOutletService)
		v1.POST("/getListUserOutlet", nUserOutletSvc.GetListUserOutletService)
		v1.POST("/dropUserOutlet", nUserOutletSvc.DropUserOutletService)
		v1.POST("/updateUserOutlet", nUserOutletSvc.UpdateUserOutletService)
	}
	{
		v1.POST("/addFeatures", nFeaturesSvc.CreateFeaturesService)
		v1.POST("/getSingleFeatures", nFeaturesSvc.GetFeaturesService)
		v1.POST("/getListFeatures", nFeaturesSvc.GetListFeaturesService)
		v1.POST("/dropFeatures", nFeaturesSvc.DropFeaturesService)
		v1.POST("/updateFeatures", nFeaturesSvc.UpdateFeaturesService)
	}
	{
		v1.POST("/addFeaturesAssignment", nFeaturesAssignmentSvc.CreateFeaturesAssignmentService)
		v1.POST("/getSingleFeaturesAssignment", nFeaturesAssignmentSvc.GetFeaturesAssignmentService)
		v1.POST("/getListFeaturesAssignment", nFeaturesAssignmentSvc.GetListFeaturesAssignmentService)
		v1.POST("/dropFeaturesAssignment", nFeaturesAssignmentSvc.DropFeaturesAssignmentService)
		v1.POST("/updateFeaturesAssignment", nFeaturesAssignmentSvc.UpdateFeaturesAssignmentService)
	}
	v0 := e.Group("/v0") //nVersion
	v0.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	v0.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Println("[Start]")
		log.Println("EndPoint :", c.Path())
		log.Println("Header :", c.Request().Header)
		log.Println("Body :", string(reqBody))
		log.Println("Response :", string(resBody))
		log.Println("[End]")
	}))
	v0.POST("/loginUserDash", nUserDashboardSvc.LoginUserDashboard)
}
