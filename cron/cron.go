package cron

import (
	"github.com/robfig/cron"
	"go-gin/models"
	"go-gin/pkg/logging"
)

func Start() {
	logging.Info("Cron Starting...")
	c := cron.New()
	c.AddFunc("* * 0 * * *", func() {
		logging.Info("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * 0 * * *", func() {
		logging.Info("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})
	c.Start()

}
