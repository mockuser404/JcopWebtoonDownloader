package main

func Log(code int, message error) {
	switch code {
	case 1:
		mw.openErrorMessBox("ERROR", message.Error())
	case 2:
		mw.openWarningMessBox("WARNING", message.Error())
	case 3:
		mw.openInfoMessBox("INFO", message.Error())
	}
}

func LoadingOn() {
	buttonLog.SetEnabled(false)
	buttonLog.SetText("Loading...")
}

func LoadingOff() {
	buttonLog.SetEnabled(true)
	buttonLog.SetText("Download")
}
