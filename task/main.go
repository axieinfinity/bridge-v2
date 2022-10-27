package task

import "github.com/axieinfinity/bridge-v2/models"

const (
	ACK_WITHDREW_TASK           = "acknowledgeWithdrew"
	DEPOSIT_TASK                = "deposit"
	WITHDRAWAL_TASK             = "withdrawal"
	VOTE_BRIDGE_OPERATORS_TASK  = "voteBridgeOperatorsBySignatures"
	RELAY_BRIDGE_OPERATORS_TASK = "relayBridgeOperators"

	STATUS_PENDING    = "pending"
	STATUS_FAILED     = "failed"
	STATUS_PROCESSING = "processing"
	STATUS_DONE       = "done"

	GATEWAY_CONTRACT        = "Gateway"
	GOVERNANCE_CONTRACT     = "Governance"
	ETH_GOVERNANCE_CONTRACT = "Governance"
	TRUSTED_ORGS_CONTRACT   = "TrustedOrgs"
	ETH_GATEWAY_CONTRACT    = "EthGateway"
	BRIDGEADMIN_CONTRACT    = "BridgeAdmin"
)

const (
	VoteStatusPending = iota
	VoteStatusApproved
	VoteStatusExecuted
)

type Tasker interface {
	collectTask(t *models.Task)
	send()
}
