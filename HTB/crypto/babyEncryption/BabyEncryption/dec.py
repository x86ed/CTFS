import string



def decrypt(file):
    #open text file in read mode
    ct = []
    with open(file, "rb") as f:
        while (byte := f.read(1)):
            char = int(byte,16)-18
            char = 179 * char % 256
            ct.append(char)
        return bytes(ct)

ct = decrypt('./msg.enc')
f = open('./msg.txt','w')
f.write(ct.hex())
f.close()
