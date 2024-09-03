INSERT INTO roles ("name") VALUES
('Normal User'),
('Admin');

INSERT INTO users (email, password, role_id) VALUES
('admin@mail.com', '$argon2id$v=19$m=65536,t=3,p=4$GmFOWlxzqhkYzqH4/M137g$nbZhFlXQQjyTRJWoylwIOsqpYB7ArI4BrYN8mzxPo+w', 2),
('user@mail.com', '$argon2id$v=19$m=65536,t=3,p=4$dlYGg+W7/l0w9t/G+AGqmg$0EfIrnZrY83+kidm9s/E+QLXavxbAqfYYqp+0t+Lsl0', 1);