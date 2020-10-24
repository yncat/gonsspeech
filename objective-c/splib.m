#import <Foundation/Foundation.h>
#import <AppKit/AppKit.h>

NSSpeechSynthesizer *speechSynth = NULL;

int NsSpeechInit(){
if(speechSynth) return 0;
speechSynth=[[NSSpeechSynthesizer alloc] initWithVoice:nil];
[speechSynth setVoice:@"com.apple.speech.synthesis.voice.kyoko"];
[speechSynth setRate:240];
[speechSynth setVolume:0.9];
return 1;
}

int NsSpeechFree(){
if(!speechSynth) return 0;
speechSynth=NULL;
return 1;
}

int NsSpeechSpeak(char *str){
if(!speechSynth) return 0;
NSString *nsstr = [NSString stringWithCString: str encoding:NSUTF8StringEncoding];
[speechSynth startSpeakingString:nsstr];
usleep(5000);
    return 1;
}

int NsSpeechStop(){
if(!speechSynth) return 0;
[speechSynth stopSpeaking];
    return 1;
}

int NsSpeechIsSpeaking(){
if(!speechSynth) return 0;
return [speechSynth isSpeaking] == YES ? 1 : 2;
}
