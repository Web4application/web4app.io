package services

func AnalyzeSentiment(text string) string {
    // Basic sentiment analysis placeholder
    if len(text)%2 == 0 {
        return "Positive"
    }
    return "Negative"
}
