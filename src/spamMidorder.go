package src

import (
	"antispam/models"
	"antispam/src/contentSpam"
	"antispam/src/picSpam"
	"antispam/src/videoSpam"
	"strings"
)

func dunGetCheckLabels(model []string) string {
	checkLabels := ""
	if len(model) == 0 {
		checkLabels = "100,200,300,500"
	} else {
		modelArrs := []string{}
		for k := range model {
			modelArrs = append(modelArrs, models.DunContentModelsMapping[model[k]])
		}
		checkLabels = strings.Join(modelArrs, ",")
	}
	return checkLabels
}

func dunGetPicCheckLabels(model []string) []string {
	checkLabels := []string{}
	if len(model) == 0 {
		checkLabels = []string{"100", "110", "200", "210", "300", "500"}
	} else {
		for k := range model {
			checkLabels = append(checkLabels, models.DunPictureModelsMapping[model[k]])
		}
	}
	return checkLabels
}

func DunContentCheckSecondStep(data []models.Data, model []string) map[string]models.ContentResult {
	checkLabels := dunGetCheckLabels(model)
	sdkResult := contentSpam.DunContentCheck(data, checkLabels)
	return sdkResult
}

func DunPictureCheckSecondStep(data []models.Data, model []string) (picSpam.DunImageCheckResponse, error) {
	checkLabels := dunGetPicCheckLabels(model)
	sdkResult, err := picSpam.DunImageCheck(data, checkLabels)
	if err != nil {
		return picSpam.DunImageCheckResponse{}, err
	}
	return sdkResult, nil
}

func ShuMeiVideoCheckSecondStep(data []models.Data, model []string) (map[string]models.ContentResult, error) {
	sdkResult, err := videoSpam.ShuMeiVideoContentCheck(data, model)
	if err != nil {
		return nil, err
	}
	return sdkResult, nil
}
