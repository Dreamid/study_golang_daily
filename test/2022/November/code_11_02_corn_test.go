package november

import (
	"github.com/robfig/cron/v3"
	"log"
	"testing"
)

func TestGolangCron(t *testing.T) {
	log.Println("---start----")
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("hello,golang!")
	})
	c.Start()

}
