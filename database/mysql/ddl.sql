create table operators (
    id int(11) primary key,
    name varchar(255),
    created_at timestamp default now(),
    updated_at timestamp
);

create table product_types (
    id int(11) primary key ,
    name varchar(255),
    created_at timestamp default now(),
    updated_at timestamp
);

create table user {
    id int,
    name text,
    is_active bool,
}
insert into user (id, name, is_active) values (1, 'fahmy', true)

create table product_descriptions
(
    id int(11),
    description text,
    created_at timestamp default now(),
    updated_at timestamp,
    primary key (id),
    foreign key (id) references products(id)
);

create table payment_methods (
    id int(11) primary key,
    name varchar(255),
    status smallint,
    created_at timestamp default now(),
    updated_at timestamp
);

create table users
(
    id int(11) primary key,
    name text,
    status smallint,
    dob date, -- date of birth
    gender char(1),
    created_at timestamp default now(),
    updated_at timestamp
);

create table products
(
    id int(11) primary key,
    product_type_id int(11) ,
    operator_id int(11),
    code varchar(50),
    name varchar(100),
    status smallint,
    created_at timestamp default now(),
    updated_at timestamp,
    constraint product_type_fk
        foreign key (product_type_id) references product_types(id),
    constraint operator_fk
        foreign key (operator_id) references operators(id)
);

create table transactions
(
    id int(11) primary key,
    user_id int(11),
    payment_method_id int(11),
    status varchar(10),
    total_qty int(11),
    total_price numeric(25,2),
    created_at timestamp default now(),
    updated_at timestamp,
    foreign key (user_id) references users(id),
    foreign key (payment_method_id) references payment_methods(id)
);

create table transaction_details
(
    transaction_id int(11),
    product_id int(11),
    status varchar(10),
    qty int(11),
    price numeric(25,2),
    created_at timestamp default now(),
    updated_at timestamp,
    primary key (transaction_id, product_id),
    foreign key (transaction_id) references transactions(id),
    foreign key (product_id) references products(id)
);