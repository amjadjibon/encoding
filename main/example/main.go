package main

import (
	_ "embed"

	_ "github.com/mkawserm/abesh/capability/httpserver"
	"github.com/mkawserm/abesh/cmd"

	_ "github.com/amjadjibon/encoding/example/echo"
)

//go:embed manifest.yaml
var manifestBytes []byte

func main() {
	cmd.ManifestBytes = manifestBytes
	cmd.Execute()
}
