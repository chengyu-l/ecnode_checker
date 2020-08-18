module github.com/chengyu-l/cfs-ecnode-checker

go 1.14

require (
	github.com/chengyu-l/chubaofs v2.0.1-0.20200818073913-4039780a1f84+incompatible
	github.com/chubaofs/chubaofs v2.1.0+incompatible // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/tiglabs/raft v0.0.0-20200304095606-b25a44ad8b33 // indirect
)

replace github.com/chubaofs/chubaofs v2.1.0+incompatible => github.com/chengyu-l/chubaofs v2.0.1-0.20200818073913-4039780a1f84+incompatible // indirect
