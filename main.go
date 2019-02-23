/****************************************************************************/
/*					Program to use in QA as the rest-stub. 					*/
/*                                                                          */
/*					 Russia, Moscow, Zoer aka R5AM, 2019.                   */
/*                                                                          */
/****************************************************************************/

package main

import (
	"github.com/Sirupsen/logrus"
	"os"
)

/* Config parameters */
type conf struct {
	DebugLevel     string `yaml:"DebugLevel"`
	ConsoleLogging bool   `yaml:"ConsoleLogging"`
	FileLogging    bool   `yaml:"FileLogging"`
	LogFilesPath   string `yaml:"LogFilesPath"`
}

var consoleLog = logrus.New()
var fileLog = logrus.New()

func main() {

	const (
		configDirName  = "/usr/local/etc/fast_fake_rest"
		configFileName = "config.yaml"
	)

	var (
		config conf
	)

	// Full configuration file name
	fullConfigFileName := configDirName + "/" + configFileName
	//log.Debugf("Full config file name: %s", fullConfigFileName)

	// Read parameters from config file
	config.GetConfigParameters(fullConfigFileName)

	// Set logging parameters
	SetLog(config)

	consoleLog.Infof("Config: %v", config)
	fileLog.Infof("Config: %v", config)
}

/* Set logging parameters */
func SetLog(config conf) {

	// Console logging
	consoleLog.Formatter = &logrus.TextFormatter{
		FullTimestamp:    false,
		QuoteEmptyFields: true,
		TimestampFormat:  "2006/01/02 15:04:05"}
	if config.ConsoleLogging == true {
		consoleLog.SetOutput(os.Stdout)
		level, _ := logrus.ParseLevel(config.DebugLevel)
		consoleLog.Level = level
	} else {
		// For silent mode
		level, _ := logrus.ParseLevel("PanicLevel")
		consoleLog.Level = logrus.Level(level)
	}

	// File logging
	if config.FileLogging == true {
		fileLog.Formatter = consoleLog.Formatter
		fullLogFileName := config.LogFilesPath + "/" + "rest.log"
		file, err := os.OpenFile(fullLogFileName, os.O_APPEND|os.O_WRONLY, 0666)
		if err == nil {
			fileLog.SetOutput(file)
			level, _ := logrus.ParseLevel(config.DebugLevel)
			fileLog.SetLevel(level)
		} else {
			fileLog.Errorf("Failed to log to file '%s', using default stderr", fullLogFileName)
		}
	}

}
