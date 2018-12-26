package twitchbot

type Trigger struct {
	Command  string
	Output   string
	Function interface{}
}

// Execute the trigger
func (t *Trigger) Execute() {
	//t.Function()
}
