package yaml

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

// Yaml represents yaml file
type Yaml struct {
	data map[string]interface{}

	// For caching
	cache sync.Map
}

// Constructs a Yaml config object from the file at path
func New(path string) (*Yaml, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Unable to find yaml file at path: %v\n%v\n", path, err)
		return nil, err
	}

	// Create our data map
	data := make(map[string]interface{})

	// Unmarshal the contents into Yaml object
	err = yaml.Unmarshal(contents, &data)
	if err != nil {
		log.Printf("Unable to umarshal yaml file at path: %v\n%v\n", path, err)
		return nil, err
	}

	y := &Yaml{data: data}
	y.cache = sync.Map{}

	return y, nil
}

// Helper function for accessing dot values in yaml file
func (y *Yaml) unwrap(key string, data map[string]interface{}) (interface{}, error) {
	// Split the key if it is dot seperated
	keys := strings.Split(key, ".")

	for i, val := range keys {
		if i == len(keys)-1 {
			return data[val], nil
		}
		values := make(map[string]interface{})
		// Make sure sub key exists
		if _, ok := data[val]; !ok {
			return nil, fmt.Errorf("Key %v could not be found\n", key)
		}

		// Convert sub key interface value to map of interfaces to be iterated
		m, ok := data[val].(map[interface{}]interface{})
		if !ok {
			return nil, fmt.Errorf("Could not convert value into map for key %v", key)
		}

		for key, value := range m {
			values[key.(string)] = value
		}
		data = values
	}

	return nil, errors.New("something unexpected happened while unwrapping value")
}

func (y *Yaml) GetString(key string) (string, bool) {
	// Before we try to unwrap in yaml file lets check our cache
	if val, ok := y.cache.Load(key); ok {
		if s, ok := val.(string); ok {
			return s, true
		}
		return "", false
	}

	// Otherwise its not in our cache so unwrap the value
	val, err := y.unwrap(key, y.data)
	if err != nil {
		return "", false
	}

	// See if what we get is actually a string
	s, ok := val.(string)
	if !ok {
		return "", false
	}

	// Insert into our cache
	y.cache.Store(key, s)
	return s, true
}

func (y *Yaml) GetInt(key string) (int, bool) {
	// Before we try to unwrap in yaml file lets check our cache
	if val, ok := y.cache.Load(key); ok {
		if i, ok := val.(int); ok {
			return i, true
		}
		return 0, false
	}

	// Otherwise its not in our cache so unwrap the value
	val, err := y.unwrap(key, y.data)
	if err != nil {
		return 0, false
	}

	// See if what we get is actually an int
	i, ok := val.(int)
	if !ok {
		return 0, false
	}

	// Insert into our cache
	y.cache.Store(key, i)
	return i, true
}

func (y *Yaml) GetTimeDuration(key string) (time.Duration, bool) {
	if val, ok := y.GetString(key); ok {
		if d, err := time.ParseDuration(val); err == nil {
			return d, true
		}
	}
	return 0, false
}

func (y *Yaml) GetBool(key string) (bool, bool) {
	// Before we try to unwrap in yaml file lets check our cache
	if val, ok := y.cache.Load(key); ok {
		if b, ok := val.(bool); ok {
			return b, true
		}
		return false, false
	}

	// Otherwise its not in our cache so unwrap the value
	val, err := y.unwrap(key, y.data)
	if err != nil {
		return false, false
	}

	// See if what we get is actually a string
	b, ok := val.(bool)
	if !ok {
		return false, false
	}

	// Insert into our cache
	y.cache.Store(key, b)
	return b, true
}
