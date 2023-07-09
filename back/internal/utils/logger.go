package utils

import "github.com/sirupsen/logrus"

func ConfigureLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
