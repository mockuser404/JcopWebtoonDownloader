package WTdown

import (
	"testing"
)

func TestNaverloadepi(t *testing.T) {
	nc := &NaverComic{TitleId :"628997",Cookies: "NID_AUT=rY4VBFsk1FLgUcF5pADmn9AH4MoZH+ps4TonQ8/iSWOQ9ZfjvE1t2Us3O8kYGOzB; NID_SES=AAABmqhFdZ8aD4+jS6WzqvavzjSmra7+itn11yCxG5pR7vt0ivKF3Pr1OYgKlW+EPZxoV5A/rCxzsayZfzekodouBGLqvu04bxQkXyE884el1S5LnU4t6ZU2wOdr+ZAUtJik/Jmp7XqGlQMgIZ12yxYIh/wLck+rNzMPQCzFVNHkiXXirqxBf/CiCmFDyg5dq+n242Rd53dhX+d2DEWhvByK/f/MwETAnKFmqS3tRXNLZ4yYrtegpq19EHjWz20p2VdXy+LOf69Osi1kySE6Td1i5OV3FJkZL/BeU9Ir0zahUz2C8BhrUxZhJ0m4+d2ddt8XLkTW4RVX4W1ayRYy7RfARUSaBDeUO43FWfdtlAVY9MF/f876ZZAnhRRmYuUlQEj2icObUmcLGTk7z4eMr01x4BpkQOtkqruNK8gqlW+sgw2BW+EyOnPqTWakukXwOt9gun4kzeKumOyGcY9dcHxXGRcB6HOJp2e6YhYYdgDGB1/RAt43eGhy2gFTWAp/0RIK/Bwa/BfsiHeReED8W/5OzaOem/5Bm4qS+6Oge1Pr63O7;"}
	nc.GetEpiData()
	code, err := nc.Download(1,1,70,"D:\\Webtoons")
	t.Log(code, err)
	t.Error("Test Error")
}

// func TestRun(t *testing.T){
// 	a := make([]string, 5)
// 	a[0] ="1"
// 	a[3] = "2"
// 	t.Error(a[1] =="")
// }
