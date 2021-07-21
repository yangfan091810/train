<?php

/**
 * 实现一个简单的支持动态扩容的数组
 */
class MyArray
{
    $public $data; //数据
    $public $length; //长度
    $public $capacity; //容量
    $public $factor = 2; //扩容倍数
    
    function __construct($capacity)
    {
        if ($capacity < 1) {
            return false;
        }
        $this->data = array();
        $this->capacity = $capacity;
        $this->length = 0;
    }

    //检测数组是否已满
    public function checkIsFull(){
        if ($this->length == ($capacity - 1)) {
            return true;
        }
        return false;
    }

    /**
     * 
     * 插入数组 
     */
    public function push($value)
    {
        if ($this->checkIsFull()) { //已经满了，需要扩容
            $this->dilatation();
        }
        $this->data[$this->length] = $value;
        $this->length += 1;
    }

    //扩容
    public function dilatation(){
        $newArr = [];
        foreach ($this->data as $key => $value) {
            $newArr[$key] = $value;
        }
        $this->data = $newArr;
        $this->capacity = $this->capacity * $this->factor;
    }

}