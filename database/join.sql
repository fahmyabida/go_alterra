-- CONTOH JOIN LEFT
/*
    Left join itu mementingkan tabel sebelah kiri
    right join itu mementingkan tabel sebelah kanan
    join itu membuat irisan antara table kanan & kiri
    */
SELECT p.name as product_name,
       pt.name as product_type_name
    FROM products p
    LEFT JOIN product_types pt ON p.product_type_id = pt.id
    WHERE pt.id is null;
