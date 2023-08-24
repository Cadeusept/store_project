CREATE TABLE transactions (
    id          SERIAL NOT NULL UNIQUE,
    user_id     INT NOT NULL,
    user_email  VARCHAR(255) NOT NULL,
    amount      INT NOT NULL,
    currency    VARCHAR(50) NOT NULL,
    created     TIMESTAMP NOT NULL,
    changed     TIMESTAMP NOT NULL,
    stat        VARCHAR(9) NOT NULL
);