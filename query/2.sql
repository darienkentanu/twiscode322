create table if not exists detail_transaksi(
    id int primary key,
    id_transaksi int not null,
    harga decimal(20,2),
    jumlah int,
    subtotal decimal(20,2)
    foreign key(id_transaksi) references transaksi(id)
);

insert into detail_transaksi ('id_transaksi', 'harga', 'jumlah', 'subtotal') values
(1, 34000, 2, 24000),
(2, 500000, 1, 50000);