package config

import "gopkg.in/yaml.v2"

type EventConfig struct {
	Name     string
	Provider string
	Template yaml.MapSlice
}

type Template struct {
}
