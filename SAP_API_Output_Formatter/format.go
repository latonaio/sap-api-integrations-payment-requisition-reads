package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-payment-requisition-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToStrategies(raw []byte, l *logger.Logger) ([]Strategies, error) {
	pm := &responses.Strategies{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Strategies. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	strategies := make([]Strategies, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		strategies = append(strategies, Strategies{
			PaymentRequisitionStrategyUUID: data.PaymentRequisitionStrategyUUID,
			PaymentRequisitionUUID:         data.PaymentRequisitionUUID,
			PaytRequisitionStrategyNumber:  data.PaytRequisitionStrategyNumber,
			PaymentReqnStrategyStatus:      data.PaymentReqnStrategyStatus,
			PaytReqnStrategyAmtInTransCrcy: data.PaytReqnStrategyAmtInTransCrcy,
			Currency:                       data.Currency,
			PaymentMethod:                  data.PaymentMethod,
			PaymentDate:                    data.PaymentDate,
			HouseBank:                      data.HouseBank,
			HouseBankAccount:               data.HouseBankAccount,
			Bank:                           data.Bank,
			BankAccount:                    data.BankAccount,
			BPBankAccountInternalID:        data.BPBankAccountInternalID,
			CreationDateTime:               data.CreationDateTime,
			ChangedOnDateTime:              data.ChangedOnDateTime,
		})
	}

	return strategies, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		item = append(item, Item{
			PaymentRequisitionItemUUID:     data.PaymentRequisitionItemUUID,
			PaymentRequisitionUUID:         data.PaymentRequisitionUUID,
			CompanyCode:                    data.CompanyCode,
			FiscalYear:                     data.FiscalYear,
			AccountingDocument:             data.AccountingDocument,
			AccountingDocumentItem:         data.AccountingDocumentItem,
			PaymentRequestAmountInPaytCrcy: data.PaymentRequestAmountInPaytCrcy,
			Currency:                       data.Currency,
		})
	}

	return item, nil
}

func ConvertToRequisitions(raw []byte, l *logger.Logger) ([]Requisitions, error) {
	pm := &responses.Requisitions{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Requisitions. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	requisitions := make([]Requisitions, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		requisitions = append(requisitions, Requisitions{
			PaymentRequisitionUUID:        data.PaymentRequisitionUUID,
			PaymentRequisitionNumber:      data.PaymentRequisitionNumber,
			CompanyCode:                   data.CompanyCode,
			Supplier:                      data.Supplier,
			PaymentRequisitionStatus:      data.PaymentRequisitionStatus,
			PlannedPaymentDate:            data.PlannedPaymentDate,
			PaymentRequisitionPriority:    data.PaymentRequisitionPriority,
			PaytRequisitionAmtInTransCrcy: data.PaytRequisitionAmtInTransCrcy,
			Currency:                      data.Currency,
			PaymentRequisitionType:        data.PaymentRequisitionType,
			PaymentMethod:                 data.PaymentMethod,
			NoteText:                      data.NoteText,
			WorkflowApproverNote:          data.WorkflowApproverNote,
			PaymentDifferenceReason:       data.PaymentDifferenceReason,
			CreationDateTime:              data.CreationDateTime,
			ChangedOnDateTime:             data.ChangedOnDateTime,  
		})
	}

	return requisitions, nil
}