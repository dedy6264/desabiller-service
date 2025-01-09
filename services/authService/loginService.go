package administrationservice

// func (svc AdministrationService) DeviceLogin(ctx echo.Context) error {
// 	var (
// 		svcName        = "DeviceLogin"
// 		resultSvc      models.ResLoginUserOutlet
// 		formatLogError = "Erorr " + svcName + " ::"
// 		uID            int
// 		oID            int
// 		mID            int
// 		cID            int
// 	)
// 	//binding n validate required value
// 	req := new(models.ReqLogin)
// 	status, err := helpers.BindValidate(req, ctx)
// 	if err != nil {
// 		log.Println("FAILLED BINDING", err.Error())
// 		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "FAILLED BINDING"+err.Error(), nil)
// 		return ctx.JSON(http.StatusNotFound, result)
// 	}
// 	resEec, err := helpers.PswEnc(req.Password)
// 	if err != nil {
// 		log.Println(formatLogError+"dcryp", err)
// 		result := helpers.ResponseJSON(false, configs.DB_NOT_FOUND, "wrong username or password", nil)
// 		return ctx.JSON(http.StatusOK, result)
// 	}
// 	// if strings.Contains(req.Username, "@") {
// 	// 	resUser, err, _ := svc.service.RepoHierarchy.GetListUser(models.ReqUserList{
// 	// 		Username:        "",
// 	// 		Password:        resEec,
// 	// 		Email:           req.Username,
// 	// 		RoleSegmentId:   0,
// 	// 		RoleSegmentName: "",
// 	// 		RoleSegmentCode: "",
// 	// 		Limit:           0,
// 	// 		Draw:            0,
// 	// 		AscDesc:         "",
// 	// 		SortBy:          "",
// 	// 	})
// 	// 	if err != nil {
// 	// 		log.Println("Not Found", err.Error())
// 	// 		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Not Found"+err.Error(), nil)
// 	// 		return ctx.JSON(http.StatusNotFound, result)
// 	// 	}
// 	// 	if len(resUser) == 0 {
// 	// 		log.Println("Not Found")
// 	// 		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Not Found", nil)
// 	// 		return ctx.JSON(http.StatusNotFound, result)
// 	// 	}
// 	// 	if resEec != resUser[0].Password {
// 	// 		result := helpers.ResponseJSON(false, configs.DB_NOT_FOUND, "wrong username or password", nil)
// 	// 		return ctx.JSON(http.StatusOK, result)
// 	// 	}
// 	// 	if resUser[0].RoleSegmentId == 1 {
// 	// 		token, err := helpers.TokenJwtGenerate(0, 0, 0, 0, "")
// 	// 		if err != nil {
// 	// 			result := helpers.ResponseJSON(false, configs.DB_NOT_FOUND, "token error", nil)
// 	// 			return ctx.JSON(http.StatusOK, result)
// 	// 		}
// 	// 		resultSvc.Token = token
// 	// 		resultSvc.OutletUsername = resUser[0].Username
// 	// 		resultSvc.Role = append(resultSvc.Role, "ALL")
// 	// 		result := helpers.ResponseJSON(true, configs.SUCCESS_CODE, "Success", resultSvc)
// 	// 		return ctx.JSON(http.StatusOK, result)
// 	// 	}
// 	// 	rolee := resUser[0].Role
// 	// 	ss := strings.Split(rolee, `"`)
// 	// 	ss = strings.Split(strings.Join(ss, ``), `[`)
// 	// 	ss = strings.Split(strings.Join(ss, ``), ` `)
// 	// 	ss = strings.Split(strings.Join(ss, ``), `]`)
// 	// 	ss = strings.Split(ss[0], `,`)
// 	// 	if resUser[0].HierarchyType == "merchant" {
// 	// 		mID = resUser[0].HierarchyId
// 	// 		resMerch, status := svc.service.RepoHierarchy.GetListMerchant(models.ReqGetListMerchant{
// 	// 			ID:           mID,
// 	// 			MerchantName: "",
// 	// 			ClientId:     0,
// 	// 			ClientName:   "",
// 	// 			Limit:        0,
// 	// 			Offset:       0,
// 	// 			OrderBy:      "",
// 	// 			StartDate:    "",
// 	// 			EndDate:      "",
// 	// 			Username:     "",
// 	// 		})
// 	// 		if !status {
// 	// 			result := helpers.ResponseJSON(false, configs.DB_NOT_FOUND, "Client Not Found", nil)
// 	// 			return ctx.JSON(http.StatusOK, result)
// 	// 		}
// 	// 		cID = resMerch[0].ClientId
// 	// 		resultSvc.MerchantId = resUser[0].HierarchyId
// 	// 		resultSvc.Nickname = resUser[0].Username
// 	// 		token, err := helpers.TokenJwtGenerate(mID, uID, oID, cID, "")
// 	// 		if err != nil {
// 	// 			result := helpers.ResponseJSON(false, configs.DB_NOT_FOUND, "token error", nil)
// 	// 			return ctx.JSON(http.StatusOK, result)
// 	// 		}
// 	// 		resultSvc.Token = token
// 	// 		resultSvc.OutletUsername = resUser[0].Username
// 	// 		resultSvc.Role = ss
// 	// 		result := helpers.ResponseJSON(true, configs.SUCCESS_CODE, "Success", resultSvc)
// 	// 		return ctx.JSON(http.StatusOK, result)
// 	// 	}
// 	// 	log.Println(formatLogError+" Not Authorized", err)
// 	// 	result := helpers.ResponseJSON(false, configs.DB_NOT_FOUND, "Not Authorized", nil)
// 	// 	return ctx.JSON(http.StatusOK, result)
// 	// }

// 	resUser, status := svc.service.RepoHierarchy.GetListUserOutlet(models.ReqGetListUserOutlet{
// 		ID:                 0,
// 		Nickname:           "",
// 		OutletUsername:     req.Username,
// 		OutletPassword:     resEec,
// 		MerchantOutletId:   0,
// 		MerchantOutletName: "",
// 		MerchantId:         0,
// 		MerchantName:       "",
// 		ClientId:           0,
// 		Limit:              0,
// 		Offset:             0,
// 		OrderBy:            "",
// 		StartDate:          "",
// 		EndDate:            "",
// 		Username:           "",
// 	})
// 	if !status {

// 		log.Println(formatLogError + " Not Found")
// 		result := helpers.ResponseJSON(false, configs.DB_NOT_FOUND, "password or username is not found", nil)
// 		return ctx.JSON(http.StatusOK, result)
// 	}
// 	if len(resUser) == 1 {
// 		for _, data := range resUser {
// 			mID = data.MerchantId
// 			uID = data.ID
// 			oID = data.MerchantOutletId
// 			cID = data.ClientId
// 			resultSvc.ClientId = data.ClientId
// 			resultSvc.ClientName = data.ClientName
// 			resultSvc.MerchantId = data.MerchantId
// 			resultSvc.MerchantName = data.MerchantName
// 			resultSvc.MerchantOutletId = data.MerchantOutletId
// 			resultSvc.MerchantOutletName = data.MerchantOutletName
// 			resultSvc.Nickname = data.Nickname
// 			resultSvc.OutletUsername = data.OutletUsername
// 			// resultSvc.Role = append(resultSvc.Role, "POS", "PRODUCT", "PRODUCTCATEGORY", "REPORTPOS")
// 			fmt.Println(resEec, "::", data.OutletPassword)
// 			if resEec != data.OutletPassword {
// 				log.Println(formatLogError + " Not match")
// 				result := helpers.ResponseJSON(false, configs.DB_NOT_FOUND, "wrong username or password", nil)
// 				return ctx.JSON(http.StatusOK, result)
// 			}
// 			data.OutletPassword = ""
// 		}
// 	}
// 	//make jwt token
// 	token, err := helpers.TokenJwtGenerate(mID, uID, oID, cID, req.DeviceSn)
// 	if err != nil {
// 		log.Println("ERRROR ", err)
// 		result := helpers.ResponseJSON(false, configs.DB_NOT_FOUND, "error generate"+err.Error(), nil)
// 		return ctx.JSON(http.StatusOK, result)
// 	}
// 	resultSvc.Token = token
// 	result := helpers.ResponseJSON(true, configs.SUCCESS_CODE, "Success", resultSvc)
// 	return ctx.JSON(http.StatusOK, result)
// }
