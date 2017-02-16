package account

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type CommandSpec struct {
	Cmd     string     `yaml:"cmd"`
	Flags   []FlagSpec `yaml:"flags"`
	Length  int        `yaml:"length"`
	Args    []ArgSpec  `yaml:"args"`
	Options []string   `yaml:"options"`
}

type FlagSpec struct {
	Reset []FlagInfo `yaml:"reset"`
	Help  []FlagInfo `yaml:"help"`
}

type FlagInfo struct {
	Value    string `yaml:"value"`
	Required bool   `yaml:"required"`
	Info     string `yaml:"info"`
}

type ArgSpec struct {
	Accname         []ArgInfo `yaml:"accname"`
	Email           []ArgInfo `yaml:"email"`
	Password        []ArgInfo `yaml:"password"`
	NewPassword     []ArgInfo `yaml:"newPwd"`
	ConfirmPassword []ArgInfo `yaml:"confirmPwd"`
}

type ArgInfo struct {
	//Validator string `yaml:"validator"`
	Response string `yaml:"response"`
	//Error     string `yaml:"error"`
}

var dir = "./cmd/amp/account/commands"

func ParseFile() ([]CommandSpec, error) {
	cmdSpecArr, err := parse(dir)
	if err != nil {
		fmt.Println("Error in ParseFile :: ", err)
		return nil, err
	}
	return cmdSpecArr, nil
}

func parse(directory string) ([]CommandSpec, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println("Error here :: ", err)
		return nil, err
	}
	var commands []CommandSpec
	for _, file := range files {
		cmdArr, err := generateCmdSpecs(path.Join(directory, file.Name()))
		if err != nil {
			fmt.Println("Error while generateCmdSpecs :: ", err)
			return nil, err
		}
		if cmdArr != nil {
			for _, cmd := range cmdArr {
				commands = append(commands, cmd)
			}
		}
	}
	return commands, nil
}

func generateCmdSpecs(fileName string) ([]CommandSpec, error) {
	var cmdArr []CommandSpec
	if filepath.Ext(fileName) != ".yml" {
		return nil, fmt.Errorf("cannot parse non-yaml file: %s", fileName)
	}
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("unable to read yaml command spec: %s. Error: %v", fileName, err)
	}
	var commandMap []CommandSpec
	if err = yaml.Unmarshal(contents, &commandMap); err != nil {
		return nil, fmt.Errorf("unable to unmarshal yaml command spec: %s. Error: %v", fileName, err)
	}
	for _, command := range commandMap {
		cmdArr = append(cmdArr, command)
	}
	return cmdArr, nil
}
