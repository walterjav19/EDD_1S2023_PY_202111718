module main

go 1.20

replace menus => ./menus

require (
	estructuras v0.0.0-00010101000000-000000000000
	menus v0.0.0-00010101000000-000000000000
	roles v0.0.0-00010101000000-000000000000
)

replace roles => ./roles

replace estructuras => ./estructuras
