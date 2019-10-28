package kafka

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
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

	var es *elasticsearch.Client
	{
		config := elasticsearch.Config{}
		config.Addresses = []string{"http://172.16.2.56:9200"}
		es, err = elasticsearch.NewClient(config)
		if err != nil {
			log.Fatalf("Error creating the client: %s", err)
		}

	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go client.Subscribe("btc", func(event broker.Event) error {
		ctx := context.Background()
		//原始信息
		if event.Message().Body == nil {
			_ = event.Ack()
			req := esapi.IndexRequest{
				Index:   "btc",
				Body:    bytes.NewReader(event.Message().Body),
			}
			_, _ = req.Do(ctx, es)
			return nil
		}

       //区块信息
		fmt.Println("-----------------------------")
		var block component.BtcBlock2
		//fmt.Println(event.Message().Body)
		err := json.Unmarshal(event.Message().Body, &block)
		if err != nil {
			_ = event.Ack()
			return nil
		}

		//fmt.Println(block)

		fmt.Println("-----------------------------")
		//返回所需要的区块信息
		b := component.BtcBlock{
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
				Bits         string `json:"bits"`
				Size         int    `json:"size"`
				ReceivedTime int    `json:"received_time"`
				RelayedBy    string `json:"relayed_by"`
				Ver          int    `json:"ver"`
			}{
				PrevBlock: block.PrevHash,
				MrklRoot:  block.MerkleRoot,
				Height:    block.Height,
				Time:      block.Time,
				MainChain: true,
				Nonce:     block.Nonce,
				Bits:      block.Bits,
				Size:      block.Size,
				Ver:       block.Version,
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
		jsonStr, err := json.Marshal(b)
		if err != nil {
			_ = event.Ack()
			return nil
		}
		fmt.Println(string(jsonStr))
		req := esapi.IndicesCreateRequest{
			Index:   "Btcblock",
			Body:    bytes.NewReader(jsonStr),
		}
		_, _ = req.Do(ctx, es)
		_ = event.Ack()
		//wg.Done()


		//返回所需要的交易信息
		list := make([]component.BtcTx, len(block.Tx))
		for index, tx := range block.Tx {
			txs := component.BtcTx{
				Hash: tx.Hash,
				Basic: component.Basic{
					BlockHash: block.Hash,
					Time:      block.Time,
					LockTime:  tx.LockTime,
					Ver:       block.Version,
					Weight:    block.Weight,
					Size:      block.Size,
				},
			}
			inputs := make([]component.Inputs, len(tx.Vin))
			out := make([]component.Out, len(tx.Vout))

			for index, vin := range tx.Vin {
				inputs[index] = component.Inputs{
					Sequence: vin.Sequence,
					Script:   "-",
					PrevOut:  component.PrevOut{},
				}
			}

			for inx, outs := range tx.Vout {
				out[inx] = component.Out{
					Script: outs.ScriptPubkey,
					Addr:   outs.Addresses,
					Value:  outs.Value,
				}
			}
			txs.Inputs = inputs
			txs.Out = out
			list[index] = txs



			for _, value := range list {
				jsonTx, err := json.Marshal(value)
				if err != nil {
					_ = event.Ack()
					return nil
				}
				req := esapi.IndexRequest{
					Index:   "BtcTx",
					Body:    bytes.NewReader(jsonTx),
				}
				_, _ = req.Do(ctx, es)
				fmt.Println(string(jsonTx))
				_ = event.Ack()
			}
		}

		//返回所需要的地址信息
		//address :=component.BtcAddress{
		//}

		return nil
	}, func(options *broker.SubscribeOptions) {
		options.Queue = "example1"
	})

	wg.Wait()
}
