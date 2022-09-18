package main

type Transport struct{}

type WriteCounter struct {
	Total      int64
	TotalStr   string
	Downloaded int64
	Percentage int
	StartTime  int64
}

type Config struct {
	Email       string
	Password    string
	Urls        []string
	Format      int
	OutPath     string
	VideoFormat int
	WantRes     string
	Token       string
}

type Args struct {
	Urls        []string `arg:"positional, required"`
	Format      int      `arg:"-f" default:"-1" help:"Track download format.\n\t\t\t 1 = 16-bit / 44.1 kHz ALAC\n\t\t\t 2 = 16-bit / 44.1 kHz FLAC\n\t\t\t 3 = 24-bit / 48 kHz MQA\n\t\t\t 4 = 360 Reality Audio / best available\n\t\t\t 5 = 150 Kbps AAC"`
	VideoFormat int      `arg:"-F" default:"-1" help:"Video download format.\n\t\t\t 1 = 480p\n\t\t\t 2 = 720p\n\t\t\t 3 = 1080p\n\t\t\t 4 = 1440p\n\t\t\t 5 = 4K / best available"`
	OutPath     string   `arg:"-o" help:"Where to download to. Path will be made if it doesn't already exist."`
}

type Auth struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type Payload struct {
	Nbf         int      `json:"nbf"`
	Exp         int      `json:"exp"`
	Iss         string   `json:"iss"`
	Aud         []string `json:"aud"`
	ClientID    string   `json:"client_id"`
	Sub         string   `json:"sub"`
	AuthTime    int      `json:"auth_time"`
	Idp         string   `json:"idp"`
	Email       string   `json:"email"`
	LegacyToken string   `json:"legacy_token"`
	LegacyUguid string   `json:"legacy_uguid"`
	Jti         string   `json:"jti"`
	Sid         string   `json:"sid"`
	Iat         int      `json:"iat"`
	Scope       []string `json:"scope"`
	Amr         []string `json:"amr"`
}

type UserInfo struct {
	Sub               string `json:"sub"`
	PreferredUsername string `json:"preferred_username"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	EmailVerified     bool   `json:"email_verified"`
}

type SubInfo struct {
	StripeMetaData struct {
		SubscriptionID      string      `json:"subscriptionId"`
		InvoiceID           string      `json:"invoiceId"`
		PaymentIntentStatus interface{} `json:"paymentIntentStatus"`
		ReturnURL           interface{} `json:"returnUrl"`
		RedirectURL         interface{} `json:"redirectUrl"`
		PaymentError        interface{} `json:"paymentError"`
	} `json:"stripeMetaData"`
	IsTrialAvailable        bool   `json:"isTrialAvailable"`
	AllowAddNewSubscription bool   `json:"allowAddNewSubscription"`
	ID                      string `json:"id"`
	LegacySubscriptionID    string `json:"legacySubscriptionId"`
	Status                  string `json:"status"`
	IsContentAccessible     bool   `json:"isContentAccessible"`
	StartedAt               string `json:"startedAt"`
	EndsAt                  string `json:"endsAt"`
	TrialEndsAt             string `json:"trialEndsAt"`
	Plan                    struct {
		ID              string      `json:"id"`
		Price           float64     `json:"price"`
		Period          int         `json:"period"`
		TrialPeriodDays int         `json:"trialPeriodDays"`
		PlanID          string      `json:"planId"`
		Description     string      `json:"description"`
		ServiceLevel    string      `json:"serviceLevel"`
		StartsAt        interface{} `json:"startsAt"`
		EndsAt          interface{} `json:"endsAt"`
	} `json:"plan"`
	Promo struct {
		ID            string      `json:"id"`
		PromoCode     string      `json:"promoCode"`
		PromoPrice    float64     `json:"promoPrice"`
		Description   string      `json:"description"`
		PromoStartsAt interface{} `json:"promoStartsAt"`
		PromoEndsAt   interface{} `json:"promoEndsAt"`
		Plan          struct {
			ID              string      `json:"id"`
			Price           float64     `json:"price"`
			Period          int         `json:"period"`
			TrialPeriodDays int         `json:"trialPeriodDays"`
			PlanID          string      `json:"planId"`
			Description     string      `json:"description"`
			ServiceLevel    string      `json:"serviceLevel"`
			StartsAt        interface{} `json:"startsAt"`
			EndsAt          interface{} `json:"endsAt"`
		} `json:"plan"`
		Gateway string `json:"gateway"`
	}
}

type StreamParams struct {
	SubscriptionID          string
	SubCostplanIDAccessList string
	UserID                  string
	StartStamp              string
	EndStamp                string
}

type Product struct {
	ProductStatusType    int           `json:"productStatusType"`
	SkuIDExt             interface{}   `json:"skuIDExt"`
	FormatStr            string        `json:"formatStr"`
	SkuID                int           `json:"skuID"`
	Cost                 int           `json:"cost"`
	CostplanID           int           `json:"costplanID"`
	Pricing              interface{}   `json:"pricing"`
	Bundles              []interface{} `json:"bundles"`
	NumPublicPricePoints int           `json:"numPublicPricePoints"`
	CartLink             string        `json:"cartLink"`
	LiveEventInfo        struct {
		IsEventLive                  bool        `json:"isEventLive"`
		EventID                      int         `json:"eventID"`
		EventStartDateStr            string      `json:"eventStartDateStr"`
		EventEndDateStr              string      `json:"eventEndDateStr"`
		TimeZoneToDisplay            interface{} `json:"timeZoneToDisplay"`
		OffsetFromLocalTimeToDisplay int         `json:"offsetFromLocalTimeToDisplay"`
		UTCoffset                    int         `json:"UTCoffset"`
		EventCode                    interface{} `json:"eventCode"`
		LinkType                     int         `json:"linkType"`
	} `json:"liveEventInfo"`
	SaleWindowInfo struct {
		IsEventSelling               bool        `json:"isEventSelling"`
		SswID                        int         `json:"sswID"`
		TimeZoneToDisplay            interface{} `json:"timeZoneToDisplay"`
		OffsetFromLocalTimeToDisplay int         `json:"offsetFromLocalTimeToDisplay"`
		SaleStartDateStr             interface{} `json:"saleStartDateStr"`
		SaleEndDateStr               interface{} `json:"saleEndDateStr"`
	} `json:"saleWindowInfo"`
	IosCost         int         `json:"iosCost"`
	IosPlanName     interface{} `json:"iosPlanName"`
	GooglePlanName  interface{} `json:"googlePlanName"`
	GoogleCost      int         `json:"googleCost"`
	NumDiscs        int         `json:"numDiscs"`
	IsSubStreamOnly int         `json:"isSubStreamOnly"`
}

type AlbArtResp struct {
	NumReviews                int         `json:"numReviews"`
	TotalContainerRunningTime int         `json:"totalContainerRunningTime"`
	HhmmssTotalRunningTime    string      `json:"hhmmssTotalRunningTime"`
	Products                  []Product   `json:"products"`
	Subscriptions             interface{} `json:"subscriptions"`
	Tracks                    []Track     `json:"tracks"`
	Pics                      []struct {
		PicID   int    `json:"picID"`
		OrderID int    `json:"orderID"`
		Height  int    `json:"height"`
		Width   int    `json:"width"`
		Caption string `json:"caption"`
		URL     string `json:"url"`
	} `json:"pics"`
	Recommendations []interface{} `json:"recommendations"`
	Reviews         struct {
		ContainerID int `json:"containerID"`
		Items       []struct {
			ReviewStatus    int    `json:"reviewStatus"`
			ReviewStatusStr string `json:"reviewStatusStr"`
			ContainerID     int    `json:"containerID"`
			ReviewID        int    `json:"reviewID"`
			ReviewerName    string `json:"reviewerName"`
			ReviewDate      string `json:"reviewDate"`
			Review          string `json:"review"`
		} `json:"items"`
		IsMoreRecords bool `json:"isMoreRecords"`
		TotalPages    int  `json:"totalPages"`
		TotalRecords  int  `json:"totalRecords"`
		NumPerPage    int  `json:"numPerPage"`
		PageNum       int  `json:"pageNum"`
	} `json:"reviews"`
	Notes []struct {
		NoteID int    `json:"noteID"`
		Note   string `json:"note"`
	} `json:"notes"`
	CategoryID       int         `json:"categoryID"`
	Labels           interface{} `json:"labels"`
	PrevContainerID  int         `json:"prevContainerID"`
	NextContainerID  int         `json:"nextContainerID"`
	PrevContainerURL string      `json:"prevContainerURL"`
	NextContainerURL string      `json:"nextContainerURL"`
	VolumeName       string      `json:"volumeName"`
	CdArtWorkList    []struct {
		DiscNumber     int    `json:"discNumber"`
		ArtWorkType    int    `json:"artWorkType"`
		ArtWorkTypeStr string `json:"artWorkTypeStr"`
		TemplateType   int    `json:"templateType"`
		ArtWorkPath    string `json:"artWorkPath"`
	} `json:"cdArtWorkList"`
	ContainerGroups         interface{}   `json:"containerGroups"`
	VideoURL                interface{}   `json:"videoURL"`
	VideoImage              interface{}   `json:"videoImage"`
	VideoTitle              interface{}   `json:"videoTitle"`
	VideoDesc               interface{}   `json:"videoDesc"`
	VodPlayerImage          string        `json:"vodPlayerImage"`
	IsInSubscriptionProgram bool          `json:"isInSubscriptionProgram"`
	SvodskuID               int           `json:"svodskuID"`
	LicensorName            string        `json:"licensorName"`
	AffID                   int           `json:"affID"`
	PageURL                 string        `json:"pageURL"`
	CoverImage              interface{}   `json:"coverImage"`
	VenueName               string        `json:"venueName"`
	VenueCity               string        `json:"venueCity"`
	VenueState              string        `json:"venueState"`
	ArtistName              string        `json:"artistName"`
	AccessList              []interface{} `json:"accessList"`
	AvailabilityType        int           `json:"availabilityType"`
	AvailabilityTypeStr     string        `json:"availabilityTypeStr"`
	Venue                   string        `json:"venue"`
	Img                     struct {
		PicID   int    `json:"picID"`
		OrderID int    `json:"orderID"`
		Height  int    `json:"height"`
		Width   int    `json:"width"`
		Caption string `json:"caption"`
		URL     string `json:"url"`
	} `json:"img"`
	ContainerID                   int         `json:"containerID"`
	ContainerInfo                 string      `json:"containerInfo"`
	PerformanceDate               string      `json:"performanceDate"`
	PerformanceDateFormatted      string      `json:"performanceDateFormatted"`
	PerformanceDateYear           string      `json:"performanceDateYear"`
	PerformanceDateShort          string      `json:"performanceDateShort"`
	PerformanceDateShortYearFirst string      `json:"performanceDateShortYearFirst"`
	PerformanceDateAbbr           string      `json:"performanceDateAbbr"`
	SongList                      interface{} `json:"songList"`
	ReleaseDate                   interface{} `json:"releaseDate"`
	ReleaseDateFormatted          string      `json:"releaseDateFormatted"`
	ActiveState                   string      `json:"activeState"`
	ContainerType                 int         `json:"containerType"`
	ContainerTypeStr              string      `json:"containerTypeStr"`
	Songs                         []Track     `json:"songs"`
	SalesLast30                   int         `json:"salesLast30"`
	SalesAllTime                  int         `json:"salesAllTime"`
	DateCreated                   string      `json:"dateCreated"`
	EpochDateCreated              float64     `json:"epochDateCreated"`
	ProductFormatList             []struct {
		PfType     int    `json:"pfType"`
		FormatStr  string `json:"formatStr"`
		SkuID      int    `json:"skuID"`
		Cost       int    `json:"cost"`
		CostplanID int    `json:"costplanID"`
		PfTypeStr  string `json:"pfTypeStr"`
		LiveEvent  struct {
			EventID                      int         `json:"eventID"`
			EventStartDateStr            interface{} `json:"eventStartDateStr"`
			EventEndDateStr              interface{} `json:"eventEndDateStr"`
			TimeZoneToDisplay            interface{} `json:"timeZoneToDisplay"`
			OffsetFromLocalTimeToDisplay int         `json:"offsetFromLocalTimeToDisplay"`
			UTCoffset                    int         `json:"UTCoffset"`
			EventCode                    interface{} `json:"eventCode"`
			LinkType                     int         `json:"linkType"`
		} `json:"liveEvent"`
		Salewindow struct {
			SswID                        int         `json:"sswID"`
			TimeZoneToDisplay            interface{} `json:"timeZoneToDisplay"`
			OffsetFromLocalTimeToDisplay int         `json:"offsetFromLocalTimeToDisplay"`
			SaleStartDateStr             interface{} `json:"saleStartDateStr"`
			SaleEndDateStr               interface{} `json:"saleEndDateStr"`
		} `json:"salewindow"`
		SkuCode         string `json:"skuCode"`
		IsSubStreamOnly int    `json:"isSubStreamOnly"`
	} `json:"productFormatList"`
	ContainsPreviewVideo  int           `json:"containsPreviewVideo"`
	ArtistID              int           `json:"artistID"`
	ContainerCategoryID   int           `json:"containerCategoryID"`
	ContainerCategoryName interface{}   `json:"containerCategoryName"`
	ContainerCode         string        `json:"containerCode"`
	ContainerIDExt        interface{}   `json:"containerIDExt"`
	ExtImage              string        `json:"extImage"`
	VideoChapters         []interface{} `json:"videoChapters"`
}

type AlbumMeta struct {
	MethodName                  string     `json:"methodName"`
	ResponseAvailabilityCode    int        `json:"responseAvailabilityCode"`
	ResponseAvailabilityCodeStr string     `json:"responseAvailabilityCodeStr"`
	Response                    AlbArtResp `json:"Response"`
}

type Token struct {
	MethodName string `json:"methodName"`
	Response   struct {
		TokenValue     string      `json:"tokenValue"`
		ReturnCode     int         `json:"returnCode"`
		ReturnCodeStr  string      `json:"returnCodeStr"`
		NnCustomerAuth interface{} `json:"nnCustomerAuth"`
	} `json:"Response"`
	ResponseAvailabilityCode    int    `json:"responseAvailabilityCode"`
	ResponseAvailabilityCodeStr string `json:"responseAvailabilityCodeStr"`
	SessionState                int    `json:"sessionState"`
	SessionStateStr             string `json:"sessionStateStr"`
}

type PlistMeta struct {
	MethodName string `json:"methodName"`
	Response   struct {
		TotalRunningTime       int    `json:"totalRunningTime"`
		HhmmssTotalRunningTime string `json:"hhmmssTotalRunningTime"`
		ID                     int    `json:"ID"`
		UserID                 int    `json:"userID"`
		Items                  []struct {
			ID                int   `json:"ID"`
			OrderID           int   `json:"orderID"`
			Track             Track `json:"track"`
			PlaylistContainer struct {
				TotalRunningTime       int         `json:"totalRunningTime"`
				HhmmssTotalRunningTime interface{} `json:"hhmmssTotalRunningTime"`
				Img                    struct {
					PicID   int    `json:"picID"`
					OrderID int    `json:"orderID"`
					Height  int    `json:"height"`
					Width   int    `json:"width"`
					Caption string `json:"caption"`
					URL     string `json:"url"`
				} `json:"img"`
				ContainerInfo          string      `json:"containerInfo"`
				Products               []Product   `json:"products"`
				VenueName              string      `json:"venueName"`
				VenueCity              string      `json:"venueCity"`
				VenueState             string      `json:"venueState"`
				ArtistName             string      `json:"artistName"`
				Venue                  string      `json:"venue"`
				ContainerID            int         `json:"containerID"`
				PerformanceDate        string      `json:"performanceDate"`
				ReleaseDate            interface{} `json:"releaseDate"`
				ContainerType          int         `json:"containerType"`
				ArtistID               int         `json:"artistID"`
				TitleType              int         `json:"titleType"`
				StrTotalRunningTime    string      `json:"strTotalRunningTime"`
				ContainerCategoryID    int         `json:"containerCategoryID"`
				ContainerCategoryName  interface{} `json:"containerCategoryName"`
				ContainerCategoryOrder int         `json:"containerCategoryOrder"`
				Availability           int         `json:"availability"`
				TicketImage            interface{} `json:"ticketImage"`
				UnavailableNote        interface{} `json:"unavailableNote"`
				Numasterisks           string      `json:"numasterisks"`
				CoverImage             interface{} `json:"coverImage"`
			} `json:"playlistContainer"`
		} `json:"items"`
		CreateDate          interface{} `json:"createDate"`
		PlayListName        string      `json:"playListName"`
		AlreadyExistsFlag   bool        `json:"alreadyExistsFlag"`
		PlayListUserInvalid bool        `json:"playListUserInvalid"`
		PlaylistImage       interface{} `json:"playlistImage"`
		NumTracks           int         `json:"numTracks"`
		GeneratedGUID       interface{} `json:"generatedGUID"`
		ShortenedLink       interface{} `json:"shortenedLink"`
	} `json:"Response"`
	ResponseAvailabilityCode    int    `json:"responseAvailabilityCode"`
	ResponseAvailabilityCodeStr string `json:"responseAvailabilityCodeStr"`
	SessionState                int    `json:"sessionState"`
	SessionStateStr             string `json:"sessionStateStr"`
}

type Track struct {
	AccessList             []interface{} `json:"accessList"`
	HhmmssTotalRunningTime string        `json:"hhmmssTotalRunningTime"`
	TrackLabel             string        `json:"trackLabel"`
	TrackURL               string        `json:"trackURL"`
	SongID                 int           `json:"songID"`
	SongTitle              string        `json:"songTitle"`
	TotalRunningTime       int           `json:"totalRunningTime"`
	DiscNum                int           `json:"discNum"`
	TrackNum               int           `json:"trackNum"`
	SetNum                 int           `json:"setNum"`
	ClipURL                string        `json:"clipURL"`
	TrackID                int           `json:"trackID"`
	TrackExclude           int           `json:"trackExclude"`
	Rootpath               interface{}   `json:"rootpath"`
	SourcePath             interface{}   `json:"sourcePath"`
	SourceFilename         interface{}   `json:"sourceFilename"`
	SourceFilePath         interface{}   `json:"sourceFilePath"`
	RootPathReal           interface{}   `json:"rootPathReal"`
	SourceFilePathReal     interface{}   `json:"sourceFilePathReal"`
	SkuIDExt               interface{}   `json:"skuIDExt"`
	TransportMethod        string        `json:"transportMethod"`
	StrTotalRunningTime    interface{}   `json:"strTotalRunningTime"`
	Products               []Product     `json:"products"`
	Subscriptions          interface{}   `json:"subscriptions"`
	AudioProduct           interface{}   `json:"audioProduct"`
	AudioLosslessProduct   interface{}   `json:"audioLosslessProduct"`
	AudioHDProduct         interface{}   `json:"audioHDProduct"`
	VideoProduct           interface{}   `json:"videoProduct"`
	LivestreamProduct      interface{}   `json:"livestreamProduct"`
	Mp4Product             interface{}   `json:"mp4Product"`
	VideoondemandProduct   interface{}   `json:"videoondemandProduct"`
	CdProduct              interface{}   `json:"cdProduct"`
	LiveHDstreamProduct    interface{}   `json:"liveHDstreamProduct"`
	HDvideoondemandProduct interface{}   `json:"HDvideoondemandProduct"`
	VinylProduct           interface{}   `json:"vinylProduct"`
	DsdProduct             interface{}   `json:"dsdProduct"`
	DvdProduct             interface{}   `json:"dvdProduct"`
	Reality360Product      interface{}   `json:"reality360Product"`
	ContainerGroups        interface{}   `json:"containerGroups"`
	IDList                 string        `json:"IDList"`
	PlayListID             int           `json:"playListID"`
	CatalogQueryType       int           `json:"catalogQueryType"`
}

type StreamMeta struct {
	StreamLink         string      `json:"streamLink"`
	Streamer           string      `json:"streamer"`
	UserID             string      `json:"userID"`
	Mason              interface{} `json:"mason"`
	SubContentAccess   int         `json:"subContentAccess"`
	StashContentAccess int         `json:"stashContentAccess"`
}

type Quality struct {
	Specs     string
	Extension string
	URL       string
	Format    int
}

type ArtistMeta struct {
	MethodName                  string `json:"methodName"`
	ResponseAvailabilityCode    int    `json:"responseAvailabilityCode"`
	ResponseAvailabilityCodeStr string `json:"responseAvailabilityCodeStr"`
	Response                    struct {
		HeaderName          interface{}  `json:"headerName"`
		Packages            interface{}  `json:"packages"`
		Containers          []AlbArtResp `json:"containers"`
		CategoryID          int          `json:"categoryID"`
		ArtistID            int          `json:"artistID"`
		ArtistName          interface{}  `json:"artistName"`
		LoadingState        int          `json:"loadingState"`
		TotalMatchedRecords int          `json:"totalMatchedRecords"`
		NnCheckSum          int          `json:"nnCheckSum"`
	} `json:"Response"`
}
