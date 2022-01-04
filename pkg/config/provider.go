package config

type ProviderConfig struct {
	Name   string
	Type   string
	Config map[string]string
}

type Provider interface {
	Init(*ProviderConfig) error
	CreateProviderTemplate(*EventConfig) (interface{}, error)
	ProcessEvent(interface{}, *map[string]interface{}) error
}
