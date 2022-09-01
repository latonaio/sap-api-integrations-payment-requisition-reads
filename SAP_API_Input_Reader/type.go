package sap_api_input_reader

type EC_MC struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	BillingDocument struct {
		BillingDocument string `json:"document_no"`
		DeliverTo       string `json:"deliver_to"`
		Quantity        string `json:"quantity"`
		PickedQuantity  string `json:"picked_quantity"`
		TotalNetAmount  string `json:"price"`
		Batch           string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema           string `json:"api_schema"`
	MaterialCode        string `json:"material_code"`
	Plant_Supplier      string `json:"plant/supplier"`
	Stock               string `json:"stock"`
	BillingDocumentType string `json:"document_type"`
	SDDocument          string `json:"document_no"`
	PlannedDate         string `json:"planned_date"`
	BillingDocumentDate string `json:"validated_date"`
	Deleted             bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Strategies    struct {
		PaymentRequisitionStrategyUUID string `json:"PaymentRequisitionStrategyUUID"`
		PaymentRequisitionUUID         string `json:"PaymentRequisitionUUID"`
		PaytRequisitionStrategyNumber  string `json:"PaytRequisitionStrategyNumber"`
		PaymentReqnStrategyStatus      string `json:"PaymentReqnStrategyStatus"`
		PaytReqnStrategyAmtInTransCrcy string `json:"PaytReqnStrategyAmtInTransCrcy"`
		Currency                       string `json:"Currency"`
		PaymentMethod                  string `json:"PaymentMethod"`
		PaymentDate                    string `json:"PaymentDate"`
		HouseBank                      string `json:"HouseBank"`
		HouseBankAccount               string `json:"HouseBankAccount"`
		Bank                           string `json:"Bank"`
		BankAccount                    string `json:"BankAccount"`
		BPBankAccountInternalID        string `json:"BPBankAccountInternalID"`
		CreationDateTime               string `json:"CreationDateTime"`
		ChangedOnDateTime              string `json:"ChangedOnDateTime"`
		Item                           struct {
			PaymentRequisitionItemUUID     string `json:"PaymentRequisitionItemUUID"`
			PaymentRequisitionUUID         string `json:"PaymentRequisitionUUID"`
			CompanyCode                    string `json:"CompanyCode"`
			FiscalYear                     string `json:"FiscalYear"`
			AccountingDocument             string `json:"AccountingDocument"`
			AccountingDocumentItem         string `json:"AccountingDocumentItem"`
			PaymentRequestAmountInPaytCrcy string `json:"PaymentRequestAmountInPaytCrcy"`
			Currency                       string `json:"Currency"`
			Requisitions                   struct {
				PaymentRequisitionUUID        string `json:"PaymentRequisitionUUID"`
				PaymentRequisitionNumber      string `json:"PaymentRequisitionNumber"`
				CompanyCode                   string `json:"CompanyCode"`
				Supplier                      string `json:"Supplier"`
				PaymentRequisitionStatus      string `json:"PaymentRequisitionStatus"`
				PlannedPaymentDate            string `json:"PlannedPaymentDate"`
				PaymentRequisitionPriority    string `json:"PaymentRequisitionPriority"`
				PaytRequisitionAmtInTransCrcy string `json:"PaytRequisitionAmtInTransCrcy"`
				Currency                      string `json:"Currency"`
				PaymentRequisitionType        string `json:"PaymentRequisitionType"`
				PaymentMethod                 string `json:"PaymentMethod"`
				NoteText                      string `json:"NoteText"`
				WorkflowApproverNote          string `json:"WorkflowApproverNote"`
				PaymentDifferenceReason       string `json:"PaymentDifferenceReason"`
				CreationDateTime              string `json:"CreationDateTime"`
				ChangedOnDateTime             string `json:"ChangedOnDateTime"`
			} `json:"Requisitions"`
		} `json:"Item"`
	} `json:"Strategies"`
	APISchema              string   `json:"api_schema"`
	Accepter               []string `json:"accepter"`
	PaymentRequisitionCode string   `json:"payment_requisition_code"`
	Deleted                bool     `json:"deleted"`
}
