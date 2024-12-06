import time
import uuid

from fastapi import FastAPI, HTTPException

app = FastAPI()

qr_store = {}

QR_CODE_EXPIRATION = 10


@app.get('/qr-code')
def qr_code():
    qr_code_uuid = str(uuid.uuid4())[:10]
    expires_at = time.time() + QR_CODE_EXPIRATION
    qr_store[qr_code_uuid] = expires_at
    return {'code': qr_code_uuid, 'expires_at': expires_at}


@app.get('/verify/{qr_code_uuid}')
def verify_qr_code(qr_code_uuid: str):
    if qr_code_uuid not in qr_store:
        raise HTTPException(status_code=400, detail='Invalid or expired QR code')

    expires_at = qr_store[qr_code_uuid]
    if time.time() > expires_at:
        del qr_store[qr_code_uuid]
        raise HTTPException(status_code=400, detail='QR Code has expired')

    del qr_store[qr_code_uuid]
    return {'message': 'QR Code is valid'}
