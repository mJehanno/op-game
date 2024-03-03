package conf

import (
	"fmt"
	"os"
	"path"
)

func GetConfFolder() (*string, error) {
	var err error
	conf, erra := os.UserConfigDir()
	if erra != nil {
		conf, err = os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("%w and %w happened", erra, err)
		}
	}

	confFolder := path.Join(conf, "mult-game")
	err = os.MkdirAll(confFolder, 0775)
	if err != nil {
		return nil, fmt.Errorf("failed to create folder in config or home directory -> %w", err)
	}

	return &confFolder, nil
}
