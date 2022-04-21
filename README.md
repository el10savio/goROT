# goROT
Run ROT13 Encryption On The Command Line

# ROT13 Encryption
ROT13 is a letter substitution cipher that replaces a letter with the 13th letter after it in the alphabet.
Since moving 13 places twices across the alphabet set returns back to the same letter, running rot13 twices yields the orginal message

# Usage

To clone and build the repo
```
git clone https://github.com/el10savio/goROT && cd goROT
go build .
```

Encrypting/Decrypting a message
```
./goROT encrypt --message='Why did the chicken cross the road?'
Jul qvq gur puvpxra pebff gur ebnq?
```
