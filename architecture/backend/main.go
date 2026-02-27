package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
    "sync"
)

// simple in-memory project descriptions; real projects would live in
// their own packages/services.
var projects = map[string]string{
    "dashboard":          "Placeholder for Cloud Infrastructure Monitoring Dashboard",
    "microservices":      "Placeholder for Microservices CI/CD Pipeline Demo",
    "aerospace":          "Aerospace Data Visualization Tool (fetch TLE by satellite name)",
    "security":           "Placeholder for GCP Security Compliance Auditor",
    "edge_compute_iot":   "Placeholder for Edge Compute IoT Platform",
    "devsecops_chatbot":  "Placeholder for DevSecOps Chatbot",
    "self_healing_infra": "Placeholder for Self‑Healing Infrastructure Demo",
    "serverless_graphql": "Placeholder for Serverless GraphQL API",
    "ai_code_review":     "Placeholder for AI‑Assisted Code Review Tool",
    "cost_prediction":    "Placeholder for Cloud Cost Prediction Model",
}

func withCORS(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        if r.Method == http.MethodOptions {
            w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
            w.WriteHeader(http.StatusOK)
            return
        }
        h(w, r)
    }
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "ok")
}

// list all available projects
func listProjects(w http.ResponseWriter, r *http.Request) {
    names := make([]string, 0, len(projects))
    for k := range projects {
        names = append(names, k)
    }
    json.NewEncoder(w).Encode(names)
}

// return info for a single project
// aerospace cache for satellite data; stores both JSON and TLE strings

type aeroEntry struct {
    Data string `json:"data"`
    TLE  string `json:"tle"`
}

var aeroCache = struct {
    sync.RWMutex
    Data map[string]aeroEntry
}{Data: make(map[string]aeroEntry)}

func loadCache() {
    bs, err := os.ReadFile("aerospace_cache.json")
    if err != nil {
        return
    }
    _ = json.Unmarshal(bs, &aeroCache.Data)
}

func saveCache() {
    aeroCache.RLock()
    defer aeroCache.RUnlock()
    bs, _ := json.Marshal(aeroCache.Data)
    _ = os.WriteFile("aerospace_cache.json", bs, 0644)
}

func projectHandler(w http.ResponseWriter, r *http.Request) {
    // expected path /projects/{name}
    parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
    if len(parts) != 2 || parts[0] != "projects" {
        http.NotFound(w, r)
        return
    }
    name := parts[1]
    desc, ok := projects[name]
    if !ok {
        http.NotFound(w, r)
        return
    }
    json.NewEncoder(w).Encode(map[string]string{
        "name":        name,
        "description": desc,
    })
}

func aerospaceDataHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        http.Error(w, "name required", http.StatusBadRequest)
        return
    }
    aeroCache.RLock()
    entry, ok := aeroCache.Data[name]
    aeroCache.RUnlock()
    if !ok || entry.Data == "" {
        http.NotFound(w, r)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(entry.Data))
}

func aerospaceFetchHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        http.Error(w, "name required", http.StatusBadRequest)
        return
    }
    url := fmt.Sprintf("https://celestrak.org/NORAD/elements/gp.php?NAME=%s&FORMAT=JSON-PRETTY", name)
    resp, err := http.Get(url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    aeroCache.Lock()
    entry := aeroCache.Data[name]
    entry.Data = string(body)
    aeroCache.Data[name] = entry
    aeroCache.Unlock()
    saveCache()
    w.Header().Set("Content-Type", "application/json")
    w.Write(body)
}

// new handlers for TLE
func aerospaceTLEHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        http.Error(w, "name required", http.StatusBadRequest)
        return
    }
    aeroCache.RLock()
    entry, ok := aeroCache.Data[name]
    aeroCache.RUnlock()
    if !ok || entry.TLE == "" {
        http.NotFound(w, r)
        return
    }
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte(entry.TLE))
}

func aerospaceTLEFetchHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        http.Error(w, "name required", http.StatusBadRequest)
        return
    }
    url := fmt.Sprintf("https://celestrak.org/NORAD/elements/gp.php?NAME=%s&FORMAT=TLE", name)
    resp, err := http.Get(url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    aeroCache.Lock()
    entry := aeroCache.Data[name]
    entry.TLE = string(body)
    aeroCache.Data[name] = entry
    aeroCache.Unlock()
    saveCache()
    w.Header().Set("Content-Type", "text/plain")
    w.Write(body)
}

func main() {
    loadCache()
    mux := http.NewServeMux()
    mux.HandleFunc("/health", withCORS(healthHandler))
    mux.HandleFunc("/projects", withCORS(listProjects))
    mux.HandleFunc("/projects/", withCORS(projectHandler))
    mux.HandleFunc("/projects/aerospace/data", withCORS(aerospaceDataHandler))
    mux.HandleFunc("/projects/aerospace/fetch", withCORS(aerospaceFetchHandler))
    mux.HandleFunc("/projects/aerospace/tle", withCORS(aerospaceTLEHandler))
    mux.HandleFunc("/projects/aerospace/tle/fetch", withCORS(aerospaceTLEFetchHandler))
    mux.HandleFunc("/", withCORS(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, `{"message":"Hello from the shared Go backend template"}`)
    }))

    http.ListenAndServe(":8080", mux)
}
