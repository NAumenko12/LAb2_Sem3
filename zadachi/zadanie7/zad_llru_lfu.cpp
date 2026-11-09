#include "zad_llru_lfu.h"
#include "../../struct_cplas/hash/hesh_table.h"
#include "../../struct_cplas/spisok/spisok2.h"
#include <iostream>
using namespace std;

void LFUSet(HashTable& cache, LList* lists, int& min_frequency, int capacity, const string& key, const string& value){
    if (capacity == 0) return;
    string saved;
    bool found;
    HGet(cache, key, saved, found);
    if (found){
        int frequency = stoi(saved.substr(saved.rfind(':') + 1));
        LDeleteValue(lists[frequency], key);
    } else if (HLength(cache) == capacity){
        string removed = lists[min_frequency].head->chislo;
        LDeleteFirst(lists[min_frequency]);
        HDelete(cache, removed);
    }
    HInsert(cache, key, value + ":1");
    LAddEnd(lists[1], key);
    min_frequency = 1;
}

void LFUGet(HashTable& cache, LList* lists, int& min_frequency,const string& key, string& value){
    string saved;
    bool found;
    HGet(cache, key, saved, found);
    if (!found){
        value = "-1";
        return;
    }
    size_t colon = saved.rfind(':');
    value = saved.substr(0, colon);
    int frequency = stoi(saved.substr(colon + 1));
    LDeleteValue(lists[frequency], key);
    LAddEnd(lists[frequency + 1], key);
    if (frequency == min_frequency && lists[frequency].head == nullptr){
        ++min_frequency;
    }
    HInsert(cache, key, value + ':' + to_string(frequency + 1));
}

void LFUDestroy(HashTable& cache, LList* lists, int q){
    for (int i = 0; i < q; ++i){
        while (lists[i + 1].head != nullptr){
            LDeleteFirst(lists[i + 1]);
        }
    }
    delete[] lists;
    HDestroy(cache);
}

int RunLFU(istream& input){
    int capacity, q;
    if (!(input >> capacity >> q) || capacity < 0 || q < 0){
        cout << "Ошибка: нужны неотрицательные числа\n";
        return 1;
    }
    HashTable cache;
    HInitTable(cache, 8);
    LList* lists = new LList[q + 1u]{};
    int min_frequency = 0;
    string result;
    int i;
    for (i = 0; i < q; ++i){
        string command, key, value;
        if (!(input >> command >> key)){
            cout << "Ошибка: запросов меньше, чем указано\n";
            break;
        }
        if (command == "SET"){
            if (!(input >> value)){
                cout << "Ошибка: нужно значение\n";
                break;
            }
            LFUSet(cache, lists, min_frequency, capacity, key, value);
        } else if (command == "GET"){
            LFUGet(cache, lists, min_frequency, key, value);
            if (!result.empty()) result += ' ';
            result += value;
        } else {
            cout << "Ошибка: неизвестная команда\n";
            break;
        }
    }
    bool correct = i == q;
    string extra;
    if (correct && (input >> extra)){
        cout << "Ошибка: лишние данные после запросов\n";
        correct = false;
    }
    if (correct) cout << "Результат: " << result << '\n';
    LFUDestroy(cache, lists, q);
    return correct ? 0 : 1;
}
