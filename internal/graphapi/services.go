package graphapi

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/clawsec/internal/constants"
)

type AIServiceInfo struct {
	Type     string `json:"type"`               // "process" or "service"
	Name     string `json:"name"`               // process comm or service name
	Category string `json:"category"`           // "llm", "agent", "training", "vector-db", "inference", "ui"
	PID      uint32 `json:"pid,omitempty"`       // for processes
	Port     uint16 `json:"port,omitempty"`      // for services
	Status   string `json:"status"`             // "running", "listening"
	Cmdline  string `json:"cmdline,omitempty"`   // full command line (truncated)
}

func (s *Server) handleServices(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var results []AIServiceInfo

	// scanning /proc 
	results = append(results, scanAIProcesses()...)

	// ports checking
	results = append(results, probeAIServicePorts()...)

	writeJSON(w, results)
}

func scanAIProcesses() []AIServiceInfo {
	entries, err := os.ReadDir("/proc")
	if err != nil {
		return nil
	}

	// Deduplicate by comm to avoid listing 50 python3 workers individually.
	// Track unique comms with their first-seen PID and total count.
	type commInfo struct {
		comm    string
		pid     uint32
		cmdline string
		count   int
	}
	seen := map[string]*commInfo{}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		pid, err := strconv.ParseUint(entry.Name(), 10, 32)
		if err != nil {
			continue
		}

		commBytes, err := os.ReadFile(fmt.Sprintf("/proc/%s/comm", entry.Name()))
		if err != nil {
			continue
		}
		comm := strings.TrimSpace(string(commBytes))

		if _, ok := constants.AIProcessNames[comm]; !ok {
			continue
		}

		if info, exists := seen[comm]; exists {
			info.count++
			continue
		}

		cmdline := readCmdline(entry.Name())

		seen[comm] = &commInfo{
			comm:    comm,
			pid:     uint32(pid),
			cmdline: cmdline,
			count:   1,
		}
	}

	results := make([]AIServiceInfo, 0, len(seen))
	for _, info := range seen {
		name := info.comm
		if info.count > 1 {
			name = fmt.Sprintf("%s (%d instances)", info.comm, info.count)
		}
		results = append(results, AIServiceInfo{
			Type:     "process",
			Name:     name,
			Category: categorizeProcess(info.comm),
			PID:      info.pid,
			Status:   "running",
			Cmdline:  info.cmdline,
		})
	}
	return results
}

func probeAIServicePorts() []AIServiceInfo {
	var results []AIServiceInfo
	for port, name := range constants.AIServicePorts {
		addr := fmt.Sprintf("127.0.0.1:%d", port)
		conn, err := net.DialTimeout("tcp", addr, 200*time.Millisecond)
		if err != nil {
			continue
		}
		conn.Close()
		results = append(results, AIServiceInfo{
			Type:     "service",
			Name:     name,
			Category: categorizeService(name),
			Port:     port,
			Status:   "listening",
		})
	}
	return results
}

func readCmdline(pidStr string) string {
	data, err := os.ReadFile(fmt.Sprintf("/proc/%s/cmdline", pidStr))
	if err != nil || len(data) == 0 {
		return ""
	}
	// cmdline is null-separated; convert to spaces and truncate 
	s := strings.ReplaceAll(string(data), "\x00", " ")
	s = strings.TrimSpace(s)
	if len(s) > 200 {
		s = s[:200] + "..."
	}
	return s
}

func categorizeProcess(comm string) string {
	switch comm {
	case "ollama", "llama", "llama-server", "llama.cpp", "vllm", "tritonserver":
		return "llm"
	case "torchrun", "accelerate":
		return "training"
	case "python", "python3", "python3.10", "python3.11", "python3.12",
		"node", "nodejs", "deno":
		return "agent"
	default:
		return "other"
	}
}

func categorizeService(name string) string {
	switch name {
	case "ollama", "vllm", "localai":
		return "llm"
	case "qdrant", "qdrant-grpc", "chromadb", "weaviate", "milvus", "elasticsearch":
		return "vector-db"
	case "gradio", "streamlit":
		return "ui"
	default:
		return "inference"
	}
}
