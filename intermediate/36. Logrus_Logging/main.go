package main

import "github.com/sirupsen/logrus"

// go mod init Logrus_Logging
// go mod tidy

func main() {

	log := logrus.New()

	// Set log level
	log.SetLevel(logrus.InfoLevel)

	// Set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	// Logging examples
	log.Info("This is an info message.")
	log.Warn("This is a warning message.")
	log.Error("This is an error message.")

	log.WithFields(logrus.Fields{
		"username": "Abhay Singh",
		"method":   "GET",
	}).Info("User logged in.")

}
