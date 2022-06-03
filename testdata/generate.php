<?php

$csv = fopen(__DIR__ . "/test.csv", "w+");

foreach (token_get_all(file_get_contents(__DIR__ . "/test.php")) as $token) {
  if (is_string($token)) {
    fputcsv($csv, [0, $token, 0]);
  } else {
    fputcsv($csv, $token);
  }
}

fclose($csv);