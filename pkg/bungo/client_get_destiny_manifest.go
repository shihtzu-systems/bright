package bungo

type GetDestinyManifest struct {
	Response struct {
		Version                  string `json:"version"`
		MobileAssetContentPath   string `json:"mobileAssetContentPath"`
		MobileGearAssetDataBases []struct {
			Version int    `json:"version"`
			Path    string `json:"path"`
		} `json:"mobileGearAssetDataBases"`
		MobileWorldContentPaths struct {
			En    string `json:"en"`
			Fr    string `json:"fr"`
			Es    string `json:"es"`
			EsMx  string `json:"es-mx"`
			De    string `json:"de"`
			It    string `json:"it"`
			Ja    string `json:"ja"`
			PtBr  string `json:"pt-br"`
			Ru    string `json:"ru"`
			Pl    string `json:"pl"`
			Ko    string `json:"ko"`
			ZhCht string `json:"zh-cht"`
			ZhChs string `json:"zh-chs"`
		} `json:"mobileWorldContentPaths"`
		JSONWorldContentPaths struct {
			En    string `json:"en"`
			Fr    string `json:"fr"`
			Es    string `json:"es"`
			EsMx  string `json:"es-mx"`
			De    string `json:"de"`
			It    string `json:"it"`
			Ja    string `json:"ja"`
			PtBr  string `json:"pt-br"`
			Ru    string `json:"ru"`
			Pl    string `json:"pl"`
			Ko    string `json:"ko"`
			ZhCht string `json:"zh-cht"`
			ZhChs string `json:"zh-chs"`
		} `json:"jsonWorldContentPaths"`
		MobileClanBannerDatabasePath string `json:"mobileClanBannerDatabasePath"`
		MobileGearCDN                struct {
			Geometry    string `json:"Geometry"`
			Texture     string `json:"Texture"`
			PlateRegion string `json:"PlateRegion"`
			Gear        string `json:"Gear"`
			Shader      string `json:"Shader"`
		} `json:"mobileGearCDN"`
		IconImagePyramidInfo []interface{} `json:"iconImagePyramidInfo"`
	} `json:"Response"`
	ErrorCode       int    `json:"ErrorCode"`
	ThrottleSeconds int    `json:"ThrottleSeconds"`
	ErrorStatus     string `json:"ErrorStatus"`
	Message         string `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}
