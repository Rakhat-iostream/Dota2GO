package main

import (
	"fmt"
	"log"
	"os"
)

type Restaurant interface {
	getMenu() map[string]int
}

type KoreanHouse struct {
	menu map[string]int
}

func (k *KoreanHouse) getMenu() map[string]int {
	k.menu["ramen"] = 1800
	k.menu["miso soup"] = 2500
	k.menu["spicy meat"] = 3000
	return k.menu
}

type KFC struct {
	menu map[string]int
}

func (k *KFC) getMenu() {
	panic("implement me")
}

type DodoPizza struct {
	menu map[string]int
}

func (k *DodoPizza) getMenu() {
	panic("implement me")
}

//Observer
type User struct {
	Name       string
	Password   string
	Address    string
	Authorized bool //used to control authorization
}

func NewAccount(name, password, address string) *User {
	return &User{Name: name,
		Password:   password,
		Address:    address,
		Authorized: false} //firstly, login
}

func (u *User) authorize(uname, upass string) error {
	if u.Name == uname {
		if u.Password == upass {
			u.Authorized = true
			return nil
		} else {
			return fmt.Errorf("wrong pass")
		}
	} else {
		return fmt.Errorf("wrong login")
	}
}

func (u *User) isAuthorized() error { //Please, use it every time
	//when you use user details
	if u.Authorized == false {
		return fmt.Errorf("not authorized")
	}
	return nil
}

func (u *User) Logout() {
	u.Authorized = false
	fmt.Println("You have logged out")
}

//TODO
//func (a *User) editAccount() {
//	fmt.Scan(a.Name)
//	fmt.Scan(a.Password)
//
//}

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

func (dF *DeliveryFacade) Login(name, pass string) error {
	fmt.Println("Processing User Details [Authorization]")
	error := dF.Account.authorize(name, pass)
	if error != nil {
		log.Fatalf("Errors: %s\n", error.Error())
	}
	fmt.Printf("Welcome, %s! You are in System.", dF.Account.Name)
	return nil
}

//TODO
func (d *DeliveryFacade) makeOrder() error {

	return nil
}

//TODO
func NewDeliveryFacade() *DeliveryFacade {
	DeliveryFacade := &DeliveryFacade{
		//Account:     nil, //after app run, register user or login
		FoodService: nil,
		Wallet:      nil,
	}
	fmt.Println("[Application Start]")
	return DeliveryFacade
}

func (dF *DeliveryFacade) RegisterUser(uname, upassword, uaddress string) {
	dF.Account = NewAccount(uname, upassword, uaddress)
	fmt.Println("You are successfully registered! Please, Login")
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

	var input string
	var choice int

	fmt.Println("Welcome to the Food Delivery Service")
	DelFacade := NewDeliveryFacade()
	DelFacade.RegisterUser("User", "Root", "NI st")

start: //authorization event
	fmt.Println("Do you have account? y/n")
	fmt.Fscan(os.Stdin, &input)

	switch {
	case input == "y":
	login:
		var name string
		var password string
		fmt.Println("Enter name")
		fmt.Fscan(os.Stdin, &name)
		fmt.Println("Enter password")
		fmt.Fscan(os.Stdin, &password)
		//Login
		DelFacade.Login(name, password)

		if DelFacade.Account.isAuthorized() != nil {
			goto login //if not authorized
		}

	home: //home/start point
		fmt.Printf("Home. Choose... \n" +
			"1 - Show me notifications \n" +
			"2 - Order food \n" +
			"3 - Account settings \n" +
			"4 - Logout\n")

		fmt.Fscan(os.Stdin, &choice)
		switch {
		case choice == 1:
			fmt.Println("All notifications here")
			goto home
		case choice == 2:
			fmt.Println("Outputting menu... Choose food ID")
			fmt.Fscan(os.Stdin, &input)
			fmt.Printf("You have choosen %s\n", input)
			goto home
		case choice == 3:
			fmt.Println("Account settings like change wallet, address")
			goto home
		case choice == 4:
			DelFacade.Account.Logout()
			goto start
		default:
			goto home
			break
		}

	case input == "n":
		fmt.Println("Account creation\n")
		var name string
		var password string
		var address string
		//registration:
		fmt.Println("Please, enter your name")
		fmt.Fscan(os.Stdin, &name)

		fmt.Println("Please, enter your password")
		fmt.Fscan(os.Stdin, &password)

		fmt.Println("Please, enter your address")
		fmt.Fscan(os.Stdin, &address)
		DelFacade.RegisterUser(name, password, address)
		goto start //redirecting after successful registration
	default:
		fmt.Println("Nothing choosen!")
		goto start
		break
	}
}
