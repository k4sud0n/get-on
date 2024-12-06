from fastapi import FastAPI
from routers.ticket import router as ticket_router

app = FastAPI(root_path='/api/v1')

app.include_router(ticket_router, prefix='/ticket')