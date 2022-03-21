package common

type RuntimeInformation struct {
	JVMInformation            JVMInformation `json:"jvmInformation"`
	OSInformation             OSInformation  `json:"osInformation"`
	MuleLicenseExpirationDate int64          `json:"muleLicenseExpirationDate"`
}
