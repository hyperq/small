package logs

import log "github.com/sirupsen/logrus"

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		TimestampFormat:        "2006-01-02 15:04:05",
		DisableLevelTruncation: true,
	})
}

//Debug Debug
func Debug(arg ...interface{}) {
	log.Debug(arg...)
}

//Error Error
func Error(arg ...interface{}) {
	log.Error(arg...)
}
