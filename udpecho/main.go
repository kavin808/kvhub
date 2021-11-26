package main

import (
	"kvhub/udpecho/tool"
	"kvhub/udpecho/udp"
	"log"
	"time"
)

func main() {
	// load config
	config, err := tool.NewConfig("./config.yaml")
	if err != nil {
		log.Printf("fail to load config %v", err)
		return
	}

	logFileName := "udpecho." + time.Now().Format("20060102150405") + ".log"
	path, err := tool.InitLog(config.LogDir, logFileName)
	if err != nil {
		log.Printf("Init log done")
	}
	log.Printf("listening %v, log path %v", config.Listen, path)

	if err := udp.ListenAndServe(config.Listen, udp.NewHandler()); err != nil {
		log.Printf("exit with err %v", err)
	}
}
