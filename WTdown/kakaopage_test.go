package WTdown

import (
	"testing"
)

func TestNaverload(t *testing.T) {
	nc := &KakaoPage{TitleId: "57132358", Cookies: "_kpawlt=''; _kpawltea=''; _kpawlst='';"}
	err := nc.GetEpiData()
	if err != nil {
		t.Log("Failed to load name of episodes")
	}
	t.Log(nc.epis, nc.EpisodeName)
	code, err := nc.Download(20, 20, 70, "D:\\Webtoons")
	t.Log(code, err)
	t.Error("Test Error")
}
