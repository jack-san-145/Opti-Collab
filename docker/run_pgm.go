package docker

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// Run_code executes the given code in the specified language using Docker containers.
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
		containerName = "gcc-container"
		fileName = "main.c"
		runCmd = fmt.Sprintf("gcc %s -o main && ./main", fileName)

	case "cpp", "c++":
		containerName = "gcc-container"
		fileName = "main.cpp"
		runCmd = fmt.Sprintf("g++ %s -o main && ./main", fileName)

	case "java":
		containerName = "java-container"

		// Detect class name for Java if public class exists
		re := regexp.MustCompile(`public\s+class\s+(\w+)`)
		matches := re.FindStringSubmatch(code)

		className := "Main"
		if len(matches) > 1 {
			className = matches[1]
		} else {
			// If no public class, remove any accidental `public` usage
			code = strings.ReplaceAll(code, "public class", "class")
		}

		fileName = fmt.Sprintf("%s.java", className)
		runCmd = fmt.Sprintf("javac %s && java %s", fileName, className)

	default:
		return "", fmt.Errorf("unsupported language: %s", language)
	}

	// Escape single quotes safely for shell
	safeCode := strings.ReplaceAll(code, `'`, `'"'"'`)

	// Build the full shell command for Docker
	cmdStr := fmt.Sprintf(`
docker exec -i %s sh -c 'cat > %s <<'EOF'
%s
EOF
%s'`, containerName, fileName, safeCode, runCmd)

	// Run the command and capture output
	cmd := exec.Command("bash", "-c", cmdStr)
	output, err := cmd.CombinedOutput()
	result := string(output)

	fmt.Printf("\n[RunCode] Language: %s\nCommand:\n%s\nOutput:\n%s\n", language, cmdStr, result)

	// If command failed, append the error for debugging
	if err != nil {
		result += fmt.Sprintf("\n[error] %v", err)
	}

	return result, nil
}
