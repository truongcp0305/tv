# =============================================================
# Test script for POST /api/stars - tử vi server
# PowerShell version - native, no external deps
# =============================================================

$baseUri = "http://localhost:4400"
$logFile = Join-Path $PSScriptRoot "log\test-stars_$(Get-Date -Format 'yyyyMMdd_HHmmss').log"

# Ensure log directory exists
New-Item -ItemType Directory -Force -Path (Split-Path $logFile -Parent) | Out-Null

function Write-Log {
    param([string]$msg, [string]$level = "INFO")
    $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    $entry = "[$timestamp] [$level] $msg"
    Add-Content -Path $logFile -Value $entry
    switch ($level) {
        "ERROR" { Write-Host $entry -ForegroundColor Red }
        "WARN"  { Write-Host $entry -ForegroundColor Yellow }
        "SUCCESS" { Write-Host $entry -ForegroundColor Green }
        default { Write-Host $entry }
    }
}

# ---- 1. Health check ----
Write-Log "=== Step 1: Health check ==="
try {
    $resp = Invoke-WebRequest -Uri "$baseUri/health" -Method Get -ErrorAction Stop
    $body = $resp.Content | ConvertFrom-Json -ErrorAction Stop
    if ($body.status -eq "ok") {
        Write-Log "Health check: OK" "SUCCESS"
    } else {
        Write-Log "Health check: failed (status != ok)" "ERROR"
    }
} catch {
    Write-Log "Health check: failed - $_" "ERROR"
    exit 1
}

# ---- Helper: send POST to /api/stars ----
function Test-StarsEndpoint {
    param(
        [hashtable] $payload
    )
    try {
        $content = $payload | ConvertTo-Json -ErrorAction Stop
        $resp = Invoke-WebRequest -Uri "$baseUri/api/stars" -Method Post -ContentType "application/json" -Body $content -ErrorAction Stop
        return $null, $resp.Content, $resp.StatusCode
    } catch {
        return $_.Exception.Message, $null, $null
    }
}

# ---- 2. Test case 1: Dương lịch ----
Write-Log "=== Step 2: POST /api/stars (dương lịch) ==="
$payload1 = @{
    "full_name" = "Test User 1"
    "day"       = 15
    "month"     = 8
    "year"      = 1990
    "hour"      = 14
    "gender"    = 1
    "is_sun"    = $true
    "calendar_mode" = ""
}
$errorMsg, $body, $code = Test-StarsEndpoint -payload $payload1
if (-not $errorMsg -and $code -eq 200) {
    Write-Log "Test 1 PASSED - status 200, HTML returned" "SUCCESS"
    # Show first 300 chars of HTML
    Write-Log "HTML (first 300 chars): $($body.Substring(0, [Math]::Min(300, $body.Length)))"
} else {
    Write-Log "Test 1 FAILED: $errorMsg (status: $code)" "ERROR"
}

# ---- 3. Test case 2: Âm lịch ----
Write-Log "=== Step 3: POST /api/stars (âm lịch) ==="
$payload2 = @{
    "full_name"       = "Test User 2"
    "day"             = 10
    "month"           = 1
    "year"            = 1990
    "hour"            = 9
    "gender"          = -1
    "is_sun"          = $false
    "calendar_mode"   = "lunar"
}
$errorMsg, $body, $code = Test-StarsEndpoint -payload $payload2
if (-not $errorMsg -and $code -eq 200) {
    Write-Log "Test 2 PASSED - status 200" "SUCCESS"
} else {
    Write-Log "Test 2 FAILED: $errorMsg (status: $code)" "ERROR"
}

# ---- 4. Test case 3: Year out of range (validation) ----
Write-Log "=== Step 4: Validation - year out of range ==="
$payload3 = @{
    "full_name" = "Bad Year"
    "day"       = 1
    "month"     = 1
    "year"      = 1800   # invalid
    "hour"      = 0
    "gender"    = 0
    "is_sun"    = $true
}
$errorMsg, $body, $code = Test-StarsEndpoint -payload $payload3
if ($code -ne 200) {
    Write-Log "Test 3 PASSED - correctly rejected (status $code)" "SUCCESS"
    if ($body) {
        Write-Log "Error response: $body"
    }
} else {
    Write-Log "Test 3 FAILED - should have been rejected" "ERROR"
}

# ---- 5. Test case 4: Day out of range ----
Write-Log "=== Step 5: Validation - day out of range ==="
$payload4 = @{
    "full_name" = "Bad Day"
    "day"       = 99   # invalid
    "month"     = 8
    "year"      = 2000
    "hour"      = 12
    "gender"    = 1
    "is_sun"    = $true
}
$errorMsg, $body, $code = Test-StarsEndpoint -payload $payload4
if ($code -ne 200) {
    Write-Log "Test 4 PASSED - correctly rejected (status $code)" "SUCCESS"
} else {
    Write-Log "Test 4 FAILED - should have been rejected" "ERROR"
}

Write-Log "=== All tests completed. Log: $logFile ===" "SUCCESS"