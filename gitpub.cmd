@echo off
set /p comment=Type Comment: 
git init
git add .
git commit -m "%comment%"
git push origin master

