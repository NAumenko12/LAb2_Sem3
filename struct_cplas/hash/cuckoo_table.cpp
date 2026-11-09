#include "cuckoo_table.h"
#include <iostream>

int CHash1(const string& key){
    int hash = 0;
    for (unsigned char byte : key){
        hash = (hash * 31 + byte) % 11;
    }
    return hash;
}

int CHash2(const string& key){
    int hash = 0;
    for (unsigned char byte : key){
        hash = (hash * 17 + byte + byte / 11) % 11;
    }
    return hash;
}

void CContains(const CuckooTable& table, const string& key, bool& found){
    int first_index = CHash1(key);
    int second_index = CHash2(key);
    found = (table.first[first_index].used && table.first[first_index].key == key) || (table.second[second_index].used && table.second[second_index].key == key);
}

void CInsert(CuckooTable& table, const string& key, bool& inserted){
    CContains(table, key, inserted);
    if (inserted) return;
    CuckooTable backup = table;
    string current = key;
    for (int i = 0; i < 20; ++i){
        int first_index = CHash1(current);
        if (!table.first[first_index].used){
            table.first[first_index].key = current;
            table.first[first_index].used = true;
            inserted = true;
            return;
        }
        string displaced = table.first[first_index].key;
        table.first[first_index].key = current;
        current = displaced;

        int second_index = CHash2(current);
        if (!table.second[second_index].used){
            table.second[second_index].key = current;
            table.second[second_index].used = true;
            inserted = true;
            return;
        }

        displaced = table.second[second_index].key;
        table.second[second_index].key = current;
        current = displaced;
    }
    table = backup;
}

void CDelete(CuckooTable& table, const string& key){
    int first_index = CHash1(key);
    int second_index = CHash2(key);
    if (table.first[first_index].used && table.first[first_index].key == key){
        table.first[first_index].used = false;
    }
    if (table.second[second_index].used && table.second[second_index].key == key){
        table.second[second_index].used = false;
    }
}

void CPrint(const CuckooTable& table){
    cout << "Индекс | первая таблица | вторая таблица\n";
    for (int i = 0; i < 11; ++i){
        const CuckooCell& first = table.first[i];
        const CuckooCell& second = table.second[i];
        cout << i << " | " << (first.used ? first.key : "-") << " | " << (second.used ? second.key : "-") << '\n';
    }
}
