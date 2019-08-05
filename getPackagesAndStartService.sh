#!/bin/bash
if [[ $(lsof -t -i:4000) ]]; then
    echo "process is already started at port 4000 so killing it..."
    kill $(lsof -t -i:4000)
fi
#todo add a code to create go directory if not present before downloading
go get -v github.com/gorilla/handlers
go get -v github.com/gorilla/mux
go get -v github.com/jinzhu/gorm
go get -v github.com/jinzhu/gorm/dialects/sqlite

echo "installed handlers, mux, gorm, gorm/dialects/sqlite packages"

nohup go run *.go >> golog.out 2>&1 &
echo "Api running on port 4000..."