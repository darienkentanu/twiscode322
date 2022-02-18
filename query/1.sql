create table if not exists data_transaksi(
    id int primary key,
    tanggal_order DATETIME,
    status_pelunasan ENUM('lunas', 'belum lunas'),
    tanggal_bayar DATETIME,
);

insert into data_transaksi('tanggal_order', 'status_pelunasan', 'tanggal_bayar') values
     ('2022-02-17 00:00:00', 'lunas', '2022-02-17 01:00:00'),
     ('2022-02-18 00:00:00' 'belum lunas', null);