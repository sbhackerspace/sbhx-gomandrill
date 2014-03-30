package messages

type MandrillMessage struct {
	Key     string   `json:"key"`
	Message *Message `json:"message"`
	Async   bool     `json:"async"`
}

// TODO(elimisteve): Add the rest of the fields below

type Message struct {
	Html      string       `json:"html"`
	Text      string       `json:"text"`
	Subject   string       `json:"subject"`
	FromEmail string       `json:"from_email"`
	To        []*Recipient `json:"to"`
}
