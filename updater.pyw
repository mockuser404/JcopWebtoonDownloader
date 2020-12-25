import ctypes, sys
import requests
import json
import wget
import pathlib
import zipfile
import os

# importing tkinter module 
from tkinter import * 
from tkinter.ttk import *
from tkinter import messagebox as mb

# creating tkinter window 
root = Tk() 
root.geometry("500x300")
root.title("Jcop Webtoon Downloader Updater")

Frame_1 = Frame()
Frame_1.pack()
# Progress bar widget 
progress = Progressbar(Frame_1, orient = HORIZONTAL, 
			length = 350, mode = 'determinate') 

log = StringVar()
# Function responsible for the updation 
# of the progress bar value 
def bar(): 
    try:
        log.set("Checking... Version")
        root.update_idletasks() 
        releases = json.loads(requests.get("https://api.github.com/repos/mynameispyo/JcopWebtoonDownloader/releases/latest").text)
        progress['value'] = 5
        root.update_idletasks() 
        if not os.path.isfile("version.txt"):
            with open("version.txt","w") as f:
                f.close()
        version =  open("version.txt", "r")
        progress['value'] = 10
        log.set("Downloading...")
        root.update_idletasks() 
    
        if releases["tag_name"] != version.read():
            version.close()
            for i in releases["assets"]:
                if i["name"] == "JcopWebtoonDownloader.zip":

                    wget.download(i["browser_download_url"],"JcopWebtoonDownloader.zip")
                    progress['value'] = 80
                    log.set("Installing...")
                    root.update_idletasks() 

                    with zipfile.ZipFile("JcopWebtoonDownloader.zip", 'r') as zip_ref:
                        zip_ref.extractall("")
                    progress['value'] = 90
                    log.set("Clearning...")
                    root.update_idletasks() 

                    os.remove("JcopWebtoonDownloader.zip")
            
                    log.set("Done")
                    progress['value'] = 100
                    mb.showinfo('Info', 'Your downloader installed successfully')
                    root.update_idletasks() 

                    return

            progress['value'] = 100
            log.set("Can't find JcopWebtoonDownloader.zip from server. Contact to Admin")
            mb.showinfo('Info', "Can't find JcopWebtoonDownloader.zip from server. Contact to Admin")
            root.update_idletasks() 
        
        else:
            mb.showinfo("Info", "You are currently using last version")
            progress['value'] = 100
            log.set("You are currently using last version")
            root.update_idletasks() 


    except Exception as e: 
        log.set("You are currently using last version")
        mb.showinfo(e)


progress.grid(row=0,column=0)

# This button will initialize 
# the progress bar 
Button(Frame_1, text = 'Start', command = bar).grid(row=0,column=1)
Label(root, textvariable=log).pack()
Frame_1.grid_columnconfigure(1, minsize=100)
# infinite loop 
mainloop() 