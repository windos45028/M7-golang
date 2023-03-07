module www

go 1.14

replace (
	backend_golang/package/lhccalendar v4.0.0+incompatible => gitlab.com/backend_golang/package/lhccalendar.git v4.0.0+incompatible
	backend_golang/package/library v3.0.0+incompatible => gitlab.com/backend_golang/package/Library.git v3.0.0+incompatible
)

require (
	backend_golang/package/lhccalendar v4.0.0+incompatible // indirect
	backend_golang/package/library v3.0.0+incompatible // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/streadway/amqp v1.0.0 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	go.mongodb.org/mongo-driver v1.10.2 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
)
