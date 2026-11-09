#include "zad_hesh_tab.h"
#include "../../struct_cplas/hash/hesh_table.h"
#include "../../struct_cplas/hash/cuckoo_table.h"
#include "../../struct_cplas/Array/massiv.h"
#include <sstream>
using namespace std;

void SplitSymbols(const string& text, Array& symbols){
    if (text.empty()){
        return;
    }
    size_t start = 0;
    for (size_t i = 1; i < text.size(); ++i){
        unsigned char byte = text[i];
        if (byte >= 128 && byte <= 191){
            continue;
        }
        MAddEnd(symbols, text.substr(start, i - start));
        start = i;
    }
    MAddEnd(symbols, text.substr(start));
}

int LongestSubstring(const Array& symbols, int& best_left){
    HashTable positions;
    HInitTable(positions, 8);
    int left = 0;
    int maximum = 0;
    best_left = 0;
    for (int right = 0; right < MLength(symbols); ++right){
        string symbol = MGet(symbols, right);
        string saved;
        bool found;
        HGet(positions, symbol, saved, found);
        if (found){
            int previous = stoi(saved);
            if (previous >= left){
                left = previous + 1;
            }
        }
        HInsert(positions, symbol, to_string(right));
        int length = right - left + 1;
        if (length > maximum){
            maximum = length;
            best_left = left;
        }
    }
    HDestroy(positions);
    return maximum;
}

void RunLongestSubstring(const string& text){
    Array symbols;
    MInitArray(symbols, text.size());
    SplitSymbols(text, symbols);
    int start;
    int length = LongestSubstring(symbols, start);
    string substring;
    for (int i = start; i < start + length; ++i){
        substring += MGet(symbols, i);
    }
    cout << "Цепочки: " << length << " (" << substring << ")\n";
    MDestroy(symbols);
}

void DemoHashTables(const string& line){
    istringstream input(line);
    HashTable chain;
    CuckooTable cuck;
    HInitTable(chain, 8);
    string key;
    while (input >> key){
        HInsert(chain, key, "value");
        bool inserted;
        CInsert(cuck, key, inserted);
        if (!inserted){
            cout << "Кукушка: цикл при добавлении ключа " << key << '\n';
        }
    }
    cout << "Цепочки:\n";
    HPrintAddresses(chain);
    cout << "Кукушка:\n";
    CPrint(cuck);
    HDestroy(chain);
}
