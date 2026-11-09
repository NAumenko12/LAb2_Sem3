#include <iostream>
#include "binary_tree.h"
using namespace std;

void BInitTree(BTree& tree){
    tree.root = nullptr;
}


void BAdd(BNode*& root, int chislo){
    if (root == nullptr){
        root = new BNode{chislo, nullptr, nullptr};
        return;
    }
    if (chislo < root->data){
        if (root->left == nullptr){
            root->left = new BNode{chislo, nullptr, nullptr};
        } else {
            BAdd(root->left, chislo);
        }
    } else if (chislo > root->data){
        if (root->right == nullptr){
            root->right = new BNode{chislo, nullptr, nullptr};
        } else {
            BAdd(root->right, chislo);
        }
    }
}

void BDelete(BNode*& root, int chislo){
    if (root == nullptr){
        return;
    }
    if (chislo < root->data){
        BDelete(root->left, chislo);
        return;
    }
    if (chislo > root->data){
        BDelete(root->right, chislo);
        return;
    }
    if (root->left != nullptr && root->right != nullptr){
        BNode* rep = root->right;
        while (rep->left != nullptr){
            rep = rep->left;
        }
        root->data = rep->data;
        BDelete(root->right, rep->data);
        return;
    }
    BNode* old = root;
    if (root->left != nullptr){
        root = root->left;
    } else {
        root = root->right;
    }
    delete old;
}

BNode* BFind(BNode* root, int chislo){
    while (root != nullptr && root->data != chislo){
        if (chislo < root->data){
            root = root->left;
        } else {
            root = root->right;
        }
    }
    return root;
}

int BHeight(BNode* root){
    if (root == nullptr){
        return 0;
    }
    int left = BHeight(root->left);
    int right = BHeight(root->right);
    if (left > right){
        return left + 1;
    }
    return right + 1;
}

void BPrintPryam(BNode* root){
    if (root == nullptr) return;
    cout << root->data << ' ';
    BPrintPryam(root->left);
    BPrintPryam(root->right);
}
void BPrintSimetrich(BNode* root){
    if (root == nullptr) return;
    BPrintSimetrich(root->left);
    cout << root->data << ' ';
    BPrintSimetrich(root->right);
}
void BPrintObrat(BNode* root){
    if (root == nullptr) return;
    BPrintObrat(root->left);
    BPrintObrat(root->right);
    cout << root->data << ' ';
}

bool BPrintLevel(BNode* root, int level){
    if (root == nullptr) return false;
    if (level == 0){
        cout << root->data << ' ';
        return true;
    }
    bool left = BPrintLevel(root->left, level - 1);
    bool right = BPrintLevel(root->right, level - 1);
    return left || right;
}
void BPrintBFS(BNode* root){
    int level = 0;
    while (BPrintLevel(root, level)){
        ++level;
    }
}

void BPrintAddresses(BNode* root){
    if (root == nullptr) return;
    cout << "Значение: " << root->data << " | адрес: " << root << " | left: " << root->left << " | right: " << root->right << '\n';
    BPrintAddresses(root->left);
    BPrintAddresses(root->right);
}

void BClear(BNode*& root){
    while (root != nullptr){
        if (root->left != nullptr){
            BNode* left = root->left;
            root->left = left->right;
            left->right = root;
            root = left;
        } else {
            BNode* old = root;
            root = root->right;
            delete old;
        }
    }
}
