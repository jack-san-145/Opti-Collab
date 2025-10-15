package docker

import (
	"fmt"
	"os/exec"
	"strings"
)

func Run_code(language, code string) (string, error) {
	var containerName, fileName, runCmd string

	switch language {
	case "python":
		containerName = "python-container"
		fileName = "main.py"
		runCmd = fmt.Sprintf("python3 %s", fileName)
	case "javascript", "js":
		containerName = "js-container"
		fileName = "main.js"
		runCmd = fmt.Sprintf("node %s", fileName)
	case "go":
		containerName = "go-container"
		fileName = "main.go"
		runCmd = fmt.Sprintf("go run %s", fileName)
	case "c":
		containerName = "cpp-container"
		fileName = "main.c"
		runCmd = fmt.Sprintf("gcc %s -o main && ./main", fileName)
	case "cpp", "c++":
		containerName = "cpp-container"
		fileName = "main.cpp"
		runCmd = fmt.Sprintf("g++ %s -o main && ./main", fileName)
	case "java":
		containerName = "java-container"
		fileName = "Main.java"
		runCmd = fmt.Sprintf("javac %s && java Main", fileName)
	default:
		return "", fmt.Errorf("unsupported language: %s", language)
	}

	// Escape single quotes to avoid breaking the shell
	safeCode := strings.ReplaceAll(code, `'`, `'"'"'`)

	// Use safe shell heredoc with literal quoting
	cmdStr := fmt.Sprintf(
		`docker exec -i %s sh -c 'cat > %s <<'EOF'
%s
EOF
%s'`,
		containerName, fileName, safeCode, runCmd,
	)

	// Execute the command
	cmd := exec.Command("bash", "-c", cmdStr)
	output, err := cmd.CombinedOutput() // Captures both stdout + stderr

	result := string(output)

	fmt.Printf("\n[RunCode] Language: %s\nCommand: %s\nOutput:\n%s\n", language, cmdStr, result)

	// Append error if command failed (non-zero exit)
	if err != nil {
		result += fmt.Sprintf("\n[error] %v", err)
	}

	return result, nil
}
