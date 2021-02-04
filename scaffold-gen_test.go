package main

import "testing"

func TestValidate1(t *testing.T) {
	// Test Validate Is True When All Flags Passed
	location := "location"
	name := "name"
	repository := "repository path"
	static := true
	flag := FlagSet{Location: &location, Name: &name, Repository: &repository, Static: &static}
	want := true
	if got := flag.validate(); got != want {
		t.Errorf("Invalid flags")
	}
}

func TestValidate2(t *testing.T) {
	// Test Validate Is True When Static Flags Not Passed
	location := "location"
	name := "name"
	repository := "repository path"
	static := false
	flag := FlagSet{Location: &location, Name: &name, Repository: &repository, Static: &static}
	want := true
	if got := flag.validate(); got != want {
		t.Errorf("Invalid flags")
	}
}

func TestSuccessMessage(t *testing.T) {
	// Test success message When all Flags passed
	location := "location"
	name := "name"
	repository := "repository path"
	static := true
	flag := FlagSet{Location: &location, Name: &name, Repository: &repository, Static: &static}
	want := "Generating scaffold for project name in location"
	if got := flag.successMessage(); got != want {
		t.Errorf("No generating message")
	}
}

func TestErrorMessage1(t *testing.T) {
	// Test error message When repository flags missing
	location := "location"
	name := "name"
	repository := ""
	static := true
	flag := FlagSet{Location: &location, Name: &name, Repository: &repository, Static: &static}
	want := "Project repository URL cannot be empty"
	flag.validate()
	if got := flag.errorMessage(); got != want {
		t.Errorf("Not showing missing Project repository URL error")
	}
}

func TestErrorMessage2(t *testing.T) {
	// Test error message When location flags missing
	location := ""
	name := "name"
	repository := ""
	static := true
	flag := FlagSet{Location: &location, Name: &name, Repository: &repository, Static: &static}
	want := "Project path cannot be empty"
	flag.validate()
	if got := flag.errorMessage(); got != want {
		t.Errorf("Not showing project path error")
	}
}

func TestErrorMessage3(t *testing.T) {
	// Test error message When name flags missing
	location := ""
	name := ""
	repository := ""
	static := true
	flag := FlagSet{Location: &location, Name: &name, Repository: &repository, Static: &static}
	want := "Project name cannot be empty"
	flag.validate()
	if got := flag.errorMessage(); got != want {
		t.Errorf("Not showing missing name error")
	}
}
