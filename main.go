package main

import (
	"fmt"
	"log"

	"github.com/sebastiengodin/alclottoscheduler/internal/auth"
	"github.com/sebastiengodin/alclottoscheduler/internal/cli"
	"github.com/sebastiengodin/alclottoscheduler/internal/jackpotsfeeds"
	"github.com/sebastiengodin/alclottoscheduler/internal/readyaml"
	"github.com/sebastiengodin/alclottoscheduler/internal/sheets"
	"github.com/sebastiengodin/alclottoscheduler/structs"
)

//"github.com/sebastiengodin/alclottoscheduler/internal/auth"

//"github.com/sebastiengodin/alclottoscheduler/internal/sheets"

func main() {

	//fbapi.GetImageList()
	//fbapi.GetVideoList()

	fmt.Println("testgo")

	//read attributes from cli
	var args structs.Args
	cli.ReadArgs(&args)

	//import settings
	var configs structs.Config
	readyaml.GetConfigs(&configs)

	//get ALC api values
	var drawGame structs.DrawGames
	jackpotsfeeds.GetFeeds(&drawGame, &args, &configs)

	//get auth
	srv := auth.GetAuth()

	//get conditions

	//return items to process from sheet
	itemsToProcess := sheets.GetSheetsData(&configs, srv)

	log.Println(itemsToProcess)

	//process items based on attributes

}
