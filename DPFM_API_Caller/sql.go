package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-freight-agreement-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-freight-agreement-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) Header(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {

	where := strings.Join([]string{
		fmt.Sprintf("WHERE header.FreightAgreement = %d ", input.Header.FreightAgreement),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	header.FreightAgreement
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_freight_agreement_header_data as header 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Item(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Item {

	where := strings.Join([]string{
		fmt.Sprintf("WHERE header.FreightAgreement = %d ", input.Header.FreightAgreement),
		fmt.Sprintf("AND item.FreightAgreementItem = %d ", input.Item.FreightAgreementItem),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
			item.FreightAgreement, item.FreightAgreementItem
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_freight_agreement_item_data as item 
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_freight_agreement_header_data as header
		ON header.FreightAgreement = item.FreightAgreement ` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
