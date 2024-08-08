package fbapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/sebastiengodin/alclottoscheduler/structs"
)

func GetImageList() {
	adAccountID := "741393411393254"
	accessToken := "EAAQmVxJJsZBoBOxyTwQkBOxpSVF83KHPSYL5evJg9sZBvZBoZBCZB0s1jWYhVuPjL7pg3PvaMB7wwegdgizgDJ1PwJiuJkUwmKZB6MRakHjJowEQpOLc21947H4W3cReP09WyXr5kiUlqLWo4bZBsCUD5nkz49YznonmgZA94XMIa0EKLZCZCODocLRPXU49Qm33gv"

	// construct the API URL
	baseURL := fmt.Sprintf("https://graph.facebook.com/v20.0/act_%s/adimages?fields=id,name,hash&access_token=%s", adAccountID, accessToken)

	if err := fetchData(baseURL); err != nil {
		log.Fatal(err)
	}

}

func GetVideoList() {
	adAccountID := "741393411393254"
	accessToken := "EAAQmVxJJsZBoBOxyTwQkBOxpSVF83KHPSYL5evJg9sZBvZBoZBCZB0s1jWYhVuPjL7pg3PvaMB7wwegdgizgDJ1PwJiuJkUwmKZB6MRakHjJowEQpOLc21947H4W3cReP09WyXr5kiUlqLWo4bZBsCUD5nkz49YznonmgZA94XMIa0EKLZCZCODocLRPXU49Qm33gv"

	// construct the API URL
	baseURL := fmt.Sprintf("https://graph.facebook.com/v20.0/act_%s/advideos?fields=id,title&access_token=%s", adAccountID, accessToken)

	if err := fetchData(baseURL); err != nil {
		log.Fatal(err)
	}
}

func fetchData(url string) error {
	// Make the HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into the struct
	var apiResponse structs.FBAPIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return err
	}

	// Output the results
	for _, image := range apiResponse.Data {
		fmt.Printf("Video ID: %s, Name: %s\n", image.ID, image.Title)
	}

	// Check if there is a next page and fetch it
	if apiResponse.Paging.Next != "" {
		return fetchData(apiResponse.Paging.Next)
	}

	return nil
}
