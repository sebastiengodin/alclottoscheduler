package cli

import (
	"os"

	"github.com/sebastiengodin/alclottoscheduler/structs"
)

func ReadArgs(args *structs.Args) {
	lotto := os.Args[1]
	args.Lotto = lotto

}
