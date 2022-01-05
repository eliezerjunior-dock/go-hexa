package core

type BalanceServicePort interface {
	ListBalance() ([]Balance ,error)
	GetBalance(account string) (Balance ,error)
	AddBalance(balance Balance) (error)
}

type BalanceRepositoryPort interface {
	ListBalance() ([]Balance ,error)
	AddBalance(balance Balance) (error)
	GetBalance(account string) (Balance ,error)
}

type BalanceClientPort interface {
 	GetRate(account string) (int32 ,error)
 }