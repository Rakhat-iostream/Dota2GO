package main

import (
	"fmt"
	"os"
)

//Abstract Factory
type Restaurant interface {
	create(factoryType string) RestaurantFactory
}

type RestaurantFactory interface {
	create(foodName string) iFood
}

type iFood interface {
	getName() string
}

type foodFactory struct {
}

func (f foodFactory) create(kh string) RestaurantFactory {
	fastFood := map[string]RestaurantFactory{
		"meal": mealFactory{},
	}
	return fastFood[kh]
}

type mealFactory struct {
}

func (mf mealFactory) create(foodName string) iFood {
	return meal{
		food: food{
			name: foodName,
		},
	}
}

type food struct {
	name string
}

type meal struct {
	food
}

func (m meal) getName() string {
	return m.food.name
}

type KfcFactory struct {
}

func (k KfcFactory) create(kfc string) iFood {
	return chicken{
		food: food{
			name: kfc,
		},
	}
}

type chicken struct {
	food
}

func (c chicken) getName() string {
	return c.food.name
}

type DodoPizzaFactory struct {
}

func (d DodoPizzaFactory) create(name string) iFood {
	return pizza{
		food: food{
			name: name,
		},
	}
}

type pizza struct {
	food
}

func (p pizza) getName() string {
	return p.food.name
}

//END OF FACTORY
//Observer
type User struct {
	Name     string
	Password string
	Address  string
}

//TODO
//func (a *User) editAccount() {
//	fmt.Scan(a.Name)
//	fmt.Scan(a.Password)
//
//}

func NewAccount(name, password, address string) *User {
	return &User{Name: name, Password: password, Address: address}
}

type Wallet struct {
	CardBalance float64
}

func (w *Wallet) AddMoneyToBalance(amount float64) {
	w.CardBalance += amount
}

func (w *Wallet) MakeTransaction(cost float64) {
	w.CardBalance -= cost
}

func NewWallet() *Wallet {
	return &Wallet{CardBalance: 0}
}

type DeliveryFacade struct {
	Account     *User
	FoodService *FoodService
	Wallet      *Wallet
}

//TODO
func (d *DeliveryFacade) makeOrder() error {

	return nil
}

//TODO
func NewDeliveryFacade() *DeliveryFacade {
	DeliveryFacade := &DeliveryFacade{
		Account:     nil,
		FoodService: nil,
		Wallet:      nil,
	}
	return DeliveryFacade
}

//Observed
type FoodService interface {
	addRestaurant(restaurant Restaurant)
	removeRestaurant(restaurant Restaurant)
	listAllRestaurants()
	addUser(user *User)
	removeUser(user *User)
	NotifyObservers()
}

func getIndexOfElementInSlice(allRestraunts []Restaurant, restaurant Restaurant) int {
	for i, v := range allRestraunts {
		if v == restaurant {
			return i
		}
	}
	return -1
}

type Glovo struct {
	restaurant []Restaurant
	users      []User
}

func (g *Glovo) addRestaurant(restaurant Restaurant) {
	panic("implement me")
}

func (g *Glovo) removeRestaurant(restaurant Restaurant) {
	panic("implement me")
}

func (g *Glovo) listAllRestaurants() {
	panic("implement me")
}

func (g *Glovo) addUser(user *User) {
	panic("implement me")
}

func (g *Glovo) removeUser(user *User) {
	panic("implement me")
}

func (g *Glovo) NotifyObservers() {
	panic("implement me")
}

type YandexFood struct {
	news  []string
	users []User
}

func (y *YandexFood) addRestaurant(restaurant Restaurant) {
	panic("implement me")
}

func (y *YandexFood) removeRestaurant(restaurant Restaurant) {
	panic("implement me")
}

func (y *YandexFood) listAllRestaurants() {
	panic("implement me")
}

func (y *YandexFood) addUser(user *User) {
	panic("implement me")
}

func (y *YandexFood) removeUser(user *User) {
	panic("implement me")
}

func (y *YandexFood) NotifyObservers() {
	panic("implement me")
}

func main() {
	var rest Restaurant = foodFactory{}
	var foodFact RestaurantFactory = rest.create("soup")
	var ramen iFood = foodFact.create("ramen")
	fmt.Println(ramen.getName()) //NULL POINTER

	var name string
	var password string
	var input string
	var choice int

	fmt.Println("Welcome to the Food Delivery Service")
auth: //authorization event
	fmt.Println("Do you have account? y/n")
	fmt.Fscan(os.Stdin, &input)

	switch {
	case input == "y":
		fmt.Println("Enter name & password")
		fmt.Fscan(os.Stdin, &name)
		fmt.Fscan(os.Stdin, &password)
		authorized := true
		if authorized {
		home: //home/start point
			fmt.Printf("Welcome %s", name+"! Choose... \n"+
				"1 - Show me notifications \n"+
				"2 - Order food \n"+
				"3 - Account settings \n")

			fmt.Fscan(os.Stdin, &choice)
			switch {
			case choice == 1:
				fmt.Println("All notifications here")
			case choice == 2:
				fmt.Println("Outputting menu... Choose food ID")
				fmt.Fscan(os.Stdin, &input)
				fmt.Printf("You have choosen %s", input)
			case choice == 3:
				fmt.Println("Account settings like change wallet, address")
			default:
				goto home
				break
			}
		} else {
			goto auth
		}
	case input == "n":
		fmt.Println("Account creation")
	default:
		fmt.Println("Nothing choosen!")
	}
}
