package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/dustin/go-humanize"
)

const userAgent = "NugsNet/3.12.3.657 (Android; 7.1.2; samsung; SM-N976N)"

var (
	jar, _ = cookiejar.New(nil)
	client = &http.Client{Jar: jar}
)

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Downloaded += uint64(n)
	percentage := float64(wc.Downloaded) / float64(wc.Total) * float64(100)
	wc.Percentage = int(percentage)
	fmt.Printf("\r%d%%, %s/%s ", wc.Percentage, humanize.Bytes(wc.Downloaded), wc.TotalStr)
	return n, nil
}

func getScriptDir() (string, error) {
	var (
		ok    bool
		err   error
		fname string
	)
	if filepath.IsAbs(os.Args[0]) {
		_, fname, _, ok = runtime.Caller(0)
		if !ok {
			return "", errors.New("Failed to get script filename.")
		}
	} else {
		fname, err = os.Executable()
		if err != nil {
			return "", err
		}
	}
	scriptDir := filepath.Dir(fname)
	return scriptDir, nil
}

func readTxtFile(path string) ([]string, error) {
	var lines []string
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

func contains(lines []string, value string) bool {
	for _, line := range lines {
		if strings.EqualFold(line, value) {
			return true
		}
	}
	return false
}

func processUrls(urls []string) ([]string, error) {
	var (
		processed []string
		txtPaths  []string
	)
	for _, url := range urls {
		if strings.HasSuffix(url, ".txt") && !contains(txtPaths, url) {
			txtLines, err := readTxtFile(url)
			if err != nil {
				return nil, err
			}
			for _, txtLine := range txtLines {
				if !contains(processed, txtLine) {
					processed = append(processed, txtLine)
				}
			}
			txtPaths = append(txtPaths, url)
		} else {
			if !contains(processed, url) {
				processed = append(processed, url)
			}
		}
	}
	return processed, nil
}

func parseCfg() (*Config, error) {
	resolveFormat := map[int]int{
		1: 2,
		2: 3,
		3: 5,
		4: 8,
		5: 10,
	}
	cfg, err := readConfig()
	if err != nil {
		return nil, err
	}
	args := parseArgs()
	if args.Format != -1 {
		cfg.Format = args.Format
	}
	if !(cfg.Format >= 1 && cfg.Format <= 5) {
		return nil, errors.New("Format must be between 1 and 5.")
	}
	cfg.Format = resolveFormat[cfg.Format]
	if args.OutPath != "" {
		cfg.OutPath = args.OutPath
	}
	if cfg.OutPath == "" {
		cfg.OutPath = "Nugs downloads"
	}
	cfg.Urls, err = processUrls(args.Urls)
	if err != nil {
		errString := fmt.Sprintf("Failed to process URLs.\n%s", err)
		return nil, errors.New(errString)
	}
	return cfg, nil
}

func readConfig() (*Config, error) {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	var obj Config
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func parseArgs() *Args {
	var args Args
	arg.MustParse(&args)
	return &args
}

func makeDirs(path string) error {
	err := os.MkdirAll(path, 0755)
	return err
}

func fileExists(path string) (bool, error) {
	f, err := os.Stat(path)
	if err == nil {
		return !f.IsDir(), nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func sanitize(filename string) string {
	regex := regexp.MustCompile(`[\/:*?"><|]`)
	sanitized := regex.ReplaceAllString(filename, "_")
	return sanitized
}

func auth(email, pwd string) (string, error) {
	const _url = "https://id.nugs.net/connect/token"
	data := url.Values{}
	data.Set("client_id", "Eg7HuH873H65r5rt325UytR5429")
	data.Set("grant_type", "password")
	data.Set("scope", "openid profile email nugsnet:api nugsnet:legacyapi offline_access")
	data.Set("username", email)
	data.Set("password", pwd)
	req, err := http.NewRequest(http.MethodPost, _url, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	do, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer do.Body.Close()
	if do.StatusCode != http.StatusOK {
		return "", errors.New(do.Status)
	}
	var obj Auth
	err = json.NewDecoder(do.Body).Decode(&obj)
	if err != nil {
		return "", err
	}
	return obj.AccessToken, nil
}

func getUserInfo(token string) (string, error) {
	const _url = "https://id.nugs.net/connect/userinfo"
	req, err := http.NewRequest(http.MethodGet, _url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("User-Agent", userAgent)
	do, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer do.Body.Close()
	if do.StatusCode != http.StatusOK {
		return "", errors.New(do.Status)
	}
	var obj UserInfo
	err = json.NewDecoder(do.Body).Decode(&obj)
	if err != nil {
		return "", err
	}
	return obj.Sub, nil
}

func getSubInfo(token string) (*SubInfo, error) {
	const _url = "https://subscriptions.nugs.net/api/v1/me/subscriptions"
	req, err := http.NewRequest(http.MethodGet, _url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("User-Agent", userAgent)
	do, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer do.Body.Close()
	if do.StatusCode != http.StatusOK {
		return nil, errors.New(do.Status)
	}
	var obj SubInfo
	err = json.NewDecoder(do.Body).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func getPlan(subInfo *SubInfo) (string, bool) {
	if !reflect.ValueOf(subInfo.Plan).IsZero() {
		return subInfo.Plan.Description, false
	} else {
		return subInfo.Promo.Plan.Description, true
	}
}

func parseTimestamps(start, end string) (string, string) {
	const layout = "01/02/2006 15:04:05"
	startTime, _ := time.Parse(layout, start)
	endTime, _ := time.Parse(layout, end)
	parsedStart := strconv.FormatInt(startTime.Unix(), 10)
	parsedEnd := strconv.FormatInt(endTime.Unix(), 10)
	return parsedStart, parsedEnd
}

func parseStreamParams(userId string, subInfo *SubInfo, isPromo bool) *StreamParams {
	startStamp, endStamp := parseTimestamps(subInfo.StartedAt, subInfo.EndsAt)
	streamParams := &StreamParams{
		SubscriptionID:          subInfo.LegacySubscriptionID,
		SubCostplanIDAccessList: subInfo.Plan.ID,
		UserID:                  userId,
		StartStamp:              startStamp,
		EndStamp:                endStamp,
	}
	if isPromo {
		streamParams.SubCostplanIDAccessList = subInfo.Promo.Plan.ID
	} else {
		streamParams.SubCostplanIDAccessList = subInfo.Plan.ID
	}
	return streamParams
}

func checkUrl(url string) string {
	const regexString = `^https://play.nugs.net/#/catalog/recording/(\d+)$`
	regex := regexp.MustCompile(regexString)
	match := regex.FindStringSubmatch(url)
	if match == nil {
		return ""
	}
	return match[1]
}

func getAlbumMeta(albumId string) (*AlbumMeta, error) {
	const _url = "https://streamapi.nugs.net/api.aspx"
	req, err := http.NewRequest(http.MethodGet, _url, nil)
	if err != nil {
		return nil, err
	}
	query := url.Values{}
	query.Set("method", "catalog.container")
	query.Set("containerID", albumId)
	query.Set("vdisp", "1")
	req.URL.RawQuery = query.Encode()
	req.Header.Add("User-Agent", userAgent)
	do, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer do.Body.Close()
	if do.StatusCode != http.StatusOK {
		return nil, errors.New(do.Status)
	}
	var obj AlbumMeta
	err = json.NewDecoder(do.Body).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func getStreamMeta(trackId, format int, streamParams *StreamParams) (string, error) {
	const _url = "https://streamapi.nugs.net/bigriver/subPlayer.aspx"
	req, err := http.NewRequest(http.MethodGet, _url, nil)
	if err != nil {
		return "", err
	}
	query := url.Values{}
	query.Set("trackID", strconv.Itoa(trackId))
	query.Set("app", "1")
	query.Set("platformID", strconv.Itoa(format))
	query.Set("subscriptionID", streamParams.SubscriptionID)
	query.Set("subCostplanIDAccessList", streamParams.SubCostplanIDAccessList)
	query.Set("nn_userID", streamParams.UserID)
	query.Set("startDateStamp", streamParams.StartStamp)
	query.Set("endDateStamp", streamParams.EndStamp)
	req.URL.RawQuery = query.Encode()
	req.Header.Add("User-Agent", "nugsnetAndroid")
	do, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer do.Body.Close()
	if do.StatusCode != http.StatusOK {
		return "", errors.New(do.Status)
	}
	var obj StreamMeta
	err = json.NewDecoder(do.Body).Decode(&obj)
	if err != nil {
		return "", err
	}
	return obj.StreamLink, nil
}

func queryQuality(streamUrl string) *Quality {
	qualityMap := map[string]Quality{
		".alac16/": {Specs: "16-bit / 44.1 kHz ALAC", Extension: ".m4a"},
		".flac16/": {Specs: "16-bit / 44.1 kHz FLAC", Extension: ".flac"},
		".mqa24/":  {Specs: "24-bit / 48 kHz MQA", Extension: ".flac"},
		".s360/":   {Specs: "360 Reality Audio", Extension: ".mp4"},
		".aac150/": {Specs: "AAC 150", Extension: ".m4a"},
	}
	for k, v := range qualityMap {
		if strings.Contains(streamUrl, k) {
			return &v
		}
	}
	return nil
}

func downloadTrack(trackPath, url string) error {
	f, err := os.Create(trackPath)
	if err != nil {
		return err
	}
	defer f.Close()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Referer", "https://play.nugs.net/")
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Range", "bytes=0-")
	do, err := client.Do(req)
	if err != nil {
		return err
	}
	defer do.Body.Close()
	if do.StatusCode != http.StatusOK && do.StatusCode != http.StatusPartialContent {
		return errors.New(do.Status)
	}
	totalBytes := uint64(do.ContentLength)
	counter := &WriteCounter{Total: totalBytes, TotalStr: humanize.Bytes(totalBytes)}
	_, err = io.Copy(f, io.TeeReader(do.Body, counter))
	fmt.Println("")
	return err
}

func main() {
	fmt.Println(`
 _____                ____                _           _         
|   | |_ _ ___ ___   |    \ ___ _ _ _ ___| |___ ___ _| |___ ___ 
| | | | | | . |_ -|  |  |  | . | | | |   | | . | .'| . | -_|  _|
|_|___|___|_  |___|  |____/|___|_____|_|_|_|___|__,|___|___|_|  
	  |___|
`)
	scriptDir, err := getScriptDir()
	if err != nil {
		panic(err)
	}
	err = os.Chdir(scriptDir)
	if err != nil {
		panic(err)
	}
	cfg, err := parseCfg()
	if err != nil {
		errString := fmt.Sprintf("Failed to parse config file.\n%s", err)
		panic(errString)
	}
	err = makeDirs(cfg.OutPath)
	if err != nil {
		errString := fmt.Sprintf("Failed to make output folder.\n%s", err)
		panic(errString)
	}
	token, err := auth(cfg.Email, cfg.Password)
	if err != nil {
		errString := fmt.Sprintf("Failed to auth.\n%s", err)
		panic(errString)
	}
	userId, err := getUserInfo(token)
	if err != nil {
		errString := fmt.Sprintf("Failed to fetch user info.\n%s", err)
		panic(errString)
	}
	subInfo, err := getSubInfo(token)
	if err != nil {
		errString := fmt.Sprintf("Failed to fetch subcription info.\n%s", err)
		panic(errString)
	}
	if !subInfo.IsContentAccessible {
		panic("Account subscription required.")
	}
	planDesc, isPromo := getPlan(subInfo)
	fmt.Println(
		"Signed in successfully - " + planDesc + "\n",
	)
	streamParams := parseStreamParams(userId, subInfo, isPromo)
	albumTotal := len(cfg.Urls)
	for albumNum, url := range cfg.Urls {
		fmt.Printf("Album %d of %d:\n", albumNum+1, albumTotal)
		albumId := checkUrl(url)
		if albumId == "" {
			fmt.Println("Invalid URL:", url)
			continue
		}
		meta, err := getAlbumMeta(albumId)
		if err != nil {
			fmt.Printf("Failed to fetch album metadata.\n%s", err)
			continue
		}
		albumFolder := meta.Response.ArtistName + " - " + strings.TrimRight(meta.Response.ContainerInfo, " ")
		fmt.Println(albumFolder)
		if len(albumFolder) > 120 {
			albumFolder = albumFolder[:120]
			fmt.Println("Album folder name was chopped as it exceeds 120 characters.")
		}
		albumPath := filepath.Join(cfg.OutPath, sanitize(albumFolder))
		err = makeDirs(albumPath)
		if err != nil {
			fmt.Println("Failed to make album folder.\n", err)
			continue
		}
		trackTotal := len(meta.Response.Tracks)
		for trackNum, track := range meta.Response.Tracks {
			trackNum++
			streamUrl, err := getStreamMeta(track.TrackID, cfg.Format, streamParams)
			if err != nil {
				fmt.Printf("Failed to fetch track stream metadata.\n%s", err)
				continue
			} else if streamUrl == "" {
				fmt.Println("The API didn't return a track stream URL.")
				continue
			}
			quality := queryQuality(streamUrl)
			if quality == nil {
				fmt.Println("The API returned an unsupported format.")
				continue
			}
			trackFname := fmt.Sprintf(
				"%02d. %s%s", trackNum, sanitize(track.SongTitle), quality.Extension,
			)
			trackPath := filepath.Join(albumPath, trackFname)
			exists, err := fileExists(trackPath)
			if err != nil {
				fmt.Printf("Failed to check if track already exists locally.\n%s", err)
				continue
			}
			if exists {
				fmt.Println("Track already exists locally.")
				continue
			}
			fmt.Printf(
				"Downloading track %d of %d: %s - %s\n", trackNum, trackTotal, track.SongTitle,
				quality.Specs,
			)
			err = downloadTrack(trackPath, streamUrl)
			if err != nil {
				fmt.Printf("Failed to download track.\n%s", err)
			}
		}
	}
}
