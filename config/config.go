package config
import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "os"
)

type Config struct {
    Host string `json:"host"`
    Port int     `json:"port"`
    Handlers []Handler `json:"handlers"`
    UseHttp2 bool `json:"useHttp2"`

}

type Handler struct {
    Method string `json:"method"`
    Path   string `json:"path"`
    Response string `json:"response"`

}


func LoadConfigFromFile(file string) (*Config,error) {
    jsonFile, err := os.Open(file)
    // if we os.Open returns an error then handle it
    if err != nil {
        return nil, err
    }
    fmt.Printf("Successfully Opened %s\n", file)
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()
    conf := &Config{} 
    byteValue, err := ioutil.ReadAll(jsonFile)
    if err != nil {
        return nil, err
    }
    json.Unmarshal(byteValue, &conf)
    fmt.Printf("%v\n", conf)
    return conf, nil

}
