// test client in C
#include <stdio.h>
#include <unistd.h>
#include "splib.h"

int main(){
printf("NsSpeech test client started.\n");
int ret;
ret=NsSpeechInit();
printf("Voice initialized ret=%d\n", ret);
if(ret==0) return 1;
printf("Speaking some text...\n");
sleep(1);
ret=NsSpeechSpeak("this is test.");
printf("ret=%d\n", ret);
while(NsSpeechIsSpeaking()!=2){
sleep(1);
}
printf("done!\n");
ret=NsSpeechFree();
printf("Voice uninitialized ret=%d\n", ret);
return 0;
}
