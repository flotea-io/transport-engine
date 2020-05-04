Run vagrant up and open in browser page https://api.engine.devel/v1/search
It is important to add ssl exception for this domain.

.
=======
Develop apps: 	
backend:	in /app/src/flt run bee run
frontend:	in /app/src/frontend npm run dev

.
=======
Domain hiearchy:

engine.devel -> nuxt frontend
phppgadmin.engine.devel -> DB admin
api.engine.devel -> API
wss.engine.devel -> web socket

.
=======
Sync params: 
# in conf/app.conf Agency can set own address to filter routers.
# remove all data before sync
go run sync.go clear-data