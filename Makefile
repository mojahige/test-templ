run:
	go run .
generate:
	mise exec go -- templ generate & npm run generate:css
dev:
	templ generate --watch --proxy="http://localhost:4000/" --cmd="go run ." & npm run dev:css
