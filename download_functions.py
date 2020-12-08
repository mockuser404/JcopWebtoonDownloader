
import requests
import os
from bs4 import BeautifulSoup
import json


class WebtoonDownload:
    def __init__(self):
        pass

    def KakaoGetImgURL(self,productId):
        result = requests.post("https://api2-page.kakao.com/api/v1/inven/get_download_data/web",
            data={
            "productId":productId,
            "deviceId": self.data['DeviceId']
            },
            cookies = self.data['cookie'],
            headers=self.data['headers']
        ).text
        output = []
        if json.loads(result).get("downloadData") != None:
            imgs = json.loads(result).get("downloadData").get("members").get("files")
            for i in imgs:
                output.append(i["secureUrl"])
        return output

    def KakaoGetTitlesURL(self,Id):
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
    
    def KakaoDownload(self):
        path = self.data['Directory']
        id = self.data['Id']
        start = self.data['Start']
        end = self.data['End']
        if self.isInputEmtpy():
            return
        try:
            if not os.path.isdir(os.path.join(path, id)):
                os.mkdir(os.path.join(path, id))

            ids = self.KakaoGetTitlesURL(id)
            
            for i in range(int(start), int(end)+1):
                c=1 #for counting
                
                try:
                    response = self.KakaoGetImgURL(str(ids[i-1]))
                except:
                    self.showErrorMessage("Wrong Webtoon Id")
                    return 
                if response == []:
                    self.showErrorMessage("Can't find images")
                    return 
                if not os.path.isdir(os.path.join(path, id, str(i))):
                    os.mkdir(os.path.join(path, id, str(i)))
                for anchor in response:
                    img_data = requests.get("http://page-edge-jz.kakao.com/sdownload/resource/"+anchor, headers=self.data['headers'],cookies = self.data['cookie']).content
                    with open(f'{path}\\{id}\\{i}\\{c}.jpg', 'wb') as handler:
                        c+=1
                        handler.write(img_data)
                    
                with open(f"{path}\\{id}\\{i}\\{i}.html","w") as htmlFile:
                    content = "<html><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><style>body, html{margin: 0;border: 0;padding: 0;}@media only screen and (max-width: 700px) {img {width: 100%;}}</style><title>Episode "+str(i)+" ("+str(id)+")</title></head><body><center>"
                    for j in range(1,c):
                        content += f"<img src='{j}.jpg'><br>"
                    content += "</body></center></html>"
                    htmlFile.write(content)
                
        except Exception as e: 
            self.showErrorMessage(str(e))
        else:
            self.showInfoMessage("Successfully download")

    def NaverDownload(self):
        path = self.data['Directory']
        id = self.data['Id']
        start = self.data['Start']
        end = self.data['End']

        if self.isInputEmtpy():
            return

        if not os.path.isdir(os.path.join(path, id)):
            try:
                os.mkdir(os.path.join(path, id))
            except:
                self.showErrorMessage("Unable Path")
        try:
            for i in range(int(start),int(end)+1):
                c=1 #for counting

                response = requests.get(f'https://comic.naver.com/webtoon/detail.nhn?titleId={id}&no={i}', cookies = self.data['cookie']).text
                soup = BeautifulSoup(response, 'html.parser')
                if soup.select('.wt_viewer img') == []:
                    self.showErrorMessage("Can't find images")
                    return 
                if not os.path.isdir(os.path.join(path, id, str(i))):
                    os.mkdir(os.path.join(path, id, str(i)))
                for anchor in soup.select('.wt_viewer img'):
                    url = anchor.get('src', '/')
                    img_data = requests.get(url, headers=self.data['headers']).content
                    with open(f'{path}\\{id}\\{i}\\{c}.jpg', 'wb') as handler:
                        c+=1
                        handler.write(img_data)
                
                with open(f"{path}\\{id}\\{i}\\{i}.html","w") as htmlFile:
                    content = "<html><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><style>body, html{margin: 0;border: 0;padding: 0;}@media only screen and (max-width: 700px) {img {width: 100%;}}</style><title>Episode "+str(i)+" ("+str(id)+")</title></head><body><center>"
                    for j in range(1,c):
                        content += f"<img src='{j}.jpg'><br>"
                    content += "</body></center></html>"
                    htmlFile.write(content)
        except Exception as e: 
            self.showErrorMessage(str(e))
        else:
            self.showInfoMessage("Successfully download")    

    def DaumGetImgURL(self,productId):
        result = requests.get("http://webtoon.daum.net/data/pc/webtoon/viewer_images/"+productId, 
            headers=self.data['headers']
        ).text
        output = []
        if json.loads(result).get("data") != None:
            imgs = json.loads(result).get("data")
            for i in imgs:
                output.append(i["url"])
        return output

    def DaumGetTitlesURL(self,Id):
        output = []
        result = requests.post("http://webtoon.daum.net/data/pc/webtoon/view/"+Id,
            headers=self.data['headers']
        ).text
        if json.loads(result).get("data") != None:
            imgs = json.loads(result).get("data").get("webtoon").get("webtoonEpisodes")
            for i in imgs:
                output.append(i["id"])
        output.sort()
        return output

    def DaumDownload(self):
        path = self.data['Directory']
        id = self.data['Id']
        start = self.data['Start']
        end = self.data['End']
        if self.isInputEmtpy():
            return

        if not os.path.isdir(os.path.join(path, id)):
            os.mkdir(os.path.join(path, id))

        ids = self.DaumGetTitlesURL(id)
        try:
            for i in range(int(start), int(end)+1):
                c=1 #for counting

                try:
                    response = self.DaumGetImgURL(str(ids[i-1]))
                except:
                    self.showErrorMessage("Wrong Webtoon Id")
                    return 
                if response == []:
                    self.showErrorMessage("Can't find images")
                    return 
                if not os.path.isdir(os.path.join(path, id, str(i))):
                    os.mkdir(os.path.join(path, id, str(i)))
                for anchor in response:
                    img_data = requests.get(str(anchor), headers=self.data['headers']).content
                    with open(f'{path}\\{id}\\{i}\\{c}.jpg', 'wb') as handler:
                        c+=1
                        handler.write(img_data)
                    
                with open(f"{path}\\{id}\\{i}\\{i}.html","w") as htmlFile:
                    content = "<html><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><style>body, html{margin: 0;border: 0;padding: 0;}@media only screen and (max-width: 700px) {img {width: 100%;}}</style><title>Episode "+str(i)+" ("+str(id)+")</title></head><body><center>"
                    for j in range(1,c):
                        content += f"<img src='{j}.jpg'><br>"
                    content += "</body></center></html>"
                    htmlFile.write(content)
        except Exception as e: 
            self.showErrorMessage(str(e))
        else:
            self.showInfoMessage("Successfully download")

    def LezhinDownload(self):
        path = self.data['Directory']
        id = self.data['Id']
        start = self.data['Start']
        end = self.data['End']
        if self.isInputEmtpy():
            return

        if not os.path.isdir(os.path.join(path, id)):
            os.mkdir(os.path.join(path, id))
        try:
            for i in range(int(start),int(end)+1):
                c=1 #for counting

                try:
                    response = json.loads(requests.get(f'https://www.lezhin.com/api/v2/inventory_groups/comic_viewer_k?alias={id}&name={i}&type=comic_episode', cookies = self.data['cookie']).text)["data"]["extra"]["episode"]["scrollsInfo"]
                except:
                    self.showErrorMessage("Wrong Webtoon Id")
                    return 
                if response == []:
                    self.showErrorMessage("Can't find images")
                    return 
                if not os.path.isdir(os.path.join(path, id, str(i))):
                    os.mkdir(os.path.join(path, id, str(i)))
                for anchor in response:
                    img_data = requests.get("https://cdn.lezhin.com/v2/"+anchor["path"]+"?access_token="+str(self.data['cookie'].get("access_token")), headers=self.data['headers'], cookies = self.data['cookie']).content
                    with open(f'{path}\\{id}\\{i}\\{c}.jpg', 'wb') as handler:
                        c+=1
                        handler.write(img_data)
                
                with open(f"{path}\\{id}\\{i}\\{i}.html","w") as htmlFile:
                    content = "<html><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><style>body, html{margin: 0;border: 0;padding: 0;}@media only screen and (max-width: 700px) {img {width: 100%;}}</style><title>Episode "+str(i)+" ("+str(id)+")</title></head><body><center>"
                    for j in range(1,c):
                        content += f"<img src='{j}.jpg'><br>"
                    content += "</body></center></html>"
                    htmlFile.write(content)
                
        except Exception as e: 
            self.showErrorMessage(str(e))
        else:
            self.showInfoMessage("Successfully download")