<h2>Project Description</h2>

Let us introduce our application, named Konis. Konis provides you to order the food and beverages product, especially coffee shop product remotely. The repository utilized Gin as a framework and Go language for development process.

Using Konis website make the ordering process being easier, which the buyer who want to get the product from the distance can order remotely. Konis provided the well-mechanism order process that make user conveniently choose and buy the product.

<h2>Tech Stack</h2>

- Backend's Programming Language: Go
- Framework : Gin-Gonic
- Package Manager : Golang PGX
- Data Migration : Golang Migrate
- RDBMS : PostgreSQL
- API Testing : ThunderClient
- Containerization : Docker

<h2>Config / installation process</h2>

<h3>1. Clone this repository</h3>

```sh
  git clone https://github.com/fajryalvin12/fgh21-go-event-organizer.git
  cd <project-name>
```

<h3>2. Open in VSCode</h3>

```sh
  code .
```

<h3>3. Install all the dependencies</h3>

```sh
  go mod tidy
```

<h3>4. Run the program</h3>

```sh
  go run main.go
```

<h2>API References</h2>

## Login

```http
  POST auth/login
```

## Register

```http
  POST auth/register
```

| Parameter                      | Type     | Description                                                                      |
| :----------------------------- | :------- | :------------------------------------------------------------------------------- |
| `users`                        | `GET`    | `Get a list of users data`                                                       |
| `users/:id`                    | `GET`    | `Select the user data according to registered id`                                |
| `users`                        | `POST`   | `Create new user data`                                                           |
| `users/:id`                    | `PATCH`  | `Edit the selected user data`                                                    |
| `users/:id`                    | `DELETE` | `Remove the selected user data`                                                  |
| `users/insertuser`             | `POST`   | `Create new user data with existing profile`                                     |
| `carts`                        | `GET`    | `Get a list of users carts`                                                      |
| `carts`                        | `DELETE` | `Remove the selected user's carts data`                                          |
| `carts/:id`                    | `POST`   | `Create new carts data`                                                          |
| `categories`                   | `GET`    | `Get a list of categories data`                                                  |
| `order-type`                   | `GET`    | `Get a list of order types data`                                                 |
| `order-type/:id`               | `GET`    | `Select the order type according to registered id`                               |
| `products/`                    | `GET`    | `Get a list of products data, including with pagination`                         |
| `products/filter/`             | `GET`    | `Get a list of filtered products data, including with pagination`                |
| `products/filter/price`        | `GET`    | `Get a list of filtered products data by price range`                            |
| `products/our-product`         | `GET`    | `Get a list of the all product, including with pagination`                       |
| `products/our-product/:id`     | `GET`    | `Get a list of the selected product by registered id, including with pagination` |
| `products/:id`                 | `GET`    | `Get a list of product data by registered id`                                    |
| `products/:id`                 | `PATCH`  | `Edit the selected product data`                                                 |
| `products/:id`                 | `DELETE` | `Remove the selected product data`                                               |
| `products/productSizes`        | `GET`    | `Get a list of product sizes data`                                               |
| `products/productSizes/:id`    | `GET`    | `Get a list of product sizes data by registered id`                              |
| `products/variant`             | `GET`    | `Get a list of product variants data`                                            |
| `products/variant/:id`         | `GET`    | `Get a list of product variants data by registered id`                           |
| `products/categoryProduct`     | `GET`    | `Get a list of product categories data`                                          |
| `products/categoryProduct/:id` | `GET`    | `Get a list of product categories data by registered id`                         |
| `carts/upload/img/:id`         | `POST`   | `Upload the product's images`                                                    |
| `profile/login`                | `GET`    | `Get a list of profile data according to registered id`                          |
| `profile`                      | `PATCH`  | `Edit the selected profile data`                                                 |
| `profile`                      | `GET`    | `Get a lists of profile data`                                                    |
| `profile`                      | `POST`   | `Create a profile, reference to registered user`                                 |
| `profile/:id`                  | `PATCH`  | `Edit the selected profile data`                                                 |
| `profile/:id`                  | `GET`    | `Get a profile data by registered user`                                          |
| `profile/:id`                  | `DELETE` | `Remove the selected profile data`                                               |
| `profile/img`                  | `PATCH`  | `Upload images for registered user`                                              |
| `profile/img/:id`              | `PATCH`  | `Upload images for registered admin`                                             |
| `promo`                        | `GET`    | `Get a lists of promo data`                                                      |
| `promo/:id`                    | `GET`    | `Get a promo data by registered id`                                              |
| `roles`                        | `GET`    | `Get a lists of roles data`                                                      |
| `roles/:id`                    | `GET`    | `Get a roles data by registered id`                                              |
| `testimonials`                 | `GET`    | `Get a lists of testimonials data`                                               |
| `transaction`                  | `POST`   | `Create a  new transactions, reference to registered user`                       |
| `transaction`                  | `GET`    | `Get a lists of transactions data, according to registered user`                 |
| `transaction/admin`            | `GET`    | `Get a lists of transactions data`                                               |
| `transaction/:id`              | `GET`    | `Get a lists of transactions data, according to registered id`                   |
| `transaction/status`           | `GET`    | `Get a lists of transactions data, according to their status`                    |
| `transaction/:id`              | `DELETE` | `Remove the selected transaction data`                                           |
| `transaction/products/:id`     | `GET`    | `Get a list of selected transaction data, reference to the product id`           |
| `transaction-status/:id`       | `GET`    | `Get a lists of transaction status data, according to registered id`             |
| `transaction`                  | `GET`    | `Get a lists of transactions status data`                                        |
| `transaction/update/:id`       | `PATCH`  | `Edit the selected transaction status data`                                      |

## Contributing

Feel free to contribute the repo for better code!

## Authors

- Fajry Alvin Hidayat (Product Owner)
- Thofikh Bisyron (Scrum Master)
- Muhammad Ashari (PIC Repository)
- Daffa Abiyyu Atha (Development Team)
- Muhammad Fariq Maasir (Development Team)
- Syarif Khalid Attamimi (Development Team)
- M. Ilyas Nazhif Azhar Al Qordhowi (Development Team)

## Feedback

If you have any feedback, please reach out to us at fajryalvin12@gmail.com
