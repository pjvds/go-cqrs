package eventstore

import (
	"github.com/op/go-logging"
	"log"
	"os"
)

var Log = logging.MustGetLogger("eventstore")

func InitLogging() {
	b := logging.NewLogBackend(os.Stdout, "", log.LstdFlags)
	b.Color = true
	logging.SetBackend(b)
}
