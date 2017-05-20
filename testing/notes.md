## 2017-05-20
- well, got ping to at least upload, but it's still not giving me much
  - I saw the message that it started the probe, but that was it
- still sorting out when I need to upload firmata to the device
  - I've been using the Arduino IDE, but I wonder if flipping between it and SublimeText might be getting in the way
- went back to try it again and now it doesn't go; just has TX & RX lit up
  - interesting - when I switch away from the console window and then back, it kills the run
  - it looks like it's initializing the firmata connection, the pins and the LED, then initializing the bot, then starting the connections, then nothing
  - blink responds the same way
  
