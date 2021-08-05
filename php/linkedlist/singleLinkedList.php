<?php

namespace linkedlist;


/**
 * 单链表
 * 
 */

class SingleLinkedList {

    //单链表头节点
    public $head;

    //单链表长度
    public $length;

    /**
     * 初始化单链表
     * 
     */
    public function __construct($head = null){
        if ($head == null) {
            $this->head = new SingleLinkedListNode();
        }else{
            $this->head = $head;
        }
        $this->length = 0;
    }

    /**
     * 获取单链表长度
     * 
     */
    public function getLength(){
        return $this->length;
    }

    /**
     * 插入数据 采用头插法 插入新数据
     * 
     * @data
     * 
     */
    public function insert($data){
        return $this->insertDataAfter($this->head, $data);
    }

    /**
     * 在某个节点后添加新的节点（直接插入数据）
     * 
     * 
     */ 
    public function insertDataAfter(SingleLinkedListNode $originNode, $data){
        if ($originNode == null) {
            return false;
        }
        //创建新的节点
        $newNode = new SingleLinkedListNode($data);
        // 新节点的下一个节点为源节点的下一个节点
        $newNode->next = $originNode->next;
        // 在originNode后插入newNode
        $originNode->next = $newNode;

        $this->length++;
        return $newNode;
    }

    /**
     * 在某个节点前插入新的节点（很少使用）
     * 
     * 
     */
    public function insertDataBefer(SingleLinkedListNode $originNode, $data){
        if ($originNode == null) {
            return false;
        }
        // 先找到originNode的前置节点，然后通过insertDataAfter插入
        $preNode = $this->getPreNode($originNode);

        $this->length++;
        return $this->insertDataAfter($preNode, $data);
    }

    /**
     * 获取某个节点的前置节点
     * 
     */
    public function getPreNode(SingleLinkedListNode $node){
        if ($node == null) {
            return false;
        }
        $curNode = $this->head;
        $preNode = $this->head;
        while ($curNode !== $node) { //条件，当前节点不等于 $node
            $preNode = $curNode;
            $curNode = $curNode->next; //$curNode 一步一步走遍所有节点
        }
        return $preNode;
    }

    /**
     * 删除节点
     * 
     * 
     */
    public function delete(SingleLinkedListNode $node){
        if ($node == null) {
            return false;
        }
        //先找到该节点的前置节点
        $preNode = $this->getPreNode($node);
        if (empty($preNode)) {
            return false;
        }
        $preNode->next = $node->next;

        $this->length--;
        return true;
    }

    /**
     * 通过索引获取节点
     * 
     * 
     */
    public function getNodeByIndex($index){
        if ($index >= $this->length) {
            return false;
        }
        $curNode = $this->head->next;
        for ($i = 0; $i < $index; $i++) { 
            $curNode = $curNode->next;
        }
        return $curNode;
    }

    /**
     * 在某个节点后插入新的节点
     * 
     */
    public function insertNodeAfter(SingleLinkedListNode $originNode, SingleLinkedListNode $node){
        if ($originNode == null) {
            return false;
        }
        $node->next = $originNode->next;
        $originNode->next = $node;
        $this->length++;
        return true;
    }

    /**
     * 输出单链表 当data的数据为可输出类型
     * 
     * 
     */
    public function printList(){
        if ($this->head->next == null) {
            return false;
        }
        $curNode = $this->head;
        // 防止链表带环，控制遍历次数
        $len = $this->getLength();
        while($curNode->next != null && $len--){
            echo $curNode->next->data. '-->';
            $curNode = $curNode->next;
        }
        echo 'NULL';
        return true;
    }


}