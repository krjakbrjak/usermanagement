package agent

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/krjakbrjak/usermanagement/generated"
)

func CheckPAMPwquality() error {
	cmd := exec.Command("dpkg", "-l")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run dpkg command: %v", err)
	}

	// Check if libpam-pwquality is listed in the dpkg output
	output := cmdOutput.Bytes()
	if !bytes.Contains(output, []byte("libpam-pwquality")) {
		return errors.New("make sure `libpam-pwquality` package is installed")
	}

	return nil
}

// parsePAMConfig reads and parses PAM configuration file
func ParsePAMConfig(filePath string, policy *generated.PasswordPolicyResponse) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Skip comments
		if strings.HasPrefix(line, "#") {
			continue
		}

		// Parsing lines with "pam_pwquality.so" keyword
		if strings.Contains(line, "pam_pwquality.so") {
			parts := strings.Fields(line)
			for _, part := range parts {
				if strings.HasPrefix(part, "minlen=") {
					minLength, minLengthErr := strconv.Atoi(strings.TrimPrefix(part, "minlen="))
					if minLengthErr != nil {
						return minLengthErr
					}
					policy.MinLength = int32(minLength)
				} else if strings.HasPrefix(part, "maxrepeat=") {
					maxDays, maxDaysErr := strconv.Atoi(strings.TrimPrefix(part, "maxdays="))
					if maxDaysErr != nil {
						return maxDaysErr
					}
					policy.MaxDays = int32(maxDays)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
