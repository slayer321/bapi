bapi 
---

bapi is a simple CRUD API build using gin.

Below is the post request which is proxy using kong

```
curl -X POST \
  http://localhost:8000/bapi/books \
  -d '{
    "title": "The Great Gatsby","author": "F. Scott Fitzgerald"
  },'
```