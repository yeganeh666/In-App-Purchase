package db

type DataBase interface {
	ConnectToDB() error
	GetAmount([]string) string
	InsertTransactions(string)
}
