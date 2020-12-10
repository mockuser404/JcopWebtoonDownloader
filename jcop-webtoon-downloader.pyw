################################################################################
##
## BY: WANDERSON M.PIMENTA
## PROJECT MADE WITH: Qt Designer and PySide2
## V: 1.0.0
##
################################################################################

import sys
import os
import platform
from PySide2 import QtCore, QtGui, QtWidgets
from PySide2.QtCore import *
from PySide2.QtGui import *
from PySide2.QtWidgets import *

# GUI FILE
from ui_main import Ui_MainWindow

# IMPORT FUNCTIONS
from download_functions import *

class MainWindow(QMainWindow, WebtoonDownload):
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
        self.ui.pushButton_open.clicked.connect(self.openDir)
        self.ui.menuSetting.triggered[QAction].connect(self.readCookie)
        self.ui.menuVersion.triggered[QAction].connect(lambda: self.showVersionMessage("1.4.2"))

        # Show Main
        self.show()

        self.isApplicationLastVersion()

    def isApplicationLastVersion(self):
        try:
            releases = json.loads(requests.get("https://api.github.com/repos/mynameispyo/JcopWebtoonDownloader/releases/latest").text)
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
    

    def getInputs(self):
        self.data['Type'] = self.ui.comboBox_webtoon_type.currentText()
        self.data['Id'] = self.ui.lineEdit_id.text()
        self.data['Start'] = self.ui.lineEdit_start.text()
        self.data['End'] = self.ui.lineEdit_end.text()
        self.data['DeviceId'] = self.ui.lineEdit_deviceid.text()
        self.data['Directory'] = self.ui.lineEdit_directory.text()


    def isInputEmtpy(self):
        if self.data['Id'] == "":
            self.showWarningMessage("No input 'Id'")
            return True

        elif self.data['Directory'] == "":
            self.showWarningMessage("No input 'Directory'")
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

    def readCookie(self):
        filename = QFileDialog.getOpenFileName(self, 'Open file', '',"Text files (*.txt)")
        filename = filename[0]
        if filename != "":
            try:
                f = open(filename, "r")
                data=f.read()
                f.close()
                self.data["cookie"] = json.loads(data)
            except:
                self.showErrorMessage("Can't read the cookie data")
                
    def openDir(self):
        fname = QFileDialog.getExistingDirectory(self)
        self.ui.lineEdit_directory.setText(str(fname))

if __name__ == "__main__":
    app = QApplication(sys.argv)
    window = MainWindow()
    sys.exit(app.exec_())
