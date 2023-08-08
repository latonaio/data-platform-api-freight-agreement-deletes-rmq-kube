package requests

type Header struct {
	FreightAgreement    int   `json:"FreightAgreement"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
