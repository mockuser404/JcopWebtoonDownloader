# -*- coding: utf-8 -*-

################################################################################
## Form generated from reading UI file 'webtoon-downloaderPRoYhY.ui'
##
## Created by: Qt User Interface Compiler version 5.15.2
##
## WARNING! All changes made in this file will be lost when recompiling UI file!
################################################################################

from PySide2.QtCore import *
from PySide2.QtGui import *
from PySide2.QtWidgets import *


class Ui_MainWindow(object):
    def setupUi(self, MainWindow):
        if not MainWindow.objectName():
            MainWindow.setObjectName(u"MainWindow")
        MainWindow.resize(342, 211)
        self.actionRead_Cookie_Data = QAction(MainWindow)
        self.actionRead_Cookie_Data.setObjectName(u"actionRead_Cookie_Data")
        self.actionVersion = QAction(MainWindow)
        self.actionVersion.setObjectName(u"actionVersion")
        self.actionNaver_Comic = QAction(MainWindow)
        self.actionNaver_Comic.setObjectName(u"actionNaver_Comic")
        self.actionKakao_Page = QAction(MainWindow)
        self.actionKakao_Page.setObjectName(u"actionKakao_Page")
        self.actionLezhin_Comics = QAction(MainWindow)
        self.actionLezhin_Comics.setObjectName(u"actionLezhin_Comics")
        self.actionDaum_Webtoon = QAction(MainWindow)
        self.actionDaum_Webtoon.setObjectName(u"actionDaum_Webtoon")
        self.actionVisit_Blog = QAction(MainWindow)
        self.actionVisit_Blog.setObjectName(u"actionVisit_Blog")
        self.actionVisit_Blog_2 = QAction(MainWindow)
        self.actionVisit_Blog_2.setObjectName(u"actionVisit_Blog_2")
        self.actionDefault_Directory = QAction(MainWindow)
        self.actionDefault_Directory.setObjectName(u"actionDefault_Directory")
        self.actionFolder_Type = QAction(MainWindow)
        self.actionFolder_Type.setObjectName(u"actionFolder_Type")
        self.actionRead_Cookie_Data_v1_4 = QAction(MainWindow)
        self.actionRead_Cookie_Data_v1_4.setObjectName(u"actionRead_Cookie_Data_v1_4")
        self.actionVisit_Blog_3 = QAction(MainWindow)
        self.actionVisit_Blog_3.setObjectName(u"actionVisit_Blog_3")
        self.actionOpen_Download_Folder = QAction(MainWindow)
        self.actionOpen_Download_Folder.setObjectName(u"actionOpen_Download_Folder")
        self.centralwidget = QWidget(MainWindow)
        self.centralwidget.setObjectName(u"centralwidget")
        self.formLayoutWidget = QWidget(self.centralwidget)
        self.formLayoutWidget.setObjectName(u"formLayoutWidget")
        self.formLayoutWidget.setGeometry(QRect(20, 10, 301, 131))
        self.formLayout_2 = QFormLayout(self.formLayoutWidget)
        self.formLayout_2.setObjectName(u"formLayout_2")
        self.formLayout_2.setContentsMargins(0, 0, 0, 0)
        self.label = QLabel(self.formLayoutWidget)
        self.label.setObjectName(u"label")
        font = QFont()
        font.setFamily(u"Arial")
        font.setPointSize(12)
        self.label.setFont(font)

        self.formLayout_2.setWidget(0, QFormLayout.LabelRole, self.label)

        self.comboBox_webtoon_type = QComboBox(self.formLayoutWidget)
        self.comboBox_webtoon_type.addItem("")
        self.comboBox_webtoon_type.addItem("")
        self.comboBox_webtoon_type.addItem("")
        self.comboBox_webtoon_type.addItem("")
        self.comboBox_webtoon_type.setObjectName(u"comboBox_webtoon_type")
        self.comboBox_webtoon_type.setFont(font)

        self.formLayout_2.setWidget(0, QFormLayout.FieldRole, self.comboBox_webtoon_type)

        self.label_2 = QLabel(self.formLayoutWidget)
        self.label_2.setObjectName(u"label_2")
        self.label_2.setFont(font)

        self.formLayout_2.setWidget(1, QFormLayout.LabelRole, self.label_2)

        self.label_4 = QLabel(self.formLayoutWidget)
        self.label_4.setObjectName(u"label_4")
        self.label_4.setFont(font)

        self.formLayout_2.setWidget(2, QFormLayout.LabelRole, self.label_4)

        self.horizontalLayout = QHBoxLayout()
        self.horizontalLayout.setObjectName(u"horizontalLayout")
        self.lineEdit_start = QLineEdit(self.formLayoutWidget)
        self.lineEdit_start.setObjectName(u"lineEdit_start")

        self.horizontalLayout.addWidget(self.lineEdit_start)

        self.label_5 = QLabel(self.formLayoutWidget)
        self.label_5.setObjectName(u"label_5")
        font1 = QFont()
        font1.setFamily(u"Arial")
        font1.setPointSize(20)
        self.label_5.setFont(font1)

        self.horizontalLayout.addWidget(self.label_5)

        self.lineEdit_end = QLineEdit(self.formLayoutWidget)
        self.lineEdit_end.setObjectName(u"lineEdit_end")

        self.horizontalLayout.addWidget(self.lineEdit_end)


        self.formLayout_2.setLayout(2, QFormLayout.FieldRole, self.horizontalLayout)

        self.label_3 = QLabel(self.formLayoutWidget)
        self.label_3.setObjectName(u"label_3")
        self.label_3.setFont(font)

        self.formLayout_2.setWidget(3, QFormLayout.LabelRole, self.label_3)

        self.lineEdit_deviceid = QLineEdit(self.formLayoutWidget)
        self.lineEdit_deviceid.setObjectName(u"lineEdit_deviceid")

        self.formLayout_2.setWidget(3, QFormLayout.FieldRole, self.lineEdit_deviceid)

        self.horizontalLayout_2 = QHBoxLayout()
        self.horizontalLayout_2.setObjectName(u"horizontalLayout_2")
        self.lineEdit_id = QLineEdit(self.formLayoutWidget)
        self.lineEdit_id.setObjectName(u"lineEdit_id")

        self.horizontalLayout_2.addWidget(self.lineEdit_id)

        self.pushButton_Search = QPushButton(self.formLayoutWidget)
        self.pushButton_Search.setObjectName(u"pushButton_Search")
        icon = QIcon()
        icon.addFile(u"img/search_icon.png", QSize(), QIcon.Normal, QIcon.Off)
        self.pushButton_Search.setIcon(icon)

        self.horizontalLayout_2.addWidget(self.pushButton_Search)


        self.formLayout_2.setLayout(1, QFormLayout.FieldRole, self.horizontalLayout_2)

        self.pushButton_download = QPushButton(self.centralwidget)
        self.pushButton_download.setObjectName(u"pushButton_download")
        self.pushButton_download.setEnabled(True)
        self.pushButton_download.setGeometry(QRect(20, 140, 301, 31))
        self.pushButton_download.setFont(font)
        self.pushButton_download.setCursor(QCursor(Qt.ArrowCursor))
        self.pushButton_download.setMouseTracking(False)
        self.pushButton_download.setCheckable(False)
        MainWindow.setCentralWidget(self.centralwidget)
        self.pushButton_download.raise_()
        self.formLayoutWidget.raise_()
        self.menubar = QMenuBar(MainWindow)
        self.menubar.setObjectName(u"menubar")
        self.menubar.setGeometry(QRect(0, 0, 342, 21))
        self.menuSetting = QMenu(self.menubar)
        self.menuSetting.setObjectName(u"menuSetting")
        self.menuGenerate_Cookie_Data = QMenu(self.menuSetting)
        self.menuGenerate_Cookie_Data.setObjectName(u"menuGenerate_Cookie_Data")
        self.menuVersion = QMenu(self.menubar)
        self.menuVersion.setObjectName(u"menuVersion")
        self.menuTool = QMenu(self.menubar)
        self.menuTool.setObjectName(u"menuTool")
        MainWindow.setMenuBar(self.menubar)
        self.statusbar = QStatusBar(MainWindow)
        self.statusbar.setObjectName(u"statusbar")
        MainWindow.setStatusBar(self.statusbar)

        self.menubar.addAction(self.menuSetting.menuAction())
        self.menubar.addAction(self.menuTool.menuAction())
        self.menubar.addAction(self.menuVersion.menuAction())
        self.menuSetting.addAction(self.menuGenerate_Cookie_Data.menuAction())
        self.menuSetting.addAction(self.actionDefault_Directory)
        self.menuGenerate_Cookie_Data.addAction(self.actionNaver_Comic)
        self.menuGenerate_Cookie_Data.addAction(self.actionKakao_Page)
        self.menuGenerate_Cookie_Data.addAction(self.actionLezhin_Comics)
        self.menuVersion.addAction(self.actionVersion)
        self.menuVersion.addAction(self.actionVisit_Blog_3)
        self.menuTool.addAction(self.actionOpen_Download_Folder)

        self.retranslateUi(MainWindow)

        QMetaObject.connectSlotsByName(MainWindow)
    # setupUi

    def retranslateUi(self, MainWindow):
        MainWindow.setWindowTitle(QCoreApplication.translate("MainWindow", u"Jcop Webtoon Downloader", None))
        self.actionRead_Cookie_Data.setText(QCoreApplication.translate("MainWindow", u"Read Cookie Data", None))
        self.actionVersion.setText(QCoreApplication.translate("MainWindow", u"Version", None))
        self.actionNaver_Comic.setText(QCoreApplication.translate("MainWindow", u"Naver Comic", None))
        self.actionKakao_Page.setText(QCoreApplication.translate("MainWindow", u"Kakao Page", None))
        self.actionLezhin_Comics.setText(QCoreApplication.translate("MainWindow", u"Lezhin Comics", None))
        self.actionDaum_Webtoon.setText(QCoreApplication.translate("MainWindow", u"Daum Webtoon", None))
        self.actionVisit_Blog.setText(QCoreApplication.translate("MainWindow", u"Visit Blog", None))
        self.actionVisit_Blog_2.setText(QCoreApplication.translate("MainWindow", u"Visit Blog", None))
        self.actionDefault_Directory.setText(QCoreApplication.translate("MainWindow", u"Default Directory", None))
        self.actionFolder_Type.setText(QCoreApplication.translate("MainWindow", u"Folder Type", None))
        self.actionRead_Cookie_Data_v1_4.setText(QCoreApplication.translate("MainWindow", u"Read Cookie Data (v1.4)", None))
        self.actionVisit_Blog_3.setText(QCoreApplication.translate("MainWindow", u"Visit Blog", None))
        self.actionOpen_Download_Folder.setText(QCoreApplication.translate("MainWindow", u"Open Download Folder", None))
        self.label.setText(QCoreApplication.translate("MainWindow", u"Type", None))
        self.comboBox_webtoon_type.setItemText(0, QCoreApplication.translate("MainWindow", u"Naver Comic", None))
        self.comboBox_webtoon_type.setItemText(1, QCoreApplication.translate("MainWindow", u"Kakao Page", None))
        self.comboBox_webtoon_type.setItemText(2, QCoreApplication.translate("MainWindow", u"Daum Webtoon", None))
        self.comboBox_webtoon_type.setItemText(3, QCoreApplication.translate("MainWindow", u"Lezhin Comics", None))

        self.label_2.setText(QCoreApplication.translate("MainWindow", u"ID", None))
        self.label_4.setText(QCoreApplication.translate("MainWindow", u"Episodes", None))
        self.label_5.setText(QCoreApplication.translate("MainWindow", u"~", None))
        self.label_3.setText(QCoreApplication.translate("MainWindow", u"DeviceID", None))
        self.pushButton_Search.setText("")
        self.pushButton_download.setText(QCoreApplication.translate("MainWindow", u"Download", None))
        self.menuSetting.setTitle(QCoreApplication.translate("MainWindow", u"Setting", None))
        self.menuGenerate_Cookie_Data.setTitle(QCoreApplication.translate("MainWindow", u"Generate Cookie Data", None))
        self.menuVersion.setTitle(QCoreApplication.translate("MainWindow", u"About", None))
        self.menuTool.setTitle(QCoreApplication.translate("MainWindow", u"Tool", None))
    # retranslateUi

