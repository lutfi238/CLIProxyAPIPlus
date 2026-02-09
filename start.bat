@echo off
setlocal ENABLEDELAYEDEXPANSION

REM === CLIProxyAPI starter ===
REM Uses config.yaml in project root by default.

set "ROOT=%~dp0"
cd /d "%ROOT%"

REM Link static assets for Go server
set "MANAGEMENT_STATIC_PATH=%ROOT%management-panel\dist"

REM Optional: override port or config path here
set "CONFIG=%ROOT%config.yaml"

REM Start server
echo Starting CLIProxyAPIPlus...
go run cmd\server\main.go -config "%CONFIG%"

REM If you prefer running built binary, comment out the line above and use:
REM .\CLIProxyAPI.exe -config "%CONFIG%"

endlocal
