# run templ generation in watch mode to detect all .templ files and 
# re-create _templ.txt files on change, then send reload event to browser. 
# Default url: http://localhost:7331
live/templ:
	go tool templ generate --watch --proxy="http://localhost:4001" --open-browser=false

# run air to detect any go file changes to re-build and re-run the server.
live/server:
	air
# run tailwindcss to generate the styles.css bundle in watch mode.
.PHONY: live/tailwind
live/tailwind:
	unbuffer npx tailwindcss -i ./ui/static/css/input.css -o ./ui/static/css/tailwind.css --watch

# run esbuild to generate the index.js bundle in watch mode.
# live/esbuild:
# 	npx --yes esbuild js/index.ts --bundle --outdir=assets/ --watch

# watch for any js or css change in the assets/ folder, then reload the browser via templ proxy.

# start all 5 watch processes in parallel.
live: 
	make -j5 live/templ live/tailwind live/server
