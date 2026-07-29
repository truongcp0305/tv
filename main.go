package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"tuvi/models"
)

const (
	indexHTMLPath = "static/index.html"
	pythonExe     = `D:\dev\letcoode\.venv\Scripts\python.exe`
	pyScript      = `D:\dev\letcoode\backend\lap_dia_ban.py`
	resultFile    = `D:\dev\letcoode\data\lap_dia_ban_output.json`
)

// --- Request / Response ---

type StarRequest struct {
	Day      int    `json:"day"`
	Month    int    `json:"month"`
	Year     int    `json:"year"`
	Hour     int    `json:"hour"`
	Gender   int    `json:"gender"`
	IsSun    string `json:"is_sun"` // receive as string, convert below
	Calendar string `json:"calendar_mode"`
	FullName string `json:"full_name"`
}

// --- Middleware ---

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		path := r.URL.Path
		log.Printf("  ▶ %s %s [%s]", r.Method, path, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("  ◀ %s %s → %.0fms", r.Method, path, time.Since(start).Seconds()*1000)
	})
}

func starsHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Parse
	var req StarRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("  ✗ JSON decode: %v", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"success": false, "error": err.Error()})
		return
	}

	// Convert IsSun from string to bool
	isSun := false
	if strings.ToLower(req.IsSun) == "true" || req.IsSun == "1" {
		isSun = true
	}

	// 2. Validate
	if req.Year < 1900 || req.Year > 2100 {
		log.Printf("  ✗ year out of range: %d", req.Year)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"success": false, "error": "year must be 1900-2100"})
		return
	}
	if req.Day < 1 || req.Day > 31 {
		log.Printf("  ✗ day out of range: %d", req.Day)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"success": false, "error": "day must be 1-31"})
		return
	}
	if req.Month < 1 || req.Month > 12 {
		log.Printf("  ✗ month out of range: %d", req.Month)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"success": false, "error": "month must be 1-12"})
		return
	}
	if req.Hour < 0 || req.Hour > 23 {
		log.Printf("  ✗ hour out of range: %d", req.Hour)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"success": false, "error": "hour must be 0-23"})
		return
	}

	log.Printf("  📥 Input: name=%q day=%d month=%d year=%d hour=%d gender=%d solar=%v",
		req.FullName, req.Day, req.Month, req.Year, req.Hour, req.Gender, isSun)

	// 3. Build payload
	payload := map[string]any{
		"ngay":      req.Day,
		"thang":     req.Month,
		"nam":       req.Year,
		"gioSinh":   req.Hour,
		"gioiTinh":  req.Gender,
		"duongLich": isSun,
		"timeZone":  7,
	}
	payloadJSON, _ := json.MarshalIndent(payload, "", "  ")
	log.Printf("  🐍 Python payload: %s", string(payloadJSON))

	// 4. Execute
	cmd := exec.Command(pythonExe, pyScript, "--input-data", string(payloadJSON), "--output-file", resultFile)
	cmd.Dir = `D:\dev\letcoode`

	log.Printf("  ⏳ Running python...")
	executeStart := time.Now()
	output, err := cmd.CombinedOutput()
	log.Printf("  ✓ Python done in %.0fms", time.Since(executeStart).Seconds()*1000)

	if err != nil {
		log.Printf("  ✗ Python error: %v", err)
		for _, line := range splitLines(string(output)) {
			log.Printf("     %s", line)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"error":   "Python failed: " + err.Error(),
		})
		return
	}

	// 5. Read result
	log.Printf("  🔍 Reading result: %s", resultFile)
	if _, st := os.Stat(resultFile); os.IsNotExist(st) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"success": false, "error": "result file not found"})
		return
	}
	data, err := os.ReadFile(resultFile)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"success": false, "error": err.Error()})
		return
	}
	log.Printf("  📄 Result size: %d bytes", len(data))

	// 6. Parse into HoroscopePage
	var page models.HoroscopePage
	if err := json.Unmarshal(data, &page); err != nil {
		log.Printf("  ✗ Unmarshal: %v | first 200: %s", err, truncate(string(data), 200))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"success": false, "error": "parse error: " + err.Error()})
		return
	}

	totalStars := 0
	for _, p := range page.TwelvePlaces {
		totalStars += len(p.Stars)
	}
	log.Printf("  ✅ Parsed: %d places, %d stars, body=%d, menh=%d",
		len(page.TwelvePlaces), totalStars, page.BodyPlace, page.DestinyPlace)

	// 7. Render HTML
	var buf strings.Builder
	if err := renderChart(&buf, &page, req); err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(500)
		buf.WriteString(`<div class="p-4 bg-red-100 text-red-700">Template error: ` + escape(err.Error()) + `</div>`)
		log.Printf("  ✗ Template error: %v", err)
	} else {
		log.Printf("  📤 Sending %d bytes HTML", buf.Len())
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(buf.String()))
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	data, err := os.ReadFile(indexHTMLPath)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// --- HTML Rendering ---

func renderChart(buf *strings.Builder, page *models.HoroscopePage, req StarRequest) error {
	buf.WriteString(`<div class="chart-grid w-full overflow-x-auto" style="min-width: min(100%, 600px)"><div class="grid grid-cols-4 gap-[1px] bg-slate-300 border border-slate-300 rounded-lg overflow-hidden">`)

	places := page.TwelvePlaces

	// Row 1: 0-3
	renderPlace(buf, places, 0)
	renderPlace(buf, places, 1)
	renderPlace(buf, places, 2)
	renderPlace(buf, places, 3)

	// Row 2: 4, center, 5
	renderPlace(buf, places, 4)
	renderCenter(buf, page, req)
	renderPlace(buf, places, 5)

	// Row 3: 6, 7
	renderPlace(buf, places, 6)
	renderPlace(buf, places, 7)

	// Row 4: 8-11
	renderPlace(buf, places, 8)
	renderPlace(buf, places, 9)
	renderPlace(buf, places, 10)
	renderPlace(buf, places, 11)

	buf.WriteString(`</div>`) 
	return nil
}

func renderPlace(buf *strings.Builder, places []models.Place, idx int) {
	sz := "min-h-[120px] xl:min-h-[160px]"
	if idx >= len(places) {
		buf.WriteString(`<div class="bg-white ` + sz + ` border border-slate-300 flex flex-col opacity-30"><div class="mx-2 mt-2 h-3 bg-slate-200 rounded w-12"></div></div>`)
		return
	}

	p := places[idx]
	buf.WriteString(`<div class="bg-white ` + sz + ` border border-slate-300 transition-all duration-200 hover:shadow-md hover:-translate-y-0.5 hover:border-indigo-400 flex flex-col">`)

	// Header: place name
	buf.WriteString(`<div class="px-2 py-1 border-b border-slate-200 flex items-center justify-between shrink-0">`)
	buf.WriteString(`<span class="text-sm font-medium text-slate-900 truncate">` + escape(p.PlaceName) + `</span>`)
	buf.WriteString(`<span class="text-xs text-slate-400 shrink-0">` + escape(p.MainPlace) + `</span>`)
	buf.WriteString(`</div>`)

	// Stars - split into 2 columns: left for StarType < 10, right for StarType >= 10
	buf.WriteString(`<div class="flex-1 px-2 py-1 overflow-hidden">`)
	buf.WriteString(`<div class="grid grid-cols-2 gap-2 h-full">`)
	
	// Left column: StarType < 10
	buf.WriteString(`<div class="space-y-0.5 overflow-y-auto">`)
	mainCount := 0
	for _, s := range p.Stars {
		isMain := s.StarType == 1
		if isMain {
			mainCount++
		}
		if s.StarType < 10 {
			color := starColor(s.StarColor)
			cls := color
			if isMain {
				cls = "font-semibold text-sm " + cls
			} else {
				cls = "text-sm opacity-70 " + cls
			}
			buf.WriteString(`<div class="` + cls + ` leading-tight">` + escape(s.Name) + `</div>`)
		}
	}
	if mainCount == 0 {
		buf.WriteString(`<div class="text-xs text-slate-400 italic text-sm">Vô chính diệu</div>`)
	}
	buf.WriteString(`</div>`)
	
	// Right column: StarType >= 10
	buf.WriteString(`<div class="space-y-0.5 overflow-y-auto">`)
	for _, s := range p.Stars {
		isMain := s.StarType == 1
		if isMain {
			mainCount++
		}
		if s.StarType >= 10 {
			color := starColor(s.StarColor)
			cls := color
			if isMain {
				cls = "font-semibold text-sm " + cls
			} else {
				cls = "text-sm opacity-70 " + cls
			}
			buf.WriteString(`<div class="` + cls + ` leading-tight">` + escape(s.Name) + `</div>`)
		}
	}
	if mainCount == 0 {
		buf.WriteString(`<div class="text-xs text-slate-400 italic text-sm">Vô chính diệu</div>`)
	}
	buf.WriteString(`</div>`)
	buf.WriteString(`</div>`)
	// Footer: Great period
	if p.GreatPeriod > 0 {
		buf.WriteString(`<div class="px-2 py-1 border-t border-slate-200 text-xs text-slate-500 shrink-0">Đại hạn: ` + fmt.Sprintf("%d", p.GreatPeriod) + `</div>`)
	}
	if p.YearPeriodName != "" {
		buf.WriteString(`<div class="px-2 py-1 border-t border-slate-200 text-xs text-slate-400 shrink-0">Tiểu hạn: ` + escape(p.YearPeriodName) + `</div>`)
	}

	buf.WriteString(`</div>`)
}

func renderCenter(buf *strings.Builder, page *models.HoroscopePage, req StarRequest) {
	buf.WriteString(`<div class="col-span-2 row-span-2 flex flex-col items-center justify-center p-4 text-slate-500 select-none bg-white min-h-[140px] xl:min-h-[200px]">`)
	buf.WriteString(`<div class="space-y-2 w-full max-w-[180px] text-center">`)
	buf.WriteString(`<span class="text-5xl opacity-15 block mb-1">☯</span>`)

	name := escape(req.FullName)
	if name == "" {
		name = "Khách"
	}

	items := [][2]string{
		{"Name", name},
		{"Gender", genderLabel(req.Gender)},
		{"Solar", fmt.Sprintf("%d/%02d/%02d", req.Year, req.Month, req.Day)},
		{"Lunar", "--"},
		{"Element", "--"},
		{"Destiny", "--"},
		{"Age", "--"},
	}
	labels := []string{"Name", "Gender", "Solar", "Lunar", "Element", "Destiny", "Age"}
	for i, item := range items {
		buf.WriteString(`<div class="text-xs"><span class="text-slate-400 block text-[10px]">` + labels[i] + `</span> ` + escape(item[1]) + `</div>`)
	}
	buf.WriteString(`</div></div>`)
}

func starColor(color string) string {
	c := strings.ToLower(strings.TrimSpace(color))
	switch c {
	case "hung", "red":
		return "text-red-600"
	case "lam", "blue":
		return "text-blue-600"
	case "xanh", "green":
		return "text-green-600"
	case "hoang", "tu", "yellow":
		return "text-yellow-600"
	case "kim", "white":
		return "text-gray-400"
	default:
		return "text-slate-700"
	}
}

func genderLabel(g int) string {
	switch g {
	case 1:
		return "Nam"
	case -1:
		return "Nữ"
	default:
		return "?"
	}
}

func escape(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, `"`, "&#34;")
	return s
}

// --- CORS ---

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// --- Main ---

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4400"
	}

	os.MkdirAll(filepath.Dir(resultFile), 0755)

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/stars", starsHandler)
	mux.HandleFunc("/health", healthHandler)

	handler := loggingMiddleware(corsMiddleware(mux))

	fmt.Println("\n🚀 Tử Vi Server")
	fmt.Printf("   Frontend: http://localhost:%s\n", port)
	fmt.Printf("   API:      POST /api/stars\n")
	fmt.Printf("   Health:   GET  /health\n")
	fmt.Printf("   Log:      server.log\n")
	fmt.Println()

	// Log to file
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetPrefix("")
	log.SetFlags(0)

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

// --- Helpers ---

func splitLines(s string) []string {
	var lines []string
	for i, l := range strings.Split(s, "\n") {
		if l != "" && i < 10 {
			lines = append(lines, l)
		}
	}
	return lines
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

