# MIDDLEWARE

`Middleware` pada Echo Framework adalah fungsi yang dipanggil `sebelum` atau `setelah` request diteruskan ke `handler`. Middleware memungkinkan kita untuk memodifikasi `request` dan `response`, melakukan `validasi`, `otentikasi`, `logging`, serta `mengubah format` data, dan lainnya. Pada dasarnya, middleware adalah `lapisan` tambahan antara client dan server yang memproses permintaan dan/atau menambahkan fitur-fitur tertentu sebelum request sampai ke handler dan response dikirim kembali ke client.

Echo Framework memiliki middleware yang sudah tersedia di dalamnya seperti:

- `Logger` Middleware, untuk menampilkan `log` setiap request yang masuk.
- `Recover` Middleware, untuk menangani `error` yang terjadi selama request dijalankan.
- `CORS` Middleware, untuk mengaktifkan `CORS` (Cross-Origin Resource Sharing) sehingga aplikasi bisa menerima request dari domain lain.
- `JWT` Middleware, untuk melakukan `autentikasi` menggunakan JSON Web Tokens.
- `Gzip` Middleware, untuk melakukan `kompresi gzip` pada response yang dikirimkan.

Selain middleware yang sudah disediakan, kita juga bisa membuat middleware custom sendiri. Middleware custom dibuat dengan mengimplementasikan fungsi yang memiliki parameter `echo.HandlerFunc`, dan mengembalikan `echo.HandlerFunc`. Fungsi tersebut nantinya akan dipanggil ketika request masuk.
