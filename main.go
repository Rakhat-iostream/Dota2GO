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

func (k *KoreanHouse) String() string {
	return fmt.Sprintf("Korean House")
}

type KFC struct {
	menu map[string]int
}

func (k *KFC) getMenu() map[string]int {
	k.menu["box master"] = 1850
	k.menu["zinger"] = 1500
	k.menu["twister"] = 1300
	return k.menu
}

func (k *KFC) String() string {
	return fmt.Sprintf("KFC")
}

type DodoPizza struct {
	menu map[string]int
}

func (d *DodoPizza) getMenu() map[string]int {
	d.menu["cheese pizza"] = 2300
	d.menu["tomato pizza"] = 1900
	d.menu["sausage pizza"] = 2600
	return d.menu
}

func (d *DodoPizza) String() string {
	return fmt.Sprintf("Dodo Pizza")
}

//Observer
type User struct {
	Name       string
	Password   string
	Address    string
	Authorized bool //used to control authorization
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

func (u *User) HandleChanges(restaurants []Restaurant) {
	fmt.Printf(
		"Hello %s \n We have some changes in application: \n =============== %s ===============\n", u.Name, restaurants)
}

func NewAccount(name, password, address string) *User {
	return &User{Name: name,
		Password:   password,
		Address:    address,
		Authorized: false} //firstly, login
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

func (dF *DeliveryFacade) Login(name, pass string) error {
	fmt.Println("Processing User Details [Authorization]")
	error := dF.Account.authorize(name, pass)
	if error != nil {
		log.Fatalf("Errors: %s\n", error.Error())
	}
	fmt.Printf("Welcome, %s! You are in System.", dF.Account.Name)
	return nil
}

func (dF *DeliveryFacade) RegisterUser(uname, upassword, uaddress string) {
	dF.Account = NewAccount(uname, upassword, uaddress)
	fmt.Println("You are successfully registered! Please, Login")
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

//Observed
type FoodService interface {
	addRestaurant(restaurant Restaurant)
	removeRestaurant(restaurant Restaurant)
	showAllRestaurants()
	addObserver(user User)
	removeObserver(user User)
	notifyObservers()
}

func getIndexOfRestaurantInSlice(allRestraunts []Restaurant, restaurant Restaurant) int {
	for i, v := range allRestraunts {
		if v == restaurant {
			return i
		}
	}
	return -1
}

func getIndexOfObserverInSlice(allUsers []User, user User) int {
	for i, v := range allUsers {
		if v == user {
			return i
		}
	}
	return -1
}

type Glovo struct {
	restaurants []Restaurant
	users       []User
}

func (g *Glovo) addRestaurant(restaurant Restaurant) {
	g.restaurants = append(g.restaurants, restaurant)
	g.notifyObservers()
}

func (g *Glovo) removeRestaurant(restaurant Restaurant) {
	counter := getIndexOfRestaurantInSlice(g.restaurants, restaurant)
	g.restaurants = append(g.restaurants[:counter], g.restaurants[counter+1:]...)
	g.notifyObservers()
}

func (g *Glovo) showAllRestaurants() {
	for _, restaurant := range g.restaurants {
		fmt.Println(restaurant)
	}
}

func (g *Glovo) addObserver(user User) {
	g.users = append(g.users, user)
	//TODO message in facade that user was added
}

func (g *Glovo) removeObserver(user User) {
	counter := getIndexOfObserverInSlice(g.users, user)
	g.users = append(g.users[:counter], g.users[counter+1:]...)
}

func (g *Glovo) notifyObservers() {
	for _, v := range g.users {
		v.HandleChanges(g.restaurants)
	}
}

type YandexFood struct {
	restaurants []Restaurant
	users       []User
}

func (y *YandexFood) addRestaurant(restaurant Restaurant) {
	y.restaurants = append(y.restaurants, restaurant)
	y.notifyObservers()
}

func (y *YandexFood) removeRestaurant(restaurant Restaurant) {
	counter := getIndexOfRestaurantInSlice(y.restaurants, restaurant)
	y.restaurants = append(y.restaurants[:counter], y.restaurants[counter+1:]...)
	y.notifyObservers()
}

func (y *YandexFood) showAllRestaurants() {
	for _, restaurant := range y.restaurants {
		fmt.Println(restaurant)
	}
}

func (y *YandexFood) addObserver(user User) {
	y.users = append(y.users, user)
	//TODO message in facade that user was added
}

func (y *YandexFood) removeObserver(user User) {
	counter := getIndexOfObserverInSlice(y.users, user)
	y.users = append(y.users[:counter], y.users[counter+1:]...)
}

func (y *YandexFood) notifyObservers() {
	for _, v := range y.users {
		v.HandleChanges(y.restaurants)
	}
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
	//application := &Glovo{}
	//
	//user1 := NewAccount("Duman", "asd", "address")
	//user2 := NewAccount("Maga", "asd", "address")
	//
	//application.addObserver(*user1)
	//application.addObserver(*user2)
	//
	//application.addRestaurant(&KoreanHouse{})

}
