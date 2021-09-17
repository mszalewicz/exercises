clear

if ( ($args[0] -eq "clean") -and (Test-Path -Path './build') )
{
    Write-Host "`n--------- Cleaning solution ---------" -ForegroundColor Cyan
    rm -r .\build\ 
}

if ( !(Test-Path -Path './build') )
{
    Write-Host "`n--- Bootstrapping CMake ---" -ForegroundColor Yellow
    mkdir build
    cd ./build  

    Write-Host "`n"
    cmake .. 
    cd ..
}

if($args -contains "release")
{
    Write-Host "`n--- Building ---`n" -ForegroundColor Yellow
    cmake --build .\build\ --config Release
}
else
{
    Write-Host "`n--- Building ---`n" -ForegroundColor Yellow
    cmake --build .\build\  
}

Write-Host "`n`n"