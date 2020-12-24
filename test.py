def divZero():
    try:
        1/0
    except ZeroDivisionError:
        print("error")
try:
    divZero()
except Exception as e: 
    print(str(e))