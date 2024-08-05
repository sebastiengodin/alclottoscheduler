package readyaml

import (
	"log"
	"os"

	"github.com/sebastiengodin/alclottoscheduler/models"

	"gopkg.in/yaml.v2"
)

func GetConfigs(configs *models.Config) {
	//read YAML file
	data, err := os.ReadFile("configs.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(data, &configs)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//add current date with Config Method
	if err = configs.PostMarshall(); err != nil {
		log.Fatalf("Post unmarshal error: %v", err)
	}

}
