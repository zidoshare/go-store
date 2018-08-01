# user authentication

every action(need authentication) must have a token signed by the server in `request.header.token`,then the server will sign a new token to the response. client get the token from `response.header.token` finally.

## token

the token just contains `id` and `role` of the user.

# database design

## Item

| column name | type         | description                       |
|:------------|:-------------|:----------------------------------|
| id          | int          | primary key and auto increment    |
| created_at  | timestamp    | created time                      |
| updated_at  | timestamp    | updated time                      |
| deleted_at  | timestamp    | deleted time,used for soft delete |
| title       | varchar(100) | item name                         |
| price       | int          | item price                        |
| reserve     | int          | item reserve                      |
| cover       | varchar(255) | cover image                       |

## ItemInfo

| column name | type          | description                       |
|:------------|:--------------|:----------------------------------|
| id          | int           | primary key and auto increment    |
| created_at  | timestamp     | created time                      |
| updated_at  | timestamp     | updated time                      |
| deleted_at  | timestamp     | deleted time,used for soft delete |
| item_id     | int           | item id                           |
| content     | varchar(3000) | the rich text                     |

