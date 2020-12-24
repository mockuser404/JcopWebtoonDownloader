################################################################################
##
## BY: Inpyo Lee
## PROJECT MADE WITH: Qt Designer and PySide2
## V: 2.1.1
##
################################################################################

import sys
import os
import subprocess
import json
from PySide2.QtCore import *
from PySide2.QtGui import *
from PySide2.QtWidgets import *

# GUI FILE
from ui_main import Ui_MainWindow

# IMPORT FUNCTIONS
from download_func import *
from cookie_saving import * 
import search_func 


version = "2.1.1"
class MainWindow(QMainWindow, WebtoonDownload, SavingCookie):
    def __init__(self):

        # Main Setup
        QMainWindow.__init__(self)
        WebtoonDownload.__init__(self)
        self.ui = Ui_MainWindow()
        self.ui.setupUi(self)


        # Downloader var
        self.data = {}
        self.data['headers'] = {'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36'}
        self.data['cookie'] = {}

        # Add function in buttons
        self.ui.pushButton_download.clicked.connect(self.download)
        self.ui.pushButton_Search.clicked.connect(self.search)
        self.ui.actionNaver_Comic.triggered.connect(self.genNaverComicCookieData)
        self.ui.actionKakao_Page.triggered.connect(self.genKakaoPageCookieData)
        self.ui.actionLezhin_Comics.triggered.connect(self.genLezhinComicsCookieData)
        self.ui.actionDefault_Directory.triggered.connect(self.defaultDirectory)
        self.ui.actionOpen_Download_Folder.triggered.connect(self.openDownloadDir)
        self.ui.actionVersion.triggered.connect(lambda: self.showVersionMessage(version))
        self.ui.actionVisit_Blog_3.triggered.connect(lambda: subprocess.call(["rundll32", "url.dll,FileProtocolHandler", "https://blog.naver.com/the3countrys/222106929101"]))

        # Show Main
        self.show()

        self.isApplicationLastVersion()
        self.loadCookie()

    
    def defaultDirectory(self):
        fname = QFileDialog.getExistingDirectory(self, dir=self.data["cookie"]["defaultDirectory"])
        if fname != "":
            with open("data.json", "w") as f:
                normal_data = self.data['cookie']
                normal_data["defaultDirectory"] = fname
                f.write(json.dumps(normal_data))
                f.close()
        
    def isApplicationLastVersion(self):
        try:
            releases = json.loads(requests.get("https://api.github.com/repos/mynameispyo/JcopWebtoonDownloader/releases/latest",timeout=5).text)
            version =  open("version.txt", "r+")
            if releases["tag_name"] != version.readline():
                self.noticeUpdatesAvaliable()
        except Exception as e: 
            self.showErrorMessage(str(e))
    
    def noticeUpdatesAvaliable(self):
        msgBox = QMessageBox()
        msgBox.setIcon(QMessageBox.Information)
        msgBox.setText("New Version is avaliable. Please update application")
        msgBox.setWindowTitle("Update Notice")
        msgBox.setStandardButtons(QMessageBox.Ok | QMessageBox.Cancel)

        returnValue = msgBox.exec()
        if returnValue == QMessageBox.Ok:
            os.startfile("updater.exe")
            sys.exit()
    
    def download(self):
        try:
            self.loadCookie()
            self.getInputs()
            if self.data['Type'] == "Kakao Page":
                WebtoonDownload.KakaoDownload(self)
            elif self.data['Type'] == "Naver Comic":
                WebtoonDownload.NaverDownload(self)
            elif self.data['Type'] == "Daum Webtoon":
                WebtoonDownload.DaumDownload(self)
            elif self.data['Type'] == "Lezhin Comics":
                WebtoonDownload.LezhinDownload(self)
            else:
                self.showWarningMessage("Invalid webtoon type")
        except Exception as e: 
            self.showErrorMessage(str(e))

    def search(self):
        try:
            self.getInputs()
            if self.data['Type'] == "Kakao Page":
                if search_func.searchKakaoId(self.data['Id']) != None:
                    self.ui.lineEdit_id.setText(search_func.searchKakaoId(self.data['Id']))
                else:
                    self.showWarningMessage("No results found for "+self.data['Id'])
            elif self.data['Type'] == "Naver Comic":
                if search_func.searchNaverId(self.data['Id']) != None:
                    self.ui.lineEdit_id.setText(search_func.searchNaverId(self.data['Id']))
                else:
                    self.showWarningMessage("No results found for "+self.data['Id'])
            elif self.data['Type'] == "Daum Webtoon":
                if search_func.searchDaumId(self.data['Id']) != None:
                    self.ui.lineEdit_id.setText(search_func.searchDaumId(self.data['Id']))
                else:
                    self.showWarningMessage("No results found for "+self.data['Id'])
            elif self.data['Type'] == "Lezhin Comics":
                if search_func.searchLezhinId(self.data['Id']) != None:
                    self.ui.lineEdit_id.setText(search_func.searchLezhinId(self.data['Id']))
                else:
                    self.showWarningMessage("No results found for "+self.data['Id'])
            else:
                self.showWarningMessage("Invalid webtoon type")
        except Exception as e: 
            self.showErrorMessage(str(e))

    def getInputs(self):
        self.data['Type'] = self.ui.comboBox_webtoon_type.currentText()
        self.data['Id'] = self.ui.lineEdit_id.text()
        self.data['Start'] = self.ui.lineEdit_start.text()
        self.data['End'] = self.ui.lineEdit_end.text()
        self.data['DeviceId'] = self.ui.lineEdit_deviceid.text()


    def isInputEmtpy(self):
        if self.data['Id'] == "":
            self.showWarningMessage("No input 'Id'")
            return True

        elif self.data['Start'] == "":
            self.showWarningMessage("No input 'Start Page'")
            return True

        elif self.data['End'] == "":
            self.showWarningMessage("No input 'End Page'")
            return True

        else:
            return False

    def showErrorMessage(self, message):
        msg = QMessageBox()
        msg.setIcon(QMessageBox.Critical)
        msg.setText("Error")
        msg.setInformativeText(message)
        msg.setWindowTitle("Error")
        msg.exec_()

    def showWarningMessage(self, message):
        msg = QMessageBox()
        msg.setIcon(QMessageBox.Warning)
        msg.setText("Warning")
        msg.setInformativeText(message)
        msg.setWindowTitle("Warning")
        msg.exec_()

    def showVersionMessage(self,message):
        msg = QMessageBox()
        msg.setIcon(QMessageBox.Information)
        msg.setText("Version")
        msg.setInformativeText(message)
        msg.setWindowTitle("Version")
        msg.exec_()
    
    def showInfoMessage(self,message):
        msg = QMessageBox()
        msg.setIcon(QMessageBox.Information)
        msg.setText("Information")
        msg.setInformativeText(message)
        msg.setWindowTitle("Information")
        msg.exec_()

    
    def openDownloadDir(self):
        subprocess.call(["rundll32", "url.dll,FileProtocolHandler", self.data["cookie"]["defaultDirectory"]])

if __name__ == "__main__":
    app = QApplication(sys.argv)
    app.setWindowIcon(QIcon('img/downloader.ico'))
    window = MainWindow()
    sys.exit(app.exec_())
