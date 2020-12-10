import ctypes, sys
import requests
import json
import wget
import pathlib


# importing tkinter module 
from tkinter import * 
from tkinter.ttk import *
from tkinter import messagebox as mb

# creating tkinter window 
root = Tk() 
root.geometry("500x300")
root.title("Jcop Webtoon Downloader Updater")
# Progress bar widget 
progress = Progressbar(root, orient = HORIZONTAL, 
			length = 400, mode = 'determinate') 

# Function responsible for the updation 
# of the progress bar value 
def bar(): 
    try:
        releases = json.loads(requests.get("https://api.github.com/repos/mynameispyo/JcopWebtoonDownloader/releases/latest").text)
        progress['value'] = 20
        root.update_idletasks() 

        version =  open("version.txt", "r")
        progress['value'] = 30
        root.update_idletasks() 
    
        if releases["tag_name"] != version.read():
            version.close()
            for i in releases["assets"]:
                if i["name"] == "jcop-webtoon-downloader.exe":
                    os.remove("jcop-webtoon-downloader.exe")
                    progress['value'] = 40
                    root.update_idletasks() 

                    wget.download(i["browser_download_url"],"jcop-webtoon-downloader.exe")
                    progress['value'] = 90
                    root.update_idletasks() 
            edit_version =  open("version.txt", "w")
            edit_version.write(releases["tag_name"])
            edit_version.close()
            progress['value'] = 100
            mb.showinfo('Info', 'Your downloader installed successfully')
            root.update_idletasks() 
        
        else:
            mb.showinfo("Info", "You are currently using last version")
            progress['value'] = 100
            root.update_idletasks() 


    except Exception as e: 
        mb.showinfo(e)

progress.pack(pady = 10) 

# This button will initialize 
# the progress bar 
Button(root, text = 'Start', command = bar).pack(pady = 10) 

# infinite loop 
mainloop() 
