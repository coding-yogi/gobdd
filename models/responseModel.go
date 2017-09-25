package models

// TokenResponse ...
type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

// APIHubErrorResponse ...
type APIHubErrorResponse struct {
	Fault Fault
}

// Fault ...
type Fault struct {
	FaultString string      `json:"faultstring"`
	Detail      FaultDetail `json:"detail"`
}

// FaultDetail ...
type FaultDetail struct {
	ErrorCode string
}

// RecommenderResponse ...
type RecommenderResponse struct {
	RestOperationStatusVOX RestOperationStatusVOX
}

// RestOperationStatusVOX ..
type RestOperationStatusVOX struct {
	CorrelationID string `json:"correlation_id"`
	Data          Data   `json:"data"`
	Status        string `json:"status"`
}

// Data ..
type Data struct {
	RestReturnData         []RestReturnData `json:"REST_RETURN_DATA"`
	RestReturnDataAsString string
}

// RestReturnData ..
type RestReturnData struct {
	Recommendations []Recommendation `json:"recommendations"`
	StudID          string           `json:"studId"`
	Error           string           `json:"error"`
}

// Recommendation ...
type Recommendation struct {
	Content   Content `json:"content"`
	Dismissed bool    `json:"dismissed"`
	ID        string  `json:"id"`
}

// Why ....
type Why struct {
	Algorithm string `json:"algorithm"`
	Topic     string `json:"topic"`
}

// Content ...
type Content struct {
	ExternalID string  `json:"externalId"`
	Topics     []Topic `json:"topics"`
	Type       string  `json:"type"`
}

// Topic ...
type Topic struct {
	ID           string `json:"id"`
	IsActive     bool   `json:"isActive"`
	IsSubscribed bool   `json:"isSubscribed"`
	Name         string `json:"name"`
	TopicType    string `json:"topicType"`
}
