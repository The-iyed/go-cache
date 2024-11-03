package transaction

type Command struct {
	Args []string
}

type Transaction struct {
	IsActive bool
	Commands []Command
}

func NewTransaction() *Transaction {
	return &Transaction{
		IsActive: true,
		Commands: make([]Command, 0),
	}
}
