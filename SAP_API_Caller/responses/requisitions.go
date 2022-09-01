package responses

type Requisitions struct {
	Count int `json:"@count"`
	Value []struct {
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
	} `json:"value"`
}