package config

import "os/exec"

func saveConfigMacOS(value, key string) error {
	cmd := exec.Command("security", "add-generic-password", "-a", key, "-s", key, "-w", value, "-U")
	return cmd.Run()
}

func saveConfigLinux(value, key string) error {
	cmd := exec.Command("secret-tool", "store", "--label=kemit", key, value)
	return cmd.Run()
}

func loadConfigMacOS(key string) ([]byte, error) {
	cmd := exec.Command("security", "find-generic-password", "-a", key, "-s", key, "-w")
	return cmd.Output()
}

func loadConfigLinux(key string) ([]byte, error) {
	cmd := exec.Command("secret-tool", "lookup", key)
	return cmd.Output()
}
