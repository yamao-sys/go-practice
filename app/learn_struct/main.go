package main

import "fmt"

type User struct {
	Name string
	Age  int
	// X, Y int
}

type T struct {
	User User
}

type Users []*User

// type Users struct {
// 	Users []*User
// }

// 構造体は値型なので、単に渡すと値渡し
func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

// Userの参照渡し
func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

// User型のオブジェクトをレシーバとしてメソッド定義
func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

// User型オブジェクトのコンストラクタ
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

// 独自型とそのレシーバをとるメソッド
type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	// var user1 User // NameとAgeの初期値はそれぞれの型のデフォルト値(Name: '', Age: 0)
	// fmt.Println(user1)

	// user1.Name = "user1"
	// user1.Age = 10
	// fmt.Println(user1)

	// user2 := User{} // 暗黙的な宣言
	// fmt.Println(user2)

	// user2.Name = "user2"
	// fmt.Println(user2)

	// user3 := User{Name: "user3", Age: 30}
	// fmt.Println(user3)

	// user4 := User{"user4", 40}
	// fmt.Println(user4)

	// // user5 := User{30, "user5"}
	// // fmt.Println(user5)

	// user6 := User{Name: "user6"}
	// fmt.Println(user6)

	// user7 := new(User) // newで生成すると、User型オブジェクトのポインタを返す
	// fmt.Println(user7)

	// user8 := &User{} // アドレス演算子でUser型のオブジェクトのポインタを返す
	// fmt.Println(user8)

	// UpdateUser(user1)
	// UpdateUser2(user8)
	// fmt.Println(user1)
	// fmt.Println(user8)

	// user1.SayName()

	// user1.SetName("A")
	// user1.SayName()

	// user1.SetName2("A")
	// user1.SayName()

	// user9 := &User{Name: "user2"}
	// user9.SetName2("B")
	// user9.SayName()

	// t := T{User: User{Name: "user1", Age: 10}}
	// fmt.Println(t)
	// fmt.Println(t.User)
	// fmt.Println(t.User.Name)
	// // fmt.Println(t.Name)
	// t.User.SayName()

	// t.User.SetName2("A")
	// t.User.SayName()

	// user1 := NewUser("user1", 10)
	// fmt.Println(user1)
	// fmt.Println(*user1)

	// user1 := User{Name: "user1", Age: 10}
	// user2 := User{Name: "user2", Age: 20}
	// user3 := User{Name: "user3", Age: 30}
	// user4 := User{Name: "user4", Age: 40}

	// users := Users{}
	// users = append(users, &user1, &user2, &user3, &user4)
	// fmt.Println(users)

	// for _, v := range users {
	// 	fmt.Println(*v)
	// }

	// users2 := make([]*User, 0)
	// users2 = append(users2, &user1, &user2, &user3, &user4)
	// for _, v := range users2 {
	// 	fmt.Println(*v)
	// }

	// m := map[int]User{
	// 	1: {Name: "user1", Age: 10},
	// 	2: {Name: "user2", Age: 20},
	// }
	// fmt.Println(m)

	// m2 := map[User]string{
	// 	{Name: "user1", Age: 10}: "Tokyo",
	// 	{Name: "user2", Age: 20}: "LA",
	// }
	// fmt.Println(m2)

	// m3 := make(map[int]User)
	// fmt.Println(m3)
	// m3[1] = User{Name: "user3", Age: 30}
	// fmt.Println(m3)

	// for _, v := range m {
	// 	fmt.Println(v)
	// }

	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	// i := 100
	// fmt.Println(mi + i) // MyIntとIntの演算はNG
	mi.Print()
}
