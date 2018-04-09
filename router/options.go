package router

import (
	"github.com/divisionone/go-api"
	"github.com/divisionone/go-micro/cmd"
	"github.com/divisionone/go-micro/registry"
)

type Options struct {
	Namespace string
	Handler   api.Handler
	Registry  registry.Registry
}

type Option func(o *Options)

func newOptions(opts ...Option) Options {
	options := Options{
		Namespace: "go.micro.api",
		Handler:   api.Default,
		Registry:  *cmd.DefaultOptions().Registry,
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

func WithHandler(h api.Handler) Option {
	return func(o *Options) {
		o.Handler = h
	}
}

func WithNamespace(ns string) Option {
	return func(o *Options) {
		o.Namespace = ns
	}
}

func WithRegistry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}
