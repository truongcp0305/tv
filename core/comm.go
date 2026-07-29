package core

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"tuvi/models"
)

const PythonExecutable = `D:\dev\letcoode\.venv\Scripts\python.exe`
const ScriptCoreFile = `D:\dev\letcoode\backend\lap_dia_ban.py`
const ResultFileFromPython = `D:\dev\letcoode\data\lap_dia_ban_output.json`

func GenerateHoroscopePage(input models.InputData) error {
	inputPayload := map[string]any{
		"ngay":      input.Day,
		"thang":     input.Month,
		"nam":       input.Year,
		"gioSinh":   input.Hour,
		"gioiTinh":  input.Gender,
		"duongLich": input.IsSun,
		"timeZone":  7,
	}

	inputJSON, err := json.Marshal(inputPayload)
	if err != nil {
		return fmt.Errorf("marshal input data: %w", err)
	}

	cmd := exec.Command(
		PythonExecutable,
		ScriptCoreFile,
		"--input-data",
		string(inputJSON),
		"--output-file",
		ResultFileFromPython,
	)
	cmd.Dir = `D:\dev\letcoode`

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("run python script: %w; output: %s", err, string(output))
	}

	return nil
}

func ParseResultFile() (models.HoroscopePage, error) {
	result := models.HoroscopePage{}

	data, err := os.ReadFile(ResultFileFromPython)
	if err != nil {
		return result, fmt.Errorf("read result file: %w", err)
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("unmarshal result file: %w", err)
	}

	return result, nil
}
