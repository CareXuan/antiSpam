package src

import (
	"antispam/models"
	"antispam/src/picSpam"
)

func ContentCheckThirdStep(re map[string]models.ContentResult) map[string]models.ContentResult {
	return re
}

func PictureCheckThirdStep(re picSpam.DunImageCheckResponse) map[string]models.ContentResult {
	antispam := re.Antispam
	result := make(map[string]models.ContentResult)
	for k := range antispam {
		var r models.ContentResult
		r.Status = models.PictureActionMapping[antispam[k].Action]
		result[antispam[k].Name] = r
	}
	return result
}

func VideoThirdStep(re map[string]models.ContentResult) map[string]models.ContentResult {
	return re
}
