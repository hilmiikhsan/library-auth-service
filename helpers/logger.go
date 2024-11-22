package helpers

import (
	"io"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func SetupLogger() {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	log.Info("Logger initiated using logrus")
	Logger = log
}

func SetupLogfile() {
	logFile, err := os.OpenFile("./logs/library_auth_service.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
