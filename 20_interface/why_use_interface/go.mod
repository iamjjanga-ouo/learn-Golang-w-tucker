module interface

go 1.17

replace (
    "interface.com/fedex" => "./fedex"
    "interface.com/koreaPost" => "./koreaPost"
)

require (
    "interface.com/fedex" v0.0.0
    "interface.com/koreaPost" v0.0.0
)