package templates

import (
	"fmt"
	"strings"
	"tuvi/models"
)

// BuildChartJSON converts HoroscopePage data into a format suitable for frontend chart display.
func BuildChartJSON(page models.HoroscopePage) string {
	result := ""
	
	// Header
	result += "<div class='text-center mb-6'>"
	result += "<h2 class='text-2xl font-semibold text-slate-900'>Lá số Tử Vi</h2>"
	result += fmt.Sprintf("<p class='text-sm text-slate-500 mt-1'>Danh sách các cung và sao</p>")
	result += "</div>"
	
	// Palace information with better structure
	for i, place := range page.TwelvePlaces {
		if i > 0 {
			result += "<hr class='my-4 border-slate-200'>"
		}
		
		result += fmt.Sprintf("<div class='mb-6 p-3 rounded-lg bg-slate-50/50 border border-slate-100'>")
		result += fmt.Sprintf("<h3 class='font-medium text-indigo-700 mb-2'>%s</h3>", place.PlaceName)
		
		// Stars section
		if len(place.Stars) > 0 {
			result += "<div class='space-y-1 ml-3'>"
			for _, star := range place.Stars {
				colorClass := getStarColorClass(star.StarColor)
				result += fmt.Sprintf("<div class='%s'>• %s</div>", colorClass, star.Name)
			}
			result += "</div>"
		} else {
			result += "<div class='text-slate-400 italic text-sm'>Vô chính diệu</div>"
		}
		
		// Special conditions
		conditions := []string{}
		if place.IsIntercept {
			conditions = append(conditions, "Triệt lộ")
		}
		if place.IsBlockade {
			conditions = append(conditions, "Tuần phá")
		}
		if len(conditions) > 0 {
			result += fmt.Sprintf("<div class='mt-2 text-xs font-medium text-amber-600'>⚠️ %s</div>", strings.Join(conditions, ", "))
		}
		
		result += "</div>"
	}
	
	return result
}

func getStarColorClass(color string) string {
	switch strings.ToLower(color) {
	case "red", "hung", "hong":
		return "text-red-600 font-medium"
	case "blue", "lam":
		return "text-blue-600 font-medium"
	case "green", "xanh":
		return "text-green-600 font-medium"
	default:
		return "text-slate-700"
	}
}
