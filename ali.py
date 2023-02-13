import hashlib
import ecdsa
import requests
import random

appId = "5dde4e1bdf9e4966b387ba58f4b3fdc3"
deviceId = "e5173011-XXXX-XXXX-XXXX-XXXXf04f3c62"
userId = "XXXXf36e37XXXX79bdXXXXdeb6203f67"
nonce = 0

# private_key = [175, 87, 171, 214, 222, 196, 127, 36, 25, 50, 237, 179, 71, 81, 49, 196,
#                250, 103, 115, 203, 138, 179, 192, 182, 43, 175, 233, 72, 200, 14, 64, 254]
# private_key = int.from_bytes(private_key, byteorder='big')
private_key = random.randint(1, 2**256-1)
ecc_pri = ecdsa.SigningKey.from_secret_exponent(
    private_key, curve=ecdsa.SECP256k1)
ecc_pub = ecc_pri.get_verifying_key()
public_key = "04"+ecc_pub.to_string().hex()


def r(appId: str, deviceId: str, userId: str, nonce: int) -> str:
    return f"{appId}:{deviceId}:{userId}:{nonce}"


def sign(appId, deviceId, userId, nonce) -> str:
    sign_dat = ecc_pri.sign(r(appId, deviceId, userId, nonce).encode('utf-8'), entropy=None,
                            hashfunc=hashlib.sha256)
    return sign_dat.hex()+"01"


signature = sign(appId, deviceId, userId, nonce)

headers = {
    "authorization": "Bearer BEARERBEARER",
    "origin": "https://www.aliyundrive.com",
    "referer": "https://www.aliyundrive.com/",
    "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.41",
    "x-canary": "client=web,app=adrive,version=v3.17.0",
    "x-device-id": deviceId,
    "x-signature": signature,
}

req = requests.post(
    "https://api.aliyundrive.com/users/v1/users/device/create_session",
    json={
        "deviceName": "Edge浏览器",
        "modelName": "Windows网页版",
        "pubKey": public_key,
    },
    headers=headers)

print(req.json())
# {'result': True, 'success': True, 'code': None, 'message': None}

req = requests.post("https://api.aliyundrive.com/v2/file/get_download_url",
                    json={"expire_sec": 14400, "drive_id": "341789",
                          "file_id": "63dd352b22e34327c0f84277b389eb381XXXXXXX"},
                    headers=headers)
print(req.json())
# {'domain_id': 'bj29', 'drive_id': '341789', 'file_id': '63dd352b22e34327c0f84277b389eb381XXXXXXX', 'revision_id': '63dd352ba95fd7...
