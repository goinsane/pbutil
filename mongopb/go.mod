module github.com/goinsane/pbutil/mongopb

go 1.13

replace github.com/goinsane/pbutil => ../

require (
	github.com/goinsane/pbutil v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.5.0
	google.golang.org/protobuf v1.26.0
)
