package WTdown

import (
	"testing"
)

func TestEpub(t *testing.T) {
	nc := &KPepub{TitleId: "57389031", Cookies: ""}
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
	_, err = nc.Download(0, 0, "D:\\Webtoons")
	if err != nil {
		t.Log(err)
	}
	// for i:=0;i<6;i++{
	// }
	t.Error("Test Error")
}
