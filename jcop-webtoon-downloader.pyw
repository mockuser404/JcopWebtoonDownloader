import tkinter
from tkinter import filedialog
import requests
import os
import json
from tkinter import ttk
from bs4 import BeautifulSoup
from tkinter import messagebox

window = tkinter.Tk()
window.title("Webtoon Downloader")
window.geometry("330x200")

headers = {'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36'}

#Frame
Frame_1 = tkinter.Frame()
Frame_1.pack()

Frame_2 = tkinter.Frame()
Frame_2.pack()

#
Path = ""
TypeOfWebtoon = tkinter.StringVar()
Start = tkinter.StringVar()
End = tkinter.StringVar()
Result = tkinter.StringVar()
Id = tkinter.StringVar()
DeviceId = tkinter.StringVar()

cookie = {}

def isInputEmtpy():
    if Id.get() == "":
        Result.set("Error - no input 'Id'")
        return True

    elif Path == "":
        Result.set("Error - no input 'Directory'")
        return True

    elif Start.get() == "":
        Result.set("Error - no input 'Start Page'")
        return True

    elif End.get() == "":
        Result.set("Error - no input 'End Page'")
        return True

    else:
        return False

def KakaoGetImgURL(productId):
    global headers, cookie
    result = requests.post("https://api2-page.kakao.com/api/v1/inven/get_download_data/web",
        data={
        "productId":productId,
        "deviceId": DeviceId.get()
        },
        cookies = cookie,
        headers=headers
    ).text
    output = []
    if json.loads(result).get("downloadData") != None:
        imgs = json.loads(result).get("downloadData").get("members").get("files")
        for i in imgs:
            output.append(i["secureUrl"])
    return output

def KakaoGetTitlesURL(Id):
    output = []
    c=0
    while True:
        ids = []
        result = requests.post("https://api2-page.kakao.com/api/v5/store/singles",
            data={
            "seriesid":Id,
            "page":str(c)
            },
        ).text
        imgs = json.loads(result).get("singles")
        if imgs != None:
            for i in imgs:
                ids.append(i["id"])
        if ids == []:
            break
        output.extend(ids)
        c+=1
    return output
 
def KakaoDownload():
    global headers, Path, cookie

    if isInputEmtpy():
        return
    try:
        if not os.path.isdir(os.path.join(Path, Id.get())):
            os.mkdir(os.path.join(Path, Id.get()))

        ids = KakaoGetTitlesURL(Id.get())
        
        for i in range(int(Start.get()), int(End.get())+1):
            c=1 #for counting
            
            try:
                response = KakaoGetImgURL(str(ids[i-1]))
            except:
                Result.set("Error - Wrong Webtoon Id")
                return 
            if response == []:
                Result.set("Error - Can't get images")
                return 
            if not os.path.isdir(os.path.join(Path, Id.get(), str(i))):
                os.mkdir(os.path.join(Path, Id.get(), str(i)))
            for anchor in response:
                img_data = requests.get("http://page-edge-jz.kakao.com/sdownload/resource/"+anchor, headers=headers,cookies = cookie).content
                with open(f'{Path}\\{Id.get()}\\{i}\\{c}.jpg', 'wb') as handler:
                    c+=1
                    handler.write(img_data)
                
            with open(f"{Path}\\{Id.get()}\\{i}\\{i}.html","w") as htmlFile:
                content = "<html><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><style>body, html{margin: 0;border: 0;padding: 0;}@media only screen and (max-width: 700px) {img {width: 100%;}}</style><title>Episode "+str(i)+" ("+str(Id.get())+")</title></head><body><center>"
                for j in range(1,c):
                    content += f"<img src='{j}.jpg'><br>"
                content += "</body></center></html>"
                htmlFile.write(content)
            
    except Exception as e: 
        Result.set(e)
    else:
        Result.set("Successfully download")

def NaverDownload():
    global headers, Path, cookie

    if isInputEmtpy():
        return

    if not os.path.isdir(os.path.join(Path, Id.get())):
        os.mkdir(os.path.join(Path, Id.get()))
    try:
        for i in range(int(Start.get()),int(End.get())+1):
            c=1 #for counting

            response = requests.get(f'https://comic.naver.com/webtoon/detail.nhn?titleId={Id.get()}&no={i}', cookies = cookie).text
            soup = BeautifulSoup(response, 'html.parser')
            if soup.select('.wt_viewer img') == []:
                Result.set("Error - Can't get images")
                return 
            if not os.path.isdir(os.path.join(Path, Id.get(), str(i))):
                os.mkdir(os.path.join(Path, Id.get(), str(i)))
            for anchor in soup.select('.wt_viewer img'):
                url = anchor.get('src', '/')
                img_data = requests.get(url, headers=headers).content
                with open(f'{Path}\\{Id.get()}\\{i}\\{c}.jpg', 'wb') as handler:
                    c+=1
                    handler.write(img_data)
            
            with open(f"{Path}\\{Id.get()}\\{i}\\{i}.html","w") as htmlFile:
                content = "<html><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><style>body, html{margin: 0;border: 0;padding: 0;}@media only screen and (max-width: 700px) {img {width: 100%;}}</style><title>Episode "+str(i)+" ("+str(Id.get())+")</title></head><body><center>"
                for j in range(1,c):
                    content += f"<img src='{j}.jpg'><br>"
                content += "</body></center></html>"
                htmlFile.write(content)
    except Exception as e: 
        Result.set(e)
    else:
        Result.set("Successfully download")    

def DaumGetImgURL(productId):
    global headers
    result = requests.get("http://webtoon.daum.net/data/pc/webtoon/viewer_images/"+productId, 
        headers=headers
    ).text
    output = []
    if json.loads(result).get("data") != None:
        imgs = json.loads(result).get("data")
        for i in imgs:
            output.append(i["url"])
    return output

def DaumGetTitlesURL(Id):
    global headers
    output = []
    result = requests.post("http://webtoon.daum.net/data/pc/webtoon/view/"+Id,
        headers=headers
    ).text
    if json.loads(result).get("data") != None:
        imgs = json.loads(result).get("data").get("webtoon").get("webtoonEpisodes")
        for i in imgs:
            output.append(i["id"])
    output.sort()
    return output

def DaumDownload():
    global headers, Path

    if isInputEmtpy():
        return

    if not os.path.isdir(os.path.join(Path, Id.get())):
        os.mkdir(os.path.join(Path, Id.get()))

    ids = DaumGetTitlesURL(Id.get())
    try:
        for i in range(int(Start.get()), int(End.get())+1):
            c=1 #for counting

            try:
                response = DaumGetImgURL(str(ids[i-1]))
            except:
                Result.set("Error - Wrong Webtoon Id")
                return 
            if response == []:
                Result.set("Error - Can't get images")
                return 
            if not os.path.isdir(os.path.join(Path, Id.get(), str(i))):
                os.mkdir(os.path.join(Path, Id.get(), str(i)))
            for anchor in response:
                img_data = requests.get(str(anchor), headers=headers).content
                with open(f'{Path}\\{Id.get()}\\{i}\\{c}.jpg', 'wb') as handler:
                    c+=1
                    handler.write(img_data)
                
            with open(f"{Path}\\{Id.get()}\\{i}\\{i}.html","w") as htmlFile:
                content = "<html><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><style>body, html{margin: 0;border: 0;padding: 0;}@media only screen and (max-width: 700px) {img {width: 100%;}}</style><title>Episode "+str(i)+" ("+str(Id.get())+")</title></head><body><center>"
                for j in range(1,c):
                    content += f"<img src='{j}.jpg'><br>"
                content += "</body></center></html>"
                htmlFile.write(content)
    except Exception as e: 
        Result.set(e)
    else:
        Result.set("Successfully download")

def LezhinDownload():
    global headers, Path, cookie

    if isInputEmtpy():
        return

    if not os.path.isdir(os.path.join(Path, Id.get())):
        os.mkdir(os.path.join(Path, Id.get()))
    try:
        for i in range(int(Start.get()),int(End.get())+1):
            c=1 #for counting

            try:
                response = json.loads(requests.get(f'https://www.lezhin.com/api/v2/inventory_groups/comic_viewer_k?alias={Id.get()}&name={i}&type=comic_episode', cookies = cookie).text)["data"]["extra"]["episode"]["scrollsInfo"]
            except:
                Result.set("Error - Wrong Webtoon Id")
                return 
            if response == []:
                Result.set("Error - Can't get images")
                return 
            if not os.path.isdir(os.path.join(Path, Id.get(), str(i))):
                os.mkdir(os.path.join(Path, Id.get(), str(i)))
            for anchor in response:
                img_data = requests.get("https://cdn.lezhin.com/v2/"+anchor["path"], headers=headers, cookies = cookie).content
                with open(f'{Path}\\{Id.get()}\\{i}\\{c}.jpg', 'wb') as handler:
                    c+=1
                    handler.write(img_data)
            
            with open(f"{Path}\\{Id.get()}\\{i}\\{i}.html","w") as htmlFile:
                content = "<html><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><style>body, html{margin: 0;border: 0;padding: 0;}@media only screen and (max-width: 700px) {img {width: 100%;}}</style><title>Episode "+str(i)+" ("+str(Id.get())+")</title></head><body><center>"
                for j in range(1,c):
                    content += f"<img src='{j}.jpg'><br>"
                content += "</body></center></html>"
                htmlFile.write(content)
            
    except Exception as e: 
        Result.set(e)
    else:
        Result.set("Successfully download")

def openDir():
    global Path
    Path = filedialog.askdirectory()

def readCookie():
    global cookie
    filename = filedialog.askopenfilename() 
    if filename != "":
        try:
            f = open(filename, "r")
            data=f.read()
            f.close()
            cookie = json.loads(data)
        except:
            Result.set("Error - Can't read the cookie data")

def DownloadWebtoon():
    if TypeOfWebtoon.get() == "Kakao Page":
        KakaoDownload()
    elif TypeOfWebtoon.get() == "Naver Comic":
        NaverDownload()
    elif TypeOfWebtoon.get() == "Daum Webtoon":
        DaumDownload()
    elif TypeOfWebtoon.get() == "Lezhin Comics":
        LezhinDownload()
    else:
        Result.set("Error - invalid webtoon type")

def showVersion():
    messagebox.showinfo("Version", "1.2.2")

menubar = tkinter.Menu(window)

SettingMenu = tkinter.Menu(menubar, tearoff=0)
SettingMenu.add_command(label="Read Cookie Data", command=readCookie)
menubar.add_cascade(label="Setting", menu=SettingMenu)

menubar.add_command(label="Version", command=showVersion)
window.config(menu=menubar)


tkinter.Label(Frame_1, text="Type of Webtoon").grid(row=0,column=0)

Webtoonchoosen = ttk.Combobox(Frame_1, width = 17, textvariable = TypeOfWebtoon) 

Webtoonchoosen['values'] = ("Naver Comic","Kakao Page", "Daum Webtoon", "Lezhin Comics") 

Webtoonchoosen.grid(row = 0, column = 1) 
Webtoonchoosen.current() 


tkinter.Label(Frame_1, text="Webtoon Id").grid(row=1,column=0)

tkinter.Entry(Frame_1, textvariable=Id).grid(row=1,column=1)



#
tkinter.Label(Frame_1, text="Directory").grid(row=2,column=0)

tkinter.Button(Frame_1, text="Open", command=openDir, width=17).grid(row=2,column=1)

#
tkinter.Label(Frame_1, text="Start Page").grid(row=3,column=0)

tkinter.Entry(Frame_1, textvariable=Start).grid(row=3,column=1)

tkinter.Label(Frame_1, text="End Page").grid(row=4,column=0)

tkinter.Entry(Frame_1, textvariable=End).grid(row=4,column=1)

tkinter.Label(Frame_1, text="DeviceId (optional)").grid(row=5,column=0)

tkinter.Entry(Frame_1, textvariable=DeviceId).grid(row=5,column=1)

#
tkinter.Button(Frame_2, text="Download", command=DownloadWebtoon).grid(row=0,column=0)

tkinter.Label(Frame_2, textvariable=Result).grid(row=1,column=0)

window.mainloop()
