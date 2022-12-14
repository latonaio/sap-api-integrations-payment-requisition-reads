package main

import (
	sap_api_caller "sap-api-integrations-payment-requisition-reads/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-payment-requisition-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Payment_Requisition_Strategies_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata4/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Strategies", "Item", "Requisitions",
		}
	}

	caller.AsyncGetPaymentRequisition(
		inoutSDC.Strategies.PaymentRequisitionUUID,
		inoutSDC.Strategies.PaymentRequisitionStrategyUUID,
		inoutSDC.Strategies.Item.PaymentRequisitionItemUUID,
		accepter,
	)
}
