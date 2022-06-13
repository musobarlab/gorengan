## List Endpoint

You can use `Insomnia` as a GraphQL Client https://insomnia.rest/

graphql endpoint http://localhost:3000/graphql.
This endpoint is secured using `basic auth`, open the `.env` for details.

### Create Product

payload
```javascript
mutation($product: ProductInput!) {
  productMutation {
    createProduct(product: $product) {
        id
        name
    }
  }
}
```

arguments

```json
{
  "product": {
    "id": "1",
    "name": "NOKIA 6",
    "category": "1",
    "quantity": 4
  }
}
```

### Get Product (by ID)

payload
```javascript
query {
    productQuery {
        product(id: "1") {
            id
            name
            quantity
            category {
            id
            name
            }
            created
        }
    }
}
```

### Get Products

payload
```javascript
query ($args: ProductsArgs!) {
    productQuery {
        products(productsArgs: $args) {
            products{
                id
                name
                category{
                    id
                    name
                }
              }
              meta{
                page
                limit
                totalPages
                totalRecords
              }
        }
    }
}
```

arguments
```json
{
  "args": {
    "limit": 3,
    "page": 1,
    "sort": "asc",
    "orderBy": "name"
  }
}
```

### Delete Product

payload
```javascript
mutation {
    productMutation {
        deleteProduct(id: "2") {
            id
            name
        }
    }
}
```

### Create Category

payload
```javascript
mutation($category: CategoryInput!) {
    categoryMutation {
        createCategory(category: $category) {
            id
            name
        }
    }
}
```

arguments
```json
{
  "category": {
    "id": "4",
    "name": "Fashion"
  }
}
```
