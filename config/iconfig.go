package config

type IConfig interface {
	Fill() error
}
