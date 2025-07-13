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