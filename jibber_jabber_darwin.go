package jibber_jabber

import (
	"os/exec"
	"strings"
)

func DetectLanguage() (language string, err error) {
	out, err := exec.Command("bash", "-c", "osascript -e 'user locale of (get system info)'").Output()
	if err != nil {
		return 
	}
	outstring := string(out[:])
	s := strings.Split(outstring, "_")
	language = s[0]
	return 
}
