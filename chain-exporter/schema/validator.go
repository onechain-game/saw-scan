package schema

import "time"

// Validator defines the structure for validator information.
type Validator struct {
	ID                      int32     `json:"id" sql:",pk"`
	Moniker                 string    `json:"moniker"`
	AccountAddress          string    `json:"account_address" sql:",notnull, unique"`
	OperatorAddress         string    `json:"operator_address" sql:",notnull, unique"`
	ConsensusAddress        string    `json:"consensus_address" sql:",notnull, unique"`
	Jailed                  bool      `json:"jailed"`
	Status                  string    `json:"status"`
	Tokens                  string    `json:"tokens"`
	VotingPower             int64     `json:"voting_power"`
	DelegatorShares         string    `json:"delegator_shares"`
	BondHeight              int64     `json:"bond_height" sql:"default:0"`
	BondIntraTxCounter      int64     `json:"bond_intra_tx_counter" sql:"default:0"`
	UnbondingHeight         int64     `json:"unbonding_height" sql:"default:0"`
	UnbondingTime           string    `json:"unbonding_time"`
	CommissionRate          string    `json:"commission_rate"`
	CommissionMaxRate       string    `json:"commission_max_rate"`
	CommissionMaxChangeRate string    `json:"commission_max_change_rate"`
	CommissionUpdateTime    string    `json:"commission_update_time"`
	Timestamp               time.Time `json:"timestamp" sql:"default:now()"`
}

// NewValidator returns a new Validator.
func NewValidator(v Validator) *Validator {
	return &Validator{
		Moniker:                 v.Moniker,
		AccountAddress:          v.AccountAddress,
		OperatorAddress:         v.OperatorAddress,
		ConsensusAddress:        v.ConsensusAddress,
		Jailed:                  v.Jailed,
		Status:                  v.Status,
		Tokens:                  v.Tokens,
		VotingPower:             v.VotingPower,
		DelegatorShares:         v.DelegatorShares,
		BondHeight:              v.BondHeight,
		BondIntraTxCounter:      v.BondIntraTxCounter,
		UnbondingHeight:         v.UnbondingHeight,
		UnbondingTime:           v.UnbondingTime,
		CommissionRate:          v.CommissionRate,
		CommissionMaxRate:       v.CommissionMaxRate,
		CommissionMaxChangeRate: v.CommissionMaxChangeRate,
		CommissionUpdateTime:    v.CommissionUpdateTime,
		Timestamp:               v.Timestamp,
	}
}
