package common

type JVMInformation struct {
	Runtime       JVMInformationRuntime       `json:"runtime"`
	Specification JVMInformationSpecification `json:"specification"`
}
