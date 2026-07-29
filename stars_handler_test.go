package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"
	"tuvi/models"
)

const PyScript = `D:\dev\letcoode\backend\lap_dia_ban.py`
const ResultFile = `D:\dev\letcoode\data\lap_dia_ban_output.json`
const PythonExe = `D:\dev\letcoode\.venv\Scripts\python.exe`

func TestDecode(t *testing.T) {
	payload := map[string]any{
		"ngay":      1,
		"thang":     1,
		"nam":       2000,
		"gioSinh":   1,
		"gioiTinh":  1,
		"duongLich": true,
		"timeZone":  7,
	}
	payloadJSON, _ := json.MarshalIndent(payload, "", "  ")
	log.Printf("  🐍 Python payload: %s", string(payloadJSON))

	cmd := exec.Command(PythonExe, PyScript, "--input-data", string(payloadJSON), "--output-file", ResultFile)
	cmd.Dir = `D:\dev\letcoode`
	log.Printf("  ⏳ Running python...")
	executeStart := time.Now()
	_, err := cmd.CombinedOutput()
	log.Printf("  ✓ Python done in %.0fms", time.Since(executeStart).Seconds()*1000)
	if err != nil {
		log.Printf("  ✗ Python error: %v", err)
		return
	}

	// 5. Read result
	log.Printf("  🔍 Reading result: %s", ResultFile)
	if _, st := os.Stat(ResultFile); os.IsNotExist(st) {
		return
	}
	data, err := os.ReadFile(ResultFile)
	if err != nil {
		return
	}
	log.Printf("  📄 Result size: %d bytes", len(data))

	// 6. Parse into HoroscopePage
	var page models.HoroscopePage
	if err := json.Unmarshal(data, &page); err != nil {
		log.Printf("  ✗ Unmarshal: %v | first 200: %s", err, string(data))
		return
	}

	totalStars := 0
	for _, p := range page.TwelvePlaces {
		totalStars += len(p.Stars)
	}
	log.Printf("  ✅ Parsed: %d places, %d stars, body=%d, menh=%d",
		len(page.TwelvePlaces), totalStars, page.BodyPlace, page.DestinyPlace)

}
