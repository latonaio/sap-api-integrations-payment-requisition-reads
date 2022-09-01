package responses

type Strategies struct {
	Count int `json:"@count"`
	Value []struct {
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
	} `json:"value"`
}