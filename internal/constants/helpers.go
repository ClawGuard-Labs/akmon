package constants

import "strings"

var AIServicePorts = map[uint16]string{
	6333:  "qdrant",
	6334:  "qdrant-grpc",
	8000:  "chromadb",
	8080:  "weaviate",
	19530: "milvus",
	9200:  "elasticsearch",
	11434: "ollama",
	8001:  "vllm",
	7860:  "gradio",
	8501:  "streamlit",
	3000:  "localai",
}

var ModelExtensions = map[string]struct{}{
	".pt": {}, ".pth": {}, ".bin": {}, ".gguf": {}, ".safetensors": {},
	".onnx": {}, ".pkl": {}, ".npz": {}, ".npy": {}, ".h5": {},
	".pb": {}, ".ckpt": {}, ".weights": {}, ".model": {}, ".tflite": {},
}


// These are process names (comm), NOT network service names:
//   - python/python3/python3.x: generic ML workload runtimes
//   - ollama, llama, llama-server, llama.cpp: local LLM servers
//   - torchrun, accelerate: PyTorch distributed training launchers
//   - vllm, tritonserver: inference servers
//   - node, nodejs, deno: JS-based AI agents
//   - cargo, rust-ai: Rust-based AI tooling
var AIProcessNames = map[string]struct{}{
	"python": {}, "python3": {}, "python3.10": {}, "python3.11": {}, "python3.12": {},
	"ollama": {}, "llama": {}, "llama-server": {}, "llama.cpp": {},
	"torchrun": {}, "accelerate": {}, "vllm": {}, "tritonserver": {},
	"node": {}, "nodejs": {}, "deno": {},
	"cargo": {}, "rust-ai": {},
}

func IsLocalhost(ip string) bool {
	return ip == "127.0.0.1" || ip == "::1" || strings.HasPrefix(ip, "127.")
}

func FileExt(path string) string {
	lastSlash := strings.LastIndex(path, "/")
	base := path[lastSlash+1:]
	lastDot := strings.LastIndex(base, ".")
	if lastDot < 0 {
		return ""
	}
	return strings.ToLower(base[lastDot:])
}

// O_WRONLY=0x1, O_RDWR=0x2, O_CREAT=0x40.
func IsWriteOpen(flags uint32) bool {
	return flags&0x1 != 0 || flags&0x2 != 0 || flags&0x40 != 0
}

func SeverityScore(sev string) int {
	switch strings.ToLower(sev) {
	case "critical":
		return 100
	case "high":
		return 70
	case "medium":
		return 40
	case "low":
		return 20
	default: // "info" or unknown
		return 10
	}
}
