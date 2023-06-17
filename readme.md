# Voice Assistant

I am creating a voice assistant that can directly speak from your microphone and provide responses generated by OPEN AI.

I am using Python to recognize incoming voice and convert it into text.

# Usage

clone this repository

You must install python before run and install needed dependency like this:
- sudo pip install speech_recognition

next

- go mod tidy
- go run main.go (if python not running automatically, run python in another terminal tab with command `pip3 main.py` instead py folder)
- if terminal says `Speak` let speak anything

NB:
Until now, I haven't found a voice recognition library in Golang, so I am using Python to convert voice into text. 
If you have and idea for recognition voice from golang, let me know, thanks


