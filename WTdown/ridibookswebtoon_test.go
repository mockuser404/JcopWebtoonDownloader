package WTdown

import (
	"log"
	"testing"
)

func TestRidi(t *testing.T) {
	nc := &RidiWT{TitleId: "4463000009", Cookies: "ridi-at=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlJTMDAwIn0.eyJjbGllbnRfaWQiOiJlUGdiS0tSeVB2ZEFGelR2RmcyRHZyUzdHZW5mc3RIZGtRMnV2Rk5kIiwiZXhwIjoxNjI1Nzk2NzY2LCJzY29wZSI6ImFsbCIsInN1YiI6Im15bmFtZWlzcHlvIiwidV9pZHgiOjUzOTUyMTUsImlhdCI6MTYyNTc5MzE2Nn0.gxB_Hkq1h31ftmIYq-9BHJ2PXkc-LbQF881mi8V8AwWBdL5PqnwnpoIvtLWaNhj7W1N7OM3lK7GTyAJaVGr2a2Y1oSismVQmuBhghQc8ZA9yZ1kiE-frzAMhmk09xba1Tg6DXIEELnnjZIN7AGx--Xay9ppYMBAG3ARBMlflBMG_mX00l05eD4JFeViRHG9W9FDQJ1alnUXl2igdfahTGpjfsXEFphzk8J0Wl_ECXE7laSsYZUTSrJJ0fotujSFyGO7JVJw1IxbUw7FtZ-l7yLO8Rf9EZOisPUrvkmtOt3Wn6-OCpvbPU-NoWm1dtOjDFQ9SLv0aj3rSiI6nX2xOTMUgzon6lzWfWC_5yWgkrEvRHqEbCTes9p5wsAQId-S-N05OD-U9cQ0mzaQQG79NyqJYy_fuKNsNQSG7IPylk1ElQzEWWa3AV7DjPzOCsYFzukw_eav7zsBL3wJqgD4X9lITjkopMQcGdhpgkxgR3sHkoDe7se9hcTTPTs9jRhUfCnIWV5Rbfc6RVUsDyShyQWypygMnKQ6ArXn2Z6ZKxa7I7SyZOiKKNy3R9pCEolFx8w_47dHBuW4EVErvswxCcK9aAd8dS2R8WXSjfDwofEYqS31JFVlVyEIRvfia8u5TsE3LKhu5QABGSFhRQGZLmH3nZtpldhyzW4MFhyExLUo;"}
	err := nc.GetEpiData()
	if err != nil {
		t.Log("Failed to load name of episodes")
	}
	log.Println(len(nc.epis), len(nc.EpisodeName))
	_, err = nc.Download(5, 5, 70, "D:\\Webtoons")
	if err != nil {
		log.Println(err)
	}
	t.Error("test")
}
