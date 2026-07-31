package core

import (
	"fmt"

	"tuvi/models"
)

type HeavenlyStem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CanNamResult struct {
	ID     int    `json:"id"`
	TenCan string `json:"tenCan"`
}

type LuuNienCung struct {
	CungSo int    `json:"cungSo"`
	TenChi string `json:"tenChi"`
}

type AnCungLuuNienResult struct {
	NamXem      int         `json:"namXem"`
	NamSinh     int         `json:"namSinh"`
	TuoiAm      int         `json:"tuoiAm"`
	ChiNamSinh  int         `json:"chiNamSinh"`
	HuongDem    string      `json:"huongDem"`
	CungLuuNien LuuNienCung `json:"cungLuuNien"`
}

type TraLuuTuHoaResult struct {
	NamXem   int               `json:"namXem"`
	CanNam   CanNamResult      `json:"canNam"`
	LuuTuHoa map[string]string `json:"luuTuHoa"`
}

var HeavenlyStems = []HeavenlyStem{
	{ID: 1, Name: "Giáp"},
	{ID: 2, Name: "Ất"},
	{ID: 3, Name: "Bính"},
	{ID: 4, Name: "Đinh"},
	{ID: 5, Name: "Mậu"},
	{ID: 6, Name: "Kỷ"},
	{ID: 7, Name: "Canh"},
	{ID: 8, Name: "Tân"},
	{ID: 9, Name: "Nhâm"},
	{ID: 10, Name: "Quý"},
}

var DanStemMap = map[int]int{
	1:  3,
	6:  3,
	2:  5,
	7:  5,
	3:  7,
	8:  7,
	4:  9,
	9:  9,
	5:  1,
	10: 1,
}

var BranchOffset = map[int]int{
	3:  0,
	4:  1,
	5:  2,
	6:  3,
	7:  4,
	8:  5,
	9:  6,
	10: 7,
	11: 8,
	12: 9,
	1:  10,
	2:  11,
}

var MapThienCanGiap = map[int]int{
	4: 3,
	9: 7,
	5: 4,
	0: 9,
	6: 6,
	1: 10,
	7: 7,
	2: 12,
	8: 6,
	3: 1,
}

var LuuThienKhoc = map[int]int{
	1:  7,
	2:  8,
	3:  9,
	4:  10,
	5:  11,
	6:  12,
	7:  1,
	8:  2,
	9:  3,
	10: 4,
	11: 5,
	12: 6,
}

var LuuThienHu = map[int]int{
	1:  7,
	2:  6,
	3:  5,
	4:  4,
	5:  3,
	6:  2,
	7:  1,
	8:  12,
	9:  11,
	10: 10,
	11: 9,
	12: 8,
}

var LuuTuHoa = map[int]map[string]string{
	1: {
		"tenCan":     "Giáp",
		"Liêm Trinh": "Hóa lộc",
		"Phá Quân":   "Hóa quyền",
		"Vũ Khúc":    "Hóa khoa",
		"Thái Dương": "Hóa kỵ",
	},
	2: {
		"tenCan":      "Ất",
		"Thiên Cơ":    "Hóa lộc",
		"Thiên Lương": "Hóa quyền",
		"Tử Vi":       "Hóa khoa",
		"Thái Âm":     "Hóa kỵ",
	},
	3: {
		"tenCan":     "Bính",
		"Thiên Đồng": "Hóa lộc",
		"Thiên Cơ":   "Hóa quyền",
		"Văn Xương":  "Hóa khoa",
		"Liêm Trinh": "Hóa kỵ",
	},
	4: {
		"tenCan":     "Đinh",
		"Thái Âm":    "Hóa lộc",
		"Thiên Đồng": "Hóa quyền",
		"Thiên Cơ":   "Hóa khoa",
		"Cự Môn":     "Hóa kỵ",
	},
	5: {
		"tenCan":    "Mậu",
		"Tham Lang": "Hóa lộc",
		"Thái Âm":   "Hóa quyền",
		"Hữu Bật":   "Hóa khoa",
		"Thiên Cơ":  "Hóa kỵ",
	},
	6: {
		"tenCan":      "Kỷ",
		"Vũ Khúc":     "Hóa lộc",
		"Tham Lang":   "Hóa quyền",
		"Thiên Lương": "Hóa khoa",
		"Văn Khúc":    "Hóa kỵ",
	},
	7: {
		"tenCan":     "Canh",
		"Thái Dương": "Hóa lộc",
		"Vũ Khúc":    "Hóa quyền",
		"Thái Âm":    "Hóa khoa",
		"Thiên Đồng": "Hóa kỵ",
	},
	8: {
		"tenCan":     "Tân",
		"Cự Môn":     "Hóa lộc",
		"Thái Dương": "Hóa quyền",
		"Văn Khúc":   "Hóa khoa",
		"Văn Xương":  "Hóa kỵ",
	},
	9: {
		"tenCan":      "Nhâm",
		"Thiên Lương": "Hóa lộc",
		"Tử Vi":       "Hóa quyền",
		"Tả Phụ":      "Hóa khoa",
		"Vũ Khúc":     "Hóa kỵ",
	},
	10: {
		"tenCan":    "Quý",
		"Phá Quân":  "Hóa lộc",
		"Cự Môn":    "Hóa quyền",
		"Thái Âm":   "Hóa khoa",
		"Tham Lang": "Hóa kỵ",
	},
}

var CungOptions = map[int]string{
	1:  "Mệnh",
	2:  "Phụ mẫu",
	3:  "Phúc đức",
	4:  "Điền trạch",
	5:  "Quan lộc",
	6:  "Nô bộc",
	7:  "Thiên di",
	8:  "Tật ách",
	9:  "Tài bạch",
	10: "Tử tức",
	11: "Phu thê",
	12: "Huynh đệ",
}

func DichCung(cungSo, buocNhay int) int {
	return mod12(cungSo-1+buocNhay) + 1
}

func CanChiNamDuongLich(namXem int) (int, int) {
	canID := mod10(namXem+6) + 1
	chiID := mod12(namXem+8) + 1
	return canID, chiID
}

func TinhTuoiAm(namXem, namSinh int) (int, error) {
	if namXem < namSinh {
		return 0, fmt.Errorf("nam_xem phai lon hon hoac bang nam_sinh")
	}
	return namXem - namSinh + 1, nil
}

func AnCungLuuNien(namXem, gioiTinh int, namSinh *int, chiNamSinh any) (AnCungLuuNienResult, error) {
	result := AnCungLuuNienResult{}

	if gioiTinh != 1 && gioiTinh != -1 {
		return result, fmt.Errorf("gioi_tinh chi nhan 1 (Nam) hoac -1 (Nu)")
	}

	if namSinh == nil && chiNamSinh == nil {
		return result, fmt.Errorf("can cung cap nam_sinh hoac chi_nam_sinh")
	}

	chiNamSinhID, err := resolveChiNamSinh(namSinh, chiNamSinh)
	if err != nil {
		return result, err
	}

	if chiNamSinhID < 1 || chiNamSinhID > 12 {
		return result, fmt.Errorf("chi_nam_sinh phai trong khoang 1..12")
	}

	if namSinh == nil {
		return result, fmt.Errorf("can nam_sinh de tinh tuoi am")
	}

	tuoiAm, err := TinhTuoiAm(namXem, *namSinh)
	if err != nil {
		return result, err
	}

	huongDem := 1
	huongDemText := "thuan"
	if gioiTinh == -1 {
		huongDem = -1
		huongDemText = "nghich"
	}

	cungLuuNienID := DichCung(chiNamSinhID, (tuoiAm-1)*huongDem)
	result = AnCungLuuNienResult{
		NamXem:     namXem,
		NamSinh:    *namSinh,
		TuoiAm:     tuoiAm,
		ChiNamSinh: chiNamSinhID,
		HuongDem:   huongDemText,
		CungLuuNien: LuuNienCung{
			CungSo: cungLuuNienID,
			TenChi: CungOptions[cungLuuNienID],
		},
	}
	return result, nil
}

func TraLuuTuHoa(namXem int) (TraLuuTuHoaResult, error) {
	canID, _ := CanChiNamDuongLich(namXem)
	thongTin, ok := LuuTuHoa[canID]
	if !ok {
		return TraLuuTuHoaResult{}, fmt.Errorf("khong tim thay du lieu luu tu hoa cho can %d", canID)
	}

	return TraLuuTuHoaResult{
		NamXem: namXem,
		CanNam: CanNamResult{
			ID:     canID,
			TenCan: thongTin["tenCan"],
		},
		LuuTuHoa: thongTin,
	}, nil
}

func GetPalaceJia(yearStemID, branchID int) (HeavenlyStem, error) {
	stemAtDan, ok := DanStemMap[yearStemID]
	if !ok {
		return HeavenlyStem{}, fmt.Errorf("invalid year stem id: %d", yearStemID)
	}

	offset, ok := BranchOffset[branchID]
	if !ok {
		return HeavenlyStem{}, fmt.Errorf("invalid branch id: %d", branchID)
	}

	stemID := mod10(stemAtDan - 1 + offset)
	return HeavenlyStems[stemID], nil
}

func GetJias(canYearID int) ([]map[string]any, error) {
	result := make([]map[string]any, 0, 12)
	for i := 1; i <= 12; i++ {
		can, err := GetPalaceJia(canYearID, i)
		if err != nil {
			return nil, err
		}
		result = append(result, map[string]any{
			"cung_so": i,
			"can_id":  can.ID,
			"ten_can": can.Name,
		})
	}
	return result, nil
}

func AnThienKhoc(yearBranchID int) int {
	return mod12(6+(yearBranchID-1)) + 1
}

func AnThienHu(yearBranchID int) int {
	return mod12(6-(yearBranchID-1)) + 1
}

func AnLuuThienMa(branchID int) (int, error) {
	switch branchID {
	case 3, 7, 11:
		return 9, nil
	case 9, 1, 5:
		return 3, nil
	case 12, 4, 8:
		return 6, nil
	case 6, 10, 2:
		return 12, nil
	default:
		return 0, fmt.Errorf("invalid branch_id")
	}
}

func GenerateYearStars(namXem int) (map[int][]models.Star, error) {
	result := make(map[int][]models.Star)
	_, chiID := CanChiNamDuongLich(namXem)

	locTonID, ok := MapThienCanGiap[mod10(namXem)]
	if !ok {
		return nil, fmt.Errorf("khong tim thay loc ton cho nam %d", namXem)
	}

	kinhDuongID := DichCung(locTonID, 1)
	daLaID := DichCung(locTonID, -1)
	thaiTueID := mod12(namXem-1984) + 1
	tangMonID := DichCung(thaiTueID, 2)
	bachHoID := DichCung(tangMonID, 6)

	thienHuID, ok := LuuThienHu[chiID]
	if !ok {
		return nil, fmt.Errorf("khong tim thay luu thien hu cho chi %d", chiID)
	}

	thienKhocID, ok := LuuThienKhoc[chiID]
	if !ok {
		return nil, fmt.Errorf("khong tim thay luu thien khoc cho chi %d", chiID)
	}

	branchID := mod12(namXem-1984) + 1
	thienMaID, err := AnLuuThienMa(branchID)
	if err != nil {
		return nil, err
	}

	result[locTonID] = append(result[locTonID], models.NewLocTonStar())
	result[kinhDuongID] = append(result[kinhDuongID], models.NewKinhDuongStar())
	result[daLaID] = append(result[daLaID], models.NewDaLaStar())
	result[thaiTueID] = append(result[thaiTueID], models.NewThaiTueStar())
	result[tangMonID] = append(result[tangMonID], models.NewTangMonStar())
	result[bachHoID] = append(result[bachHoID], models.NewBachHoStar())
	result[thienHuID] = append(result[thienHuID], models.NewThienHuStar())
	result[thienKhocID] = append(result[thienKhocID], models.NewThienKhocStar())
	result[thienMaID] = append(result[thienMaID], models.NewThienMaStar())

	return result, nil
}

func resolveChiNamSinh(namSinh *int, chiNamSinh any) (int, error) {
	switch value := chiNamSinh.(type) {
	case nil:
		if namSinh == nil {
			return 0, fmt.Errorf("can cung cap nam_sinh hoac chi_nam_sinh")
		}
		_, chiID := CanChiNamDuongLich(*namSinh)
		return chiID, nil
	case int:
		return value, nil
	case int8:
		return int(value), nil
	case int16:
		return int(value), nil
	case int32:
		return int(value), nil
	case int64:
		return int(value), nil
	case uint:
		return int(value), nil
	case uint8:
		return int(value), nil
	case uint16:
		return int(value), nil
	case uint32:
		return int(value), nil
	case uint64:
		return int(value), nil
	case string:
		return 0, fmt.Errorf("chi_nam_sinh dang string chua duoc ho tro trong ban dich nay")
	default:
		return 0, fmt.Errorf("chi_nam_sinh kieu khong ho tro: %T", chiNamSinh)
	}
}

func mod12(value int) int {
	result := value % 12
	if result < 0 {
		result += 12
	}
	return result
}

func mod10(value int) int {
	result := value % 10
	if result < 0 {
		result += 10
	}
	return result
}

func BuildCungFromModel(page models.HoroscopePage) []map[string]any {
	result := make([]map[string]any, 0, len(page.TwelvePlaces))
	for _, place := range page.TwelvePlaces {
		stars := make([]map[string]any, 0, len(place.Stars))
		for _, star := range place.Stars {
			stars = append(stars, map[string]any{
				"saoID":         star.StarId,
				"saoTen":        star.Name,
				"saoLoai":       star.StarType,
				"saoDacTinh":    star.StarFeature,
				"saoNguHanh":    star.StarFE,
				"saoHuong":      star.StarDirection,
				"saoAmDuong":    int(star.YinYang),
				"vongTrangSinh": star.IsLiveCycle,
				"cssSao":        star.StarColor,
			})
		}

		result = append(result, map[string]any{
			"cungSo":      place.PlaceId,
			"hanhCung":    place.PlaceFE,
			"cungSao":     stars,
			"cungAmDuong": place.YinYang,
			"cungTen":     place.PlaceName,
			"cungThan":    place.IsBody,
			"cungDaiHan":  place.GreatPeriod,
			"cungTieuHan": place.YearPeriodName,
		})
	}
	return result
}
