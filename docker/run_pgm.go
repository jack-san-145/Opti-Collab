package docker

import (
	"fmt"
	"os/exec"
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

	// Create the bash script: write code to file, then execute it
	cmdStr := fmt.Sprintf(`docker exec -i %s bash -c $'cat << "EOF" > %s
%s
EOF
%s'`, containerName, fileName, code, runCmd)

	cmd := exec.Command("bash", "-c", cmdStr)
	output, err := cmd.CombinedOutput() // captures both stdout + stderr

	result := string(output)

	// Log everything (optional for debugging)
	fmt.Printf("\n[RunCode] Language: %s\nCommand: %s\nOutput:\n%s\n", language, cmdStr, result)

	// Always return combined output; if error occurs, return it as well
	if err != nil {
		// Append the error message if execution failed (e.g. non-zero exit)
		result += fmt.Sprintf("\n[error] %v", err)
	}

	return result, nil
}
