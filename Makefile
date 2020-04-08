OUT_DIR := ./out

clean:
	rm -rf $(OUT_DIR)/*

build: clean
	./build/build.sh ./cmd/main.go ./out/nurse-schedule

run:
	go run ./cmd/main.go
