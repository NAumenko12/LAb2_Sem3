#include "zad_array.h"
#include <algorithm>
#include <climits>
#include <iostream>
using namespace std;

int MaxPyramid(Array& blocks){
    int count = MLength(blocks);
    sort(blocks.data, blocks.data + count);
    int total = 0;
    int i = 0;
    while (i < count){
        string block = MGet(blocks, i);
        string width = block.substr(0, block.find(' '));
        int best = 0;
        while (i < count){
            block = MGet(blocks, i);
            size_t space = block.find(' ');
            if (block.substr(0, space) != width){
                break;
            }
            int height = stoi(block.substr(space + 1));
            if (height > best){
                best = height;
            }
            ++i;
        }
        if (best > INT_MAX - total){
            return -1;
        }
        total += best;
    }
    return total;
}

int RunPyramid(istream& input){
    int n;
    if (!(input >> n) || n < 1 || n > 100000){
        cout << "Ошибка: количество блоков от 1 до 100000\n";
        return 1;
    }
    Array blocks;
    MInitArray(blocks, n);
    for (int i = 0; i < n; ++i){
        int width, height;
        if (!(input >> width >> height) || width < 1 || height < 1){
            cout << "Ошибка: нужны положительные ширина и высота\n";
            MDestroy(blocks);
            return 1;
        }
        MAddEnd(blocks, to_string(width) + " " + to_string(height));
    }
    string extra;
    if (input >> extra){
        cout << "Ошибка: лишние данные после блоков\n";
        MDestroy(blocks);
        return 1;
    }
    int height = MaxPyramid(blocks);
    if (height < 0){
        cout << "Ошибка: сумма высот не помещается в int\n";
    } else {
        cout << "Максимальная высота: " << height << '\n';
    }
    MDestroy(blocks);
    return height < 0 ? 1 : 0;
}
