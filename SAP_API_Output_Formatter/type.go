package sap_api_output_formatter

type PaymentRequisition struct {
	ConnectionKey         string `json:"connection_key"`
	Result                bool   `json:"result"`
	RedisKey              string `json:"redis_key"`
	Filepath              string `json:"filepath"`
	APISchema             string `json:"api_schema"`
	PaymentRequisitionCode string `json:"supplier_quotation_code"`
	Deleted               bool   `json:"deleted"`
}

type Strategies struct {
		PaymentRequisitionStrategyUUID string    `json:"PaymentRequisitionStrategyUUID"`
		PaymentRequisitionUUID         string    `json:"PaymentRequisitionUUID"`
		PaytRequisitionStrategyNumber  string    `json:"PaytRequisitionStrategyNumber"`
		PaymentReqnStrategyStatus      string    `json:"PaymentReqnStrategyStatus"`
		PaytReqnStrategyAmtInTransCrcy int       `json:"PaytReqnStrategyAmtInTransCrcy"`
		Currency                       string    `json:"Currency"`
		PaymentMethod                  string    `json:"PaymentMethod"`
		PaymentDate                    string    `json:"PaymentDate"`
		HouseBank                      string    `json:"HouseBank"`
		HouseBankAccount               string    `json:"HouseBankAccount"`
		Bank                           string    `json:"Bank"`
		BankAccount                    string    `json:"BankAccount"`
		BPBankAccountInternalID        string    `json:"BPBankAccountInternalID"`
		CreationDateTime               string    `json:"CreationDateTime"`
		ChangedOnDateTime              string    `json:"ChangedOnDateTime"`
}

type Item struct {
		PaymentRequisitionItemUUID     string `json:"PaymentRequisitionItemUUID"`
		PaymentRequisitionUUID         string `json:"PaymentRequisitionUUID"`
		CompanyCode                    string `json:"CompanyCode"`
		FiscalYear                     string `json:"FiscalYear"`
		AccountingDocument             string `json:"AccountingDocument"`
		AccountingDocumentItem         string `json:"AccountingDocumentItem"`
		PaymentRequestAmountInPaytCrcy int    `json:"PaymentRequestAmountInPaytCrcy"`
		Currency                       string `json:"Currency"`
}

type Requisitions struct {
		PaymentRequisitionUUID        string    `json:"PaymentRequisitionUUID"`
		PaymentRequisitionNumber      string    `json:"PaymentRequisitionNumber"`
		CompanyCode                   string    `json:"CompanyCode"`
		Supplier                      string    `json:"Supplier"`
		PaymentRequisitionStatus      string    `json:"PaymentRequisitionStatus"`
		PlannedPaymentDate            string    `json:"PlannedPaymentDate"`
		PaymentRequisitionPriority    string    `json:"PaymentRequisitionPriority"`
		PaytRequisitionAmtInTransCrcy int       `json:"PaytRequisitionAmtInTransCrcy"`
		Currency                      string    `json:"Currency"`
		PaymentRequisitionType        string    `json:"PaymentRequisitionType"`
		PaymentMethod                 string    `json:"PaymentMethod"`
		NoteText                      string    `json:"NoteText"`
		WorkflowApproverNote          string    `json:"WorkflowApproverNote"`
		PaymentDifferenceReason       string    `json:"PaymentDifferenceReason"`
		CreationDateTime              string    `json:"CreationDateTime"`
		ChangedOnDateTime             string    `json:"ChangedOnDateTime"`  
}