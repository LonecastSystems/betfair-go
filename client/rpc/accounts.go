package rpc

func (client *JsonRpcClient) GetAccountFunds(params AccountDetailsParams) (WalletResult, error) {
	json := WalletResult{}

	if err := GetAccounts(client, 1, "getAccountFunds", params, &json); err != nil {
		return WalletResult{}, err
	}

	return json, nil
}
