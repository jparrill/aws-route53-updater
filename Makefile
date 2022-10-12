test:
	gotestsum --format testname

coverage:
	gotestsum -- -coverprofile=cover.out ./...
	
