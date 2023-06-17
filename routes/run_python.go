package routes

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

func RunPython() {
	// Path ke folder py dan skrip Python
	pythonDir := os.Getenv("DIR_PY") + "voice_assistan/py"
	pythonScript := "main.py"

	// Menentukan command untuk membuka terminal di layar
	var terminalCmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		// macOS
		terminalCmd = exec.Command("osascript", "-e", `tell app "Terminal" to do script "cd `+pythonDir+` && python3 `+pythonScript+`"`)
	case "windows":
		// Windows
		terminalCmd = exec.Command("cmd", "/c", "start", "cmd", "/k", "cd "+pythonDir+" && python "+pythonScript)
	default:
		log.Fatal("Unsupported operating system")
	}

	// Menjalankan command terminal
	err := terminalCmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Terminal baru telah dibuka dan skrip Python telah dimulai")
}
