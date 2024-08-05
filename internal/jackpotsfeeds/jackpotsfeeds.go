package jackpotsfeeds

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/sebastiengodin/alclottoscheduler/models"
)

func GetFeeds(drawGames *models.DrawGames, args *models.Args, configs *models.Config) {
	var url string

	if args.Lotto == "LottoMax" {
		url = configs.Sources.LottoMaxUrl
	} else {
		url = configs.Sources.Lotto649Url
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("could not fetch feeds data")
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("could not read feeds data")
	}

	if err := json.Unmarshal(data, &drawGames); err != nil {
		log.Fatalf("could not parse JSON %v", err)
	}

	//fmt.Println(drawGames[0].NextDraw)
}
