<?php

/**
 * 实现一个大小固定的有序数组，支持动态增删改操作
 */
class MyArrayTwo
{

    public $data;
    public $length;
    public $capacity; 
    
    function __construct($capacity)
    {
        if ($capacity < 1) {
            return false;
        }

        $this->data = [];
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

    //有序数组，插入值，需要从新排序，假设数组升序排列
    public function insert($value){
        if ($this->checkIsFull()) {
            return false;
        }
        if ($this->length > 0) {
            for ($i = $this->length-1; $i >= 0 ; $i--) { 
                if ($this->data[$i] < $value) {
                    $this->data[$i+1] = $value;
                    break;
                }else{
                    $this->data[$i+1] = $this->data[$i]; //数组元素向后移动一位
                    if ($i == 0) {
                        $this->data[$i] = $value;
                    }
                }
            }
        }else{
            $this->data[0] = $value;
        }
        
        $this->length++;
    }

    //删除一个
    public function delete($index){
        for ($i = 0; $i < $this->length; $i++) { 
            if ($i >= $index) {
                $this->data[$i] = $this->data[$i+1];
            }
        }
    }

    public function print(){
        print_r($this->data);
    }

}