package component

//BTC交易
type BtcTx struct {
	Hash  string `json:"hash"`
	Basic struct {
		BlockHash string `json:"block_hash"`
		Time      int    `json:"time"`
		VinSz     int    `json:"vin_sz"`
		VoutSz    int    `json:"vout_sz"`
		LockTime  int    `json:"lock_time"`
		RelayedBy string `json:"relayed_by"`
		Ver       int    `json:"ver"`
		Weight    int    `json:"weight"`
		Size      int    `json:"size"`
		Inputs    []struct {
			Sequence int64  `json:"sequence"`
			Script   string `json:"script"`
			PrevOut  struct {
				Script string `json:"script"`
				Addr   string `json:"addr"`
				Value  int64  `json:"value"`
			} `json:"prev_out"`
		} `json:"inputs"`
		Out []struct {
			Script string `json:"script"`
			Addr   string `json:"addr"`
			Value  int64  `json:"value"`
		} `json:"out"`
	} `json:"basic"`
	Financial struct {
		TxVolume int64 `json:"tx_volume"`
		Fee      int   `json:"fee"`
	} `json:"financial"`
	UtxoAddress struct {
		UtxoConsumed int `json:"utxo_consumed"`
		UtxoCreate   int `json:"utxo_create"`
	} `json:"utxo_address"`
	Privacy struct {
		NewAddress    int `json:"new_address"`
		AddressReused int `json:"address_reused"`
	} `json:"privacy"`
}

//BTC交易副本　　因为kafka返回的数据中标签和所需要的不一致，因此克隆副本后再赋值给预定义字段的结构体
type BtcTx2 struct {
	Hash  string `json:"hash"`
	Basic struct {
		BlockHash string `json:"hash"`
		Time      int    `json:"time"`
		VinSz     int    `json:"vin_sz"`
		VoutSz    int    `json:"vout_sz"`
		LockTime  int    `json:"lock_time"`
		RelayedBy string `json:"relayed_by"`
		Ver       int    `json:"ver"`
		Weight    int    `json:"weight"`
		Size      int    `json:"size"`
		Inputs    []struct {
			Sequence int64  `json:"sequence"`
			Script   string `json:"script"`
			PrevOut  struct {
				Script string `json:"script"`
				Addr   string `json:"addr"`
				Value  int64  `json:"value"`
			} `json:"prev_out"`
		} `json:"inputs"`
		Out []struct {
			Script string `json:"script"`
			Addr   string `json:"addr"`
			Value  int64  `json:"value"`
		} `json:"out"`
	} `json:"basic"`
	Financial struct {
		TxVolume int64 `json:"tx_volume"`
		Fee      int   `json:"fee"`
	} `json:"financial"`
	UtxoAddress struct {
		UtxoConsumed int `json:"utxo_consumed"`
		UtxoCreate   int `json:"utxo_create"`
	} `json:"utxo_address"`
	Privacy struct {
		NewAddress    int `json:"new_address"`
		AddressReused int `json:"address_reused"`
	} `json:"privacy"`
}

//BTC区块
type BtcBlock struct {
	Hash  string `json:"hash"`
	Basic struct {
		PrevBlock    string `json:"prev_block"`
		MrklRoot     string `json:"mrkl_root"`
		Height       int    `json:"height"`
		Time         int    `json:"time"`
		NTx          int    `json:"n_tx"`
		MainChain    bool   `json:"main_chain"`
		Fee          int    `json:"fee"`
		Nonce        int64  `json:"nonce"`
		Bits         string    `json:"bits"`
		Size         int    `json:"size"`
		ReceivedTime int    `json:"received_time"`
		RelayedBy    string `json:"relayed_by"`
		Ver          int    `json:"ver"`
	} `json:"basic"`
	Transaction struct {
		TxTotalVolume   int64 `json:"tx_total_volume"`
		TxAverageVolume int64 `json:"tx_average_volume"`
	} `json:"transaction"`
	Fee struct {
		FeeTotalVolume   int `json:"fee_total_volume"`
		FeeAverageVolume int `json:"fee_average_volume"`
	} `json:"fee"`
	UtxoAddress struct {
		UtxoConsumed        int     `json:"utxo_consumed"`
		AddressUtxoConsumed int     `json:"address_utxo_consumed"`
		UtxoCreated         int     `json:"utxo_created"`
		AddressUtxoCreated  int     `json:"address_utxo_created"`
		NewAddress          int     `json:"new_address"`
		AddressReused       float64 `json:"address_reused"`
	} `json:"utxo_address"`
}

//BTC区块副本　　因为kafka返回的数据中标签和所需要的不一致，因此克隆副本后再赋值给预定义字段的结构体
type BtcBlock2 struct {
	Hash         string `json:"hash"`
	StrippedSize int    `json:"stripped_size"`
	Size         int    `json:"size"`
	Weight       int    `json:"weight"`
	Height       int    `json:"height"`
	Version      int    `json:"version"`
	MerkleRoot   string `json:"merkle_root"`
	Tx           []struct {
		TxID     string `json:"tx_id"`
		Hash     string `json:"hash"`
		Version  int    `json:"version"`
		Size     int    `json:"size"`
		Vsize    int    `json:"vsize"`
		Weight   int    `json:"weight"`
		LockTime int    `json:"lock_time"`
		Vin      []struct {
			Coinbase string `json:"coinbase"`
			Sequence int64  `json:"sequence"`
		} `json:"vin"`
		Vout []struct {
			Value              float64  `json:"value"`
			IsCoinbase         bool     `json:"is_coinbase"`
			ScriptPubkey       string   `json:"script_pubkey"`
			RequiredSignatures int      `json:"required_signatures"`
			Type               string   `json:"type"`
			Addresses          []string `json:"addresses"`
		} `json:"vout"`
	} `json:"tx"`
	Time       int     `json:"time"`
	MedianTime int     `json:"median_time"`
	Nonce      int64   `json:"nonce"`
	Bits       string  `json:"bits"`
	Difficulty float64 `json:"difficulty"`
	ChainWork  string  `json:"chain_work"`
	PrevHash   string  `json:"prev_hash"`
}





















//BTC地址
type BtcAddress struct {
	Addr  string `json:"addr"`
	Basic struct {
		FirstSeen string `json:"first_seen"`
		LastSeen  string `json:"last_seen"`
		InTxs     int    `json:"in_txs"`
		OutTxs    int    `json:"out_txs"`
		InVolume  int64  `json:"in_volume"`
		OutVolume int64  `json:"out_volume"`
		Balance   int    `json:"balance"`
		Tag       string `json:"tag"`
		TagLink   string `json:"tag_link"`
	} `json:"basic"`
	TransactionSummary struct {
		Monthly []struct {
			Interval  string `json:"interval"`
			InTxs     int    `json:"in_txs"`
			OutTxs    int    `json:"out_txs"`
			InVolume  int64  `json:"in_volume"`
			OutVolume int64  `json:"out_volume"`
			Daily     []struct {
				Interval  int `json:"interval"`
				InTxs     int `json:"in_txs"`
				OutTxs    int `json:"out_txs"`
				InVolume  int `json:"in_volume"`
				OutVolume int `json:"out_volume"`
			} `json:"daily,omitempty"`
		} `json:"monthly"`
	} `json:"transaction_summary"`
	TransactionDetail []struct {
		Hash string `json:"hash"`
	} `json:"transaction_detail"`
}

//BTC地址    因为kafka返回的数据中标签和所需要的不一致，因此克隆副本后再赋值给预定义字段的结构体
type BtcAddress2 struct {
	Addr  string `json:"addr"`
	Basic struct {
		FirstSeen string `json:"first_seen"`
		LastSeen  string `json:"last_seen"`
		InTxs     int    `json:"in_txs"`
		OutTxs    int    `json:"out_txs"`
		InVolume  int64  `json:"in_volume"`
		OutVolume int64  `json:"out_volume"`
		Balance   int    `json:"balance"`
		Tag       string `json:"tag"`
		TagLink   string `json:"tag_link"`
	} `json:"basic"`
	TransactionSummary struct {
		Monthly []struct {
			Interval  string `json:"interval"`
			InTxs     int    `json:"in_txs"`
			OutTxs    int    `json:"out_txs"`
			InVolume  int64  `json:"in_volume"`
			OutVolume int64  `json:"out_volume"`
			Daily     []struct {
				Interval  int `json:"interval"`
				InTxs     int `json:"in_txs"`
				OutTxs    int `json:"out_txs"`
				InVolume  int `json:"in_volume"`
				OutVolume int `json:"out_volume"`
			} `json:"daily,omitempty"`
		} `json:"monthly"`
	} `json:"transaction_summary"`
	TransactionDetail []struct {
		Hash string `json:"hash"`
	} `json:"transaction_detail"`
}

//BTC节点信息
type BtcNode struct {
	Addr         string `json:"addr"`
	FirstSeen    string `json:"first_seen"`
	LastSeen     string `json:"last_seen"`
	LatestHeight int    `json:"latest_height"`
	Nodes        struct {
		Basic struct {
			BitcoinAddress   string `json:"bitcoin_address"`
			Height           int    `json:"height"`
			Hostname         string `json:"hostname"`
			City             string `json:"city"`
			CountryCode      string `json:"country_code"`
			Timezone         string `json:"timezone"`
			ASN              string `json:"ASN"`
			OrganizationName string `json:"organization_name"`
		} `json:"basic"`
		Coordinates struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"coordinates"`
		UserAgents struct {
			Version int    `json:"version"`
			Client  string `json:"client"`
		} `json:"user_agents"`
	} `json:"nodes"`
}
