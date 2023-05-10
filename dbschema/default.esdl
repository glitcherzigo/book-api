module default {
    type Book {
        required property title -> str;
        required property author -> str;
        required property genre -> str;
        required property release_date -> datetime;
        required property quantity -> int16;
        property sales -> bigint;
    }
}
