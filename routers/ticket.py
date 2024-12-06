import uuid
import time

from fastapi import APIRouter, HTTPException
from pydantic import BaseModel, NaiveDatetime

router = APIRouter()

qr_store = {}

QR_CODE_EXPIRATION = 10


class Ticket(BaseModel):
    name: str
    description: str
    time: NaiveDatetime
    expiredAt: NaiveDatetime


# 티켓 만들기
@router.post('')
def make(ticket: Ticket):
    print(ticket)


@router.get('/qr-code')
def qr_code():
    qr_code_uuid = str(uuid.uuid4())[:10]
    expires_at = time.time() + QR_CODE_EXPIRATION
    qr_store[qr_code_uuid] = expires_at
    return {'code': qr_code_uuid, 'expires_at': expires_at}


@router.get('/verify/{qr_code_uuid}')
def verify_qr_code(qr_code_uuid: str):
    if qr_code_uuid not in qr_store:
        raise HTTPException(status_code=400, detail='Invalid or expired QR code')

    expires_at = qr_store[qr_code_uuid]
    if time.time() > expires_at:
        del qr_store[qr_code_uuid]
        raise HTTPException(status_code=400, detail='QR Code has expired')

    del qr_store[qr_code_uuid]
    return {'message': 'QR Code is valid'}