DROP TABLE IF EXISTS urls;
CREATE TABLE ulrs (
                       short VARCHAR(10) PRIMARY KEY,
                       long TEXT UNIQUE
);