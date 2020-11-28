# -*- coding: utf-8 -*-

################################################################################
## Form generated from reading UI file 'webtoon-downloader-1XkGGIL.ui'
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
        MainWindow.resize(341, 241)
        self.actionRead_Cookie_Data = QAction(MainWindow)
        self.actionRead_Cookie_Data.setObjectName(u"actionRead_Cookie_Data")
        self.actionVersion = QAction(MainWindow)
        self.actionVersion.setObjectName(u"actionVersion")
        self.centralwidget = QWidget(MainWindow)
        self.centralwidget.setObjectName(u"centralwidget")
        self.formLayoutWidget = QWidget(self.centralwidget)
        self.formLayoutWidget.setObjectName(u"formLayoutWidget")
        self.formLayoutWidget.setGeometry(QRect(20, 10, 301, 152))
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

        self.label_6 = QLabel(self.formLayoutWidget)
        self.label_6.setObjectName(u"label_6")
        self.label_6.setFont(font)

        self.formLayout_2.setWidget(4, QFormLayout.LabelRole, self.label_6)

        self.horizontalLayout_2 = QHBoxLayout()
        self.horizontalLayout_2.setObjectName(u"horizontalLayout_2")
        self.lineEdit_directory = QLineEdit(self.formLayoutWidget)
        self.lineEdit_directory.setObjectName(u"lineEdit_directory")
        self.lineEdit_directory.setFrame(True)

        self.horizontalLayout_2.addWidget(self.lineEdit_directory)

        self.pushButton_open = QPushButton(self.formLayoutWidget)
        self.pushButton_open.setObjectName(u"pushButton_open")
        self.pushButton_open.setFont(font)

        self.horizontalLayout_2.addWidget(self.pushButton_open)


        self.formLayout_2.setLayout(4, QFormLayout.FieldRole, self.horizontalLayout_2)

        self.lineEdit_id = QLineEdit(self.formLayoutWidget)
        self.lineEdit_id.setObjectName(u"lineEdit_id")

        self.formLayout_2.setWidget(1, QFormLayout.FieldRole, self.lineEdit_id)

        self.pushButton_download = QPushButton(self.centralwidget)
        self.pushButton_download.setObjectName(u"pushButton_download")
        self.pushButton_download.setGeometry(QRect(20, 170, 301, 31))
        self.pushButton_download.setFont(font)
        MainWindow.setCentralWidget(self.centralwidget)
        self.menubar = QMenuBar(MainWindow)
        self.menubar.setObjectName(u"menubar")
        self.menubar.setGeometry(QRect(0, 0, 341, 21))
        self.menuSetting = QMenu(self.menubar)
        self.menuSetting.setObjectName(u"menuSetting")
        self.menuVersion = QMenu(self.menubar)
        self.menuVersion.setObjectName(u"menuVersion")
        MainWindow.setMenuBar(self.menubar)
        self.statusbar = QStatusBar(MainWindow)
        self.statusbar.setObjectName(u"statusbar")
        MainWindow.setStatusBar(self.statusbar)

        self.menubar.addAction(self.menuSetting.menuAction())
        self.menubar.addAction(self.menuVersion.menuAction())
        self.menuSetting.addAction(self.actionRead_Cookie_Data)
        self.menuVersion.addAction(self.actionVersion)

        self.retranslateUi(MainWindow)

        QMetaObject.connectSlotsByName(MainWindow)
    # setupUi

    def retranslateUi(self, MainWindow):
        MainWindow.setWindowTitle(QCoreApplication.translate("MainWindow", u"Webtoon Downloader", None))
        self.actionRead_Cookie_Data.setText(QCoreApplication.translate("MainWindow", u"Read Cookie Data", None))
        self.actionVersion.setText(QCoreApplication.translate("MainWindow", u"Version", None))
        self.label.setText(QCoreApplication.translate("MainWindow", u"Type", None))
        self.comboBox_webtoon_type.setItemText(0, QCoreApplication.translate("MainWindow", u"Naver Comic", None))
        self.comboBox_webtoon_type.setItemText(1, QCoreApplication.translate("MainWindow", u"Kakao Page", None))
        self.comboBox_webtoon_type.setItemText(2, QCoreApplication.translate("MainWindow", u"Daum Webtoon", None))
        self.comboBox_webtoon_type.setItemText(3, QCoreApplication.translate("MainWindow", u"Lezhin Comics", None))

        self.label_2.setText(QCoreApplication.translate("MainWindow", u"ID", None))
        self.label_4.setText(QCoreApplication.translate("MainWindow", u"Episodes", None))
        self.label_5.setText(QCoreApplication.translate("MainWindow", u"~", None))
        self.label_3.setText(QCoreApplication.translate("MainWindow", u"DeviceID", None))
        self.label_6.setText(QCoreApplication.translate("MainWindow", u"Directory", None))
        self.pushButton_open.setText(QCoreApplication.translate("MainWindow", u"Open", None))
        self.pushButton_download.setText(QCoreApplication.translate("MainWindow", u"Download", None))
        self.menuSetting.setTitle(QCoreApplication.translate("MainWindow", u"Setting", None))
        self.menuVersion.setTitle(QCoreApplication.translate("MainWindow", u"About", None))
    # retranslateUi

