# Nugs-Downloader
Nugs downloader written in Go.
![](https://i.imgur.com/NOsQjnP.png)
![](https://i.imgur.com/BEudufy.png)
[Windows, Linux, macOS, and Android binaries](https://github.com/Sorrow446/Nugs-Downloader/releases)

# Setup
Input credentials into config file.
Configure any other options if needed.
|Option|Info|
| --- | --- |
|email|Email address.
|password|Password.
|format|Track download quality. 1 = 16-bit / 44.1 kHz ALAC, 2 = 16-bit / 44.1 kHz FLAC, 3 = 24-bit / 48 kHz MQA, 4 = 360 Reality Audio / best available, 5 = 150 Kbps AAC.
|videoFormat|Video download format. 1 = 480p, 2 = 720p, 3 = 1080p, 4 = 1440p, 5 = 4K / best available. **FFmpeg needed, see below.**
|outPath|Where to download to. Path will be made if it doesn't already exist.

**FFmpeg is needed for ts -> MP4 losslessly for videos.**  
[Windows (gpl)](https://github.com/BtbN/FFmpeg-Builds/releases)    
Linux: `sudo apt install ffmpeg`    
Termux `pkg install ffmpeg`

## Supported Media
|Type|URL example|
| --- | --- |
|Album|`https://play.nugs.net/#/catalog/recording/23329`
|Catalog playlist|`https://2nu.gs/3PmqXLW`
|User playlist|`https://play.nugs.net/#/playlists/playlist/1215400`
|Video|`https://play.nugs.net/#/videos/artist/1437/Bachman%20&%20Turner/container/26695`

# Usage
Args take priority over the config file.

Download two albums:   
`nugs_dl_x64.exe https://play.nugs.net/#/catalog/recording/23329 https://play.nugs.net/#/catalog/recording/23790`

Download a single album and from two text files:   
`nugs_dl_x64.exe https://play.nugs.net/#/catalog/recording/23329 G:\1.txt G:\2.txt`

Download a catalog playlist and video:
`nugs_dl_x64.exe https://play.nugs.net/#/playlists/playlist/1215400 https://play.nugs.net/#/videos/artist/1437/Bachman%20&%20Turner/container/26695`

```
 _____                ____                _           _
|   | |_ _ ___ ___   |    \ ___ _ _ _ ___| |___ ___ _| |___ ___
| | | | | | . |_ -|  |  |  | . | | | |   | | . | .'| . | -_|  _|
|_|___|___|_  |___|  |____/|___|_____|_|_|_|___|__,|___|___|_|
          |___|

Usage: nugs_dl_x64.exe [--format FORMAT] [--videoformat VIDEOFORMAT] [--outpath OUTPATH] URLS [URLS ...]

Positional arguments:
  URLS

Options:
  --format FORMAT, -f FORMAT
                         Track download format.
                         1 = 16-bit / 44.1 kHz ALAC
                         2 = 16-bit / 44.1 kHz FLAC
                         3 = 24-bit / 48 kHz MQA
                         4 = 360 Reality Audio / best available
                         5 = 150 Kbps AAC [default: -1]
  --videoformat VIDEOFORMAT, -F VIDEOFORMAT
                         Video download format.
                         1 = 480p
                         2 = 720p
                         3 = 1080p
                         4 = 1440p
                         5 = 4K / best available [default: -1]
  --outpath OUTPATH, -o OUTPATH
                         Where to download to. Path will be made if it doesn't already exist.
  --help, -h             display this help and exit
  ```
 
# Disclaimer
- I will not be responsible for how you use Nugs Downloader.    
- Nugs brand and name is the registered trademark of its respective owner.    
- Nugs Downloader has no partnership, sponsorship or endorsement with Nugs.
