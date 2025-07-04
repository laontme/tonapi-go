// Code generated by ogen, DO NOT EDIT.

package tonapi

// OperationName is the ogen operation name
type OperationName = string

const (
	AccountDnsBackResolveOperation                     OperationName = "AccountDnsBackResolve"
	AddressParseOperation                              OperationName = "AddressParse"
	BlockchainAccountInspectOperation                  OperationName = "BlockchainAccountInspect"
	DecodeMessageOperation                             OperationName = "DecodeMessage"
	DnsResolveOperation                                OperationName = "DnsResolve"
	DownloadBlockchainBlockBocOperation                OperationName = "DownloadBlockchainBlockBoc"
	EmulateMessageToAccountEventOperation              OperationName = "EmulateMessageToAccountEvent"
	EmulateMessageToEventOperation                     OperationName = "EmulateMessageToEvent"
	EmulateMessageToTraceOperation                     OperationName = "EmulateMessageToTrace"
	EmulateMessageToWalletOperation                    OperationName = "EmulateMessageToWallet"
	ExecGetMethodForBlockchainAccountOperation         OperationName = "ExecGetMethodForBlockchainAccount"
	ExecGetMethodWithBodyForBlockchainAccountOperation OperationName = "ExecGetMethodWithBodyForBlockchainAccount"
	GaslessConfigOperation                             OperationName = "GaslessConfig"
	GaslessEstimateOperation                           OperationName = "GaslessEstimate"
	GaslessSendOperation                               OperationName = "GaslessSend"
	GetAccountOperation                                OperationName = "GetAccount"
	GetAccountDiffOperation                            OperationName = "GetAccountDiff"
	GetAccountDnsExpiringOperation                     OperationName = "GetAccountDnsExpiring"
	GetAccountEventOperation                           OperationName = "GetAccountEvent"
	GetAccountEventsOperation                          OperationName = "GetAccountEvents"
	GetAccountExtraCurrencyHistoryByIDOperation        OperationName = "GetAccountExtraCurrencyHistoryByID"
	GetAccountInfoByStateInitOperation                 OperationName = "GetAccountInfoByStateInit"
	GetAccountJettonBalanceOperation                   OperationName = "GetAccountJettonBalance"
	GetAccountJettonHistoryByIDOperation               OperationName = "GetAccountJettonHistoryByID"
	GetAccountJettonsBalancesOperation                 OperationName = "GetAccountJettonsBalances"
	GetAccountJettonsHistoryOperation                  OperationName = "GetAccountJettonsHistory"
	GetAccountMultisigsOperation                       OperationName = "GetAccountMultisigs"
	GetAccountNftHistoryOperation                      OperationName = "GetAccountNftHistory"
	GetAccountNftItemsOperation                        OperationName = "GetAccountNftItems"
	GetAccountNominatorsPoolsOperation                 OperationName = "GetAccountNominatorsPools"
	GetAccountPublicKeyOperation                       OperationName = "GetAccountPublicKey"
	GetAccountSeqnoOperation                           OperationName = "GetAccountSeqno"
	GetAccountSubscriptionsOperation                   OperationName = "GetAccountSubscriptions"
	GetAccountTracesOperation                          OperationName = "GetAccountTraces"
	GetAccountsOperation                               OperationName = "GetAccounts"
	GetAllAuctionsOperation                            OperationName = "GetAllAuctions"
	GetAllRawShardsInfoOperation                       OperationName = "GetAllRawShardsInfo"
	GetBlockchainAccountTransactionsOperation          OperationName = "GetBlockchainAccountTransactions"
	GetBlockchainBlockOperation                        OperationName = "GetBlockchainBlock"
	GetBlockchainBlockTransactionsOperation            OperationName = "GetBlockchainBlockTransactions"
	GetBlockchainConfigOperation                       OperationName = "GetBlockchainConfig"
	GetBlockchainConfigFromBlockOperation              OperationName = "GetBlockchainConfigFromBlock"
	GetBlockchainMasterchainBlocksOperation            OperationName = "GetBlockchainMasterchainBlocks"
	GetBlockchainMasterchainHeadOperation              OperationName = "GetBlockchainMasterchainHead"
	GetBlockchainMasterchainShardsOperation            OperationName = "GetBlockchainMasterchainShards"
	GetBlockchainMasterchainTransactionsOperation      OperationName = "GetBlockchainMasterchainTransactions"
	GetBlockchainRawAccountOperation                   OperationName = "GetBlockchainRawAccount"
	GetBlockchainTransactionOperation                  OperationName = "GetBlockchainTransaction"
	GetBlockchainTransactionByMessageHashOperation     OperationName = "GetBlockchainTransactionByMessageHash"
	GetBlockchainValidatorsOperation                   OperationName = "GetBlockchainValidators"
	GetChartRatesOperation                             OperationName = "GetChartRates"
	GetDnsInfoOperation                                OperationName = "GetDnsInfo"
	GetDomainBidsOperation                             OperationName = "GetDomainBids"
	GetEventOperation                                  OperationName = "GetEvent"
	GetExtraCurrencyInfoOperation                      OperationName = "GetExtraCurrencyInfo"
	GetItemsFromCollectionOperation                    OperationName = "GetItemsFromCollection"
	GetJettonAccountHistoryByIDOperation               OperationName = "GetJettonAccountHistoryByID"
	GetJettonHoldersOperation                          OperationName = "GetJettonHolders"
	GetJettonInfoOperation                             OperationName = "GetJettonInfo"
	GetJettonInfosByAddressesOperation                 OperationName = "GetJettonInfosByAddresses"
	GetJettonTransferPayloadOperation                  OperationName = "GetJettonTransferPayload"
	GetJettonsOperation                                OperationName = "GetJettons"
	GetJettonsEventsOperation                          OperationName = "GetJettonsEvents"
	GetMarketsRatesOperation                           OperationName = "GetMarketsRates"
	GetMultisigAccountOperation                        OperationName = "GetMultisigAccount"
	GetMultisigOrderOperation                          OperationName = "GetMultisigOrder"
	GetNftCollectionOperation                          OperationName = "GetNftCollection"
	GetNftCollectionItemsByAddressesOperation          OperationName = "GetNftCollectionItemsByAddresses"
	GetNftCollectionsOperation                         OperationName = "GetNftCollections"
	GetNftHistoryByIDOperation                         OperationName = "GetNftHistoryByID"
	GetNftItemByAddressOperation                       OperationName = "GetNftItemByAddress"
	GetNftItemsByAddressesOperation                    OperationName = "GetNftItemsByAddresses"
	GetOpenapiJsonOperation                            OperationName = "GetOpenapiJson"
	GetOpenapiYmlOperation                             OperationName = "GetOpenapiYml"
	GetOutMsgQueueSizesOperation                       OperationName = "GetOutMsgQueueSizes"
	GetPurchaseHistoryOperation                        OperationName = "GetPurchaseHistory"
	GetRatesOperation                                  OperationName = "GetRates"
	GetRawAccountStateOperation                        OperationName = "GetRawAccountState"
	GetRawBlockProofOperation                          OperationName = "GetRawBlockProof"
	GetRawBlockchainBlockOperation                     OperationName = "GetRawBlockchainBlock"
	GetRawBlockchainBlockHeaderOperation               OperationName = "GetRawBlockchainBlockHeader"
	GetRawBlockchainBlockStateOperation                OperationName = "GetRawBlockchainBlockState"
	GetRawBlockchainConfigOperation                    OperationName = "GetRawBlockchainConfig"
	GetRawBlockchainConfigFromBlockOperation           OperationName = "GetRawBlockchainConfigFromBlock"
	GetRawConfigOperation                              OperationName = "GetRawConfig"
	GetRawListBlockTransactionsOperation               OperationName = "GetRawListBlockTransactions"
	GetRawMasterchainInfoOperation                     OperationName = "GetRawMasterchainInfo"
	GetRawMasterchainInfoExtOperation                  OperationName = "GetRawMasterchainInfoExt"
	GetRawShardBlockProofOperation                     OperationName = "GetRawShardBlockProof"
	GetRawShardInfoOperation                           OperationName = "GetRawShardInfo"
	GetRawTimeOperation                                OperationName = "GetRawTime"
	GetRawTransactionsOperation                        OperationName = "GetRawTransactions"
	GetReducedBlockchainBlocksOperation                OperationName = "GetReducedBlockchainBlocks"
	GetStakingPoolHistoryOperation                     OperationName = "GetStakingPoolHistory"
	GetStakingPoolInfoOperation                        OperationName = "GetStakingPoolInfo"
	GetStakingPoolsOperation                           OperationName = "GetStakingPools"
	GetStorageProvidersOperation                       OperationName = "GetStorageProviders"
	GetTonConnectPayloadOperation                      OperationName = "GetTonConnectPayload"
	GetTraceOperation                                  OperationName = "GetTrace"
	GetWalletInfoOperation                             OperationName = "GetWalletInfo"
	GetWalletsByPublicKeyOperation                     OperationName = "GetWalletsByPublicKey"
	ReindexAccountOperation                            OperationName = "ReindexAccount"
	SearchAccountsOperation                            OperationName = "SearchAccounts"
	SendBlockchainMessageOperation                     OperationName = "SendBlockchainMessage"
	SendRawMessageOperation                            OperationName = "SendRawMessage"
	StatusOperation                                    OperationName = "Status"
	TonConnectProofOperation                           OperationName = "TonConnectProof"
)
