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
	Cost     float64
}

func NewMeal(mname string, mcost float64) *Meal {
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
type Restaurant struct {
	RestName string
	Menu     []Meal
}

func (r *Restaurant) setName(name string) {
	r.RestName = name
}

func (r *Restaurant) setMenu(meals []Meal) {
	r.Menu = meals
}

func (r *Restaurant) getName() string {
	return r.RestName
}

func (r *Restaurant) getMenu() []Meal {
	return r.Menu
}

//Concrete product
type KFC struct {
	Restaurant
}

func NewKFC() IRestaurant {
	return &KFC{Restaurant: Restaurant{RestName: "KFC",
		Menu: []Meal{*NewMeal("Coca-Cola", 250),
			*NewMeal("Fried_Chicken", 1200),
			*NewMeal("BoxMaster", 1450),
			*NewMeal("Twister", 1150),
			*NewMeal("Zinger", 1250),
		}}}
}

func (k *KFC) String() string {
	return fmt.Sprintf("KFC \n")
}

//Concrete product 2
type BurgerKing struct {
	Restaurant
}

func NewBurgerKing() IRestaurant {
	return &BurgerKing{Restaurant{RestName: "Burger King",
		Menu: []Meal{*NewMeal("Sprite", 200),
			*NewMeal("Cheese_Burger", 350),
			*NewMeal("Crispy_Chicken", 650),
			*NewMeal("Nuggets", 900),
			*NewMeal("Wopper", 1100),
		}}}
}

func (b *BurgerKing) String() string {
	return fmt.Sprintf("Burger_King \n")
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
		fmt.Printf("%s - %.2f \n", val.MealName, val.Cost)
	}
}

//Factory Pattern END

func getMealByName(mealName, restaurantName string) (Meal, error) {
	menuOfRest := getRestaurantsWithMenu(restaurantName)
	menuItems := menuOfRest.getMenu()
	for _, val := range menuItems {
		if val.MealName == mealName {
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
		"Hello %s \n There are our available restaurants: \n %s \n", u.Name, restaurants)
}

func NewAccount(name, password, address string) *User {
	return &User{Name: name,
		Password:   password,
		Address:    address,
		Authorized: false} //firstly, login
}

//kind of Strategy pattern for paying
type Payer interface {
	AddMoney(amount float64)
	Pay(amount float64) error
}

type Wallet struct {
	Cash float64
	Card Card
}

func (w *Wallet) AddMoney(amount float64) {
	w.Cash += amount
}

func (w *Wallet) Pay(amount float64) error {
	if w.Cash < amount {
		return fmt.Errorf("not enough cash in Wallet")
	}
	w.Cash -= amount
	return nil
}

type Card struct {
	Owner   string
	Balance float64
}

func (c *Card) AddMoney(amount float64) {
	c.Balance += amount
}

func (c *Card) Pay(amount float64) error {
	if c.Balance < amount {
		return fmt.Errorf("not enough balance")
	}
	c.Balance -= amount
	return nil
}

func Buy(p Payer, cartSum float64) { //Метод Buy скорее всего реализация в Фасаде в makeOrder()
	switch p.(type) {
	case *Wallet:
		fmt.Println("Okay,here you need to pay... ")
		//TODO
		fmt.Println("Payment was made")

	case *Card:
		debitCard, _ := p.(*Card)
		fmt.Println("Please, wait. Making transaction with your Card", debitCard.Owner)
		err := p.Pay(cartSum)
		if err != nil {
		addingMoney:
			fmt.Printf("You need to add %.2f more money to balance, choose amount:", cartSum-debitCard.Balance)
			var amount float64
			fmt.Fscan(os.Stdin, &amount)
			if amount < cartSum-debitCard.Balance {
				goto addingMoney
			}
			debitCard.AddMoney(amount)
			fmt.Println("Money was successfully added to balance")
			fmt.Printf("You have: %.2f \n", debitCard.Balance)
			fmt.Println("Payment was made")
			_ = p.Pay(cartSum)
			fmt.Printf("Now you have: %.2f \n", debitCard.Balance)
		}
	default:
		fmt.Println("Something new!")
	}
}

func ShowAllDeliveryServices() {
	fmt.Println("Glovo")
	fmt.Println("YandexFood")
}

func ShowAllCouriers() {
	fmt.Println("beginner")
	fmt.Println("experienced")
	fmt.Println("master")
}

//End of Strategy

type DeliveryFacade struct {
	User            User
	DeliveryService DeliveryService
	Wallet          Wallet
	Card            Card
}

func (dF *DeliveryFacade) Login(name, pass string) error {
	fmt.Println("Processing User Details [Authorization]")
	err := dF.User.authorize(name, pass)
	if err != nil {
		log.Fatalf("Errors: %s\n", err.Error())
	}
	fmt.Printf("Welcome, %s! You are in System.", dF.User.Name)
	return nil
}

func (dF *DeliveryFacade) RegisterUser(uname, upassword, uaddress string) {
	dF.User = *NewAccount(uname, upassword, uaddress)
	dF.DeliveryService.addObserver(dF.User)
	fmt.Println("You are successfully registered! Please, Login")
}

//TODO
func (dF *DeliveryFacade) makeOrder(card Card, cart []Meal, cartSum float64) error {
	dF.Card = card
	var input string
choosingCourier:
	fmt.Println("Please choose your courier:")
	ShowAllCouriers()
	var courier Courier
	fmt.Fscan(os.Stdin, &input)
	switch input {
	case "beginner":
		courier = &OrdinaryCourier{}
	case "experienced":
		courier = NewExperiencedCourier(&OrdinaryCourier{})
		cartSum = cartSum + cartSum*0.1
	case "master":
		courier = NewMagisterCourier(NewExperiencedCourier(&OrdinaryCourier{}))
		cartSum = cartSum + cartSum*0.15
	default:
		goto choosingCourier
	}
	fmt.Println("Courier's been chosen")
	var orderMakingAmountOfTimeInMinutes int
	for range cart {
		orderMakingAmountOfTimeInMinutes += 5
	}
	Buy(&dF.Card, cartSum)
	fmt.Printf("We recieved your order, it'll be ready in %d minutes \n", orderMakingAmountOfTimeInMinutes)
	fmt.Printf("Courier is delivering your order to %v address \n", dF.User.Address)
	fmt.Println("Hello, Look at the check")
	for _, meal := range cart {
		fmt.Printf("Name: %s Cost: %.2f \n", meal.MealName, meal.Cost)
	}
	fmt.Printf("Total sum of the check with courier additional percentage: %.2f \n", cartSum)
	fmt.Println(courier.GiveOrderToClient())
	return nil
}

//TODO
func NewDeliveryFacade(deliveryService *DeliveryService) *DeliveryFacade {
	DeliveryFacade := &DeliveryFacade{
		DeliveryService: *deliveryService,
	}
	return DeliveryFacade
}

//Observed
type DeliveryService interface {
	showAllRestaurants()
	addRestaurant(restaurant IRestaurant)
	removeRestaurant(restaurant IRestaurant)
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

func (g *Glovo) showAllRestaurants() {
	for _, restaurant := range g.restaurants {
		fmt.Println(restaurant)
	}
}

func (g *Glovo) addRestaurant(restaurant IRestaurant) {
	g.restaurants = append(g.restaurants, restaurant)
}

func (g *Glovo) removeRestaurant(restaurant IRestaurant) {
	counter := getIndexOfRestaurantInSlice(g.restaurants, restaurant)
	g.restaurants = append(g.restaurants[:counter], g.restaurants[counter+1:]...)
}

func (g *Glovo) addObserver(user User) {
	g.users = append(g.users, user)
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
}

func (y *YandexFood) removeRestaurant(restaurant IRestaurant) {
	counter := getIndexOfRestaurantInSlice(y.restaurants, restaurant)
	y.restaurants = append(y.restaurants[:counter], y.restaurants[counter+1:]...)
}

func (y *YandexFood) showAllRestaurants() {
	for _, restaurant := range y.restaurants {
		fmt.Println(restaurant)
	}
}

func (y *YandexFood) addObserver(user User) {
	y.users = append(y.users, user)
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

func CalculateCartTotalSum(cart []Meal) float64 {
	var sum float64
	for _, meal := range cart {
		sum = sum + meal.Cost
	}
	return sum
}

//Decorator Courier

//Component
type Courier interface {
	GiveOrderToClient() string
}

//Target
type OrdinaryCourier struct {
}

func (g *OrdinaryCourier) GiveOrderToClient() string {
	return "Bonjour, here is your order! \n"
}

//Decorators:

type ExperiencedCourier struct {
	Courier Courier
}

func (s *ExperiencedCourier) GiveOrderToClient() string {
	return s.Courier.GiveOrderToClient() + s.sayBonAppetit()
}

func (s *ExperiencedCourier) sayBonAppetit() string {
	return "Bon appetit!\n"
}

func NewExperiencedCourier(courier Courier) *ExperiencedCourier {
	return &ExperiencedCourier{Courier: courier}
}

type MagisterOfAllCouriers struct {
	Courier Courier
}

func (t *MagisterOfAllCouriers) GiveOrderToClient() string {
	return t.Courier.GiveOrderToClient() + t.be4sv()
}

func (t *MagisterOfAllCouriers) be4sv() string {
	return "Za svoyu uletnost' deneg ne beru! (Slamming the door)\n"
}

func NewMagisterCourier(courier Courier) *MagisterOfAllCouriers {
	return &MagisterOfAllCouriers{Courier: courier}
}

// End of decorator

func main() {

	deliveryServiceGlovo := &Glovo{}
	deliveryServiceGlovo.addRestaurant(NewBurgerKing())
	deliveryServiceGlovo.addRestaurant(NewKFC())

	deliveryServiceYandex := &YandexFood{}
	deliveryServiceYandex.addRestaurant(NewKFC())

	var input string
	var choice int
	var myService DeliveryService
	fmt.Println("[Application Start]")
choosingDeliveryService:
	fmt.Println("Please, choose delivery service from list:")
	ShowAllDeliveryServices()
	fmt.Fscan(os.Stdin, &input)

	switch input {
	case "Glovo":
		myService = &Glovo{}
		myService.addRestaurant(NewKFC())
		myService.addRestaurant(NewBurgerKing())
		fmt.Println("Welcome to the Glovo!")
	case "YandexFood":
		myService = &YandexFood{}
		myService.addRestaurant(NewBurgerKing())
		fmt.Println("Welcome to the Yandex!")
	default:
		goto choosingDeliveryService
	}

	deliveryServiceFacade := NewDeliveryFacade(&myService)

start:
	fmt.Println("Do you have an account? y/n")
	fmt.Fscan(os.Stdin, &input)

	switch input {
	case "y":
	login:
		var name string
		var password string
		fmt.Println("Enter name")
		fmt.Fscan(os.Stdin, &name)
		fmt.Println("Enter password")
		fmt.Fscan(os.Stdin, &password)

		deliveryServiceFacade.Login(name, password)

		if deliveryServiceFacade.User.isAuthorized() != nil {
			goto login
		}

	home:
		fmt.Printf("Home. Choose... \n" +
			"1 - Show me notifications \n" +
			"2 - Order food \n" +
			"3 - Logout\n")

		fmt.Fscan(os.Stdin, &choice)
		switch choice {
		case 1:
			fmt.Println("All notifications here")
			myService.notifyObservers()
			goto home
		case 2:
		showingMenu:
			fmt.Println("Which Restaurant Menu do you want to see")
			myService.showAllRestaurants()
			fmt.Fscan(os.Stdin, &input)
			var restName string
			switch input {
			case "KFC":
				MenuOfRestaurant("KFC")
				restName = "KFC"
			case "Burger_King":
				MenuOfRestaurant("Burger King")
				restName = "Burger King"
			default:
				fmt.Println("Choose from the list of Restaurants")
				goto showingMenu
			}
			var cart []Meal
		ordering:
			fmt.Println("Please, type the name of Food")
			fmt.Fscan(os.Stdin, &input)
			meal, err := getMealByName(input, restName)
			if err != nil {
				goto ordering
			}
			cart = append(cart, meal)
		choosingOneMore:
			fmt.Println("Do you want to add one more meal?\n" +
				"1 - Yes\n" +
				"2 - No")
			fmt.Fscan(os.Stdin, &choice)
			switch choice {
			case 1:
				goto ordering
			case 2:
				cartSum := CalculateCartTotalSum(cart)
				var payingMethod string
			paymentChoice:
				fmt.Println("Will you pay in cash or by card? 1/2?")
				fmt.Fscan(os.Stdin, &payingMethod)
				switch {
				case payingMethod == "1":
					myWallet := deliveryServiceFacade.Wallet
					myWallet.AddMoney(2000)
					//TODO
				case payingMethod == "2":
					myCard := &Card{Balance: deliveryServiceFacade.Wallet.Card.Balance, Owner: deliveryServiceFacade.User.Name}
					myCard.AddMoney(2000)
					deliveryServiceFacade.makeOrder(*myCard, cart, cartSum)
				default:
					goto paymentChoice
				}
			default:
				goto choosingOneMore
			}
			goto home
		case 3:
			deliveryServiceFacade.User.Logout()
			goto choosingDeliveryService
		default:
			goto home
		}

	case "n":
		fmt.Println("User creation")
		var name string
		var password string
		var address string

		fmt.Println("Please, enter your name")
		fmt.Fscan(os.Stdin, &name)

		fmt.Println("Please, enter your password")
		fmt.Fscan(os.Stdin, &password)

		fmt.Println("Please, enter your address")
		fmt.Fscan(os.Stdin, &address)
		deliveryServiceFacade.RegisterUser(name, password, address)
		goto start
	default:
		fmt.Println("Nothing chosen!")
		goto start
	}
}
