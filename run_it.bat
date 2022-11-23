@REM  run server
    @REM  use cmd.exe with the /K flag (instead of /C) to leave the server window open
 start cmd.exe /C "go run .\main.go server"

@REM run client
 go run .\main.go client
