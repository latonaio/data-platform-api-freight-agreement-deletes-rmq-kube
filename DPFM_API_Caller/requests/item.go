package requests

type Item struct {
	FreightAgreement     int   `json:"FreightAgreement"`
	FreightAgreementItem int   `json:"FreightAgreementItem"`
	IsMarkedForDeletion  *bool `json:"IsMarkedForDeletion"`
}
