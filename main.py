import hashlib

def getHash(input, zeroCount):
    zeros = "0" * zeroCount
    postfix = 0
    while True:
        postfix  += 1
        stri = input + str(postfix)
        encodedHash = hashlib.sha256(stri.encode()).hexdigest()
        if encodedHash.startswith(zeros):
            return { "postfix": postfix, "hash": encodedHash };

# GET THESE VALUES FROM THE RESPONSE WHEN FETCHING
# THE CHALLENGE AT .../serviceapi/pow/challenge/...
input = "f02b931c-52f0-4507-9406-f1221678dc16"
zeroCount = 4
# RETURNS THE POSTFIX AND HASH SOLUTION
print(getHash(input, zeroCount))