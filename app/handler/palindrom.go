package handler

import (
	"net/http"
	"transfer-pinnacle/app/helper"

	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	hello := "Hello Go Developer"
	ctx.JSON(http.StatusOK, hello)

}

func GetPalindrom(ctx *gin.Context) {
	input1 := Data1{Input: "abcdcba"}
	input2 := Data2{Input: "test"}

	inputString1 := helper.InterfaceToString(input1.Input)
	inputString2 := helper.InterfaceToString(input2.Input)

	getPalindrom1 := helper.IsPalindrome(inputString1)
	getPalindrom2 := helper.IsPalindrome(inputString2)

	arr := map[string]string{
		"input 1 is : ": getPalindrom1,
		"input 2 is : ": getPalindrom2,
	}

	ctx.JSON(http.StatusOK, arr)
}
