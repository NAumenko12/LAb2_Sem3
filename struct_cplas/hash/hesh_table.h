#pragma once
#include <iostream>
#include <string>
using namespace std;

struct HNode{
    string key;
    string value;
    HNode* next;
};

struct HashTable{
    HNode** data = nullptr;
    int size = 0;
    int mesto = 0;
};

void HInitTable(HashTable& table, int mesto);
int HHash(const string& key, int mesto);
void HResize(HashTable& table);
void HInsert(HashTable& table, const string& key, const string& value);
void HGet(const HashTable& table, const string& key, string& value, bool& found);
void HContains(const HashTable& table, const string& key, bool& found);
void HDelete(HashTable& table, const string& key);
int HLength(const HashTable& table);
void HPrint(const HashTable& table);
void HPrintAddresses(const HashTable& table);
void HClear(HashTable& table);
void HDestroy(HashTable& table);
