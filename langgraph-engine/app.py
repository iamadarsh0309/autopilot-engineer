from fastapi import FastAPI, Request
import uvicorn

app = FastAPI()

@app.get("/health")
def health():
    return {"status": "ok", "service": "langgraph-engine"}

@app.post("/analyze")
async def analyze(request: Request):
    data = await request.json()
    print("[LangGraph Engine] Received data for analysis:", data)
    return {"engine_response": f"Processed data: {data} "}

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8001)