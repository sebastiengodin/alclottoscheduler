package structs

import (
	"time"
)

type Args struct {
	Lotto string
}

type Config struct {
	Settings struct {
		TestMode    bool      `yaml:"testMode"`
		TestDate    time.Time `yaml:"testDate"`
		CurrentDate time.Time
	}
	Sheets struct {
		SpreadsheetId string `yaml:"spreadsheetId"`
		SheetName     string `yaml:"sheetName"`
		LoadRange     string `yaml:"loadRange"`
	}
	Sources struct {
		LottoMaxUrl string `yaml:"lottomaxURL"`
		Lotto649Url string `yaml:"lotto649URL"`
	}
}

// Config functions
func (c *Config) PostMarshall() error {

	c.Settings.CurrentDate = time.Now().UTC()

	return nil
}

type Sheet struct {
	Status                 string    //col 0 in sheet
	Lotto                  string    //col 1
	Type                   string    //col 2
	StartDate              time.Time //col 3
	EndDate                time.Time //col 4
	DayOfDraw              string    //col 5
	CampaignId             int64     //col 6
	AdId                   int64     //col 7
	PageId                 int64     //col 8
	InstagramActorId       int64     //col 9
	StartRange             int64     //col 10
	EndRange               int64     //col 11
	HighJackpotAddedBudget float64   //col 12
	PrimaryText            string    //col 13
	Headline               string    //col 14
	Link                   string    //col 15
	ActionType             string    //col 16
	MediaSearchPattern     string    //col 17
}

// Sheets functions
// get day of the week number. Week starting on Monday
func (s *Sheet) GetDrawDayOfWeek() int64 {
	switch s.DayOfDraw {
	case "Monday":
		return 0
	case "Tuesday":
		return 1
	case "Wednesday":
		return 2
	case "Thursday":
		return 3
	case "Friday":
		return 4
	case "Saturday":
		return 5
	case "Sunday":
		return 6
	}
	return 99
}

type DrawGames []struct {
	Draw struct {
		ProviderDrawID  string        `json:"providerdrawId"`
		BonusNumber     string        `json:"bonus_number"`
		PrizePayouts    []PrizePayout `json:"prize_payouts"`
		Tag             string        `json:"tag"`
		TagPrizePayouts []PrizePayout `json:"tag_prize_payouts"`
		WinningNumbers  []string      `json:"winning_numbers"`
	} `json:"draw"`
	DrawDate         string            `json:"draw_date"`
	Game             string            `json:"game"`
	GuaranteedDraws  interface{}       `json:"guaranteed_draws"`
	LastEditDate     string            `json:"last_edit_date"`
	NextDraw         NextDraw          `json:"next_draw"`
	PromotionalDraws []PromotionalDraw `json:"promotional_draws"`
	StandardBalls    int               `json:"standard_balls"`
	JackpotBalls     int               `json:"jackpot_balls"`
	JackpotBallDrawn bool              `json:"jackpot_ball_drawn"`
}

type PrizePayout struct {
	AtlanticBreakdowns     *[]AtlanticBreakdown `json:"atlantic_breakdowns"`
	GuaranteedPrizeEnglish interface{}          `json:"guaranteed_prize_english"`
	GuaranteedPrizeFrench  interface{}          `json:"guaranteed_prize_french"`
	GuaranteedPrizeType    interface{}          `json:"guaranteed_prize_type"`
	NumberOfPrizes         int                  `json:"number_of_prizes"`
	PrizeValue             float64              `json:"prize_value"`
	RegionBreakdowns       []RegionBreakdown    `json:"region_breakdowns"`
	Type                   string               `json:"type"`
}

type AtlanticBreakdown struct {
	City           string `json:"city"`
	Province       string `json:"province"`
	Online         bool   `json:"online"`
	NumberOfPrizes int    `json:"number_of_prizes"`
}

type RegionBreakdown struct {
	Region         string `json:"region"`
	NumberOfPrizes int    `json:"number_of_prizes"`
}

type NextDraw struct {
	ProviderDrawId                    string      `json:"providerDrawId"`
	DrawDate                          string      `json:"draw_date"`
	Jackpot                           float64     `json:"jackpot"`
	EstimatedNumberOfPromotionalDraws int         `json:"estimated_number_of_promotional_draws"`
	GuaranteedPrizeEnglish            interface{} `json:"guaranteed_prize_english"`
	GuaranteedPrizeFrench             interface{} `json:"guaranteed_prize_french"`
	GuaranteedPrizeType               string      `json:"guaranteed_prize_type"`
	EstimatedJackpot                  interface{} `json:"estimated_jackpot"`
	ClassicJackpot                    interface{} `json:"classic_jackpot"`
	StandardBalls                     interface{} `json:"standard_balls"`
	JackpotBalls                      interface{} `json:"jackpot_balls"`
}

type PromotionalDraw struct {
	WinningNumbers []string      `json:"winning_numbers"`
	PrizePayouts   []PrizePayout `json:"prize_payouts"`
	BonusNumber    interface{}   `json:"bonus_number"`
}

type FBAPIResponse struct {
	Data []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Hash  string `json:"hash"`
		Title string `json:"title"`
	} `json:"data"`
	Paging struct {
		Next string `json:"next"` // URL for the next page of results
	} `json:"paging"`
}
