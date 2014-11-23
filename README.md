Message Queue in GO
===


Add new Message
---

Push "hello, world" to the pool `hello`

curl -v -X POST http://localhost:8080/push/hello -d "hello, world"


Read Messages
---

Get first channel message in the queue 

curl -v http://localhost:8080/peek/hello

