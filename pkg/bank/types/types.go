package types

//Money тип это деньги
type Money int64
 
//Currency тип валюты там наши или русские или всеми любимые америкосы
type Currency string


//Коды валют 
const(
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EURO Currency = "EURO"
)

//PAN  пан?! Кажется это номер карты, если, конечно память не изменяет.
type PAN string 

//MinBalance это минимальный баланс на карте 
type MinBalance Money

//Card тип - структура 
type Card struct{
	ID 			int 
	PAN 			PAN
	Balance 		Money //Использовали Noney
	Currency 		Currency
	Color 			string
	Name 			string 
	Active 			bool
}