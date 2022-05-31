# ownpage

## About
Trivial CMS  
Statically rendered with Go Templates  
Admin features use TS, JSX  
Uses contenteditable for editing  
Postgres DB  
Register, login, bcrypt passwords, cookie sessions  
CRUD Posts, Events, Pages, Site  

## Screenshot

![ownpage screnshot](/ownpage.png)
 
## How to run
```
go run main.go
```

## For development
```
esbuild static/ts/main.tsx --bundle --outfile=static/js/main.js --sourcemap --watch --jsx
```

```
sass ./static/scss/main.scss ./static/css/main.css --watch
```


