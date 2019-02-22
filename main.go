/****************************************************************************/
/*					Program to use in QA as the rest-stub. 					*/
/*                                                                          */
/*					 Russia, Moscow, Zoer aka R5AM, 2019.                   */
/*                                                                          */
/****************************************************************************/

package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func main() {

	const (
		configDirName  = "/usr/local/etc/fast_fake_rest"
		configFileName = "config.yaml"
	)

	// Full configuration file name
	fullConfigFileName := configDirName + "/" + configFileName
	log.Debugf("Full config file name: %s", fullConfigFileName)

	// Read parameters from config file

	// Set logging parameters
	SetLog(log.DebugLevel) // TODO: Уровень брать из конфига
}

/* Set logging parameters */
func SetLog(debugLevel log.Level) {
	log.SetOutput(os.Stdout)
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006/01/02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true
	log.SetLevel(debugLevel)
}
