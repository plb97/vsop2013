// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package configuration

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
)

type Config struct {
	InputDir  string    `json:"inputDir"`
	OutputDir string    `json:"outputDir"`
	TchebNmax int       `json:"tchebNmax"`
	SQL       SqlConfig `json:"sqlConfig"`
}

//func (r Vsop2013Config) String() string {
//	return fmt.Sprintf("[%s %s %v]",InputDir, OutputDir, SQL)
//}
type SqlConfig struct {
	Host string `json:"hostname"`
	Port int    `json:"port"`
	User string `json:"username"`
	Pwd  string `json:"password"`
	Db   string `json:"database"`
}

//func (r SqlConfig) String() string {
//	return fmt.Sprintf("[%s %d %s %s %s]",Host,Port,User,Pwd,Db)
//}

func LoadConfig() Config {
	configPtr := flag.String("config", "./config.json", "config.json file path")
	flag.Parse()
	x, _ := os.Getwd()
	println("*configPtr",*configPtr,x)
	config := LoadConfigFile(*configPtr)
	return config
}

func LoadConfigFile(path string) Config {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	return config
}

var Configuration = LoadConfig()
