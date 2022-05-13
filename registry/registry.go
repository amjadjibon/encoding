package registry

import (
	"sync"

	"github.com/amjadjibon/encoding/iface"
)

var (
	encodingRegistryOnce sync.Once
	encodingRegistryIns  *encodingRegistry
)

type encodingRegistry struct {
	registry map[string]iface.IEncoding
}

func (g *encodingRegistry) setup() {
	g.registry = make(map[string]iface.IEncoding)
}

func (g *encodingRegistry) AddEncoding(name string, encoding iface.IEncoding) {
	g.registry[name] = encoding
}
func (g *encodingRegistry) GetEncoding(name string) iface.IEncoding {
	return g.registry[name]
}

func EncodingRegistry() *encodingRegistry {
	return encodingRegistryIns
}

func init() {
	encodingRegistryOnce.Do(func() {
		encodingRegistryIns = &encodingRegistry{}
		encodingRegistryIns.setup()
	})
}
