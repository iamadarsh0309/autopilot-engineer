# langgraph-engine/core/analyzer.py
from textblob import TextBlob

def analyze_text(input_data: dict) -> dict:
    """
    Perform basic text analysis: sentiment, word count, summary.
    """
    text = input_data.get("text", "").strip()
    if not text:
        return {"error": "No text provided"}

    blob = TextBlob(text)

    # Simple NLP metrics
    sentiment = blob.sentiment.polarity
    word_count = len(text.split())

    # Tiny summary logic (for demo)
    sentences = text.split(".")
    summary = sentences[0] if sentences else text

    return {
        "summary": summary.strip(),
        "word_count": word_count,
        "sentiment_score": sentiment,
    }
