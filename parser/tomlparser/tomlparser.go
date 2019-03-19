package tomlparser

import (
	"io"

	"github.com/BurntSushi/toml"
)

// Parse parses the toml data sent over a Reader and marshalls it into the interface
func Parse(r io.Reader, v interface{}) (interface{}, error) {
	data, err := toml.DecodeReader(r, v)
	if err != nil {
		return nil, err
	}

	return data, nil
}
