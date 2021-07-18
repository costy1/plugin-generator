package main

import (
	"flag"

	"github.com/costy1/plugin-generator/pkg/plugin"
)

func main() {
	var pluginName string
	flag.StringVar(&pluginName, "name", "PluginTest", "the name for the plugin")
	flag.Parse()

	plugin.Generate(pluginName)
}
