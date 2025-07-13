CRUD ke-4

POST
curl -X POST http://localhost:8080/orders \
-H "Content-Type: application/json" \
-d '{
  "customer_id": 1,
  "item": "Keyboard Mechanical",
  "quantity": 2,
  "status": "pending"
}'

GET
curl http://localhost:8080/orders

curl http://localhost:8080/orders/1


PUT
curl -X PUT http://localhost:8080/orders/1 \
-H "Content-Type: application/json" \
-d '{
  "customer_id": 1,
  "item": "Keyboard",
  "quantity": 1,
  "status": "paid"
}'

DELETE
curl -X DELETE http://localhost:8080/orders/1