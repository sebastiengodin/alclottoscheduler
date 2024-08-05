package cli

import (
	"os"

	"github.com/sebastiengodin/alclottoscheduler/models"
)

func ReadArgs(args *models.Args) {
	lotto := os.Args[1]
	args.Lotto = lotto

}
