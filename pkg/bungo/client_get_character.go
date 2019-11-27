package bungo

import "time"

type GetCharacter struct {
	Response struct {
		Inventory struct {
			Data struct {
				Items []struct {
					ItemHash              int64     `json:"itemHash"`
					ItemInstanceID        string    `json:"itemInstanceId,omitempty"`
					Quantity              int       `json:"quantity"`
					BindStatus            int       `json:"bindStatus"`
					Location              int       `json:"location"`
					BucketHash            int       `json:"bucketHash"`
					TransferStatus        int       `json:"transferStatus"`
					Lockable              bool      `json:"lockable"`
					State                 int       `json:"state"`
					DismantlePermission   int       `json:"dismantlePermission"`
					IsWrapper             bool      `json:"isWrapper"`
					OverrideStyleItemHash int       `json:"overrideStyleItemHash,omitempty"`
					ExpirationDate        time.Time `json:"expirationDate,omitempty"`
				} `json:"items"`
			} `json:"data"`
			Privacy int `json:"privacy"`
		} `json:"inventory"`
		Character struct {
			Data struct {
				MembershipID             string      `json:"membershipId"`
				MembershipType           int         `json:"membershipType"`
				CharacterID              string      `json:"characterId"`
				DateLastPlayed           time.Time   `json:"dateLastPlayed"`
				MinutesPlayedThisSession string      `json:"minutesPlayedThisSession"`
				MinutesPlayedTotal       string      `json:"minutesPlayedTotal"`
				Light                    int         `json:"light"`
				Stats                    map[int]int `json:"stats"`
				RaceHash                 int         `json:"raceHash"`
				GenderHash               int64       `json:"genderHash"`
				ClassHash                int         `json:"classHash"`
				RaceType                 int         `json:"raceType"`
				ClassType                int         `json:"classType"`
				GenderType               int         `json:"genderType"`
				EmblemPath               string      `json:"emblemPath"`
				EmblemBackgroundPath     string      `json:"emblemBackgroundPath"`
				EmblemHash               int         `json:"emblemHash"`
				EmblemColor              struct {
					Red   int `json:"red"`
					Green int `json:"green"`
					Blue  int `json:"blue"`
					Alpha int `json:"alpha"`
				} `json:"emblemColor"`
				LevelProgression struct {
					ProgressionHash     int `json:"progressionHash"`
					DailyProgress       int `json:"dailyProgress"`
					DailyLimit          int `json:"dailyLimit"`
					WeeklyProgress      int `json:"weeklyProgress"`
					WeeklyLimit         int `json:"weeklyLimit"`
					CurrentProgress     int `json:"currentProgress"`
					Level               int `json:"level"`
					LevelCap            int `json:"levelCap"`
					StepIndex           int `json:"stepIndex"`
					ProgressToNextLevel int `json:"progressToNextLevel"`
					NextLevelAt         int `json:"nextLevelAt"`
				} `json:"levelProgression"`
				BaseCharacterLevel int     `json:"baseCharacterLevel"`
				PercentToNextLevel float64 `json:"percentToNextLevel"`
			} `json:"data"`
			Privacy int `json:"privacy"`
		} `json:"character"`
		Equipment struct {
			Data struct {
				Items []struct {
					ItemHash            int64  `json:"itemHash"`
					ItemInstanceID      string `json:"itemInstanceId"`
					Quantity            int    `json:"quantity"`
					BindStatus          int    `json:"bindStatus"`
					Location            int    `json:"location"`
					BucketHash          int    `json:"bucketHash"`
					TransferStatus      int    `json:"transferStatus"`
					Lockable            bool   `json:"lockable"`
					State               int    `json:"state"`
					DismantlePermission int    `json:"dismantlePermission"`
					IsWrapper           bool   `json:"isWrapper"`
				} `json:"items"`
			} `json:"data"`
			Privacy int `json:"privacy"`
		} `json:"equipment"`
	} `json:"Response"`
	ErrorCode       int    `json:"ErrorCode"`
	ThrottleSeconds int    `json:"ThrottleSeconds"`
	ErrorStatus     string `json:"ErrorStatus"`
	Message         string `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}
