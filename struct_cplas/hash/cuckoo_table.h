#pragma once
#include <string>
using namespace std;

struct CuckooCell{
    string key;
    bool used = false;
};

struct CuckooTable{
    CuckooCell first[11];
    CuckooCell second[11];
};

int CHash1(const string& key);
int CHash2(const string& key);
void CInsert(CuckooTable& table, const string& key, bool& inserted);
void CContains(const CuckooTable& table, const string& key, bool& found);
void CDelete(CuckooTable& table, const string& key);
void CPrint(const CuckooTable& table);
