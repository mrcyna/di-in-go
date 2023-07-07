package contracts

type EmailInterface interface {
	Send(body string, to string)
}
