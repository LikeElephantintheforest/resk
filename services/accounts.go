package services

import (
	"github.com/shopspring/decimal"
	"time"
)

type AccountService interface {
	CreateAccount(request AccountCreateDTO) (*AccountDTO, error)
	Transfer(request AccountTransferDTO) (TransferStatus, error)
	StoreValue(request AccountTransferDTO) (TransferStatus, error)
	GetEnvelopeAccountByUserId(userId string) *AccountDTO
}

//账户交易参与者
type TradeParticipator struct {
	AccountNo string
	UserId    string
	Username  string
}

//账户转账
type AccountTransferDTO struct {
	TradeNo     string
	TradeBody   TradeParticipator
	TradeTarget TradeParticipator
	AmountStr   string
	ChangeType  ChangeType
	ChangeFlag  ChangeFlag
	Desc        string
}

//账户创建对象
type AccountCreateDTO struct {
	UserId       string
	Username     string
	AccountName  string
	AccountType  int
	CurrencyCode string
	Amount       string
}

//账户信息
type AccountDTO struct {
	AccountNo    string          //账户编号,账户唯一标识
	AccountName  string          //账户名称,用来说明账户的简短描述,账户对应的名称或者命名，比如xxx积分、xxx零钱
	AccountType  int             //账户类型，用来区分不同类型的账户：积分账户、会员卡账户、钱包账户、红包账户
	CurrencyCode string          //货币类型编码：CNY人民币，EUR欧元，USD美元 。。。
	UserId       string          //用户编号, 账户所属用户
	Username     string          //用户名称
	Balance      decimal.Decimal //账户可用余额
	Status       int             //账户状态，账户状态：0账户初始化，1启用，2停用
	CreatedAt    time.Time       //创建时间
	UpdatedAt    time.Time       //更新时间
}
