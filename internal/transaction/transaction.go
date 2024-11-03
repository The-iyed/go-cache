package transaction

type Command struct {
	Name string
	Args []string
}

type Transaction struct {
	Commands []Command
	IsActive bool
}

func NewTransaction() *Transaction {
	return &Transaction{
		IsActive: true,
		Commands: []Command{},
	}
}
