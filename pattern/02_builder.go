package pattern

/*
Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

Применимость:
	- Когда необходимо избавиться от конструктора с огромным количеством опциональных параметров
	- Когда необходимо создавать разные представления какого-то объекта
	- Когда необходимо создать сложный составной объект

Плюсы:
	- Позволяет создавать объекты пошагово
	- Позволяет использовать один и тот же код для создания различных объектов
	- Изолирует сложный код сборки объекта от его основной бизнес-логики.

Минусы:
	- Усложняет код программы из-за введения дополнительных классов
*/

type Response struct {
	user    User
	profile Profile
	address Address
}

type User struct {
	email string
}

type Profile struct {
	balance float64
}

type Address struct {
	city string
}

type IBuilder interface {
	SetEmail(string)
	SetBalance(float64)
	SetCity(string)
	Build() *Response
}

type ResponseBuilder struct {
	email   string
	balance float64
	city    string
}

func NewResponseBuilder() IBuilder {
	return &ResponseBuilder{}
}

func (rb *ResponseBuilder) SetEmail(email string) {
	rb.email = email
}

func (rb *ResponseBuilder) SetBalance(balance float64) {
	rb.balance = balance
}

func (rb *ResponseBuilder) SetCity(city string) {
	rb.city = city
}

func (rb *ResponseBuilder) Build() *Response {

	return &Response{User{email: rb.email}, Profile{balance: rb.balance}, Address{city: rb.city}}
}

func builderPattern() {
	builder := NewResponseBuilder()
	builder.SetEmail("example@example.com")
	builder.SetBalance(123.4)
	_ = builder.Build()
}
