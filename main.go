package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/helphone/importer/job"
	"github.com/robfig/cron"
)

func refresh() {
	firstPass, err := job.IsDatabaseEmpty()
	if err != nil {
		log.Info("Database failed")
		return
	}
	needRefresh, err := job.UpdateSource()
	if (needRefresh == true || firstPass == true) && err == nil {
		log.Infof("needRefresh is %v and firstPass is %v", needRefresh, firstPass)
		job.Refresh()
	}
}

func main() {
	log.Info("Importer stared")

	c := cron.New()
	c.AddFunc("@every 1m", refresh)
	c.Start()

	refresh()

	select {}
}
