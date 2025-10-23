<?php

define("MAX_STABLE_CHANGE", 3);

function readRowsFromFile($filename) {
    $lines = file($filename, FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);
    $lines = array_map(function($line) {
        $string_list = explode(' ', $line);
        return array_map(function($string) {
            return intval($string);
        }, $string_list);
    }, $lines);
    return $lines;
}

function rowStable($row) {
    $prev = $row[0];
    $increasing = $row[1] > $prev;

    for ($i = 1; $i < count($row); ++$i) {
        $curr_increasing = $row[$i] > $prev;

        if ($increasing != $curr_increasing) {
            return false;
        }

        $change = abs($row[$i] - $prev);

        if ($change > MAX_STABLE_CHANGE || $change <= 0) {
            return false;
        }

        $prev = $row[$i];
    }

    return true;
}

// Example usage
/* $filename = 'input.txt'; */
$filename = 'input.txt';
$rows = readRowsFromFile($filename);

$stable_count = 0;

foreach ($rows as $row) {
    if (rowStable($row)) {
        echo "Row is stable: ";
        print_r($row);
        echo "\n";
        $stable_count++;
    } else {
        echo "Row is not stable: ";
        print_r($row);
        echo "\n";
    }
}

echo "Number of stable rows: $stable_count\n";
