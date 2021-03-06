// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package vsop2013

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"fmt"
	"os"
)

type Config struct {
	InputDir  string    `json:"inputDir"`
	OutputDir string    `json:"outputDir"`
	TchebNmax int       `json:"tchebNmax"`
}
func (c Config) String() string {
	return fmt.Sprintf("config.json\n{\n\t\"inputDir\":\"%v\",\n\t\"outputDir\":\"%v\",\n\t\"tchebNmax\":%v\n}",c.InputDir,c.OutputDir,c.TchebNmax)
}

func LoadConfig() Config {
	configPtr := flag.String("config", "./config.json", "config.json.txt file path")
	flag.Parse()
	config := LoadConfigFile(*configPtr)
	return config
}

func LoadConfigFile(path string) Config {
	var config Config
	file, err := ioutil.ReadFile(path)
	if err != nil {
		const (
			tcheb = 14
			input  = "/input"
			output = "/output"
		)
		var wd string
		wd, err = os.Getwd()
		if err != nil {
			panic(err)
		}
		config = Config{InputDir:wd+input, OutputDir:wd+output,TchebNmax:tcheb}
		fmt.Println(config.String())
	} else {
		err = json.Unmarshal(file, &config)
		if err != nil {
			panic(err)
		}
	}
	return config
}

var Configuration Config = LoadConfig()
