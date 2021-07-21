<?php

include "./two.php";

$capacity = 5;
$obj = new MyArrayTwo($capacity);

$obj->insert(5);
$obj->insert(2);
$obj->insert(10);
$obj->insert(8);
$obj->insert(5);

$obj->print();

$obj->delete(2);

$obj->print();