INSERT INTO "users" (name,username,password_hash) VALUES (
    'asadbek',
    'asadbek33',
    'a21code'
);

INSERT INTO "users" (name,username,password_hash,is_admin) VALUES (
    'Admin',
    'admin',
    'admin123',
     true
);

INSERT INTO "users" (name,username,password_hash) VALUES (
    'Nodir',
    'nod32',
    'nod32nodir'
);

INSERT INTO "users" (name,username,password_hash) VALUES (
    'Bobur',
    'sher03',
    'Bobursher'
);

INSERT INTO "products" ("product_name","category_id" ,"price","color","count", "image_url") VALUES ('scarf',1,20,'blue',20,'assets/images/img.jpg');
INSERT INTO "products" ("product_name","category_id" ,"price","color","count", "image_url") VALUES ('suit1',2,200,'dark',5,'assets/images/img2.jpg');
INSERT INTO "products" ("product_name","category_id" ,"price","color","count", "image_url") VALUES ('army boots',3,70,'orange',10,'assets/images/img1.jpg');






INSERT INTO "categories" ("name") VALUES ('accessories');
INSERT INTO "categories" ("name") VALUES ('suits');
INSERT INTO "categories" ("name") VALUES ('shoes');


INSERT INTO "sizes" ("size_num") VALUES (28);
INSERT INTO "sizes" ("size_num") VALUES (30);
INSERT INTO "sizes" ("size_num") VALUES (32);
INSERT INTO "sizes" ("size_num") VALUES (34);
INSERT INTO "sizes" ("size_num") VALUES (36);
INSERT INTO "sizes" ("size_num") VALUES (38);
INSERT INTO "sizes" ("size_num") VALUES (40);
