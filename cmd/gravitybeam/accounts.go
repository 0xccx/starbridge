package main

import (
	"github.com/stellar/go/txnbuild"
)

func Accounts(tx *txnbuild.Transaction) []string {
	accountsMap := map[string]bool{}

	accountsMap[tx.SourceAccount().AccountID] = true

	for _, op := range tx.Operations() {
		opSource := op.GetSourceAccount()
		if opSource != "" {
			accountsMap[opSource] = true
		}

		switch o := op.(type) {
		case *txnbuild.CreateClaimableBalance:
			for _, claimant := range o.Destinations {
				// TODO: Assess the predicate portion of the claimant?
				accountsMap[claimant.Destination] = true
			}
		}
	}

	accounts := make([]string, 0, len(accountsMap))
	for a := range accountsMap {
		accounts = append(accounts, a)
	}

	return accounts
}