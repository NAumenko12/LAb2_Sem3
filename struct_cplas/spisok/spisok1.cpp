#include <iostream>
#include <string>
#include "spisok1.h"
using namespace std;

void FInitList(FList& list){
    list.head = nullptr;
}

void FAddAfter(FNode* flist, string chislo){
    FNode* new_node = new FNode{chislo, flist->next};
    flist->next = new_node;
}
void FAddFirst(FNode*& flist, string chislo){
    FNode* new_node = new FNode{chislo, flist};
    flist = new_node;
}
void FAddEnd(FNode*& flist, string chislo){
    FNode* new_node = new FNode{chislo, nullptr};
    FNode* cur_node = flist;
    if (flist == nullptr){
        flist = new_node;
        return;
    }
    while(cur_node->next !=nullptr){
        cur_node = cur_node->next;
    }
    cur_node->next = new_node;
}
void FAddBefore(FNode*& flist, FNode* target, string chislo){
    FNode* cur_node =flist;
    FNode* new_node = new FNode{chislo, target};
    while(cur_node->next != target){
        cur_node = cur_node->next;
    }
    cur_node->next=new_node;
}

void FAddIndex(FNode*& flist, string chislo, int index){
    FNode* cur_node = flist;
    if (index == 0){
        FAddFirst(flist, chislo);
        return;
    }
    for(int i=0; i<index-1;++i){
        cur_node = cur_node->next;
    }
    FNode* new_node = new FNode{chislo, cur_node->next};
    cur_node->next = new_node;
}

void FDeleteAfter(FNode* flist){
    FNode* delete_node = flist->next;
    flist->next = delete_node->next;
    delete delete_node;
}
void FDeleteFirst(FNode*& flist){
    FNode* delete_node = flist;
    flist = flist->next;
    delete delete_node;
}
void FDeleteEnd(FNode*& flist){
    if (flist->next == nullptr){
        FDeleteFirst(flist);
        return;
    }
    FNode* cur_node = flist;
    while (cur_node->next->next != nullptr){
        cur_node = cur_node->next;
    }
    FDeleteAfter(cur_node);
}
void FDeleteBefore(FNode*& flist, FNode* target){
    if (flist->next == target){
        FDeleteFirst(flist);
        return;
    }
    FNode* cur_node = flist;
    while (cur_node->next->next != target){
        cur_node = cur_node->next;
    }
    FDeleteAfter(cur_node);
}
void FDeleteIndex(FNode*& flist, int index){
    if (index == 0){
        FDeleteFirst(flist);
        return;
    }
    FNode* cur_node = flist;
    for(int i =0; i< index -1;++i){
        cur_node = cur_node->next;
    }
    FDeleteAfter(cur_node);
}
void FDeleteValue(FNode*& flist, string chislo){
    if (flist->chislo == chislo){
        FDeleteFirst(flist);
        return;
    }
    FNode* cur_node = flist;
    while (cur_node->next != nullptr && cur_node->next->chislo != chislo){
        cur_node = cur_node->next;
    }
    if (cur_node->next != nullptr){
        FDeleteAfter(cur_node);
    }
}

void FPrintBack(FNode* flist){
    if (flist == nullptr) return;
    FPrintBack(flist->next);
    cout << flist->chislo << "\n";
}
void FPrint(FNode* flist){
    while(flist != nullptr){
        cout << flist->chislo<<"\n";
        flist = flist->next;
    }
}
void FPrintAddresses(FNode* flist){
    while (flist != nullptr){
        cout << "Узел: " << flist << ", chislo: " << flist->chislo << ", next: " << flist->next << '\n';
        flist = flist->next;
    }
}
void FFind(FNode* flist, string chislo){
    while (flist != nullptr){
        if (flist->chislo == chislo){
            cout << chislo << " найден, адрес: " << flist << "\n";
            return;
        }
        flist = flist->next;
    }
    cout << chislo << " не найден\n";
}

