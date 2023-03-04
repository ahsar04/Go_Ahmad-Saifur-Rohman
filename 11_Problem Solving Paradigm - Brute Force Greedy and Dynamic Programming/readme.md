# Golang Problem Solving Paradigm - Brute Force, Greedy, and Dynamic Programming # 
- `Brute Force`
Brute Force adalah metode pemecahan masalah dengan cara mencoba `semua kemungkinan` solusi yang ada hingga ditemukan solusi yang benar atau optimal. Metode ini paling mudah diimplementasikan, namun sangat tidak `efisien` karena harus mengecek semua kemungkinan solusi yang ada.
Dalam Golang, Brute Force dapat diimplementasikan dengan cara melakukan `nested loop` untuk mencoba semua kemungkinan solusi yang ada.
- `Greedy`
Greedy adalah metode pemecahan masalah dengan cara memilih solusi yang `paling optimal` pada saat itu tanpa mempertimbangkan konsekuensi di masa depan. Metode ini lebih cepat dibandingkan Brute Force, namun hasilnya `tidak selalu optimal`.
- `Dynamic Programming`
Dynamic Programming adalah metode pemecahan masalah dengan cara `memecah` masalah besar menjadi submasalah yang lebih kecil dan menyelesaikannya secara terpisah. Solusi dari submasalah tersebut kemudian digunakan untuk menyelesaikan masalah yang `lebih besar`.
Dalam Golang, Dynamic Programming dapat diimplementasikan dengan cara membuat sebuah `fungsi rekursif` yang memecah masalah menjadi submasalah yang lebih kecil dan menyimpan solusi dari submasalah tersebut.