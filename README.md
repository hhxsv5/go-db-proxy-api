Database Proxy
==============

A simple API is used to connect the database which is in private/intranet network.

## Usage

1. Install

```Shell
//cd your-root-folder-of-project
//Create the file glide.yaml if not exist
//touch glide.yaml
glide get github.com/hhxsv5/go-db-proxy-api#~1.0.0
```

2. Config
- config of database: config/db.toml
- config of http server: config/http.toml


3. Run

```Go
var handlers = map[string]func(w http.ResponseWriter, r *http.Request){}
handlers["/"] = func(w http.ResponseWriter, r *http.Request) {
    //rand.Seed(time.Now().UnixNano())
    //cellphone := "1878020" + strconv.Itoa(int(1000+rand.Int31n(9999-1000)))
    //mydefault.CreateUser(cellphone)

    users := mydefault.GetUsersByIds([]uint64{})

    var (
        rsp string
        it  time.Time
        mt  time.Time
    )

    for i, u := range users {
        it = time.Unix(u.InsertTime, 0)
        mt = time.Unix(u.ModifyTime, 0)
        rsp += fmt.Sprintf("%d: %d, %s, %s, %s\n", i, u.Id, u.Cellphone, it.Format(time.RFC3339), mt.Format("2006-01-02 15:04:05"))
    }

    io.WriteString(w, rsp)
}

p := godpa.NewProxy(handlers)
p.Run()
```


## License

[MIT](https://github.com/hhxsv5/go-db-proxy-api/blob/master/LICENSE)
