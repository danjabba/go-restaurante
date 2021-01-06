package tools

// ResponseStructure ...
type ResponseStructure struct {
	Message string      `bson:"message" json:"message"`
	Success bool        `bson:"success" json:"success"`
	Data    interface{} `bson:"data" json:"data"`
}

// CreateResponse ...
func CreateResponse(data interface{}, message string, success bool) ResponseStructure {

	return ResponseStructure{
		Message: message,
		Data:    data,
		Success: success,
	}

}
