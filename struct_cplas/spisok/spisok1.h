#pragma once
#include <iostream>
#include <string>
using namespace std;

struct FNode{
    string chislo;
    FNode* next;
};
struct FList{
    FNode* head;
};

void FInitList(FList& list);

void FAddAfter(FNode* flist, string chislo);
void FAddFirst(FNode*& flist, string chislo);
void FAddEnd(FNode*& flist, string chislo);
void FAddBefore(FNode*& flist, FNode* target, string chislo);
void FAddIndex(FNode*& flist, string chislo, int index);
void FDeleteAfter(FNode* flist);
void FDeleteFirst(FNode*& flist);
void FDeleteEnd(FNode*& flist);
void FDeleteBefore(FNode*& flist, FNode* target);
void FDeleteIndex(FNode*& flist, int index);
void FDeleteValue(FNode*& flist, string chislo);
void FPrintBack(FNode* flist);
void FPrint(FNode* flist);
void FPrintAddresses(FNode* flist);
void FFind(FNode* flist, string chislo);
