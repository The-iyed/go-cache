package transaction

func (t *Transaction) AddCommand(name string, args []string) {
	if t.IsActive {
		t.Commands = append(t.Commands, Command{Name: name, Args: args})
	}
}
