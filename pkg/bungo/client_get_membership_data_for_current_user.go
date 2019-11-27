package bungo

import "time"

type GetMembershipDataForCurrentUser struct {
	Response struct {
		DestinyMemberships []struct {
			LastSeenDisplayName       string `json:"LastSeenDisplayName"`
			LastSeenDisplayNameType   int    `json:"LastSeenDisplayNameType"`
			IconPath                  string `json:"iconPath"`
			CrossSaveOverride         int    `json:"crossSaveOverride"`
			ApplicableMembershipTypes []int  `json:"applicableMembershipTypes"`
			IsPublic                  bool   `json:"isPublic"`
			MembershipType            int    `json:"membershipType"`
			MembershipID              string `json:"membershipId"`
			DisplayName               string `json:"displayName"`
		} `json:"destinyMemberships"`
		BungieNetUser struct {
			MembershipID        string    `json:"membershipId"`
			UniqueName          string    `json:"uniqueName"`
			DisplayName         string    `json:"displayName"`
			ProfilePicture      int       `json:"profilePicture"`
			ProfileTheme        int       `json:"profileTheme"`
			UserTitle           int       `json:"userTitle"`
			SuccessMessageFlags string    `json:"successMessageFlags"`
			IsDeleted           bool      `json:"isDeleted"`
			About               string    `json:"about"`
			FirstAccess         time.Time `json:"firstAccess"`
			LastUpdate          time.Time `json:"lastUpdate"`
			Context             struct {
				IsFollowing  bool `json:"isFollowing"`
				IgnoreStatus struct {
					IsIgnored   bool `json:"isIgnored"`
					IgnoreFlags int  `json:"ignoreFlags"`
				} `json:"ignoreStatus"`
			} `json:"context"`
			PsnDisplayName       string    `json:"psnDisplayName"`
			ShowActivity         bool      `json:"showActivity"`
			Locale               string    `json:"locale"`
			LocaleInheritDefault bool      `json:"localeInheritDefault"`
			ShowGroupMessaging   bool      `json:"showGroupMessaging"`
			ProfilePicturePath   string    `json:"profilePicturePath"`
			ProfileThemeName     string    `json:"profileThemeName"`
			UserTitleDisplay     string    `json:"userTitleDisplay"`
			StatusText           string    `json:"statusText"`
			StatusDate           time.Time `json:"statusDate"`
			BlizzardDisplayName  string    `json:"blizzardDisplayName"`
			SteamDisplayName     string    `json:"steamDisplayName"`
		} `json:"bungieNetUser"`
	} `json:"Response"`
	ErrorCode       int    `json:"ErrorCode"`
	ThrottleSeconds int    `json:"ThrottleSeconds"`
	ErrorStatus     string `json:"ErrorStatus"`
	Message         string `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}
