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
if not exist CLIProxyAPI.exe (
    echo Membangun binary CLIProxyAPIPlus... proses ini hanya dilakukan sekali.
    go build -o CLIProxyAPI.exe cmd\server\main.go
)

echo Starting CLIProxyAPIPlus...
.\CLIProxyAPI.exe -config "%CONFIG%"

REM If you want to re-compile after making code changes, delete CLIProxyAPI.exe first
REM or run: go build -o CLIProxyAPI.exe cmd\server\main.go

endlocal
