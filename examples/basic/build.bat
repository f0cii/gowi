@if exist "main-res.syso" (
    @del "main-res.syso"
)

windres -o main-res.syso main.rc
go build

pause
