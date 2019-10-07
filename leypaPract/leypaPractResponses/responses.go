package leypaPractResponses

import (
	"encoding/json"
	"errors"
	// "github.com/gogo/protobuf/test/data"
	"kiplog"
	"leypaPractBuilder"
	"net/http"
	"strings"
)

const SUCCESS = "success"

func WriteSuccessResponse(w *(http.ResponseWriter), data string) {

	r := leypaPractBuilder.SuccessResponse{Message: SUCCESS, Data: data}

	jsonObj, err := json.MarshalIndent(r, "", "    ")

	if err != nil {
		kiplog.HTTPLog(err.Error())
		WriteFailureResponse(w, http.StatusBadRequest, errors.New("Error processing request. Please check method prerequisites."))
		return
	}

	(*w).Write(jsonObj)

}

func WriteSuccessResponseList(w *(http.ResponseWriter), jsonContent []byte) {

	listEntries := leypaPractBuilder.ListEntry{}
	listEntries.Entries = make([]leypaPractBuilder.DataObjInsideListResponse, 0)

	err := json.Unmarshal(jsonContent, &listEntries)
	if err != nil {
		kiplog.HTTPLog(err.Error())
		WriteFailureResponse(w, http.StatusBadRequest, errors.New("Error processing request. Please check method prerequisites."))

		return
	}

	if len(listEntries.Entries) == 0 {
		if !strings.Contains(string(jsonContent), "null") {
			kiplog.HTTPLog(string(jsonContent))
			(*w).WriteHeader(http.StatusBadRequest)
			(*w).Write(jsonContent)
			return
		}
	}
	r := leypaPractBuilder.SuccessResponseDataObjList{Message: SUCCESS, Data: listEntries}

	jsonObj, err := json.MarshalIndent(r, "", "    ")

	if err != nil {
		kiplog.HTTPLog(err.Error())
		WriteFailureResponse(w, http.StatusBadRequest, errors.New("Error processing request. Please check method prerequisites."))

		return
	}

	(*w).Write(jsonObj)
}

func WriteSuccessResponseStatus(w *(http.ResponseWriter), jsonContent []byte) {

	dataObj := leypaPractBuilder.DataObjInsideStatusResponse{}
	kiplog.HTTPLog(string(jsonContent))

	err := json.Unmarshal(jsonContent, &dataObj)

	if dataObj.Hash == "" {
		kiplog.HTTPLog(string(jsonContent))
		(*w).WriteHeader(http.StatusBadRequest)
		(*w).Write(jsonContent)
		return
	}
	if err != nil {
		kiplog.HTTPLog(err.Error())
		WriteFailureResponse(w, http.StatusBadRequest, errors.New("Error processing request. Please check method prerequisites."))

		return
	}
	r := leypaPractBuilder.SuccessResponseDataObjStatus{Message: SUCCESS, Data: dataObj}

	jsonObj, err := json.MarshalIndent(r, "", "    ")

	if err != nil {
		kiplog.HTTPLog(err.Error())
		WriteFailureResponse(w, http.StatusBadRequest, errors.New("Error processing request. Please check method prerequisites."))

		return
	}

	(*w).Write(jsonObj)
}

func WriteSuccessResponseDataObjAdd(w *(http.ResponseWriter), jsonContent []byte) {

	dataObj := leypaPractBuilder.DataObjInsideAddResponse{}
	err := json.Unmarshal(jsonContent, &dataObj)
	if err != nil {
		kiplog.HTTPLog(err.Error())
		WriteFailureResponse(w, http.StatusBadRequest, errors.New("Error processing request. Please check method prerequisites."))

		return
	}
	r := leypaPractBuilder.SuccessResponseDataObjAdd{Message: SUCCESS, Data: dataObj}

	jsonObj, err := json.MarshalIndent(r, "", "    ")

	if err != nil {
		kiplog.HTTPLog(err.Error())
		WriteFailureResponse(w, http.StatusBadRequest, errors.New("Error processing request. Please check method prerequisites."))

		return
	}

	(*w).Write(jsonObj)

}
func WriteFailureResponse(w *(http.ResponseWriter), code int, err error) {

	r := leypaPractBuilder.FailureResponse{Message: http.StatusText(code), Data: err.Error()}

	jsonObj, err := json.MarshalIndent(r, "", "    ")

	if err != nil {
		kiplog.HTTPLog(err.Error())
		WriteFailureResponse(w, http.StatusBadRequest, errors.New("Error processing request. Please check method prerequisites."))

		return
	}
	(*w).WriteHeader(code)
	(*w).Write(jsonObj)

}
