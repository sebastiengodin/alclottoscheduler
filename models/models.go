package models

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

// Config methods
func (c *Config) PostMarshall() error {

	c.Settings.CurrentDate = time.Now().UTC()

	return nil
}

type Sheet struct {
	Lotto                  string
	Type                   string
	DayOfDraw              string
	AdId                   int64
	PageId                 int64
	InstagramActorId       int64
	StartRange             int64
	EndRange               int64
	HighJackpotAddedBudget float64
	PrimaryText            string
	Headline               string
	Link                   string
	ActionType             string
}

// add day of the week with a number method

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
