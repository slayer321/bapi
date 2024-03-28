curl -X POST \
  http://localhost:8000/bapi/books \
  -d '{
    "title": "The Great Gatsby","author": "F. Scott Fitzgerald"
  },
  {
    "title": "The Alchemist","author": "Paulo Coelho"
  },
  '




curl -i -s -X POST http://localhost:8001/services \
 --data name=example_service \
 --data url='http://192.168.1.116:7777'




curl -X POST \
  http://192.168.1.116:7777/books \
  -d '{
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald"
  }'


curl -X POST http://localhost:8001/upstreams/example_upstream/targets \
 --data target='httpbun.com:80'
curl -X POST http://localhost:8001/upstreams/example_upstream/targets \
 --data target='httpbin.org:80'
curl -X POST http://localhost:8001/upstreams/example_upstream/targets \
 --data target='192.168.1.116:7777'

