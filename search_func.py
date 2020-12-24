import requests
from bs4 import BeautifulSoup
import urllib
import json

def searchKakaoId(search):
    response = requests.get("https://page.kakao.com/search?word=" + urllib.parse.quote(search.encode('utf8'))).text
    soup = BeautifulSoup(response, 'html.parser')
    items = soup.find_all("a", attrs={"class": "css-4cffwv"})
    if len(items) < 1:
        return None
    else:
        return items[0].get("href").split("?")[1].split("=")[1]

def searchNaverId(search):
    response = requests.get("https://comic.naver.com/search.nhn?keyword=" + urllib.parse.quote(search.encode('utf8'))).text
    soup = BeautifulSoup(response, 'html.parser')
    try:
        items = soup.select(".resultList")[0].li.h5.a.get("href")
        return items.split("?")[1].split("=")[1]
    except Exception:
        return None

def searchLezhinId(search):
    headers = {'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36'}
    res = requests.get("https://dondog.lezhin.com/search?&v=2&type=comic&lang=ko&q="+ urllib.parse.quote(search.encode('utf8')), headers=headers).text
    items = json.loads(res)["sections"][0]["items"]
    if len(items) < 1:
        return None
    else:
        return items[0]["alias"]

def searchDaumId(search):
    headers = {'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36'}
    res = requests.get("http://webtoon.daum.net/data/pc/search/suggest?q="+ urllib.parse.quote(search.encode('utf8')), headers=headers).text
    items = json.loads(res)["data"]
    if len(items) < 1:
        return None
    else:
        return items[0]["nickname"]

if __name__ == "__main__":
    print(searchNaverId(""))