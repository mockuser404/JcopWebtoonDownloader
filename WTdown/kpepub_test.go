package WTdown

import (
	"testing"
)

func TestEpub(t *testing.T) {
	nc := &KPepub{TitleId: "57352442", Cookies: ""}
	err := nc.GetEpiData()
	if err != nil {
		t.Log("Failed to load name of episodes")
	}
	// t.Log(nc.epis, nc.EpisodeName)
	// epubViewerId, err := nc.getTextURL("57375482")
	// if err != nil {
	// 	t.Log(err)
	// }
	// log.Println(epubViewerId)
	code, err := nc.Download(88, 88, "D:\\Webtoons")
	if err != nil {
		t.Log(code)
	}
	// for i:=0;i<6;i++{
	// }
	t.Error("Test Error")
}
