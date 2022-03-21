-- DML
-- 1.a Insert 5 operators pada table operators.
insert into operators (id, name) values (1, 'fahmy');
insert into operators (id, name) values (2, 'abida');
insert into operators (id, name) values (3, 'asa');
insert into operators (id, name) values (4, 'firdausi');
insert into operators (id, name) values (5, 'eric');
-- 1.b Insert 3 product type.
insert into product_types(id, name) values (1, 'food');
insert into product_types(id, name) values (2, 'drink');
insert into product_types(id, name) values (3, 'snack');
-- 1.c Insert 2 product dengan product type id = 1, dan operators id = 3.
insert into products (id, product_type_id, operator_id, code, name, status)
    values (1, 1, 3, 'Food1', 'Roti Tawar', 1);
insert into products (id, product_type_id, operator_id, code, name, status)
    values (2, 1, 3, 'Food2', 'Taican', 1);
-- 1.d Insert 3 product dengan product type id = 2, dan operators id = 1.
insert into products (id, product_type_id, operator_id, code, name, status)
    values (3, 2, 1, 'Drink1', 'Aqua', 1);
insert into products (id, product_type_id, operator_id, code, name, status)
    values (4, 2, 1, 'Drink2', 'Le Mineral', 1);
insert into products (id, product_type_id, operator_id, code, name, status)
    values (5, 2, 1, 'Drink3', 'Ades', 1);
-- 1.e Insert 3 product dengan product type id = 3, dan operators id = 4.
insert into products (id, product_type_id, operator_id, code, name, status)
    values (6, 3, 4, 'Snack1', 'Kuwaci', 1);
insert into products (id, product_type_id, operator_id, code, name, status)
    values (7, 3, 4, 'Snack2', 'Yupi', 1);
insert into products (id, product_type_id, operator_id, code, name, status)
    values (8, 3, 4, 'Snack3', 'Goodtime', 1);
-- 1.f Insert product description pada setiap product.
insert into product_descriptions (id, description) values (1, 'Tepung tapi enak');
insert into product_descriptions (id, description) values (2, 'Raw meat with sauce');
insert into product_descriptions (id, description) values (3, 'minuman menyegarkan');
insert into product_descriptions (id, description) values (4, 'ada manis manis nya');
insert into product_descriptions (id, description) values (5, 'minuman kesegaran');
insert into product_descriptions (id, description) values (6, 'bukanya susah, isisnya ga seberapa');
insert into product_descriptions (id, description) values (7, 'kenyal kenyal');
insert into product_descriptions (id, description) values (8, 'roti tapi punya item item');
-- 1.g Insert 3 payment methods.
insert into payment_methods (id, name, status)
    values(1, 'Debit ATM', 1);
insert into payment_methods (id, name, status)
    values(2, 'Cash', 1);
insert into payment_methods (id, name, status)
    values(3, 'Paylater', 1);
-- 1.h Insert 5 user pada tabel user.
insert into users (id, name, status, dob, gender) values(1, 'Galang', 1, '1990-01-01', 'M');
insert into users (id, name, status, dob, gender) values(2, 'Ilmi', 1, '1990-01-02', 'F');
insert into users (id, name, status, dob, gender) values(3, 'Phil', 1, '1990-01-03', 'M');
insert into users (id, name, status, dob, gender) values(4, 'Melati', 1, '1990-01-04', 'F');
insert into users (id, name, status, dob, gender) values(5, 'As ad', 1, '1990-01-05', 'M');
-- 1.i Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
insert into transactions(id, user_id, payment_method_id, status, total_qty, total_price)
    values (1, 1, 1, 1, 3, 3000);
insert into transactions(id, user_id, payment_method_id, status, total_qty, total_price)
    values (2, 1, 1, 1, 3, 3000);
-- 1.j Insert 3 product di masing-masing transaksi.
insert into transaction_details (transaction_id, product_id, status, qty, price)
    values (1, 1, 'ok', 1, 1000);
insert into transaction_details (transaction_id, product_id, status, qty, price)
    values (1, 2, 'ok', 1, 1000);
insert into transaction_details (transaction_id, product_id, status, qty, price)
    values (1, 3, 'ok', 1, 1000);
insert into transaction_details (transaction_id, product_id, status, qty, price)
    values (2, 1, 'ok', 1, 1000);
insert into transaction_details (transaction_id, product_id, status, qty, price)
    values (2, 2, 'ok', 1, 1000);
insert into transaction_details (transaction_id, product_id, status, qty, price)
    values (2, 3, 'ok', 1, 1000);


-- 2.a Tampilkan nama user / pelanggan dengan gender Laki-laki / M. -- karena name tidak maka pakai id
SELECT name, gender FROM users
    WHERE gender = 'M';
-- 2.b Tampilkan product dengan id = 3.
SELECT * FROM products
    WHERE id = 3;
-- 2.c Tampilkan data pelanggan yang created_ at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata a'.
SELECT * FROM users
    WHERE created_at >= '2022-03-09' AND name like '%a%';
-- 2.d Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(*) as jumlah_pelangan_perempuan FROM users
    WHERE gender = 'F';
-- 2.e Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT * FROM users
    ORDER BY name;

-- JOIN, UNION, SUB-QUERY, FUNCTION
-- 1. Gabungkan data transaksi dari user id 1 dan user id 2.
SELECT * FROM transactions
    WHERE user_id IN (1, 2);
-- 2. Tampilkan jumlah harga transaksi user id 1.
SELECT SUM(total_price) as jumlah_harga_transaksi FROM transactions
    WHERE user_id = 1;
-- 3.