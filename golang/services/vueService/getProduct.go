package vueservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (svc VueService) GetProduct(ctx echo.Context) error {
	// var (
	// 	svcName = " Getproduct"
	// 	errSvc  = "Error " + svcName
	// )
	type resVueProduct struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Img  string `json:"img"`
	}
	var respp []resVueProduct

	for i := 0; i < 9; i++ {
		resp := resVueProduct{
			Id:   i,
			Name: "HP" + strconv.Itoa(i),
			Img:  "https://www.exabytes.co.id/blog/wp-content/uploads/2020/08/fotoproduk2.jpeg",
		}
		respp = append(respp, resp)
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", respp)
	return ctx.JSON(http.StatusOK, result)
}
func (svc VueService) GetProductSingle(ctx echo.Context) error {
	// var (
	// 	svcName = " Getproduct"
	// 	errSvc  = "Error " + svcName
	// )
	type req struct {
		Id int `json:"id"`
	}
	reqq := new(req)
	ctx.Bind(reqq)
	type photo struct {
		Id  interface{} `json:"id"`
		Img string      `json:"img"`
	}
	type resVueProduct struct {
		Id    int     `json:"id"`
		Name  string  `json:"name"`
		Photo []photo `json:"photo"`
	}
	a, _ := strconv.Atoi(ctx.Param("id"))
	if a == 0 || a != 2 {
		fmt.Println("::::::::MBUH ", ctx.Param("id"))
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "Failed", nil)
		return ctx.JSON(http.StatusOK, result)
	}

	var photoar []photo
	var aa []string
	aa = append(aa, "https://plus.unsplash.com/premium_photo-1674688194029-17dda3aaf779?q=80&w=2680&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D")
	aa = append(aa, "https://www.exabytes.co.id/blog/wp-content/uploads/2020/08/fotoproduk2.jpeg")

	for i := 0; i < len(aa); i++ {
		p := photo{
			Id:  i,
			Img: aa[i],
		}
		photoar = append(photoar, p)
	}
	resp := resVueProduct{
		Id:    2,
		Name:  "HP",
		Photo: photoar,
	}

	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", resp)
	return ctx.JSON(http.StatusOK, result)
}
