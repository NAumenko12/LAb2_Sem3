#pragma once
#include <iostream>
#include <string>
#include <stdexcept>
using namespace std;

struct SNode{
    string chislo;
    SNode* next;
};
struct Stack{
    SNode* head;
};

void SInitStack(Stack* stack);
bool SIsEmptyStack(Stack* stack);
void SPush(Stack& stack, string chislo);
void SPop(Stack& stack);
string STop(const Stack& stack);
void SPrint(SNode* flist);
void SPrintAddresses(const Stack& stack);
void SClear(Stack& stack);
