package lib

import (
	"fmt"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

/*var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`*/

// T YAML struct test
type T struct {
	A string
	B struct {
			C int   `yaml:"c"`
			D        []int `yaml:",flow"`
	}
}

// GetLocalValue test using local function
func GetLocalValue() string {
	return "Hello from this main package"
}

// TestYaml test yaml function
func TestYaml() {
	t := T{}

	yamlFile, errYamlFile := ioutil.ReadFile("./lib/test.yaml")
    if errYamlFile != nil {
        log.Printf("yamlFile.Get err   #%v ", errYamlFile)
    }
    
	err := yaml.Unmarshal(yamlFile, &t)
	if err != nil {
			log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)
	log.Printf(t.A)

	d, err := yaml.Marshal(&t)
	if err != nil {
			log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})
    
	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
			log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
			log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}