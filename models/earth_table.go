package models

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Thông tin lá số
type HoroscopePage struct {
	LunaMonth int `json:"thangSinhAmLich"`
	LunaHour  int `json:"gioSinhAmLich"`
	// Thông tin 12 cung
	TwelvePlaces []Place `json:"thapNhiCung"`
	// Cung thân
	BodyPlace int `json:"cungThan"`
	// Cung mệnh
	DestinyPlace int `json:"cungMenh"`
	// Cung nô bộc
	FriendPlace int `json:"cungNoboc"`
	// Cung tật ách
	ObstaclePlace int `json:"cungTatAch"`
}

// Thông tin cung
type Place struct {
	PlaceId int `json:"cungSo"`
	// Ngũ hành cung
	PlaceFE   string `json:"hanhCung"`
	Stars     []Star `json:"cungSao"`
	YinYang   int    `json:"cungAmDuong"`
	PlaceName string `json:"cungTen"`
	MainPlace string `json:"cungChu"`
	IsBody    bool   `json:"cungThan"`
	// Cung đại hạn năm
	GreatPeriod int `json:"cungDaiHan"`
	// Tên chi cung tiểu hạn
	YearPeriodName string `json:"cungTieuHan"`
	IsIntercept    bool   `json:"trietLo"`
	IsBlockade     bool   `json:"tuanTrung"`
}

type Star struct {
	StarId        int         `json:"saoID"`
	Name          string      `json:"saoTen"`
	StarFE        string      `json:"saoNguHanh"`
	StarType      int         `json:"saoLoai"`
	StarDirection string      `json:"saoHuong"`
	YinYang       FlexibleInt `json:"saoAmDuong"`
	IsLiveCycle   int         `json:"vongTrangSinh"`
	StarColor     string      `json:"cssSao"`
	StarFeature   string      `json:"saoDacTinh"`
}

type FlexibleInt int

func (f *FlexibleInt) UnmarshalJSON(data []byte) error {
	trimmed := strings.TrimSpace(string(data))
	if trimmed == "" || trimmed == "null" {
		*f = 0
		return nil
	}

	if trimmed[0] == '"' {
		var text string
		if err := json.Unmarshal(data, &text); err != nil {
			return err
		}
		if text == "" {
			*f = 0
			return nil
		}
		value, err := strconv.Atoi(text)
		if err != nil {
			return err
		}
		*f = FlexibleInt(value)
		return nil
	}

	var value int
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FlexibleInt(value)
	return nil
}

type YearStar struct {
	Name    string `json:"name"`
	PlaceId int    `json:"PlaceId"`
}
