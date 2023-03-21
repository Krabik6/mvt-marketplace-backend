create table sales
(
    id              serial primary key,
    "NftId"         integer not null,
    "NftCollection" text    not null,
    "Seller"        text    not null,
    "Cost"          integer not null,
    "Token"         text    not null,
    constraint uniqNft unique ("NftCollection", "NftId")
);