from fastapi import FastAPI

app = FastAPI()


@app.get("/")
async def root():
    return {'message':
    'Hello, passhold is an API implementation of a password manager'}