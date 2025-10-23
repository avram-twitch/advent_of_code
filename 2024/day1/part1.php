<?php
function readTwoColumnFile($filename) {
    $column1 = [];
    $column2 = [];

    // Open the file and read line by line
    $lines = file($filename, FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);

    foreach ($lines as $line) {
        // Split the line by whitespace
        $parts = preg_split('/\s+/', trim($line));
        
        // Ensure we have two columns
        if (count($parts) == 2) {
            $column1[] = intval($parts[0]);
            $column2[] = intval($parts[1]);
        }
    }

    return [$column1, $column2];
}

// Example usage
$filename = 'input.txt';
list($first_column, $second_column) = readTwoColumnFile($filename);

sort($first_column);
sort($second_column);

$diffs = [];

for ($i = 0; $i < count($first_column); ++$i) {
    $sum = abs($first_column[$i] - $second_column[$i]);
    echo "comparing {$first_column[$i]} to {$second_column[$i]}, got {$sum}\n";

    $diffs[] = $sum;
}

/* print_r($second_column); */
print_r(array_sum($diffs));
?>
