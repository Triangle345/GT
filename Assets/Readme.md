GT/Assets Readme
======

This directory is not implicitly used anywhere. To use it, simply add images to the Images folder here, and use the -Assets tag when running in go.

i.e.
- Add test.png to GT/Assets/Images/
- from the Examples directory run:
  - go run SimpleWindow.go -Assets "../Assets"
    
*This tag will allow you to override the local Assets default ("./Assets")*