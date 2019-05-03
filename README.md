# go-blog-api

## Development

Start your webdev server in auto refresh mode.
Download air binary to your bin folder in the GO-Workspace.

[https://github.com/cosmtrek/air/releases](https://github.com/cosmtrek/air/releases)

And execute...

    air -c air.conf

Start RethinkDB with Docker WebUI under [http://localhost:8080/#dashboard](http://localhost:8080/#dashboard):

    docker run -p 8080:8080 -p 28015:28015 -p 29015:29015 --name rethink1 rethinkdb