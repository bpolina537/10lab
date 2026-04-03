from fastapi import FastAPI, HTTPException
import httpx

app = FastAPI()

GO_SERVICE_URL = "http://localhost:8080/process"

@app.get("/ping")
async def ping():
    return {"message": "pong from FastAPI"}

@app.post("/call-go")
async def call_go(data: dict):
    async with httpx.AsyncClient() as client:
        try:
            response = await client.post(
                GO_SERVICE_URL,
                json={"data": data.get("data", "")},
                timeout=5.0
            )
            response.raise_for_status()
            return response.json()
        except httpx.RequestError:
            raise HTTPException(status_code=503, detail="Go service unavailable")
        except httpx.HTTPStatusError:
            raise HTTPException(status_code=502, detail="Go service error")