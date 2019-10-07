package leypaPractBuilder

type SuccessResponse struct {
	Message string
	Data    string
}

type SuccessResponseDataObjAdd struct {
	Message string
	Data    DataObjInsideAddResponse
}

type SuccessResponseDataObjList struct {
	Message string
	Data    ListEntry
}

type SuccessResponseDataObjStatus struct {
	Message string
	Data    DataObjInsideStatusResponse
}

type DataObjInsideAddResponse struct {
	Hash string
	Name string
	Size string
}

type DataObjInsideListResponse struct {
	Name string
	Type int
	Size int64
	Hash string
}

type ListEntry struct {
	Entries []DataObjInsideListResponse
}

type DataObjInsideStatusResponse struct {
	Hash           string
	Size           int
	CumulativeSize uint64
	Blocks         int
	Type           string
}

type FailureResponse struct {
	Message string
	Data    string
}
