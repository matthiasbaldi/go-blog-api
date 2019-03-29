# go-blog-api

## Development

Start your webdev server in auto refresh mode:

    gin -p 4000 run main.go

Start RethinkDB with Docker WebUI under [http://localhost:8080/#dashboard](http://localhost:8080/#dashboard):

    docker run -p 8080:8080 -p 28015:28015 -p 29015:29015 --name rethink1 rethinkdb