import os, json
from PySide2.QtCore import *
from PySide2.QtGui import *
from PySide2.QtWidgets import *

class SavingCookie:
    def genNaverComicCookieData(self):
        require_cookie = ["NID_AUT","NID_SES"]
        result = {}
        for name in require_cookie:
            text, ok = QInputDialog.getText(self, 'Naver', name)
            result[name] = text
            if not ok:
                return
        self.data["cookie"]["naver"] = result
        self.saveCookie()
        
    def genKakaoPageCookieData(self):
        require_cookie = ["_kawlp","_kawlptea","_kawlt"]
        result = {}
        for name in require_cookie:
            text, ok = QInputDialog.getText(self, 'Kakao', name)
            result[name] = text
            if not ok:
                return
        self.data["cookie"]["kakao"] = result
        self.saveCookie()

    def genLezhinComicsCookieData(self):
        require_cookie = ["x-lz-locale", "access_token"]
        result = {}
        for name in require_cookie:
            text, ok = QInputDialog.getText(self, 'Lezhin', name)
            result[name] = text
            if not ok:
                return
        self.data["cookie"]["lezhin"] = result
        self.saveCookie()

    def loadCookie(self):
        if not os.path.isfile("data.json"):
            with open("data.json", "w") as f:
                if not os.path.isdir(os.path.join(os.path.expandvars("%userprofile%"),"Documents", "Jcop Webtoon Downloader")):
                    os.makedirs(os.path.join(os.path.expandvars("%userprofile%"),"Documents", "Jcop Webtoon Downloader"))
                f.write(json.dumps({
                    "naver":{"NID_AUT":"","NID_SES":""},
                    "kakao":{"_kawlp":"","_kawlptea":"","_kawlt":""}, 
                    "lezhin":{"x-lz-locale":"ko_KR", "access_token":""},
                    "defaultDirectory": os.path.join(os.path.expandvars("%userprofile%"),"Documents", "Jcop Webtoon Downloader"),
                }))
                f.close()

        with open("data.json", "r") as f:
            data = f.read()
            f.close()
        try:
            data= json.loads(data)
        except Exception as e: 
            self.showErrorMessage(str(e))
        self.data['cookie'] = data
    
    def saveCookie(self):
        with open("data.json", "w") as f:
            f.write(json.dumps(self.data['cookie']))
            f.close()