import random,string
def Encrypt(input,shift_int = random.randint(2,50)):
    result = ""
    alphabet =list(string.ascii_letters) + list(string.digits) + [",",'.','?',"!"," "]
    for char in input:
        try:
            result += alphabet[alphabet.index(char) + shift_int]
        except IndexError:
            result += alphabet[int(alphabet.index(char) + shift_int) -len(alphabet)]
    return result,shift_int
def Decrypt(input,shift_int = 0):
    if shift_int != 0:
        result = ""
        alphabet =list(string.ascii_letters) + list(string.digits) + [",",'.','?',"!"," "]
        for char in input:
            try:
                result += alphabet[alphabet.index(char) - shift_int]
            except IndexError:
                result += alphabet[int(alphabet.index(char) - shift_int) -len(alphabet)]
        return result
    else:
        pass
print(Encrypt("If he had anything confidential to say, he wrote it in cipher, that is, by so changing the order of the letters of the alphabet, that not a word could be made out.",2))
print(Decrypt("Khjg jcf cpavjkpi eqphkfgpvkcn vq uca, jg ytqvg kv kp ekrjgt, vjcv ku, da uq ejcpikpi vjg qtfgt qh vjg ngvvgtu qh vjg cnrjcdgv, vjcv pqv c yqtf eqwnf dg ocfg qwv.",2))