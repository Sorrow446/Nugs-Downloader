# Nugs-Downloader
Go port of my Nugs tool.
![](https://i.imgur.com/62HwfYd.png)
[Windows, Linux and macOS binaries](https://github.com/Sorrow446/Nugs-Downloader/releases)

# Setup
Input credentials into config file.
Configure any other options if needed.
|Option|Info|
| --- | --- |
|email|Email address.
|password|Password.
|format|Download quality. 1 = 16-bit / 44.1 kHz ALAC, 2 = 16-bit / 44.1 kHz FLAC, 3 = 24-bit / 48 kHz MQA (or next best), 4 =  360 Reality Audio (or next best). 5 = AAC 150.
|outPath|Where to download to. Path will be made if it doesn't already exist.

# Usage
Args take priority over the config file.

Download two albums:   
`nugs_dl_x64.exe https://play.nugs.net/#/catalog/recording/23329 https://play.nugs.net/#/catalog/recording/23790`

Download a single album and from two text files:   
`nugs_dl_x64.exe https://play.nugs.net/#/catalog/recording/23329 G:\1.txt G:\2.txt`

```
 _____                ____                _           _
|   | |_ _ ___ ___   |    \ ___ _ _ _ ___| |___ ___ _| |___ ___
| | | | | | . |_ -|  |  |  | . | | | |   | | . | .'| . | -_|  _|
|_|___|___|_  |___|  |____/|___|_____|_|_|_|___|__,|___|___|_|
          |___|

Usage: nugs_dl_x64.exe [--format FORMAT] [--outpath OUTPATH] URLS [URLS ...]

Positional arguments:
  URLS

Options:
  --format FORMAT, -f FORMAT [default: -1]
  --outpath OUTPATH, -o OUTPATH
  --help, -h             display this help and exit
  ```
 
# Disclaimer
- I will not be responsible for how you use Nugs Downloader.    
- Nugs brand and name is the registered trademark of its respective owner.    
- Nugs Downloader has no partnership, sponsorship or endorsement with Nugs.
