package WTdown

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestMain(t *testing.T) {
	// client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.lezhin.com/ko/comic/girlwetwall", nil)
	if err != nil {
		t.Error(err)
	}
	// req.AddCookie(&http.Cookie{Name: "RSESSION", Value: "RjFzb1pFamh6ZmxBekpsYStmUmhsN0FBQmQ1RFUvQWZyL2ZsNUZiZjZ5REdEVnJBWSsvSllTNVRSNkRqNU9pOWlmMHBTT0dVQlZialJmdXZmeGlLbHhhaiswNFp4UUZlZkRVZjJsa2pFN0tLMXJhUjA0czg2Tkg1VFdoSTRBUUc1bUFwcDE4ZmJKV0JQb1dDRjdZQlNEVllINzlFMVlMcldOcTJ6Qjl0ODduLzJ3b3ZreEJtTkZTVWpHdHVuOTRXcFRiWDZPSlVCYk11dUt4SzNkR3NScWZ3dUNCUDZ3NkM5c1ZwT1g4eTdWNk1WRnUwbkpPVXIvVnlKeXpiNUNFVDI4b0xKK0ozTjZQc2x6aEJMUTNrZ2ZRYmtvNDJzb3d3WUJnSzRlMEQ4aHlEd0lzT0Z1d09nS3ZXZXU4ajJvRCstLWdmQ1BYVXRoMklGRFVuMXFKUDh6NHc9PQ%3D%3D--3979488d310d732c04f99a28c2f22668fa9cdb14"})
	req.Header.Add("Cookie", "RSESSION=SFY0RC8xQ2MxdldJeHV6SEZnY0tmV0p0dlFzVkZ5bS9ub1o2TlRLNFJIazZRL1BKZjExUndBd2hFbnN3a3VJK0VOcExvUERRZXlxeGxod1JpQy9sOE9PS3ZuRm8xcDkrMkgySEJpVTVPVnUvN0ptTFc4N3B4YkxDZjdOc3UwRk03dk1sR09JV0ZxdS9QV3hwQWFIUDFWbEJuOUFhSWViNEpONDU5NGdybWpLY3ZLWFFBNHZ3MCtSekxJUzJKRW4xanhzSWNWMXU2ejZjbkp5elJjNFRHcitoOE9IQjUyVGlWK21Jdkhsc1BlMWViV3lLeXI5TU5RTlVPMUk0QjhmRmthZzFHd0ZuMlUxekFvZ0hYWENqbnIxMHRYa1NjYm41ZmJ0dmEwenl2cURKSVdyeGdKWGNGaHdRTHp5SGFadTgtLWJ4bzFoQ3p2eWoyU2pWMEFBcUlQSXc9PQ%3D%3D--109ebab11ad726bf9dbdc074e973e9a06ecf6d4c")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	// resp, _ := client.Do(reqest)
	// 	fmt.Println(resp.Status)
	// }

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	log.Println(resp.Header)
	if err != nil {
		t.Error(err)
	}

	source, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	log.Println("RESULT:", string(source))
	t.Error("test")
}
