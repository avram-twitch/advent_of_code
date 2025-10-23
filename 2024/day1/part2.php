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
/* $filename = 'input.txt'; */
$filename = 'input.txt';
list($first_column, $second_column) = readTwoColumnFile($filename);

$frequencies = [];

// Gather number frequencies from second column
for ($i = 0; $i < count($second_column); ++$i) {
    $curr_key = strval($second_column[$i]);
    if (!array_key_exists($curr_key, $frequencies)) {
        $frequencies[$curr_key] = 0;
    }
    $frequencies[$curr_key]++;
}

// Sum up similarity score based on frequencies
$sum = 0;

for ($i = 0; $i < count($first_column); ++$i) {
    $curr_num = $first_column[$i];
    $curr_frequency = $frequencies[$curr_num];
    $mult = $curr_num * $curr_frequency;

    $sum += $mult; 
}

print_r($frequencies);
print_r($sum);
?>
