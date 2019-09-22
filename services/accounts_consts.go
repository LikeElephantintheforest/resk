package services

//转账状态
type TransferStatus int8

const (
	//转账失败
	TransferStatusFailure TransferStatus = -1
	//余额不足
	TransferStatusSufficientFunds TransferStatus = 0
	//转账成功
	TransferStatusSuccess TransferStatus = 1
)

//转账类型
type ChangeType int8

const (
	//创建账户
	AccountCreated ChangeType = 0
	//储值
	AccountStoreValue ChangeType = 1
	//支出
	EnvelopeOutgoing ChangeType = -2
	//收入
	EnvelopeIncoming ChangeType = 2
	//红包退款
	EnvelopeExpiredRefund ChangeType = 3
)

//资金交易变化标识
type ChangeFlag int8

const (
	//创建账户 = 0
	FlagAccountCreated ChangeFlag = 0
	//支出 = -1
	FlagTransferOut ChangeFlag = -1
	//收入 = 1
	FlagTransferIn ChangeFlag = 1
)
