package main

import (
	"fmt"
	"log"
	"os"
)

//Factory Pattern START
//Useful unit
type Meal struct {
	MealName string
	Cost     int
}

func NewMeal(mname string, mcost int) *Meal {
	return &Meal{MealName: mname,
		Cost: mcost}
}

//Product Interface
type IRestaurant interface {
	setName(name string)
	setMenu(meallist []Meal)
	getName() string
	getMenu() []Meal
}

//Concrete template of product
type Restoran struct {
	RestName string
	Menu     []Meal
}

func (r *Restoran) setName(name string) {
	r.RestName = name
}

func (r *Restoran) setMenu(meals []Meal) {
	r.Menu = meals
}

func (r *Restoran) getName() string {
	return r.RestName
}

func (r *Restoran) getMenu() []Meal {
	return r.Menu
}

//Concrete product
type KFC struct {
	Restoran
}

func NewKFC() IRestaurant {
	return &KFC{Restoran: Restoran{RestName: "KFC",
		Menu: []Meal{*NewMeal("Coca-Cola", 250),
			*NewMeal("Fried Chicken", 1200),
			*NewMeal("Shaurma", 700)}}}
}

//Concrete product 2
type BurgerKing struct {
	Restoran
}

func NewBurgerKing() IRestaurant {
	return &BurgerKing{Restoran{RestName: "Burger King",
		Menu: []Meal{*NewMeal("Sprite", 200),
			*NewMeal("Grill", 900),
			*NewMeal("Cheese Burger", 450)}}}
}

//Factory of Restaurants
func getRestaurantsWithMenu(restName string) IRestaurant {
	switch restName {
	case "KFC":
		return NewKFC()
	case "Burger King":
		return NewBurgerKing()
	default:
		return nil
	}
}

//For understandable representation
func MenuOfRestaurant(restaurantName string) {
	menuRest := getRestaurantsWithMenu(restaurantName)
	fmt.Println("Menu of " + menuRest.getName())
	menuItems := menuRest.getMenu()
	fmt.Println("Meal Name - Cost")
	for _, val := range menuItems {
		fmt.Printf("%s - %d \n", val.MealName, val.Cost)
	}
}

//Factory Pattern END

func getMealByName(mealname, restaunrantName string) (Meal, error) {
	menuOfRest := getRestaurantsWithMenu(restaunrantName)
	menuItems := menuOfRest.getMenu()
	for _, val := range menuItems {
		if val.MealName == mealname {
			return val, nil
		}
	}
	return Meal{}, fmt.Errorf("no meal with such name found")
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

func (u *User) HandleChanges(restaurants []IRestaurant) {
	fmt.Printf(
		"Hello %s \n We have some changes in application: \n =============== %s ===============\n", u.Name, restaurants)
}

func NewAccount(name, password, address string) *User {
	return &User{Name: name,
		Password:   password,
		Address:    address,
		Authorized: false} //firstly, login
}

//kind of Strategy pattern for paying
type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
	card Card
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Not enough cash in Wallet")
	}
	w.Cash -= amount
	return nil
}

type Card struct {
	Owner   string
	Balance int
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Not enough balance")
	}
	c.Balance -= amount
	return nil
}

func Buy(p Payer, cartSum int) { //Метод Buy скорее всего реализация в Фасаде в makeOrder()
	switch p.(type) {
	case *Wallet:
		fmt.Println("Okay,here you need to pay... ")

	case *Card:
		debitCard, _ := p.(*Card)
		fmt.Println("Please, wait. Making transaction with your Card", debitCard.Owner)
		err := p.Pay(cartSum)
		if err != nil {
			panic(err)
		}
		fmt.Println("Payment was made")

	default:
		fmt.Println("Something new!")
	}

}

//End of Strategy

type DeliveryFacade struct {
	Account     *User
	FoodService *FoodService
	Wallet      *Wallet
	Card        *Card
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
		Card:        nil,
	}
	fmt.Println("[Application Start]")
	return DeliveryFacade
}

//Observed
type FoodService interface {
	addRestaurant(restaurant IRestaurant)
	removeRestaurant(restaurant IRestaurant)
	showAllRestaurants()
	addObserver(user User)
	removeObserver(user User)
	notifyObservers()
}

func getIndexOfRestaurantInSlice(allRestraunts []IRestaurant, restaurant IRestaurant) int {
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
	restaurants []IRestaurant
	users       []User
}

func (g *Glovo) addRestaurant(restaurant IRestaurant) {
	g.restaurants = append(g.restaurants, restaurant)
}

func (g *Glovo) removeRestaurant(restaurant IRestaurant) {
	counter := getIndexOfRestaurantInSlice(g.restaurants, restaurant)
	g.restaurants = append(g.restaurants[:counter], g.restaurants[counter+1:]...)
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
	restaurants []IRestaurant
	users       []User
}

func (y *YandexFood) addRestaurant(restaurant IRestaurant) {
	y.restaurants = append(y.restaurants, restaurant)
	y.notifyObservers()
}

func (y *YandexFood) removeRestaurant(restaurant IRestaurant) {
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

func CalculateCartTotalSum(cart []Meal) int {
	sum := 0
	for _, val := range cart {
		sum = sum + val.Cost
	}
	return sum
}

func main() {

	deliveryServiceGlovo := &Glovo{}
	deliveryServiceGlovo.addRestaurant(NewBurgerKing())
	deliveryServiceGlovo.addRestaurant(NewKFC())

	deliveryServiceYandex := &YandexFood{}
	deliveryServiceYandex.addRestaurant(NewKFC())

	DelFacade := NewDeliveryFacade()

	var input string
	var choice int
	var myService FoodService

	fmt.Println("Please, choose delivery service")
	switch input {

	}
	fmt.Println("Welcome to the Food Delivery Service")

	//DelFacade.RegisterUser("User", "Root", "NI st")

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
			myService.notifyObservers()
			goto home
		case choice == 2:
		showmenu:
			fmt.Println("Which Restaurant Menu do you want to see\n" +
				"1 - KFC \n" +
				"2 - Burger King")
			fmt.Fscan(os.Stdin, &choice)
			var restName string
			switch choice {
			case 1:
				MenuOfRestaurant("KFC")
				restName = "KFC"
			case 2:
				MenuOfRestaurant("Burger King")
				restName = "Burger King"
			default:
				fmt.Println("Choose from the list of Restaurants")
				goto showmenu
			}
			korzina := []Meal{}
		ordering:
			fmt.Println("Please, type the name of Food")
			fmt.Fscan(os.Stdin, &input)
			meal, err := getMealByName(input, restName)
			if err != nil {
				goto ordering
			}
			korzina = append(korzina, meal)
		choosingOneMore:
			fmt.Println("Do you want to add one more meal?\n" +
				"1 - Yes\n" +
				"2 - No")
			fmt.Fscan(os.Stdin, &choice)
			switch choice {
			case 1:
				goto ordering
			case 2:
				//Making order
				cartSum := CalculateCartTotalSum(korzina)
				var payMethod string
			paymentChoice:
				fmt.Println("Will you pay in cash or by card? 1/2?\n")
				fmt.Fscan(os.Stdin, &payMethod)
				switch {
				case payMethod == "1":
					myWallet := DelFacade.Wallet
					Buy(myWallet, cartSum)
				case payMethod == "2":
					myCard := &Card{Balance: DelFacade.Wallet.card.Balance, Owner: DelFacade.Account.Name}
					Buy(myCard, cartSum)
				default:
					goto paymentChoice
				}

			default:
				goto choosingOneMore
			}

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

	/*
		kfc := getRestaurantsWithMenu("KFC")
		fmt.Println("Menu for "+kfc.getName())
		fmt.Println(kfc.getMenu())
	*/
}
