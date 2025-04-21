package cmd_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestMainCLI(t *testing.T) {
	// Skip if running in short mode
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// get the project directory
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	// binary will be in the root directory of the project
	binPath := filepath.Join(filepath.Dir(filepath.Dir(pwd)), "interview-cli")

	// test the binary exists
	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		// build the binary if it doesn't exist
		buildCmd := exec.Command("go", "build", "-o", binPath, "../cmd/interview-cli")
		buildCmd.Dir = filepath.Dir(filepath.Dir(pwd)) // set working directory to project root

		if output, err := buildCmd.CombinedOutput(); err != nil {
			t.Fatalf("Failed to build binary: %v\nOutput: %s", err, output)
		}
	}

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"Help Command", []string{"-help"}, false},
		{"List Command", []string{"-list"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binPath, tt.args...)
			output, err := cmd.CombinedOutput()

			// check error expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("Command execution error = %v, wantErr %v\nOutput: %s", err, tt.wantErr, output)
			}
		})
	}
}

