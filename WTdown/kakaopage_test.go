package WTdown

import (
	"testing"
)

func TestNaverload(t *testing.T) {
	nc := &KakaoPage{TitleId: "55479899", Cookies: "_kpawlt=qIrp9ItF_RTcFM37jY6Qd2Hj__0Vb0oSQRTkset1L7tVT_9u4UoVfur5QLlW1ToX-5xvE3FHZ02HSltb-bsyHA1q_C7Ce68Sl5szAQ0w1Z55e_NIVReb9jq5UmAFfHij; _kpawltea=1624649908855; _kpawlst=VFzltox6LWOqBiny9T_xe4huMxj8TodvB_AzUDXDUIYjm0hI8d96M9OveZP4i5TWDKrap3r1CTIbKLM1Hjzeqvoa40Q4srU6H9QYSMhv5e0;"}
	err := nc.GetEpiData()
	if err != nil {
		t.Log("Failed to load name of episodes")
	}
	t.Log(nc.epis, nc.EpisodeName)
	code, err := nc.Download(20, 20, 70, "D:\\Webtoons")
	t.Log(code, err)
	t.Error("Test Error")
}
