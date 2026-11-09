#pragma once
#include <iostream>
using namespace std;


struct BNode{
    int data;
    BNode* left;
    BNode* right;
};
struct BTree{
    BNode* root;
};

void BInitTree(BTree& tree);
void BAdd(BNode*& root, int chislo);
void BDelete(BNode*& root, int chislo);
BNode* BFind(BNode* root, int chislo);
int BHeight(BNode* root);
void BPrintPryam(BNode* root);
void BPrintSimetrich(BNode* root);
void BPrintObrat(BNode* root);
bool BPrintLevel(BNode* root, int level);
void BPrintBFS(BNode* root);
void BPrintAddresses(BNode* root);
void BClear(BNode*& root);
