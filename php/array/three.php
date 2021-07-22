<?php

//两个有序数组合并成一个

$a = [2,5,8,10];
$b = [1,2,6,9,18];

$arr = combine($a, $b);

print_r($arr);

function combine($a, $b){
    foreach($a as $k => $v){
        $len = count($b);
        for ($i=$len-1; $i >=0 ; $i--) { 
            if ($b[$i] < $v) {
                $b[$i+1] = $v;
                break;
            }else{
                $b[$i+1] = $b[$i];
            }
        }
    }
    return $b;
}



