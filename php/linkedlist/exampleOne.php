<?php

namespace linkedlist;

require_once './singleLinkedListNode.php';
require_once './singleLinkedList.php';


$list = new SingleLinkedList();

$list->insert('a');

$list->insert('b');
$list->insert('c');
$list->insert('f');
$list->insert('c');
$list->insert('b');
$list->insert('a');


$list->printList();

var_dump(isPalindrome($list));


/**
 * 判断字符串是否是回文
 * 
 * 
 */
function isPalindrome(SingleLinkedList $list){
    if ($list->getLength() <= 1) {
        return true;
    }

    $pre = null;
    $slow = $list->head->next;
    $fast = $list->head->next;
    $remainNode = null;

    while($fast->next != null && $fast != null){
        $fast = $fast->next->next; //快指针每次走两步

        //单链表 反转 关键代码，需要三个指针
        $remainNode = $slow->next; //
        $slow->next = $pre; //将当前慢指针节点的next ，指向之前的指针
        $pre = $slow; //将前一个指针重新赋值，向前走一步，反转的部分
        $slow = $remainNode; //将当前慢指针，向前移动一步
    }

    //链表长度为偶数的情况
    if ($fast != null) {
        $slow = $slow->next; //slow 需要向前走一步，然后
    }

    //开始逐个比较
    while ($slow != null) {
        if ($slow->data != $pre->data) {
            return false;
        }
        $slow = $slow->next;
        $pre = $pre->next;
    }
    return true;
}