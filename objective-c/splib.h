#ifndef __SPLIB__
#define __SPLIB__

int NsSpeechInit();
int NsSpeechFree();
int NsSpeechSpeak(char *str);
int NsSpeechStop();
int NsSpeechIsSpeaking();
#endif