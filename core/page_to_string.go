package core

import (
	"fmt"
	"strings"
	"tuvi/models"
)

// GetPlaceIdInThree tra ve danh sach cung so trong tam hop voi cung so cho truoc.
func GetPlaceIdInThree(placeId int) []int {
	normalized := mod12(placeId)
	if normalized == 0 {
		normalized = 12
	}

	switch normalized {
	case 1:
		return []int{5, 9}
	case 2:
		return []int{6, 10}
	case 3:
		return []int{7, 11}
	case 4:
		return []int{8, 12}
	case 5:
		return []int{9, 1}
	case 6:
		return []int{10, 2}
	case 7:
		return []int{11, 3}
	case 8:
		return []int{12, 4}
	case 9:
		return []int{1, 5}
	case 10:
		return []int{2, 6}
	case 11:
		return []int{3, 7}
	case 12:
		return []int{4, 8}
	default:
		return []int{}
	}
}

// GetPlaceIdOffsite tra ve cung doi voi cung so cho truoc.
func GetPlaceIdOffsite(placeId int) int {
	return (placeId+6-1)%12 + 1
}

// GetLeftRight tra ve 2 cung ke voi cung so cho truoc.
func GetLeftRight(placeId int) []int {
	return []int{
		(placeId+11-1)%12 + 1,
		(placeId+1-1)%12 + 1,
	}
}

func GetPlaceIdFromPlaceName(placeOptions int, page models.HoroscopePage) (int, error) {
	placeName := models.MAIN_PLACE_OPTIONS[placeOptions]
	for _, place := range page.TwelvePlaces {
		if place.MainPlace == placeName {
			return place.PlaceId, nil
		}
	}

	fmt.Printf("place name %s not found in horoscope page", placeName)
	return 0, fmt.Errorf("place name %s not found in horoscope page", placeName)
}

func GetPlaceDetailString(placeId int, page models.HoroscopePage) string {
	isIntercept := false
	isBlockade := false
	starNames := []string{}
	mainStarNames := []string{}

	mainStraStr := ""
	subStarStr := ""
	animalName := ""
	for _, place := range page.TwelvePlaces {
		if place.PlaceId == placeId {
			animalName = place.PlaceName
			if place.IsIntercept {
				isIntercept = true
			}
			if place.IsBlockade {
				isBlockade = true
			}
			for _, star := range place.Stars {
				if star.StarType == 1 {
					mainStarNames = append(mainStarNames, star.Name)
				} else {
					starNames = append(starNames, star.Name)
				}
			}
		}
	}

	if len(mainStarNames) == 0 {
		mainStraStr = fmt.Sprintf("Cung vô chính diệu tại %s", animalName)
	} else if len(mainStarNames) >= 2 {
		mainStraStr = fmt.Sprintf("Chính tinh %s đồng cung tại %s", strings.Join(mainStarNames, ", "), animalName)
	} else {
		mainStraStr = fmt.Sprintf("Chính tinh %s đơn thủ tại %s", mainStarNames[0], animalName)
	}

	subStarStr = fmt.Sprintf("Các sao phụ tinh: %s", strings.Join(starNames, ", "))
	if isIntercept {
		subStarStr += ", Triệt lộ không vong"
	}
	if isBlockade {
		subStarStr += ", Tuần trung không vong"
	}

	return fmt.Sprintf("%s. %s.", mainStraStr, subStarStr)
}

func GetOffsitePlaceDetailString(placeId int, page models.HoroscopePage) string {
	offsitePlaceId := GetPlaceIdOffsite(placeId)
	detail := GetPlaceDetailString(offsitePlaceId, page)
	for _, place := range page.TwelvePlaces {
		if place.PlaceId == offsitePlaceId {
			return fmt.Sprintf("(Cung  %s) %s", place.MainPlace, detail)
		}
	}
	return detail
}

func GetThreeDetailString(placeId int, page models.HoroscopePage) string {
	details := []string{}
	placeIdInThree := GetPlaceIdInThree(placeId)

	for _, pid := range placeIdInThree {
		for _, place := range page.TwelvePlaces {
			if place.PlaceId == pid {
				namne := place.MainPlace
				placeDetail := GetPlaceDetailString(pid, page)
				details = append(details, fmt.Sprintf("(Cung %s) %s", namne, placeDetail))
			}
		}
	}

	return strings.Join(details, ";")
}

func GetLeftRightDetailString(placeId int, page models.HoroscopePage) string {
	details := []string{}
	placeIdLeftRight := GetLeftRight(placeId)

	for _, pid := range placeIdLeftRight {
		for _, place := range page.TwelvePlaces {
			if place.PlaceId == pid {
				namne := place.MainPlace
				placeDetail := GetPlaceDetailString(pid, page)
				details = append(details, fmt.Sprintf("(Cung %s) %s", namne, placeDetail))
			}
		}
	}
	return strings.Join(details, ";")

}

const SYSTEM_PORMT = "Bạn là trợ lý tử vi. Nhiệm vụ: Luận giải 1 cung với dữ liệu json đầu vào. Trả lời chi tiết, Giải thích tính chất theo các sao hoặc bộ cách, không nói chung chung. Không cần nói giảm nếu là điều xấu."

func BuildMessageDetailPlace(mainId int, gender int, page models.HoroscopePage) string {
	placeId, err := GetPlaceIdFromPlaceName(mainId, page)
	if err != nil {
		return ""
	}

	placeName := models.MAIN_PLACE_OPTIONS[mainId]
	placeDetail := GetPlaceDetailString(placeId, page)
	three := GetThreeDetailString(placeId, page)
	opposite := GetOffsitePlaceDetailString(placeId, page)
	nextPlace := GetLeftRightDetailString(placeId, page)
	g := "Nam"
	if gender != 1 {
		g = "Nữ"
	}

	return fmt.Sprintf(
		"Hãy Luận Giải cung này theo bố cục:\n"+
			"1) Tổng quan cung theo bộ sao chính tinh\n"+
			"2) Từng bộ cách sao phụ tinh bạn biết\n"+
			"3) Giải thích ngắn ngọn các sao khi ở trong cung\n"+
			"4) Kết luận tổng quan\n"+
			"Dữ liệu các cung của người %s mệnh này:\n"+
			"Các sao tại cung %s: %s\n"+
			"Các sao trong 2 cung còn lại tạo thành tam hợp: %s\n"+
			"Các sao cung đối diện chiếu về: %s.\n"+
			"Các sao tại 2 cung liền kề: %s\n",
		g,
		placeName,
		placeDetail,
		three,
		opposite,
		nextPlace,
	)
}

func BuildMessageAppearance(gender int, page models.HoroscopePage) string {
	// placeName := models.MAIN_PLACE_OPTIONS[1]
	placeDetail := GetPlaceDetailString(1, page)
	three := GetThreeDetailString(1, page)
	opposite := GetOffsitePlaceDetailString(1, page)

	g := "Nam"
	if gender != 1 {
		g = "Nữ"
	}

	return fmt.Sprintf(
		"Mục đích giúp tôi biết được bề ngoài người (phụ nữ) này nhìn như thế nào: Hãy Luận Giải cung mệnh này theo bố cục:\n"+
			"1) Tổng quan nổi bật tính cách người này theo chính tinh\n"+
			"2) Tổng quan diện tướng người này\n"+
			"3) Kết luận tổng quan\n"+
			"Dữ liệu các cung của người %s mệnh này:\n"+
			"Các sao tại cung mệnh: %s\n"+
			"Các sao trong 2 cung còn lại tạo thành tam hợp: %s\n"+
			"Các sao cung đối diện chiếu về: %s.\n",
		g,
		placeDetail,
		three,
		opposite,
	)
}
