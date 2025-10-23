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
    $skip_unsafe_level = false;
    $skip_count = 0;
    $ever_skipped = false;
    $increasing = $row[1] > $prev;

    for ($i = 1; $i < count($row); ++$i) {
        $curr_increasing = $row[$i] > $prev;

        if ($increasing != $curr_increasing || $row[$i] == $prev) {
            if ($skip_unsafe_level) {
                return false;
            } else {
                $skip_count++;
                $skip_unsafe_level = true;
                $ever_skipped = $ever_skipped || true;

                if ($i == 1) {
                    $increasing = $row[$i + 1] > $row[0];
                }
                continue;
            }
        }

        $change = abs($row[$i] - $prev);

        if ($change > MAX_STABLE_CHANGE || $change <= 0) {
            if ($skip_unsafe_level) {
                return false;
            } else {
                $skip_count++;
                $skip_unsafe_level = true;
                $ever_skipped = $ever_skipped || true;

                if ($i == 1) {
                    $increasing = $row[$i + 1] > $row[0];
                }
                continue;
            }
        }

        if (!$skip_unsafe_level) {
            // Only update $prev if we are not skipping unsafe levels
            $prev = $row[$i];
        }
        $skip_unsafe_level = false;
    }

    if ($skip_count > 1) {
        return false;
    }

    if ($ever_skipped) {
        echo "Row is stable after skipping unsafe levels: ";
        echo "Skipped $skip_count unsafe levels\n";
        print_r($row);
        echo "\n";
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
        /* echo "Row is stable: "; */
        /* print_r($row); */
        /* echo "\n"; */
        $stable_count++;
    } else {
        /* echo "Row is not stable: "; */
        /* print_r($row); */
        echo "\n";
    }
}

echo "Number of stable rows: $stable_count\n";
