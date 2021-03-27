package controllers

import (
	"net/http"

	conf "serasa/config"

	models "serasa/models"
	services "serasa/services"
	utils "serasa/utils"
)

var negativationService = services.NegativationService{}
var config = conf.Config{}

func init() {
	config.Read()

	negativationService.Server = config.Server
	negativationService.Database = config.Database
	negativationService.Connect()
}

func List(w http.ResponseWriter, r *http.Request) {
	filters := models.FilterNegativation{}
	params := r.URL.Query()

	if len(params["companyName"]) > 0 {
		companyName := params["companyName"][0]
		filters.CompanyName = &companyName
	}

	if len(params["companyDocument"]) > 0 {
		companyDocument := params["companyDocument"][0]
		filters.CompanyDocument = &companyDocument
	}

	if len(params["customerDocument"]) > 0 {
		customerDocument := params["customerDocument"][0]
		filters.CustomerDocument = &customerDocument
	}

	negativations, err := negativationService.List(filters)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
  if len(negativations) == 0 {
    utils.ResponseWithJSON(w, http.StatusNoContent, []models.Negativation{})
  }

	utils.ResponseWithJSON(w, http.StatusOK, negativations)
}

func Create(w http.ResponseWriter, r *http.Request) {
	negativations, err := utils.ListNegativationsByAPI()
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := negativationService.Create(negativations); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJSON(w, http.StatusCreated, negativations)
}
