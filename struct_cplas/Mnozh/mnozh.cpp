#include "mnozh.h"

void SetInit(Set& set){
    HInitTable(set.table, 8);
}

bool SetAt(const Set& set, const string& value){
    bool found;
    HContains(set.table, value, found);
    return found;
}

void SetAdd(Set& set, const string& value){
    HInsert(set.table, value, "");
}

void SetDel(Set& set, const string& value){
    HDelete(set.table, value);
}

int SetSize(const Set& set){
    return HLength(set.table);
}

void SetElements(const Set& set, Array& elements){
    for (int i = 0; i < set.table.mesto; ++i){
        HNode* node = set.table.data[i];
        while (node != nullptr){
            MAddEnd(elements, node->key);
            node = node->next;
        }
    }
}

void SetClear(Set& set){
    HClear(set.table);
}

void SetDestroy(Set& set){
    HDestroy(set.table);
}
