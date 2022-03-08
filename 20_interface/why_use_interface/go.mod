module main

go 1.17

replace (
    "main.com/fedex" => "./fedex"
    "main.com/koreaPost" => "./koreaPost"
)

require (
    "main.com/fedex" v0.0.0
    "main.com/koreaPost" v0.0.0
)