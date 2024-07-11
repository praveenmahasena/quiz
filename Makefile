all: test
all: vet
all: package
all: package_race


test: vet
test: base_test

base_test:
	go test ./... -v

vet:
	go vet ./...

package: quiz

package_race: quiz_race

quiz:
	go build -o ./bin/quiz .

quiz_race:
	go build --race -o ./bin/quiz_race .

