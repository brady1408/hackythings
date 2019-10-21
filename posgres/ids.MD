# Creating the ID Generation in the Database

## The Magical Query

```sql
create schema shard_1;
create sequence shard_1.global_id_sequence;

CREATE OR REPLACE FUNCTION shard_1.id_generator(OUT result bigint) AS $$
DECLARE
    our_epoch bigint := 1546300800000; -- (2019-01-01 00:00:00.000)
    seq_id bigint;
    now_millis bigint;
    -- the id of this DB shard, must be set for each
    -- schema shard you have - you could pass this as a parameter too
    shard_id int := 1;
BEGIN
    SELECT nextval('shard_1.global_id_sequence') % 1024 INTO seq_id;

    SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000) INTO now_millis;
    result := (now_millis - our_epoch) << 23;
    result := result | (shard_id << 10);
    result := result | (seq_id);
END;
$$ LANGUAGE PLPGSQL;

select shard_1.id_generator();
```

## Breaking Down the Magical Query

```sql
-- our_epoch bigint := 1314220021721;
-- seq_id bigint;
-- now_millis bigint;
-- the id of this DB shard, must be set for each
-- schema shard you have - you could pass this as a parameter too
-- shard_id int := 1;

SELECT 1 % 1024; --> 1

SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000); --> 1571373386441
SELECT 1571373386441 - 1314220021721; --> 257153364720 --> 0000000000000000000000000011101111011111100010001111001011110000
SELECT 257153364720 << 23; --> 2157158772517109760 --> 0001110111101111110001000111100101111000000000000000000000000000
SELECT 1 << 10; --> 1024 --> 0000010000000000
SELECT 2157158772517109760 | 1024; --> 2157158772517110784 --> 0001110111101111110001000111100101111000000000000000010000000000
SELECT 2157158772517110784 | 1;    --> 2157158772517110785 --> 0001110111101111110001000111100101111000000000000000010000000001
```

## Example of a function to extract the service id

```sql
CREATE OR REPLACE FUNCTION nav_service_id_1.extract_service_id(id bigint) returns bigint AS $$
BEGIN
return ((id::bit(64) & B'0000000000000000000000000000000000000000011111111111110000000000') >> 10)::bigint;
-- Could also use the following but I worry about the readability.
-- return ((id & 8387584) >> 10);
END;
$$ LANGUAGE PLPGSQL;
```

## Example of a function to extract the UTC created time

```sql
CREATE OR REPLACE FUNCTION nav_service_id_1.extract_id_time(id bigint) returns timestamp as $$
BEGIN
return TIMESTAMP WITH TIME ZONE 'epoch' + (((id >> 23)+1314220021721)/1000) * INTERVAL '1 second';
END;
$$ LANGUAGE PLPGSQL;
```
