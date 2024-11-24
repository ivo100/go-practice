package main

/*
#cgo CFLAGS: -framework CoreAudio -framework AudioToolbox -framework CoreServices
#include <CoreAudio/CoreAudio.h>
#include <AudioToolbox/AudioToolbox.h>
#include <CoreServices/CoreServices.h>

static void playSystemSound() {
    SystemSoundID soundID = 0;
    // Play system sound (example: beep sound)
    AudioServicesCreateSystemSoundID(CFURLCreateWithString(NULL, CFSTR("file:///System/Library/Sounds/Ping.aiff"), NULL), &soundID);
    AudioServicesPlaySystemSound(soundID);
}
*/
import "C"
import "fmt"

func main() {
	fmt.Println("Playing system sound...")
	// Call the C function to play the system sound
	C.playSystemSound()
}
