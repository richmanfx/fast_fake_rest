/****************************************************************************/
/*					Program to use in QA as the rest-stub. 					*/
/*                                                                          */
/*					 Russia, Moscow, Zoer aka R5AM, 2019.                   */
/*                                                                          */
/****************************************************************************/

package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/kataras/muxie"
	"net/http"
	"os"
	"time"
)

/* Config parameters */
type Conf struct {
	TcpPort        string     `yaml:"TcpPort"`
	DebugLevel     string     `yaml:"DebugLevel"`
	ConsoleLogging bool       `yaml:"ConsoleLogging"`
	FileLogging    bool       `yaml:"FileLogging"`
	LogFilesPath   string     `yaml:"LogFilesPath"`
	Rest           []struct { // `yaml:"rest"`
		EndPoint string `yaml:"EndPoint"`
		Response string `yaml:"Response"`
	}
}

var consoleLog = logrus.New()
var fileLog = logrus.New()

func main() {

	const (
		configDirName  = "/usr/local/etc/fast_fake_rest"
		configFileName = "config.yaml"
	)

	var (
		config Conf
	)

	// Full configuration file name
	fullConfigFileName := configDirName + "/" + configFileName
	//log.Debugf("Full config file name: %s", fullConfigFileName)

	// Read parameters from config file
	config.GetConfigParameters(fullConfigFileName)

	// Set logging parameters
	SetLog(config)

	// Для дебага
	//consoleLog.Infof("Config: %v", config)
	//fileLog.Infof("Config: %v", config)

	// Adding RESTs from config
	for index := 0; index < len(config.Rest); index++ {
		rests = append(rests, Rest{EndPoint: config.Rest[index].EndPoint, Response: config.Rest[index].Response})
	}

	// Adding RESTs from code
	rests = append(rests, Rest{EndPoint: "/v1/endpoint777", Response: "{response777: \"7777777\"}"})

	// Running
	//router := mux.NewRouter()
	router := muxie.NewMux()
	//router.HandleFunc("/", getRestsList).Methods("GET")

	router.HandleFunc("/", getRestsList)
	err := http.ListenAndServe(":"+config.TcpPort, router)
	if err != nil {
		consoleLog.Fatal("Error listen on the TCP network address")
	}

}

/* Set logging parameters */
func SetLog(config Conf) {

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
		fullLogFileName := config.LogFilesPath + "/" + "rest-" + time.Now().Format("20060102") + ".log"
		file, err := os.OpenFile(fullLogFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err == nil {
			fileLog.SetOutput(file)
			level, _ := logrus.ParseLevel(config.DebugLevel)
			fileLog.SetLevel(level)
		} else {
			fileLog.Errorf("Failed to log to file '%s', using default stderr", fullLogFileName)
		}
	}

}
