import Foundation

final class EmbeddingsService {
    static let shared = EmbeddingsService()
    private init() {}

    /// Creates an embedding via OpenAI embeddings endpoint.
    /// Uses KeychainHelper.loadAPIKey() to read the key saved in Settings.
    func createEmbedding(text: String, model: String = "text-embedding-3-small", completion: @escaping (Result<[Double], Error>) -> Void) {
        guard let apiKey = KeychainHelper.loadAPIKey(), !apiKey.isEmpty else {
            completion(.failure(NSError(domain: "Embeddings", code: 401, userInfo: [NSLocalizedDescriptionKey: "Missing API key"])))
            return
        }
        guard let url = URL(string: "https://api.openai.com/v1/embeddings") else {
            completion(.failure(NSError(domain: "Embeddings", code: 0, userInfo: [NSLocalizedDescriptionKey: "Bad URL"])))
            return
        }

        var req = URLRequest(url: url)
        req.httpMethod = "POST"
        req.addValue("Bearer \(apiKey)", forHTTPHeaderField: "Authorization")
        req.addValue("application/json", forHTTPHeaderField: "Content-Type")

        let payload: [String: Any] = [
            "model": model,
            "input": text
        ]

        do {
            req.httpBody = try JSONSerialization.data(withJSONObject: payload)
        } catch {
            completion(.failure(error))
            return
        }

        URLSession.shared.dataTask(with: req) { data, _, err in
            if let err = err { completion(.failure(err)); return }
            guard let data = data else {
                completion(.failure(NSError(domain: "Embeddings", code: -1, userInfo: [NSLocalizedDescriptionKey: "No data"])))
                return
            }
            do {
                if let json = try JSONSerialization.jsonObject(with: data) as? [String: Any],
                   let arr = json["data"] as? [[String: Any]],
                   let embedding = arr.first?["embedding"] as? [Double] {
                    completion(.success(embedding))
                } else {
                    completion(.failure(NSError(domain: "Embeddings", code: -2, userInfo: [NSLocalizedDescriptionKey: "Malformed response"])))
                }
            } catch {
                completion(.failure(error))
            }
        }.resume()
    }
}