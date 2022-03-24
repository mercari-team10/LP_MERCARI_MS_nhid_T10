package models

type Record struct {
	RecordId    string   `json:"_id" bson:"_id"`
	NHID        string   `json:"nhid"`
	RecordType  string   `json:"recordType"`
	IssuerType  string   `json:"issuerType"`
	IssuerId    string   `json:"issuerId"`
	Attachments []string `json:"attachments"`
	Comments    string   `json:"comments"`
	Timestamp   int64    `json:"timestamp"`
}

type AddRecordInput struct {
	NHID        string   `json:"nhid" binding:"required"`
	RecordType  string   `json:"recordType" binding:"required"`
	Attachments []string `json:"attachments" binding:"required"`
	Comments    string   `json:"comments"`
}

type FindRecordsInput struct {
	NHID      string `json:"nhid" binding:"required"`
	StartDate int64  `json:"startDate"`
	EndDate   int64  `json:"endDate"`
}
