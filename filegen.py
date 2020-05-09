import os
import random

def genSizeFile(fileName, fileSize):
    filePath="Data"+fileName+".csv"

    ds=0
    with open(filePath, "w", encoding="utf-8") as f:
        while ds<fileSize:
            f.write(str(round(random.uniform(-1000, 1000),2)))
            f.write("\n")
            ds=os.path.getsize(filePath)
    print(os.path.getsize(filePath))


#genSizeFile("10M",10*1024*1024)
#genSizeFile("1M",1024*1024)
#genSizeFile("5M",5*1024*1024)
#genSizeFile("50M",50*1024*1024)
#genSizeFile("100M",100*1024*1024)
genSizeFile("200M",200*1024*1024)
genSizeFile("300M",300*1024*1024)  
