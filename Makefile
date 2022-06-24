createDb:
	createdb --username=postgres --owner=postgres go_finance

postgres:
	# docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres:14-alpine
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres