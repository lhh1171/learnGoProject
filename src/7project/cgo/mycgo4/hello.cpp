#include<iostream>
using namespace std;
extern "C"{
    #include"hello.h"
}

void sayHello(char *s){
    cout<<s<<endl;
}