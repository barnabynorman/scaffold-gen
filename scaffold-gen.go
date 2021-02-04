package main

import (
	"flag"
	"fmt"
)

type FlagSet struct {
	Location     *string `json:"location"`
	Name         *string `json:"name"`
	Repository   *string `json:"repository"`
	Static       *bool   `json:"static"`
	ErrorMessage string
}

func (f *FlagSet) validate() bool {
	nameLen := len(*f.Name)
	if nameLen == 0 {
		f.ErrorMessage = "Project name cannot be empty"
		return false
	}

	locationLen := len(*f.Location)
	if locationLen == 0 {
		f.ErrorMessage = "Project path cannot be empty"
		return false
	}

	repositoryLen := len(*f.Repository)
	if repositoryLen == 0 {
		f.ErrorMessage = "Project repository URL cannot be empty"
		return false
	}

	return true
}

func (f *FlagSet) successMessage() string {
	return fmt.Sprintf("Generating scaffold for project %s in %s", *f.Name, *f.Location)
}

func (f *FlagSet) errorMessage() string {
	return fmt.Sprintf("%s", f.ErrorMessage)
}

func main() {
	location := flag.String("d", "", "Project location on disk")
	name := flag.String("n", "", "Project name")
	repository := flag.String("r", "", "Project remote repository URL")
	static := flag.Bool("s", false, "Project will have static assets or not")

	flag.Parse()

	flag := FlagSet{Location: location, Name: name, Repository: repository, Static: static}
	res := flag.validate()

	if res {
		msg := flag.successMessage()
		fmt.Println(msg)
	} else {
		msg := flag.errorMessage()
		fmt.Println(msg)
	}

}
