package sheets

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/sebastiengodin/alclottoscheduler/models"

	"google.golang.org/api/sheets/v4"
)

func GetSheetsData(configs *models.Config, srv *sheets.Service) []models.Sheet {

	spreadsheetId := configs.Sheets.SpreadsheetId
	mrange := fmt.Sprintf("%s!%s", configs.Sheets.SheetName, configs.Sheets.LoadRange)

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, mrange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	var itemsToProcess []models.Sheet

	if len(resp.Values) == 0 {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	} else {
		for i, r := range resp.Values {

			//preprocess
			adid, err := convertIntefaceToInt(r[3])
			if err != nil {
				log.Fatalf("The Spreadsheet does not contain a valid Page ID on row %v", i+1)
			}

			pageid, err := convertIntefaceToInt(r[4])
			if err != nil {
				log.Fatalf("The Spreadsheet does not contain a valid Instagram Page ID on row %v", i+1)
			}

			instagramid, err := convertIntefaceToInt(r[5])
			if err != nil {
				log.Fatalf("The Spreadsheet does not contain a valid Start Range on row %v", i+1)
			}

			startRange, err := convertIntefaceToInt(r[6])
			if err != nil {
				log.Fatalf("The Spreadsheet does not contain a valid End Range on row %v", i+1)
			}

			endRange, err := convertIntefaceToInt(r[7])
			if err != nil {
				log.Fatalf("The Spreadsheet does not contain a valid End Range on row %v", i+1)
			}

			highJackpotAddedBudget, err := convertIntefaceToFloat(r[8])
			if err != nil {
				log.Fatalf("The Spreadsheet does not contain a valid Added Budget on row %v", i+1)
			}

			row := models.Sheet{
				Lotto:                  convertInterfaceToStr(r[0]),
				Type:                   convertInterfaceToStr(r[1]),
				DayOfDraw:              convertInterfaceToStr(r[2]),
				AdId:                   adid,
				PageId:                 pageid,
				InstagramActorId:       instagramid,
				StartRange:             startRange,
				EndRange:               endRange,
				HighJackpotAddedBudget: highJackpotAddedBudget,
				PrimaryText:            convertInterfaceToStr(r[9]),
				Headline:               convertInterfaceToStr(r[10]),
				Link:                   convertInterfaceToStr(r[11]),
				ActionType:             convertInterfaceToStr(r[12]),
			}

			itemsToProcess = append(itemsToProcess, row)
		}
	}

	return itemsToProcess
}

//inner functions

func convertIntefaceToInt(v interface{}) (int64, error) {
	var val int64

	if strVal, ok := v.(string); ok {
		val, err := strconv.ParseInt(strVal, 10, 64)
		if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
		}
		return val, nil
	}

	return val, errors.New("not a valid int64")
}

func convertIntefaceToFloat(v interface{}) (float64, error) {
	var val float64

	if v == "" {
		return 0, nil
	}

	if strVal, ok := v.(string); ok {
		val, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
		}
		return val, nil
	}

	return val, errors.New("not a valid float number")
}

func convertInterfaceToStr(s interface{}) string {
	if strVal, ok := s.(string); ok {
		return strVal
	} else {
		log.Panicln("Not a valid string")
		return ""
	}
}
