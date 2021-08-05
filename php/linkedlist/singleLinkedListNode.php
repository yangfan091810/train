<?php

namespace linkedlist;

/**
 * 单链表节点类
 *  
 */

class SingleLinkedListNode {

    //节点数据域
    public $data;

    //节点指针域，指向下一个节点
    public $next;

    public function __construct($data = null){
        $this->data = $data;
        $this->next = null;
    }

}