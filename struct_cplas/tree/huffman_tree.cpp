#include "huffman_tree.h"
#include <utility>
using namespace std;

bool HufIsLeaf(HufNode* node){
    return node->left == nullptr && node->right == nullptr;
}

void HufCountFrequencies(const string& text, int frequency[]){
    for (unsigned char symbol : text){
        ++frequency[symbol];
    }
}

HufNode* HufBuildTree(const int frequency[]){
    HufNode* nodes[256];
    int count = 0;
    for (int i = 0; i < 256; ++i){
        if (frequency[i] > 0){
            nodes[count] = new HufNode{(unsigned char)i, frequency[i], nullptr, nullptr};
            ++count;
        }
    }
    if (count == 0) return nullptr;
    while (count > 1){
        int min1 = -1;
        int min2 = -1;
        for (int i = 0; i < count; ++i){
            if (min1 == -1 || nodes[i]->weight < nodes[min1]->weight){
                min2 = min1;
                min1 = i;
            } else if (min2 == -1 || nodes[i]->weight < nodes[min2]->weight){
                min2 = i;
            }
        }
        HufNode* left = nodes[min1];
        HufNode* right = nodes[min2];
        HufNode* parent = new HufNode{0, left->weight + right->weight, left, right};
        if (min1 > min2){
            swap(min1, min2);
        }
        nodes[min1] = parent;
        for (int i = min2; i < count - 1; ++i){
            nodes[i] = nodes[i + 1];
        }
        --count;
    }
    return nodes[0];
}

void HufBuildCodes(HufNode* node, string code, Array& codes){
    if (node == nullptr) return;
    if (HufIsLeaf(node)){
        if (code.empty()){
            code = "0";
        }
        string entry = string(1, node->symbol) + ":" + code;
        MAddEnd(codes, entry);
        return;
    }
    HufBuildCodes(node->left, code + "0", codes);
    HufBuildCodes(node->right, code + "1", codes);
}

string HufEncode(const string& text, const Array& codes){
    string bits;
    for (char c : text){
        for (int i = 0; i < MLength(codes); ++i){
            string entry = MGet(codes, i);
            if (entry[0] == c){
                bits += entry.substr(2);
                break;
            }
        }
    }
    return bits;
}

bool HufDecode(HufNode* root, const string& bits, string& text){
    text = "";
    if (root == nullptr) return bits.empty();
    if (HufIsLeaf(root)){
        for (char b : bits){
            if (b != '0'){
                return false;
            }
            text += root->symbol;
        }
        return true;
    }
    HufNode* cur_node = root;
    for (char b : bits){
        if (b == '0'){
            cur_node = cur_node->left;
        } else if (b == '1'){
            cur_node = cur_node->right;
        } else {
            return false;
        }
        if (HufIsLeaf(cur_node)){
            text += cur_node->symbol;
            cur_node = root;
        }
    }
    return cur_node == root;
}

void HufDeleteTree(HufNode* node){
    if (node == nullptr) return;
    HufDeleteTree(node->left);
    HufDeleteTree(node->right);
    delete node;
}
