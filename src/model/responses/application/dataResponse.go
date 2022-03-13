package application

func NewDataResponse(data []Response) *DataResponse {
	return &DataResponse{
		Data: data,
	}

}

type DataResponse struct {
	Data []Response `json:"data"`
}
