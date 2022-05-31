# ownpage

Trivial CMS  
Uses contenteditable for editing  
Statically rendered  
Admin features use TS, JSX  
Postgres DB  
Register, login, bcrypt passwords, cookie sessions
CRUD Posts, Events, Pages, Site


![ownpage screnshot](/ownpage.png)
 
## To run:
```
go run main.go
```

```
esbuild static/ts/main.tsx --bundle --outfile=static/js/main.js --sourcemap --watch --jsx
```

```
sass ./static/scss/main.scss ./static/css/main.css --watch
```


