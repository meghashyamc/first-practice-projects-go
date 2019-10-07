package leypaPractHandlers

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"kiplog"
	"leypaPractResponses"
	"leypaPractUtil"
	"mime/multipart"
	"net/http"
)

var argError = errors.New("Error in arguments provided.")

func Add(w http.ResponseWriter, req *http.Request) {

	processFileUploadRequest(&w, req, leypaPractUtil.AddUrl)

}

func Remove(w http.ResponseWriter, req *http.Request) {

	processRequest("PUT", &w, req, leypaPractUtil.RemoveUrl, leypaPractUtil.RemoveCode, 1)
}
func Status(w http.ResponseWriter, req *http.Request) {

	processRequest("GET", &w, req, leypaPractUtil.StatusUrl, leypaPractUtil.StatusCode, 1)

}

func List(w http.ResponseWriter, req *http.Request) {

	processRequest("GET", &w, req, leypaPractUtil.ListUrl, leypaPractUtil.ListCode, 0)

}

func Mkdir(w http.ResponseWriter, req *http.Request) {

	processRequest("GET", &w, req, leypaPractUtil.MkdirUrl, leypaPractUtil.MkdirCode, 1)

}

func Read(w http.ResponseWriter, req *http.Request) {

	processRequest("GET", &w, req, leypaPractUtil.ReadUrl, leypaPractUtil.ReadCode, 1)

}

func Cat(w http.ResponseWriter, req *http.Request) {

	processRequest("GET", &w, req, leypaPractUtil.CatUrl, leypaPractUtil.CatCode, 1)

}

func Copy(w http.ResponseWriter, req *http.Request) {

	processRequest("PUT", &w, req, leypaPractUtil.CopyUrl, leypaPractUtil.CopyCode, 2)

}

func Move(w http.ResponseWriter, req *http.Request) {

	processRequest("PUT", &w, req, leypaPractUtil.MoveUrl, leypaPractUtil.MoveCode, 2)

}

func processFileUploadRequest(w *(http.ResponseWriter), req *http.Request, url string) {

	file, fileHeader, err := req.FormFile("arg")

	if err != nil {

		leypaPractResponses.WriteFailureResponse(w, http.StatusBadRequest, argError)
		return
	}

	fileName := fileHeader.Filename
	body, writer := getBufferToSend("arg", fileName, file)

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		leypaPractResponses.WriteFailureResponse(w, http.StatusBadRequest, err)

		return
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())

	request.Header.Add("Content-Type", "multipart/form-data")
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		leypaPractResponses.WriteFailureResponse(w, http.StatusBadRequest, err)
		return
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		leypaPractResponses.WriteFailureResponse(w, http.StatusBadRequest, err)

		return
	}

	leypaPractResponses.WriteSuccessResponseDataObjAdd(w, content)

}
func processRequest(reqType string, w *(http.ResponseWriter), req *http.Request, url string, urlCode uint8, compArgs uint8) {

	var value string
	if compArgs == 0 || compArgs == 1 {
		value = req.FormValue("arg")
	} else {
		value = req.FormValue("arg1")
	}

	if value == "" {

		if compArgs == 0 {

			makeSimpleGetRequest(w, url, urlCode)
			return
		}

		leypaPractResponses.WriteFailureResponse(w, http.StatusBadRequest, argError)

		return
	}

	request, err := http.NewRequest(reqType, url, nil)
	if err != nil {
		leypaPractResponses.WriteFailureResponse(w, http.StatusBadRequest, err)

		return
	}

	q := request.URL.Query()
	q.Add("arg", value)

	if compArgs == 2 {

		value = req.FormValue("arg2")
		q.Add("arg", value)
	}

	request.URL.RawQuery = q.Encode()

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		leypaPractResponses.WriteFailureResponse(w, http.StatusBadRequest, err)
		return
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		leypaPractResponses.WriteFailureResponse(w, http.StatusBadRequest, err)

		return
	}

	switch urlCode {

	case (leypaPractUtil.ListCode):
		leypaPractResponses.WriteSuccessResponseList(w, content)
		return
	case (leypaPractUtil.StatusCode):
		leypaPractResponses.WriteSuccessResponseStatus(w, content)
		return
	case (leypaPractUtil.CatCode):
		(*w).Write(content)
		return

	case (leypaPractUtil.MkdirCode):
		if len(content) != 0 {
			(*w).WriteHeader(http.StatusBadRequest)

			(*w).Write(content)
			return
		}
		leypaPractResponses.WriteSuccessResponse(w, "Directory created successfully.")
		return

	case (leypaPractUtil.ReadCode):
		(*w).Write(content)
		return

	case (leypaPractUtil.RemoveCode):
		if len(content) != 0 {
			(*w).WriteHeader(http.StatusBadRequest)

			(*w).Write(content)
			return
		}

		leypaPractResponses.WriteSuccessResponse(w, "Directory/file removed successfully.")
		return

	case (leypaPractUtil.CopyCode):
		if len(content) != 0 {
			(*w).WriteHeader(http.StatusBadRequest)

			(*w).Write(content)
			return
		}
		leypaPractResponses.WriteSuccessResponse(w, "Directory/file copied successfully.")
		return

	case (leypaPractUtil.MoveCode):
		if len(content) != 0 {
			(*w).WriteHeader(http.StatusBadRequest)

			(*w).Write(content)
			return
		}

		leypaPractResponses.WriteSuccessResponse(w, "Directory/file moved successfully.")
		return

	}
}

func makeSimpleGetRequest(w *(http.ResponseWriter), url string, urlCode uint8) {

	response, err := http.Get(url)
	if err != nil {
		leypaPractResponses.WriteFailureResponse(w, http.StatusBadRequest, err)

		return
	}

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		leypaPractResponses.WriteFailureResponse(w, http.StatusBadRequest, err)

		return
	}

	if urlCode == leypaPractUtil.ListCode {

		leypaPractResponses.WriteSuccessResponseList(w, content)
		return
	}

	(*w).Write(content)

}

func getBufferToSend(arg string, fileName string, file multipart.File) (*bytes.Buffer, *(multipart.Writer)) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(arg, fileName)

	if err != nil {
		kiplog.HTTPLog(err.Error())
	}

	io.Copy(part, file)
	writer.Close()

	return body, writer
}
