package docker

import (
	"fmt"
	"os/exec"
)

func Run_code(language, code string) (string, error) {
	// Map language to container & file info
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

	// Use cat <<EOF to write multi-line code inside container
	cmdStr := fmt.Sprintf(`docker exec -i %s bash -c $'cat << EOF > %s
%s
EOF
%s'`, containerName, fileName, code, runCmd)

	cmd := exec.Command("bash", "-c", cmdStr)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
