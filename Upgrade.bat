@echo off
git pull && powershell -command "Stop-service -Force -name "RepostingRobot" -ErrorAction SilentlyContinue; go build; Start-service -name "RepostingRobot""
:: Hail Hydra
