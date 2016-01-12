package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/helphone/importer/job"
	"github.com/robfig/cron"
)

func refresh() {
	job.UpdateSource()
	job.Refresh()
}

func main() {
	log.Info("Importer stared")

	c := cron.New()
	c.AddFunc("@every 1h", refresh)
	c.Start()

	refresh()

	select {}
}
