#pragma once
#include "../Array/massiv.h"

struct HufNode{
    unsigned char symbol;
    int weight;
    HufNode* left;
    HufNode* right;
};

void HufCountFrequencies(const string& text, int frequency[]);
HufNode* HufBuildTree(const int frequency[]);
void HufBuildCodes(HufNode* node, string code, Array& codes);
string HufEncode(const string& text, const Array& codes);
bool HufDecode(HufNode* root, const string& bits, string& text);
void HufDeleteTree(HufNode* node);
