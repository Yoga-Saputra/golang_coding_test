package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"transfer-pinnacle/app/helper"
	"transfer-pinnacle/config"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var languges []LanguangeEntity

func GetLanguage(ctx *gin.Context) {
	var res map[string]interface{}
	byteJsonFile, _ := config.LoadJsonFile()
	json.Unmarshal([]byte(byteJsonFile), &res)

	appeared := helper.GetApeared(res["appeared"])
	lang := fmt.Sprintf("%v", res["language"])

	sliceCreated := helper.GetSliceCreated(res["created"])

	sliceInfluencedBy := helper.GetInfluenced(res["relation"], "influenced-by")

	sliceInfluences := helper.GetInfluenced(res["relation"], "influences")

	relation := Relation{
		InfluencedBy: sliceInfluencedBy,
		Influences:   sliceInfluences,
	}

	data := LanguangeEntity{
		Appeared:       appeared,
		Languages:      lang,
		Created:        sliceCreated,
		Functional:     reflect.ValueOf(res["functional"]).Bool(),
		ObjectOriented: reflect.ValueOf(res["object-oriented"]).Bool(),
		Relation:       relation,
	}

	config.Loggers("info", data)

	ctx.JSON(http.StatusOK, data)
}

func GetLanguages(ctx *gin.Context) {
	var getLanguages []LanguagesOnlyEntity

	var payload []map[string]interface{}
	inputBytes, _ := json.Marshal(languges)
	json.Unmarshal([]byte(inputBytes), &payload)

	for _, value := range payload {
		lang := fmt.Sprintf("%v", value["language"])
		langs := LanguagesOnlyEntity{
			Languages: lang,
		}
		getLanguages = append(getLanguages, langs)
	}

	if len(getLanguages) == 0 {
		ctx.JSON(http.StatusNotFound, "no data available")
		return
	}

	ctx.JSON(http.StatusOK, getLanguages)
}

func GetLanguageById(ctx *gin.Context) {
	var data LanguangeEntity
	var input GetLanguangeDetailInput

	err := ctx.ShouldBindUri(&input)
	if err != nil {
		panic(err)
	}

	if *input.ID == 0 {
		ctx.JSON(http.StatusNotFound, "ID cannot be zero")
		return
	}

	var payload []map[string]interface{}
	inputBytes, _ := json.Marshal(languges)
	json.Unmarshal([]byte(inputBytes), &payload)

	var id []int
	for key, value := range payload {
		if key+1 == *input.ID {
			appeared := helper.GetApeared(value["appeared"])
			lang := fmt.Sprintf("%v", value["language"])

			sliceCreated := helper.GetSliceCreated(value["created"])

			sliceInfluencedBy := helper.GetInfluenced(value["relation"], "influenced-by")

			sliceInfluences := helper.GetInfluenced(value["relation"], "influences")

			relation := Relation{
				InfluencedBy: sliceInfluencedBy,
				Influences:   sliceInfluences,
			}

			langData := LanguangeEntity{
				Appeared:       appeared,
				Languages:      lang,
				Created:        sliceCreated,
				Functional:     reflect.ValueOf(value["functional"]).Bool(), // string to bool
				ObjectOriented: reflect.ValueOf(value["object-oriented"]).Bool(),
				Relation:       relation,
			}
			data = langData
			id = append(id, *input.ID)
		}

	}

	if id == nil {
		ctx.JSON(http.StatusOK, "no data available")
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func PostLanguage(ctx *gin.Context) {
	var input []LanguangeEntity

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		panic(err)
	}

	var payload []map[string]interface{}
	inputBytes, _ := json.Marshal(input)
	json.Unmarshal([]byte(inputBytes), &payload)

	for _, value := range payload {
		appeared := helper.GetApeared(value["appeared"])
		lang := fmt.Sprintf("%v", value["language"])

		sliceCreated := helper.GetSliceCreated(value["created"])

		sliceInfluencedBy := helper.GetInfluenced(value["relation"], "influenced-by")

		sliceInfluences := helper.GetInfluenced(value["relation"], "influences")

		relation := Relation{
			InfluencedBy: sliceInfluencedBy,
			Influences:   sliceInfluences,
		}

		langData := LanguangeEntity{
			Appeared:       appeared,
			Languages:      lang,
			Created:        sliceCreated,
			Functional:     reflect.ValueOf(value["functional"]).Bool(), // string to bool
			ObjectOriented: reflect.ValueOf(value["object-oriented"]).Bool(),
			Relation:       relation,
		}
		languges = append(languges, langData)
	}

	ctx.JSON(http.StatusOK, languges)
}

func UpdateLanguage(ctx *gin.Context) {
	var input LanguangeEntity
	var param GetLanguangeDetailInput

	err := ctx.ShouldBindUri(&param)
	if err != nil {
		panic(err)
	}

	err = ctx.ShouldBindBodyWith(&input, binding.JSON)
	if err != nil {
		panic(err)
	}

	var payload []map[string]interface{}
	inputBytes, _ := json.Marshal(languges)
	json.Unmarshal([]byte(inputBytes), &payload)

	var id int

	for key, value := range payload {
		if key == *param.ID {
			appeared := helper.GetApeared(input.Appeared)
			lang := fmt.Sprintf("%v", input.Languages)

			sliceCreated := helper.GetSliceCreated(input.Created)

			rel, _ := helper.StructToMap(input.Relation)
			sliceInfluencedBy := helper.GetInfluenced(rel, "influenced-by")

			sliceInfluences := helper.GetInfluenced(rel, "influences")

			relation := Relation{
				InfluencedBy: sliceInfluencedBy,
				Influences:   sliceInfluences,
			}

			languges[key].Appeared = appeared
			languges[key].Languages = lang
			languges[key].Created = sliceCreated
			languges[key].Functional = reflect.ValueOf(input.Functional).Bool()         // string to bool
			languges[key].ObjectOriented = reflect.ValueOf(input.ObjectOriented).Bool() // string to bool
			languges[key].Relation = relation

			config.Loggers("info", value)

			id = key
		}
	}

	if id == 0 {
		ctx.JSON(http.StatusNotFound, "no data available")
		return
	}

	ctx.JSON(http.StatusOK, "data already updated")
}

func DeleteLanguage(ctx *gin.Context) {
	var param GetLanguangeDetailInput

	err := ctx.ShouldBindUri(&param)
	if err != nil {
		panic(err)
	}

	var payload []map[string]interface{}
	inputBytes, _ := json.Marshal(languges)
	json.Unmarshal([]byte(inputBytes), &payload)

	var id int

	for key, value := range payload {
		if key == *param.ID {
			languges = append(languges[:key], languges[key+1:]...)
			id = key

			config.Loggers("info", value)
		}
	}

	if id == 0 {
		ctx.JSON(http.StatusNotFound, "data not available")
		return
	}

	ctx.JSON(http.StatusOK, "data already delete")
}
