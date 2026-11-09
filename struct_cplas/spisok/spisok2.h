#pragma once
#include <iostream>
#include <string>
using namespace std;

struct LNode{
    string chislo;
    LNode* next, *prev;
};
struct LList{
    LNode* head;
    LNode* tail;
};

void LInitList(LList& list);

void LPrint(LNode* flist);
void LPrintBac(LNode* last);
void LAddAfter(LList& list, LNode* flist, string chislo);
void LAddFirst(LList& list, string chislo);
void LAddEnd(LList& list, string chislo);
void LAddBefore(LList& list, LNode* flist, string chislo);
void LAddIndex(LList& list, string chislo, int index);
void LDeleteFirst(LList& list);
void LDeleteEnd(LList& list);
void LDeleteAfter(LList& list, LNode* flist);
void LDeleteBefore(LList& list, LNode* flist);
void LDeleteIndex(LList& list, int index);
void LDeleteNode(LList& list, LNode* flist);
void LDeleteValue(LList& list, string chislo);
void LFind(LNode* flist, string chislo);
void LPrintAddresses(LNode* first);
