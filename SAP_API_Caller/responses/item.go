package responses

type Item struct {
	Count int `json:"@count"`
	Value []struct {
		PaymentRequisitionItemUUID     string `json:"PaymentRequisitionItemUUID"`
		PaymentRequisitionUUID         string `json:"PaymentRequisitionUUID"`
		CompanyCode                    string `json:"CompanyCode"`
		FiscalYear                     string `json:"FiscalYear"`
		AccountingDocument             string `json:"AccountingDocument"`
		AccountingDocumentItem         string `json:"AccountingDocumentItem"`
		PaymentRequestAmountInPaytCrcy int    `json:"PaymentRequestAmountInPaytCrcy"`
		Currency                       string `json:"Currency"`
	} `json:"value"`
}