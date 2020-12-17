package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	// "gopkg.in/yaml.v2"
	"sigs.k8s.io/yaml"
)

func main() {
	var yamlFile string
	flag.StringVar(&yamlFile, "f", "", "absolute path to file")
	flag.Parse()

	data, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	parsedYaml := YamlToGo{}
	if err := yaml.UnmarshalStrict(data, &parsedYaml); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// won't get here using "sigs.k8s.io/yaml"
	fmt.Printf("%#v\n", parsedYaml)
}

type YamlToGo struct {
	ArrayOfThings []Thing `yaml:"arrayOfThings,omitempty"`
}

type Thing struct {
	Name         string `yaml:"name,omitempty"`
	AttributeOne string `yaml:"attributeOne,omitempty"`
	AttributeTwo string `yaml:"attributeTwo,omitempty"`
}
