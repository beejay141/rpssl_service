package services

import (
	"fmt"
	"math"
	"rpssl_service/config"
	"rpssl_service/infrastructures/helpers"
	"rpssl_service/utils"
)

type IRandomNumberService interface {
	GetRandomNumber(limit int) (int,error)
}

type RandomNumberService struct {
	RestRequestHelperService helpers.IRestRequestHelper
}

func (service RandomNumberService) GetRandomNumber(limit int) (int,error) {
    status,byteData,err := service.RestRequestHelperService.Get(fmt.Sprintf("%s/random",config.GetString("random_number_base_url","")),nil)
	if status != 200 {
		return 1, err
	}

	var responseData map[string]float64
	if err = utils.DecodeJSON(byteData,&responseData); err != nil {
		return 1, err
	}

	number := math.Ceil((responseData["random_number"]/100) * float64(limit))
	if number == 0 {
		return 1,nil
	}

	return int(number),nil
}
