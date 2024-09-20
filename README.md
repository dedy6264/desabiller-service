Desabiller adalah mini service yang digunakan untuk menyediakan layanan pembelian dan pembayaran produk biller seperti;
1. Pulsa
2. Tagilan Listrik
3. Token PLN dsb

Service ini menggunakan bahasa Go dalam pembuatannya dan menggunakan framework Echo untuk mempermudah pembuatannya.
Dalam aplikasi ini telah tersedia setingan koneksi database yang basis utama menggunakan Postgre Sql.

Service ini mengimplementasikan multi provider, sehingga memungkinkan untuk integrasi dengan berbagai provider demi memberikan harga yang bersain dan kecepatan proses yang cepat.
harga jual bisa dilakukan murkup atau penyesuaian sehingga bisa menentukan margin atau keuntungan, 
hanya saja untuk harga jual ke merchant masih menggunakan single segment, jadi semua merchant akan mendapatkan harga jual yang sama.

untuk menggunakannya pengguna atau merchant hanya membutuhkan token untuk mengakses semua produk dan melakukan penjualan. sehingga aktifitas lebih aman dan otentik.
