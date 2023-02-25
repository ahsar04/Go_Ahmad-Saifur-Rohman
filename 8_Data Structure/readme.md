#section 8 Data Structure#
mempelaraji tentang funtion (kumpulan code yang dibuat agar dapat dijalankan secara dinamis dengan tujuan menghindari pengulangan dalam menuliskan algoritma) dan macam-macam struktur data seperti array, slice, dan map dimana terdapat poin penting untuk membedakan ketiganya yaitu:

- array = kumpulan elemen dengan `tipe data yang sama` dan `panjang index` yang sudah ditentukan contoh `arr:=[3][3]int{{1,2,3},{4,5,6},{9,8,9}}`
- slice = hampir sama dengan array tetapi memiliki `panjang index yang dinamis` contoh `arr:=[][]int{{1,2,3},{4,5,6},{9,8,9}}`
- map = berbeda dengan array dan slice, map memiliki parameter berupa key untuk mengakses data yang diperlukan. tipe data ini biasa digunakan untuk Result dari sebuah API atau biasa disebut JSON contoh `var chickens = []map[string]string{{"name": "chicken blue",   "gender": "male"},{"name": "chicken red",    "gender": "male"},{"name": "chicken yellow", "gender": "female"},}`
