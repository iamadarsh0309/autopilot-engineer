from fastapi import FastAPI, HTTPException, Request
import uvicorn
from core.analyzer import analyze_text
from pydantic import BaseModel

app = FastAPI(title="LangGraph Engine", version="1.0")

class AnalyzeRequest(BaseModel):
    text: str = None
    repo_summary: dict = None


@app.get("/health")
def health():
    return {"status": "ok", "service": "langgraph-engine"}


@app.post("/analyze")
async def analyze(req: AnalyzeRequest):
    try:
        # --- Case 1: Repo analysis ---
        if req.repo_summary:
            summary = req.repo_summary
            return {
                "engine_response": f"Repo has {summary['total_files']} files and uses {list(summary['languages'].keys())}"
            }

        # --- Case 2: Text analysis ---
        if req.text:
            result = analyze_text(req.dict())
            if "error" in result:
                raise HTTPException(status_code=400, detail=result["error"])
            return {"engine_response": result}

        # --- Case 3: Invalid ---
        raise HTTPException(status_code=400, detail="Invalid input: expected 'text' or 'repo_summary'")

    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8001)
