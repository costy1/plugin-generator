package plugin

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Plugin struct {
	Name string
}

func Generate(pluginName string) {
	plugin := Plugin{
		Name: pluginName,
	}
	createPluginBoilerplate("../../pkg/plugin/templates", plugin)
}

func createPluginBoilerplate(templatesDirectoryPath string, plugin Plugin) {
	os.RemoveAll(plugin.Name)
	err := os.MkdirAll("output/"+plugin.Name, 0777)
	if err != nil {
		panic(err)
	}
	parseAndCreateStructure(templatesDirectoryPath, "output/"+plugin.Name, plugin, true)
}

func parseAndCreateStructure(templatesDirectoryPath string, destinationPath string, plugin Plugin, firstCall bool) {
	if !firstCall {
		err := os.MkdirAll(destinationPath, 0777)
		if err != nil {
			panic(err)
		}
	}
	files, err := ioutil.ReadDir(templatesDirectoryPath)
	if err != nil {
		panic(err)
	}

	t, _ := template.ParseGlob(templatesDirectoryPath + "/*.*")
	for _, file := range files {
		if file.IsDir() {
			parseAndCreateStructure(filepath.Join(templatesDirectoryPath, file.Name()), filepath.Join(destinationPath, file.Name()), plugin, false)
		} else {
			var b bytes.Buffer
			_ = t.ExecuteTemplate(&b, file.Name(), plugin)
			fileName := file.Name()
			fileName = strings.Replace(fileName, "{{ .Name }}", plugin.Name, -1)
			os.WriteFile(filepath.Join(destinationPath, fileName), b.Bytes(), 0777)
		}
	}
}
