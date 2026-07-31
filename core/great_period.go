package core

import "tuvi/models"

func MapJias(year int, page *models.HoroscopePage) ([]map[string]any, error) {
	jiasID, _ := CanChiNamDuongLich(year)
	result := make([]map[string]any, 0, 12)
	for i := range page.TwelvePlaces {
		if i == 0 {
			continue
		}
		can, err := GetPalaceJia(jiasID, i)
		if err != nil {
			return nil, err
		}
		page.TwelvePlaces[i].Jia = can.Name
	}
	return result, nil
}
