package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vmware/harbor/tests/apitests/apilib"
	"testing"
	"time"
)

func TestAddProject(t *testing.T) {

	fmt.Println("\nTesting Add Project(ProjectsPost) API")
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
	fmt.Printf("\n")

}

func TestProGet(t *testing.T) {
	fmt.Println("\nTest for Project GET API")
	assert := assert.New(t)

	apiTest := newHarborAPI()
	var result []apilib.SearchProject

	//----------------------------case 1 : Response Code=200----------------------------//
	fmt.Println("case 1: respose code:200")
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

	//----------------------------case 2 : Response Code=401:is_public=0----------------------------//
	fmt.Println("case 2: respose code:401,isPublic = 0")
	httpStatusCode, result, err = apiTest.ProjectsGet("library", 0)
	if err != nil {
		t.Error("Error while search project by proName and isPublic", err.Error())
		t.Log(err)
	} else {
		assert.Equal(int(401), httpStatusCode, "httpStatusCode should be 200")
	}
	fmt.Printf("\n")
}

func TestToggleProjectPublicity(t *testing.T) {
	fmt.Println("\nTest for Project PUT API: Update properties for a selected project")
	assert := assert.New(t)

	apiTest := newHarborAPI()

	admin := &usrInfo{"admin", "Harbor12345"}
	prjUsr := &usrInfo{"unknown", "unknown"}

	//-------------------case1: Response Code=200------------------------------//
	fmt.Println("case 1: respose code:200")
	httpStatusCode, err := apiTest.ToggleProjectPublicity(*admin, "1", true)
	if err != nil {
		t.Error("Error while search project by proId", err.Error())
		t.Log(err)
	} else {
		assert.Equal(int(200), httpStatusCode, "httpStatusCode should be 200")
	}
	//-------------------case2: Response Code=401 User need to log in first. ------------------------------//
	fmt.Println("case 2: respose code:401, User need to log in first.")
	httpStatusCode, err = apiTest.ToggleProjectPublicity(*prjUsr, "1", true)
	if err != nil {
		t.Error("Error while search project by proId", err.Error())
		t.Log(err)
	} else {
		assert.Equal(int(401), httpStatusCode, "httpStatusCode should be 401")
	}
	//-------------------case2: Response Code=400 Invalid project id------------------------------//
	fmt.Println("case 3: respose code:400, Invalid project id")
	httpStatusCode, err = apiTest.ToggleProjectPublicity(*admin, "cc", true)
	if err != nil {
		t.Error("Error while search project by proId", err.Error())
		t.Log(err)
	} else {
		assert.Equal(int(400), httpStatusCode, "httpStatusCode should be 400")
	}
	//-------------------case4: Response Code=404 Not found the project------------------------------//
	fmt.Println("case 4: respose code:404, Not found the project")
	httpStatusCode, err = apiTest.ToggleProjectPublicity(*admin, "0", true)
	if err != nil {
		t.Error("Error while search project by proId", err.Error())
		t.Log(err)
	} else {
		assert.Equal(int(404), httpStatusCode, "httpStatusCode should be 404")
	}

	fmt.Printf("\n")
}
func TestProjectLogsFilter(t *testing.T) {
	fmt.Println("\nTest for search access logs filtered by operations and date time ranges..")
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
	fmt.Printf("\n")
}
