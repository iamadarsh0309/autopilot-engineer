from fastapi import FastAPI, HTTPException, Request
import uvicorn
from core.analyzer import analyze_text
from pydantic import BaseModel


app = FastAPI(title="LangGraph Engine", version="1.0")

class AnalyzeRequest(BaseModel):
    text: str


@app.get("/health")
def health():
    return {"status": "ok", "service": "langgraph-engine"}

@app.post("/analyze")
def analyze(req: AnalyzeRequest):
    try:
        result = analyze_text(req.dict())
        if "error" in result:
            raise HTTPException(status_code=400, detail=result["error"])
        return {"engine_response": result}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8001)