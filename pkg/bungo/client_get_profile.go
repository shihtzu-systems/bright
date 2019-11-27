package bungo

import "time"

type GetProfile struct {
	Response struct {
		Profile struct {
			Data struct {
				UserInfo struct {
					CrossSaveOverride         int    `json:"crossSaveOverride"`
					ApplicableMembershipTypes []int  `json:"applicableMembershipTypes"`
					IsPublic                  bool   `json:"isPublic"`
					MembershipType            int    `json:"membershipType"`
					MembershipID              string `json:"membershipId"`
					DisplayName               string `json:"displayName"`
				} `json:"userInfo"`
				DateLastPlayed time.Time     `json:"dateLastPlayed"`
				VersionsOwned  int           `json:"versionsOwned"`
				CharacterIds   []string      `json:"characterIds"`
				SeasonHashes   []interface{} `json:"seasonHashes"`
			} `json:"data"`
			Privacy int `json:"privacy"`
		} `json:"profile"`
	} `json:"Response"`
	ErrorCode       int    `json:"ErrorCode"`
	ThrottleSeconds int    `json:"ThrottleSeconds"`
	ErrorStatus     string `json:"ErrorStatus"`
	Message         string `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}
