package Lesson_19_Structures_and_Methods

import "fmt"

type Account struct {
	Number  string
	Owner   string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
	fmt.Printf("Баланс после пополнения: %.2f\n", a.Balance)
}

func (a *Account) Withdraw(amount float64) {
	if a.Balance < amount {
		fmt.Println("Недостаточно средств")
	} else {
		a.Balance -= amount
		fmt.Printf("Баланс после снятия: %.2f\n", a.Balance)
	}
}

// ------------------------------------------------------------------//
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Position string
}

func (p *Person) PrintInfo() {
	fmt.Printf("Имя: %s, Возраст: %d\n", p.Name, p.Age)
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Имя: %s, Возраст: %d, Должность: %s\n", e.Name, e.Age, e.Position)
}

// ------------------------------------------------------------------//

/*
	Задача 1: Управление счетами клиентов

Вы разрабатываете систему управления банковскими счетами. Каждый счет имеет номер, владельца и баланс.
Создайте структуру Account, которая содержит:

	Number string — номер счета
	Owner string — владелец
	Balance float64 — баланс

Добавьте методы:

	Deposit(amount float64) — пополнение счета.
	Withdraw(amount float64) — снятие денег (если сумма доступна).

Пример использования:

	acc := Account{"123456", "Иван Петров", 1000.00}
	acc.Deposit(500)
	acc.Withdraw(300)

Ожидаемый вывод:

	Баланс после пополнения: 1500.00
	Баланс после снятия: 1200.00

Если попытаться снять больше, чем есть на счету:

	Недостаточно средств
*/
func Task_One() {
	Bond := Account{"007", "Bond", 377.50}

	fmt.Println(Bond.Balance)
	Bond.Deposit(22.5)
	Bond.Withdraw(100.00)
	Bond.Withdraw(500.00)
	fmt.Println(Bond.Balance)
}

/*
Задача 2: Расширение структуры через эмбеддинг

	Создайте две структуры:
		1. Person с полями Name string и Age int.
		2. Employee, которая встраивает Person и добавляет поле Position string.
	Добавьте методы:
			PrintInfo() для Person, который выводит имя и возраст.
			PrintInfo() для Employee, который также выводит должность.
	Пример использования:
		e := Employee{Person{"Анна", 30}, "Разработчик"}
		e.PrintInfo()
	Ожидаемый вывод:
		Имя: Анна, Возраст: 30, Должность: Разработчик
*/
func Task_Two() {
	pavel := Person{"Паша", 40}
	e := Employee{Person{"Анна", 30}, "Разработчик"}

	pavel.PrintInfo()
	e.PrintInfo()
}
