package constants

// ContractStatus enum
type ContractStatus string

const (

	// ContractDraft ContractStatus
	ContractDraftStatus ContractStatus = "draft"

	// ContractConclude  created
	ContractConcludeStatus ContractStatus = "concluded"

	// ContractChecking contractStatus
	ContractCheckingStatus ContractStatus = "checking"

	// ContractCancelStatus contractStatus
	ContractCancelledStatus ContractStatus = "cancelled"

	// ContractExpiredStatus contractStatus
	ContractExpiredStatus ContractStatus = "expired"
)
