package WTdown

import "testing"

func TestDaum(t *testing.T) {
	nc := &DaumWebtoon{TitleId: "nonhyeonbusiness", Cookies: "_kpawlt=z7nSK735CBsLFEHPXrHy7z9tDQBDpksa4XEPlc4YqC0T1LVa1-IIddmu3bp67OG-lXX1dcz8wLuShknpqfzyNOax_jaq7OX7bbfihOqgeQ1xNU28fGWSXtSqZrBV0J1Z; _kpawltea=1624560687655; _kpawlst=BwuvAZ8Bm-39c4hsowOpUGOUK-SYuTf6h8yJ-CnRPGSRH3xh-zMFqAO9KKs1oX5EFs5d1rcDUfRRq3elTGUkaw;"}
	err := nc.GetEpiData()
	if err != nil {
		t.Log("Failed to load name of episodes")
	}
	t.Log(nc.epis, nc.EpisodeName)
	code, err := nc.Download(1, 1, 70, "D:\\Webtoons")
	t.Log(code, err)
	t.Error("Test Error")
}
