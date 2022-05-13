VERSION = "0.0.1"
change-version:
	@echo $(VERSION)>VERSION
	@echo "package constant\n\n// Version constant of encoding\nconst Version = \"$(VERSION)\"">constant/version.go
	@git add VERSION
	@git add constant/version.go
	@git commit -m "v$(VERSION)"
	@git tag -a "v$(VERSION)" -m "v$(VERSION)"
	@git push origin
	@git push origin "v$(VERSION)"

update-module:
	get get github.com/mkawserm/abesh

run-example:
	go run example/main/main.go embedded run
