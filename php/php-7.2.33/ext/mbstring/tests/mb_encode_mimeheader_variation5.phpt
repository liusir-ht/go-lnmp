--TEST--
Test mb_encode_mimeheader() function : usage variations - Pass different data types to $indent arg
--SKIPIF--
<?php
extension_loaded('mbstring') or die('skip');
function_exists('mb_encode_mimeheader') or die("skip mb_encode_mimeheader() is not available in this build");
if (PHP_INT_SIZE != 8) die('skip 64-bit only');
?>
--FILE--
<?php
/* Prototype  : string mb_encode_mimeheader
 * (string $str [, string $charset [, string $transfer_encoding [, string $linefeed [, int $indent]]]])
 * Description: Converts the string to MIME "encoded-word" in the format of =?charset?(B|Q)?encoded_string?=
 * Source code: ext/mbstring/mbstring.c
 */

/*
 * Pass different data types to $indent argument to see how mb_encode_mimeheader() behaves
 */

echo "*** Testing mb_encode_mimeheader() : usage variations ***\n";

mb_internal_encoding('utf-8');

// Initialise function arguments not being substituted
$str = base64_decode('zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868zrXOvc6/LiAwMTIzNDU2Nzg5Lg==');
$charset = 'utf-8';
$transfer_encoding = 'B';
$linefeed = "\r\n";

//get an unset variable
$unset_var = 10;
unset ($unset_var);

// get a class
class classA
{
  public function __toString() {
    return "Class A object";
  }
}

// heredoc string
$heredoc = <<<EOT
hello world
EOT;

// get a resource variable
$fp = fopen(__FILE__, "r");

// unexpected values to be passed to $indent argument
$inputs = array(

       // int data
/*1*/  0,
       1,
       12345,
       -2345,

       // float data
/*5*/  10.5,
       -10.5,
       12.3456789000e10,
       12.3456789000E-10,
       .5,

       // null data
/*10*/ NULL,
       null,

       // boolean data
/*12*/ true,
       false,
       TRUE,
       FALSE,

       // empty data
/*16*/ "",
       '',

       // string data
/*18*/ "string",
       'string',
       $heredoc,

       // object data
/*21*/ new classA(),

       // undefined data
/*22*/ @$undefined_var,

       // unset data
/*23*/ @$unset_var,

       // resource variable
/*24*/ $fp
);

// loop through each element of $inputs to check the behavior of mb_encode_mimeheader()
$iterator = 1;
foreach($inputs as $input) {
  echo "\n-- Iteration $iterator --\n";
  var_dump( mb_encode_mimeheader($str, $charset, $transfer_encoding, $linefeed, $input));
  $iterator++;
};

fclose($fp);

echo "Done";
?>
--EXPECTF--
*** Testing mb_encode_mimeheader() : usage variations ***

-- Iteration 1 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 2 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 3 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 4 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 5 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66?=
 =?UTF-8?B?zrXOr868zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 6 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 7 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 8 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 9 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 10 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 11 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 12 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 13 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 14 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 15 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 16 --

Warning: mb_encode_mimeheader() expects parameter 5 to be integer, string given in %s on line %d
NULL

-- Iteration 17 --

Warning: mb_encode_mimeheader() expects parameter 5 to be integer, string given in %s on line %d
NULL

-- Iteration 18 --

Warning: mb_encode_mimeheader() expects parameter 5 to be integer, string given in %s on line %d
NULL

-- Iteration 19 --

Warning: mb_encode_mimeheader() expects parameter 5 to be integer, string given in %s on line %d
NULL

-- Iteration 20 --

Warning: mb_encode_mimeheader() expects parameter 5 to be integer, string given in %s on line %d
NULL

-- Iteration 21 --

Warning: mb_encode_mimeheader() expects parameter 5 to be integer, object given in %s on line %d
NULL

-- Iteration 22 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 23 --
string(115) "=?UTF-8?B?zpHPhc+Ez4wgzrXOr869zrHOuSDOtc67zrvOt869zrnOus+MIM66zrXOr868?=
 =?UTF-8?B?zrXOvc6/LiAwMTIzNDU2Nzg5Lg==?="

-- Iteration 24 --

Warning: mb_encode_mimeheader() expects parameter 5 to be integer, resource given in %s on line %d
NULL
Done
