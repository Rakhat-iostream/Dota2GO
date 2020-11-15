package main

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

}
