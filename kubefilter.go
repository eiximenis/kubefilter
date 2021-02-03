package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"reflect"
	"strings"
)

type ResourceDescription struct {
	Key string `yaml: "key"`
}

func main() {
	clargs := parseCommandLines()
	scanner := bufio.NewScanner(os.Stdin)
	var sb strings.Builder
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
		sb.WriteRune('\n')
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading stdin: %v", err)
		os.Exit(1)
	}

	kubeOutput := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(sb.String()), &kubeOutput)
	if err != nil {
		log.Fatalf("error parsing yaml: %v", err)
	}
	prunedOutput := pruneYaml(kubeOutput, clargs)
	finalYaml, _ := yaml.Marshal(prunedOutput)
	fmt.Printf("%s\n", string(finalYaml))
}

func pruneYaml(yaml map[interface{}]interface{}, clargs *commandLineArgs)   map[interface{}]interface{}  {
	pruneMap(yaml, "", clargs)
	return yaml
}

func pruneMap(valueMap map[interface{}]interface{}, baseKey string, clargs *commandLineArgs) {
	for k, v := range valueMap {
		fullKey := k.(string)
		if baseKey != "" {
			fullKey = baseKey + "." + fullKey
		}
		if clargs.HasDebugLevel() {
			log.Printf("Processing key %s...", fullKey)
		}
		var vtype = reflect.TypeOf(v)
		if vtype == nil {
			if clargs.removeNil {
				if clargs.HasDebugLevel()  {
					log.Printf("Removed due null value\n")
				}
				delete(valueMap, k)
			}
		} else {
			pruneValue(valueMap, v, k, fullKey, clargs)
		}
	}
}

func pruneValue(owner map[interface{}]interface{}, value interface{}, key interface{}, fullKeyName string, clargs *commandLineArgs) {
	prunned := pruneLeafValue(owner, key, fullKeyName, clargs)
	if !prunned {
		switch value.(type) {
		case map[interface{}]interface{}:
			pruneMap(value.(map[interface{}]interface{}), fullKeyName, clargs)
			break
		case []interface{}:
			// todo: Insepct slice items
			break;
		}
	}

}

func pruneLeafValue(owner map[interface{}]interface{}, key interface{}, fullKeyName string, clargs *commandLineArgs) bool {
	switch fullKeyName {
	case "metadata.managedFields", "metadata.creationTimestamp", "status", "metadata.resourceVersion", "metadata.selfLink", "metadata.uid":
		if clargs.HasDebugLevel() {
			log.Printf("Deleted")
		}
		delete(owner, key)
		return true
	case "metadata.ownerReferences":
		if clargs.removeOwnerRefs {
			delete(owner, key)
		}
		return clargs.removeOwnerRefs
	}
	_,needToDelete :=findKey(clargs.additionalKeys, fullKeyName)
	if needToDelete  {
		if clargs.HasDebugLevel() {
			log.Printf("Deleted due to remove-keys value")
		}
		delete (owner, key)
	}
	return false
}

func findKey(slice []string, key string) (int, bool) {
	for idx, item := range slice {
		if item == key {
			return idx, true
		}
	}
	return -1, false
}
