#!/bin/bash

# Membuat folder dengan perintah pertama 
main_folder="$1_$at$(date +'%a %b %d %T %Z %Y')"
mkdir "$main_folder"

# Pindah ke folder utama
cd "$main_folder"

# Membuat folder about_me
mkdir about_me
cd about_me

# Membuat subfolder personal
mkdir personal
cd personal

# Membuat file facebook.txt dengan URL dari argumen kedua
echo "$2" > facebook.txt

# Kembali ke folder about_me
cd ..

# Membuat subfolder professional
mkdir professional
cd professional

# Membuat file linkedin.txt dengan URL dari argumen ketiga
echo "$3" > linkedin.txt

# Kembali ke folder utama
cd ../..

# Membuat folder my_friends
mkdir my_friends
cd my_friends

# Mengambil daftar nama teman-teman menggunakan command curl dari URL file yang diberikan
curl -o list_of_my_friends.txt https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt

# Kembali ke folder utama
cd ..

# Membuat folder my_system_info
mkdir my_system_info
cd my_system_info

# Membuat file about_this_laptop.txt dengan nama pengguna dan hasil dari command uname -a
echo "Nama Pengguna: $USER" > about_this_laptop.txt
uname -a >> about_this_laptop.txt

# Membuat file internet_connection.txt dengan hasil keluaran ping ke google.com sebanyak 3 kali
ping -c 3 google.com > internet_connection.txt

# Kembali ke folder utama
cd ..

# Menampilkan pesan bahwa script selesai dijalankan
echo "Script telah selesai dijalankan."
