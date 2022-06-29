GOOS=js GOARCH=wasm go build -o=html\\game.wasm .
zip -r html.zip .\\html
