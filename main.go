package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/helphone/importer/job"
	"github.com/robfig/cron"
)

func refresh() {
	job.PullRepo()
	job.Refresh()
}

func main() {
	log.Info("Importer stared")

	c := cron.New()
	c.AddFunc("@every 1m", refresh)
	c.Start()

	refresh()

	select {}
}
