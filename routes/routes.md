### /signup

```bash
curl --include -X POST http://localhost:4000/signup \
    -H "Content-Type: application/json" \
    -d '{
    "name": "Test",
    "email": "test@gmail.com",
    "password": "password"
  }'
```
