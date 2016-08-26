package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vmware/harbor/tests/apitests/apilib"
	"testing"
	"time"
)

func TestAddProject(t *testing.T) {

	fmt.Println("Testing Add Project(ProjectsPost) API")
	assert := assert.New(t)

	apiTest := newHarborAPI()

	//prepare for test

	admin := &usrInfo{"admin", "Harbor12345"}

	prjUsr := &usrInfo{"unknown", "unknown"}

	var project apilib.Project
	project.ProjectName = "test_project"
	project.Public = true

	//case 1: admin not login, expect project creation fail.

	result, err := apiTest.ProjectsPost(*prjUsr, project)
	if err != nil {
		t.Error("Error while creat project", err.Error())
		t.Log(err)
	} else {
		assert.Equal(result, int(401), "Case 1: Project creation status should be 401")
		//t.Log(result)
	}

	//case 2: admin successful login, expect project creation success.
	fmt.Println("case 2: admin successful login, expect project creation success.")

	prjUsr = admin

	result, err = apiTest.ProjectsPost(*prjUsr, project)
	if err != nil {
		t.Error("Error while creat project", err.Error())
		t.Log(err)
	} else {
		assert.Equal(result, int(201), "Case 2: Project creation status should be 201")
		//t.Log(result)
	}

	//case 3: duplicate project name, create project fail
	fmt.Println("case 3: duplicate project name, create project fail")

	result, err = apiTest.ProjectsPost(*prjUsr, project)
	if err != nil {
		t.Error("Error while creat project", err.Error())
		t.Log(err)
	} else {
		assert.Equal(result, int(409), "Case 3: Project creation status should be 409")
		//t.Log(result)
	}

}

func TestProGet(t *testing.T) {
	fmt.Println("Test for Project GET API")
	assert := assert.New(t)

	apiTest := newHarborAPI()
	var result []apilib.SearchProject

	httpStatusCode, result, err := apiTest.ProjectsGet("library", 1)
	if err != nil {
		t.Error("Error while search project by proName and isPublic", err.Error())
		t.Log(err)
	} else {
		assert.Equal(int(200), httpStatusCode, "httpStatusCode should be 200")
		assert.Equal(int64(0), result[0].Id, "Project id should be equal")
		assert.Equal("library", result[0].Name, "Project name should be library")
		assert.Equal(int32(1), result[0].Public, "Project public status should be 1 (true)")
	}

}

func TestToggleProjectPublicity(t *testing.T) {
	fmt.Println("Test for Project PUT API: Update properties for a selected project")
	assert := assert.New(t)

	apiTest := newHarborAPI()

	admin := &usrInfo{"admin", "Harbor12345"}

	//-------------------case1: Response Code=200------------------------------//
	httpStatusCode, err := apiTest.ToggleProjectPublicity(*admin, "1", true)
	if err != nil {
		t.Error("Error while search project by proId", err.Error())
		t.Log(err)
	} else {
		assert.Equal(int(200), httpStatusCode, "httpStatusCode should be 200")
	}
	//-------------------case2: Response Code=404 Not found the project------------------------------//
	httpStatusCode, err = apiTest.ToggleProjectPublicity(*admin, "0", true)
	if err != nil {
		t.Error("Error while search project by proId", err.Error())
		t.Log(err)
	} else {
		assert.Equal(int(404), httpStatusCode, "httpStatusCode should be 401")
	}
	/*
		//-------------------case3: Response Code=400 Invalid project id ------------------------------//
		httpStatusCode, result, err = apiTest.ProjectGetByPk("cc")
		if err != nil {
			t.Error("Error while search project by proId", err.Error())
			t.Log(err)
		} else {
			assert.Equal(int(400), httpStatusCode, "httpStatusCode should be 400")
		}
	*/

}
func TestProjectLogsFilter(t *testing.T) {
	fmt.Println("Test for search access logs filtered by operations and date time ranges..")
	assert := assert.New(t)

	apiTest := newHarborAPI()
	admin := &usrInfo{"admin", "Harbor12345"}
	endTimestamp := time.Now().Unix()
	startTimestamp := endTimestamp - 3600
	accessLog := &apilib.AccessLogFilter{
		Username:       "admin",
		Keywords:       "",
		BeginTimestamp: startTimestamp,
		EndTimestamp:   endTimestamp,
	}
	projectId := "1"
	httpStatusCode, _, err := apiTest.ProjectLogsFilter(*admin, projectId, *accessLog)
	if err != nil {
		t.Error("Error while search access logs")
		t.Log(err)
	} else {
		assert.Equal(int(200), httpStatusCode, "httpStatusCode should be 200")
	}

}
