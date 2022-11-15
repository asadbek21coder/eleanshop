CREATE TABLE "users" (
    "id" serial not null primary key,
    "name" varchar(128) not null,
    "username" varchar(64) not null unique,
    "password_hash" varchar(128) not null
);

CREATE TABLE "categories" (
    "id" serial not null primary key,
    name varchar(128) not null,
    UNIQUE(name)
);

CREATE TABLE "sizes" (
    "id" serial not null primary key,
    "size_num" int not null,
    UNIQUE("size_num")
);

CREATE TABLE "products" (
    "id" serial not null primary key,
    "product_name" varchar(256) not null,
    "category_id" int not null references categories(id),
    "price" bigint not null,
    "color" varchar(64),
    "count" int not null,
    UNIQUE ("product_name", "price", "color") 

);

CREATE TABLE "product_sizes" (
    "id" serial not null primary key,
    "product_id" int not null references products(id),
    "size_id" int not null references sizes(id)
);

CREATE TABLE "feedbacks" (
    "id" serial not null primary key,
    "user_id" int not null references users(id),
    "phone_number" varchar(128) not null,
    "email" varchar(128) not null,
    "text" varchar(256) not null,
    "product_id" int not null references products(id)
);

CREATE TABLE "questions" (
    "id" serial not null primary key,
    "name" varchar(128) not null,
    "phone_number" varchar(128) not null,
    "time" varchar(128) not null,
    "text" varchar(256) not null
);