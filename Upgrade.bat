@echo off
git pull && powershell -command "Stop-service -Force -name "Punk05Robot" -ErrorAction SilentlyContinue; go build; Start-service -name "Punk05Robot""
:: Hail Hydra
