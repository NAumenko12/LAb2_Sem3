#include "zad_mnozh.h"
#include <sstream>
#include <iostream>
#include <string>
using namespace std;

void ReadSetValues(const string& values, Set& set, int& result){
    result = -1;
    if (values.find_first_not_of(" \t\r\n\v\f") == string::npos){
        result = 0;
        return;
    }
    if (values.back() == ','){
        return;
    } 
    istringstream list(values);
    string part;
    while (getline(list, part, ',')){
        istringstream input(part);
        string number, extra;
        if (!(input >> number)){
            return;
        } 
        if (input >> extra){
            return;
        } 
        for (char c : number){
            if (c < '0' || c > '9'){
                return;
            } 
        }
        size_t first = 0;
        while (first + 1 < number.size() && number[first] == '0'){
            ++first;
        }
        SetAdd(set, number.substr(first));
    }
    result = 0;
}

void ReadSets(const string& text, Set* sets, int& count){
    count = -1;
    istringstream input(text);
    char bracket;
    if (!(input >> bracket) || bracket != '{') return;
    if (!(input >> bracket)) return;

    int read = 0;
    while (bracket == '{'){
        string values;
        if (!getline(input, values, '}') || input.eof()) return;
        int result;
        ReadSetValues(values, sets[read], result);
        if (result == -1) return;
        ++read;

        if (!(input >> bracket)) return;
        if (bracket == '}') break;
        if (bracket != ',') return;

        if (!(input >> bracket) || bracket != '{') return;
    }

    if (bracket != '}') return;
    if (input >> bracket) return;
    count = read;
}

int CountCommon(const Set& first, const Set& second){
    Array elements;
    MInitArray(elements, SetSize(first));
    SetElements(first, elements);
    int common = 0;
    for (int i = 0; i < MLength(elements); ++i){
        string number = MGet(elements, i);
        if (SetAt(second, number)){
            ++common;
        }
    }
    MDestroy(elements);
    return common;
}

void PrintSetInline(const Set& set){
    Array elements;
    MInitArray(elements, SetSize(set));
    SetElements(set, elements);
    cout << '{';
    for (int i = 0; i < MLength(elements); ++i){
        if (i > 0) cout << ", ";
        cout << MGet(elements, i);
    }
    cout << '}';
    MDestroy(elements);
}

void PrintIntersectionPair(const Set& first, const Set& second, int common){
    cout << '{';
    PrintSetInline(first);
    cout << ", ";
    PrintSetInline(second);
    cout << "}, количество общих элементов: " << common;
    if (common > 0){
        Array elements;
        MInitArray(elements, SetSize(first));
        SetElements(first, elements);
        cout << " (";
        bool first_value = true;
        for (int i = 0; i < MLength(elements); ++i){
            string value = MGet(elements, i);
            if (SetAt(second, value)){
                if (!first_value) cout << ", ";
                cout << value;
                first_value = false;
            }
        }
        cout << ')';
        MDestroy(elements);
    }
    cout << '\n';
}

int RunIntersections(const string& text){
    int places = 0;
    for (char c : text){
        if (c == '{') ++places;
    }
    Set* sets = new Set[places];
    for (int i = 0; i < places; ++i){
        SetInit(sets[i]);
    }
    int count;
    ReadSets(text, sets, count);
    if (count < 2){
        if (count < 0){
            cerr << "Ошибка: нужна запись вида {{1,2,3},{2,3,4}}\n";
        } else {
            cerr << "Ошибка: нужно хотя бы два множества\n";
        }
    } else {
        int maximum = -1;
        for (int i = 0; i < count; ++i){
            for (int j = i + 1; j < count; ++j){
                int common = CountCommon(sets[i], sets[j]);
                if (common > maximum){
                    maximum = common;
                }
            }
        }
        for (int i = 0; i < count; ++i){
            for (int j = i + 1; j < count; ++j){
                if (CountCommon(sets[i], sets[j]) == maximum){
                    PrintIntersectionPair(sets[i], sets[j], maximum);
                }
            }
        }
    }
    for (int i = 0; i < places; ++i){
        SetDestroy(sets[i]);
    }
    delete[] sets;
    if (count < 2) return 1;
    return 0;
}
