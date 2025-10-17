package schedulers

import "errors"

func List() map[string]string {
	return map[string]string{
		"scx_rusty":    "ghcr.io/schedkit/scx_rusty:latest",
		"scx_rustland": "ghcr.io/schedkit/scx_rustland:latest",
		"scx_bpfland":  "ghcr.io/schedkit/scx_bpfland:latest",
		"gthulhu":      "ghcr.io/schedkit/gthulhu:latest",
	}
}

func GetScheduler(id string) (string, error) {
	var image string

	for key, value := range List() {
		if key == id {
			image = value
		}
	}

	if len(image) == 0 {
		return "", errors.New("scheduler not found")
	}

	return image, nil
}
