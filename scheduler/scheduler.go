package scheduler

import (
	"log"
	"time"
	"url-shortener/config"
	"url-shortener/model"
	"url-shortener/service"
)

type Scheduler struct {
	UrlService service.UrlService
}

func (s *Scheduler) Run() {

	for {
		expired := s.UrlService.CheckForExpired()
		for i := range expired {
			url := expired[i]
			if err := s.UrlService.DeleteShortUrl(model.DeleteUrlRequest{
				Short: url.Short,
			}); err != nil {
				log.Println("failed to delete expired url: ", err.Message)
			}
		}

		time.Sleep(time.Duration(config.CleanupInterval) * time.Minute)
	}

}
