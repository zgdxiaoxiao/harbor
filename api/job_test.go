package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJobsGet(t *testing.T) {
	var httpStatusCode int
	var err error

	assert := assert.New(t)
	apiTest := newHarborAPI()

	jobsParams := &JobsParams{PolicyID: "1"}

	fmt.Println("Testing Jobs Get API")

	//-------------------case 1 : response code = 200------------------------//
	fmt.Println("case 1 : response code = 200")
	httpStatusCode, err = apiTest.ListJobs(*admin, *jobsParams)
	if err != nil {
		t.Error("Error while get jobs", err.Error())
		t.Log(err)
	} else {
		assert.Equal(int(200), httpStatusCode, "httpStatusCode should be 200")
	}
	//-------------------case 2 : response code = 401------------------------//
	fmt.Println("case 2 : response code = 401,User need to login first.")
	httpStatusCode, err = apiTest.ListJobs(*unknownUsr, *jobsParams)
	if err != nil {
		t.Error("Error while get jobs", err.Error())
		t.Log(err)
	} else {
		assert.Equal(int(401), httpStatusCode, "httpStatusCode should be 401")
	}
	//-------------------case 3 : response code = 400------------------------//
	fmt.Println("case 3 : response code = 400,invalid policy_id")
	jobsParams = &JobsParams{PolicyID: "cc"}
	httpStatusCode, err = apiTest.ListJobs(*admin, *jobsParams)
	if err != nil {
		t.Error("Error while get jobs", err.Error())
		t.Log(err)
	} else {
		assert.Equal(int(400), httpStatusCode, "httpStatusCode should be 400")
	}

}
func TestJobLogsGet(t *testing.T) {
	var httpStatusCode int
	var err error
	var id string

	assert := assert.New(t)
	apiTest := newHarborAPI()

	fmt.Println("Testing JobLogs Get API")

	//-------------------case 1 : response code = 200------------------------//
	fmt.Println("case 1 : response code = 200")
	id = "1"
	httpStatusCode, err = apiTest.GetJobLogsByID(*admin, id)
	if err != nil {
		t.Error("Error while get jobLogs", err.Error())
		t.Log(err)
	} else {
		assert.Equal(int(200), httpStatusCode, "httpStatusCode should be 200")
	}
	//-------------------case 2 : response code = 404------------------------//
	fmt.Println("case 2 : response code = 404,page not found")
	id = "111"
	httpStatusCode, err = apiTest.GetJobLogsByID(*admin, id)
	if err != nil {
		t.Error("Error while get jobLogs", err.Error())
		t.Log(err)
	} else {
		assert.Equal(int(404), httpStatusCode, "httpStatusCode should be 404")
	}

}
