
  <!-- create category sementara -->
  curl -X POST http://localhost:8080/api/categories -d '{"name": "Elektronik", "description": "Peralatan elektronik kantor"}'

  <!-- Get All categories sementara -->
  curl -X GET http://localhost:8080/api/categories

  <!-- Get category by id sementara -->
  curl -X GET http://localhost:8080/api/categories/1

  <!-- Update category by id sementara -->
  curl -X PUT http://localhost:8080/api/categories/1 -d '{"name" : "Peralatan Elektronik", "description" : "Peralatan elektronik yang digunakan di kantor"}'

  <!-- Delete category by id sementara -->
  curl -X DELETE http://localhost:8080/api/categories/3

  <!-- create item sementara request body form value-->
  curl -X POST http://localhost:8080/api/items \
  -F "name=Item 1" \
  -F "category=1" \
  -F "price=100" \
  -F "purchase_date=2024-01-12" \
  -F "photo_url=@./uploads/onepiece.jpg"

  <!-- read all item sementara -->
  curl -X GET http://localhost:8080/api/items

  <!-- read item by id sementara -->
  curl -X GET http://localhost:8080/api/items/1

  <!-- update item by id sementara -->
  curl -X PUT http://localhost:8080/api/items/1 -d '{"name" : "update 1"}'

  <!-- delete item by id sementara -->
  curl -X DELETE http://localhost:8080/api/items/2

  <!-- read investment by id -->
  curl -X GET http://localhost:8080/api/items/investment/1

  <!-- read all investment -->
  curl -X GET http://localhost:8080/api/items/investment