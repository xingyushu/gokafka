package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
	"testing"
	"xinyu/go_splitter/broker"
	"xinyu/go_splitter/component"
)

func Test_kBroker_Subscribe(t *testing.T) {
	config := DefaultClusterConfig
	config.Version = sarama.V0_10_2_1

	client := NewBroker(ClusterConfig(config))

	err := client.Connect()
	if err != nil {
		return
	}

	//err = client.Publish("btc", &broker.Message{
	//	Body: []byte{},
	//})
	//
	//err = client.Publish("btc", &broker.Message{
	//	Body: []byte{},
	//})
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go client.Subscribe("btc", func(event broker.Event) error {

		fmt.Println("-----------------------------")
		var block component.BtcBlock2
		//err := json.Unmarshal(event.Message().Body, &block)
		//if err != nil {
		//	_ = event.Ack()
		//	return nil
		//}
		
		//fmt.Println(block)

		fmt.Println("-----------------------------")
		b :=component.BtcBlock{
			Hash: block.Hash,
			Basic: struct {
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
			}{
				PrevBlock:block.PrevHash,
				MrklRoot:block.MerkleRoot,
				Height:block.Height,
				Time:block.Time,
				MainChain:true,
				Nonce:block.Nonce,
				Bits:block.Bits,
				Size:block.Size,
				Ver:block.Version,

			},
			Transaction: struct {
				TxTotalVolume   int64 `json:"tx_total_volume"`
				TxAverageVolume int64 `json:"tx_average_volume"`
			}{},
			Fee: struct {
				FeeTotalVolume   int `json:"fee_total_volume"`
				FeeAverageVolume int `json:"fee_average_volume"`
			}{},
			UtxoAddress: struct {
				UtxoConsumed        int     `json:"utxo_consumed"`
				AddressUtxoConsumed int     `json:"address_utxo_consumed"`
				UtxoCreated         int     `json:"utxo_created"`
				AddressUtxoCreated  int     `json:"address_utxo_created"`
				NewAddress          int     `json:"new_address"`
				AddressReused       float64 `json:"address_reused"`
			}{},
		}
		fmt.Println("this is the real block we want")
		fmt.Println(b)
		_ = event.Ack()
		//wg.Done()
		return nil
	}, func(options *broker.SubscribeOptions) {
		options.Queue = "example1"
	})

	//err = client.Publish("btc", &broker.Message{
	//	Body: []byte{},
	//})
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	wg.Wait()
}
