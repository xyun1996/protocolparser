
from os import O_CREAT, O_TRUNC
import pandas as pd
url = 'https://zh.wikipedia.org/wiki/IP%E5%8D%8F%E8%AE%AE%E5%8F%B7%E5%88%97%E8%A1%A8'
tables = pd.read_html(url)
print("table number", len(tables))
protocollist = tables[2]


#print(protocollist)
fs = open("ipheaderprotocol.txt", "w+")
for index, row in protocollist.iterrows():
    fs.write(str(row[0])+"\t\t"+str(row[2])+"\n")
fs.close()