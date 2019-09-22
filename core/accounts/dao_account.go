package accounts

import (
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
)

type AccountDao struct {
	runner *dbx.TxRunner //仅可本包访问
}

//查询数据库持久化单条记录
func (dao *AccountDao) GetOne(accountNo string) *Account {

	a := &Account{AccountNo: accountNo}
	ok, err := dao.runner.GetOne(a)

	if err != nil {
		logrus.Error(err)
		return nil
	}

	if !ok {
		return nil
	}

	return a
}

//通过用户ID和账户类型查询账户信息
func (dao *AccountDao) GetByUserId(userId string, accountType int) *Account {

	a := &Account{}
	sql := " select * from account where user_id = ? and account_type = ?"
	ok, err := dao.runner.Get(a, sql, userId, accountType)

	if err != nil {
		logrus.Error(err)
		return nil
	}

	if !ok {
		return nil
	}

	return a

}

//账户数据的插入
func (dao *AccountDao) Insert(a *Account) (id int64, err error) {

	rs, err := dao.runner.Insert(a)

	if err != nil {
		return 0, err
	}

	return rs.LastInsertId()
}

//账户余额的更新
func (dao *AccountDao) UpdateBalance(accountNo string, amount decimal.Decimal) (row int64, err error) {

	sql := "update account " +
		" set balance=balance+CAST(? AS DECIMAL(30,6))" +
		" where account_no=? " +
		" and balance>=-1*CAST(? AS DECIMAL(30,6)) "

	rs, err := dao.runner.Exec(sql, amount, accountNo, amount)

	if err != nil {
		return 0, err
	}

	return rs.RowsAffected()

}

//账户状态的更新操作
func (dao *AccountDao) UpdateStatus(accountNo string, status int) (rows int64, err error) {

	sql := "update account set status=? " +
		" where account_no=? "

	rs, err := dao.runner.Exec(sql, status, accountNo)

	if err != nil {
		return 0, err
	}

	return rs.RowsAffected()

}
