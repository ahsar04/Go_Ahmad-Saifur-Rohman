package main

// gunakn camel case untuk nama variabel
type user struct {
	id       int
	username int
	password int
}

// penulisan harus menggunakan camelCase (userService)
type userservice struct {
	// penamaan variabel harus jelas User [] User
	t []user
}

// penulisan harus menggunakan camelCase (getAllUser)
// func (u *UserService) GetAllUsers() []User {
func (u userservice) getallusers() []user {
	// u.User
	return u.t
}

// penulisan harus menggunakan camelCase (getUserBy)
// func (u *UserService) GetUserByID(id int) *User {
func (u userservice) getuserbyid(id int) user {
	// for i := range u.Users {
	// 	if u.Users[i].ID == id {
	// 		return &u.Users[i]
	// 	}
	// }
	// return nil
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}
	return user{}
}

// terdapat 4 kekurangan dalam penulisan kode diatas antara lain
// 1. nama variabel tidak ditulis mengunakan camel case
// 2. Nama variabel yang tidak jelas: Dalam kode yang kompleks, penggunaan nama variabel yang jelas dan deskriptif akan sangat membantu membaca dan memahami kode dengan lebih mudah.
// 3. Penamaan struct dan method yang tidak konsisten: Penamaan yang konsisten membantu memahami struktur dan hubungan antara tipe data dan method.
// 4. Penulisan fungsi yang tidak optimal: Fungsi yang ditulis secara optimal meningkatkan kinerja dan memudahkan untuk diubah atau ditambahkan fungsionalitas di kemudian hari.
